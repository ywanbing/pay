package lklpay

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/imroc/req/v3"
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
	Appid           string `json:"appid,omitempty" `            // 接入方唯一编号
	SerialNo        string `json:"serial_no,omitempty"`         // 证书序列号
	MerchantNo      string `json:"merchant_no,omitempty"`       // 商户号
	TermNo          string `json:"term_no,omitempty"`           // 终端号(收银台可不送
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

// Client pay client
type Client struct {
	ctx context.Context // 上下文
	cfg Config          // 配置
	log log.Logger      // logger

	isProd bool                // 是否生产环境
	valid  *validator.Validate // 参数校验
	cli    *req.Client
}

// New a pay client
func New(cfg Config, options ...Option) *Client {
	client := &Client{
		ctx:    context.Background(),
		cfg:    cfg,
		log:    log.DefLogger(),
		isProd: false,
		valid:  validator.New(),
	}

	for _, option := range options {
		option(client)
	}

	// 初始化http客户端
	client.initHttpClient()

	return client
}

func (c *Client) initHttpClient() {
	cli := req.NewClient()
	if c.isProd {
		cli.SetBaseURL(ProdUrl)
	} else {
		cli.SetBaseURL(TestUrl)
	}

	cli.SetLogger(c.log)
	cli.SetCommonContentType("application/json")
	c.cli = cli
}
