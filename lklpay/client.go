package lklpay

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"hash"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imroc/req/v3"
	"github.com/ywanbing/pay/lklpay/common"
	"github.com/ywanbing/pay/log"
)

const (
	// TestUrl 测试环境地址
	TestUrl = "https://test.wsmsd.cn/sit"
	// ProdUrl 生产环境地址
	ProdUrl = "https://s2.lakala.com"
)

// Config 拉卡拉SDK配置
type Config struct {
	// TermNo          string `json:"term_no,omitempty"`           // 终端号(收银台可不送
	Appid           string `json:"appid,omitempty" `            // 接入方唯一编号
	SerialNo        string `json:"serial_no,omitempty"`         // 证书序列号
	SyncPublicKey   string `json:"sync_public_key,omitempty"`   // 异步通知验签证书公钥key (如果配置，就不会读取文件
	SignPrivateKey  string `json:"sign_private_key,omitempty"`  // 加签证书私钥key(如果配置，就不会读取文件
	SyncPubicPath   string `json:"sync_pubic_path,omitempty"`   // 异步通知验签证书公钥路径(拉卡拉分配
	SignPrivatePath string `json:"sign_private_path,omitempty"` // 加签证书私钥路径(接入方生成
}

type Option func(client *Client)

// WithLogger set custom logger
func WithLogger(logger log.Logger) Option {
	return func(client *Client) {
		client.log = logger
	}
}

// WithIsProd set isProd
func WithIsProd(isProd bool) Option {
	return func(client *Client) {
		client.isProd = isProd
	}
}

// WithDev set dev
func WithDev(dev bool) Option {
	return func(client *Client) {
		client.dev = dev
	}
}

// WithContext set context
func WithContext(ctx context.Context) Option {
	return func(client *Client) {
		client.ctx = ctx
	}
}

// WhitValid set valid
func WhitValid(valid *validator.Validate) Option {
	return func(client *Client) {
		client.valid = valid
	}
}

// WithHash set hash type
func WithHash(hashType crypto.Hash) Option {
	return func(client *Client) {
		client.hashType = hashType
	}
}

// WithNonceStrLen set nonceStr len
func WithNonceStrLen(nonceStrLen int) Option {
	return func(client *Client) {
		client.nonceStrLen = nonceStrLen
	}
}

// WithVerifyResp set verify resp
func WithVerifyResp(verifyResp bool) Option {
	return func(client *Client) {
		client.verifyResp = verifyResp
	}
}

// WithVerifyHttps set verify https
func WithVerifyHttps() Option {
	return func(client *Client) {
		client.verifyHttps = true
	}
}

// Client pay client
type Client struct {
	ctx context.Context // 上下文
	cfg Config          // 配置
	log log.Logger      // logger

	isProd      bool                // 是否生产环境
	dev         bool                // 是否调试模式
	nonceStrLen int                 // 随机字符串长度,默认12
	verifyResp  bool                // 是否响应验签
	verifyHttps bool                // 是否验证https
	valid       *validator.Validate // 参数校验
	cli         *req.Client

	lklCertificate *x509.Certificate // 拉卡拉公钥证书
	privateKey     *rsa.PrivateKey   // 自己的私钥
	hashType       crypto.Hash       // hash 类型
	hash           hash.Hash         // hash 计算
	mu             sync.Mutex        // 锁（由于签名需要一个一个的签名，所以需要加锁，既然签名都只能一个，那么我们hash也复用一个）
}

// New a pay client
func New(cfg Config, options ...Option) *Client {
	client := &Client{
		ctx:         context.Background(),
		cfg:         cfg,
		log:         log.DefLogger(),
		isProd:      false,
		verifyResp:  false,
		verifyHttps: false,
		nonceStrLen: 12,
		valid:       validator.New(),
		hashType:    crypto.SHA256,
	}

	for _, option := range options {
		option(client)
	}

	// 初始化客户端
	client.init()

	return client
}

