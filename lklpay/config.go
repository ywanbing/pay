package lklpay

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