func (c *Client) init() {
	c.initHttpClient()
	c.hash = c.hashType.New()

	// 解析公钥证书
	lklCertificate, err := c.getLklCertificate()
	if err != nil {
		panic(err)
	}
	c.lklCertificate = lklCertificate

	// 解析私钥
	privateKey, err := c.getPrivateKey()
	if err != nil {
		panic(err)
	}
	c.privateKey = privateKey
}

func (c *Client) initHttpClient() {
	cli := req.NewClient()
	if c.isProd {
		cli.SetBaseURL(ProdUrl)
	} else {
		cli.SetBaseURL(TestUrl)
	}

	cli.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !c.verifyHttps,
	})
	if c.dev {
		cli.DevMode()
	}
	cli.SetLogger(c.log)
	cli.SetCommonContentType("application/json")
	c.cli = cli
}

// getRsaSign 获取签名字符串
func (c *Client) getRsaSign(body []byte) (auth string, err error) {
	var (
		appid    = c.cfg.Appid
		ts       = time.Now().Unix()
		nonceStr = common.RandomString(c.nonceStrLen)
		serialNo = c.cfg.SerialNo
	)
	if appid == "" || nonceStr == "" || serialNo == "" {
		return "", fmt.Errorf("签名缺少必要的参数")
	}

	validStr := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", appid, serialNo, ts, nonceStr, body)
	c.log.Debugf("valid: %s", validStr)
	c.mu.Lock()
	defer c.mu.Unlock()
	defer c.hash.Reset()

	// 计算签名
	sign, err := common.Sign([]byte(validStr), c.privateKey, c.hashType, c.hash)
	if err != nil {
		return "", err
	}

	c.log.Debugf("sign: %s", sign)

	// 拼接签名
	return fmt.Sprintf(common.AuthFormat, common.Algorism_SHA256, appid, serialNo, ts, nonceStr, sign), nil
}

// getRsaSign 获取签名字符串
func (c *Client) getRsaNotifySign(body []byte) (auth string, err error) {
	var (
		ts       = time.Now().Unix()
		nonceStr = common.RandomString(c.nonceStrLen)
	)
	if nonceStr == "" {
		return "", fmt.Errorf("签名缺少必要的参数")
	}

	validStr := fmt.Sprintf("%d\n%s\n%s\n", ts, nonceStr, body)
	c.log.Debugf("valid: %s", validStr)
	c.mu.Lock()
	defer c.mu.Unlock()
	defer c.hash.Reset()

	// 计算签名
	sign, err := common.Sign([]byte(validStr), c.privateKey, c.hashType, c.hash)
	if err != nil {
		return "", err
	}

	c.log.Debugf("sign: %s", sign)

	// 拼接签名
	return fmt.Sprintf(common.AuthNotifyFormat, common.Algorism_SHA256, ts, nonceStr, sign), nil
}

// VerifyResponseSign 验证响应签名（直接送上来的header和body）
func (c *Client) VerifyResponseSign(header http.Header, body []byte) error {
	appid := header.Get(common.Lklapi_Appid)
	serialNo := header.Get(common.Lklapi_Serial)
	ts := header.Get(common.Lklapi_Timestamp)
	nonce := header.Get(common.Lklapi_Nonce)
	sign := header.Get(common.Lklapi_Sign)
	return c.VerifySign(appid, serialNo, ts, nonce, string(body), sign)
}

// VerifySignForAuth 验证请求auth签名
func (c *Client) VerifySignForAuth(auth string, body []byte) error {
	var (
		ts, nonce, sign string
	)

	split := strings.Split(auth, " ")
	if len(split) != 2 {
		return fmt.Errorf("auth format error")
	}

	params := strings.Split(split[1], ",")
	for _, param := range params {
		kv := strings.SplitN(param, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("auth format error")
		}
		switch kv[0] {
		case common.AuthTimestamp:
			ts = strings.Trim(kv[1], `"`)
		case common.AuthNonceStr:
			nonce = strings.Trim(kv[1], `"`)
		case common.AuthSignature:
			sign = strings.Trim(kv[1], `"`)
		}
	}

	return c.VerifyNotifySign(ts, nonce, string(body), sign)
}

// VerifySignForReq 验证请求request签名
func (c *Client) VerifySignForReq(req *http.Request, body []byte) error {
	auth := req.Header.Get(common.Authorization)
	if auth == "" {
		return fmt.Errorf("request header Authorization is empty")
	}
	return c.VerifySignForAuth(auth, body)
}

func (c *Client) VerifySign(appid, serialNo, ts, nonce, body, sign string) error {
	signature, _ := base64.StdEncoding.DecodeString(sign)

	// 计算签名
	// 将获取到的appId、证书序列号、时间戳、随机字符串、报文body拼接。拼接报文一共有5行，每一行为一个参数。行尾以\n（换行符，ASCII编码值为0x0A）结束，包括最后一行：
	// ${Lklapi-appid}\n+${Lklapi-serialNo}\n+${Lklapi-timeStamp}\n+${Lklapi-nonceStr}\n+${body}\n

	// 计算签名
	validStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", appid, serialNo, ts, nonce, body)

	err := c.lklCertificate.CheckSignature(x509.SHA256WithRSA, []byte(validStr), signature)
	if err != nil {
		c.log.Errorf("verify sign error: %v \n appid: %s, serialNo: %s, ts: %s, nonce: %s, sign: %s, body: %s",
			err, appid, serialNo, ts, nonce, sign, body)
		return err
	}
	return nil
}

func (c *Client) VerifyNotifySign(ts, nonce, body, sign string) error {
	signature, _ := base64.StdEncoding.DecodeString(sign)

	// 计算签名
	// 将获取到的appId、证书序列号、时间戳、随机字符串、报文body拼接。拼接报文一共有5行，每一行为一个参数。行尾以\n（换行符，ASCII编码值为0x0A）结束，包括最后一行：
	// ${Lklapi-appid}\n+${Lklapi-serialNo}\n+${Lklapi-timeStamp}\n+${Lklapi-nonceStr}\n+${body}\n

	// 计算签名
	validStr := fmt.Sprintf("%s\n%s\n%s\n", ts, nonce, body)

	err := c.lklCertificate.CheckSignature(x509.SHA256WithRSA, []byte(validStr), signature)
	if err != nil {
		c.log.Errorf("verify sign error: %v \n ts: %s, nonce: %s, sign: %s, body: %s",
			err, ts, nonce, sign, body)
		return err
	}
	return nil
}

// ReSet 重新设置配置
func (c *Client) ReSet(cfg Config, options ...Option) {
	c.cfg = cfg

	for _, option := range options {
		option(c)
	}

	c.initHttpClient()
	c.hash = c.hashType.New()

	// 解析公钥证书
	lklCertificate, err := c.getLklCertificate()
	if err != nil {
		panic(err)
	}
	c.lklCertificate = lklCertificate

	// 解析私钥
	privateKey, err := c.getPrivateKey()
	if err != nil {
		panic(err)
	}
	c.privateKey = privateKey
}

func (c *Client) getLklCertificate() (*x509.Certificate, error) {
	content := ""
	if c.cfg.SyncPublicKey != "" {
		content = c.cfg.SyncPublicKey
	} else if c.cfg.SyncPubicPath != "" {
		// 读取证书文件
		publicBytes, err := os.ReadFile(c.cfg.SyncPubicPath)
		if err != nil {
			return nil, err
		}
		content = string(publicBytes)
	} else {
		return nil, fmt.Errorf("未设置拉卡拉公钥证书")
	}

	return common.ParseCertificate([]byte(content))
}

func (c *Client) getPrivateKey() (*rsa.PrivateKey, error) {
	privateKey := ""
	if c.cfg.SignPrivateKey != "" {
		privateKey = c.cfg.SignPrivateKey
	} else if c.cfg.SignPrivatePath == "" {
		// 读取证书文件
		privateBytes, err := os.ReadFile(c.cfg.SignPrivatePath)
		if err != nil {
			return nil, err
		}
		privateKey = string(privateBytes)
	} else {
		return nil, fmt.Errorf("未设置签名私钥")
	}

	return common.ParsePrivateKey([]byte(privateKey))
}
