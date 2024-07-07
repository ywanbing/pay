package common

const (
	Algorism_SHA256 = "LKLAPI-SHA256withRSA"

	// AuthFormat 签名格式(${Algorism}+空格+appid=“${appid}“,serial_no=“${serialNo}“,timestamp=“${timeStamp}“,nonce_str=“${nonceStr}“,signature=“${signature}“)
	AuthFormat = `%s appid="%s",serial_no="%s",timestamp="%d",nonce_str="%s",signature="%s"`

	// ReplyOrderNotifySuccess 收银台回调通知成功
	ReplyOrderNotifySuccess = `{"code":"SUCCESS","message":"执行成功"}`
)

// 收银台的支付方式
const (
	ALIPAY    = "ALIPAY"    // 支付宝
	WECHAT    = "WECHAT"    // 微信
	UNION     = "UNION"     // 银联云闪付
	CARD      = "CARD"      // POS刷卡交易
	LKLAT     = "LKLAT"     // 线上转账
	QUICK_PAY = "QUICK_PAY" // 快捷支付
	EBANK     = "EBANK"     // 网银支付
	UNION_CC  = "UNION_CC"  // 银联支付
	BESTPAY   = "BESTPAY"   // 翼支付
	HB_FQ     = "HB_FQ"     // 花呗分期
	UNION_FQ  = "UNION_FQ"  // 银联聚分期
)

// CounterParam 收银台参数格式化
const CounterParam = `{"pay_mode":"%s"}`

// 业务类型控制参数
const (
	UPCARD     = "UPCARD"     // 刷卡
	SCPAY      = "SCPAY"      // 扫码
	CRDFLG_D   = "CRDFLG_D"   // 借记卡
	CRDFLG_C   = "CRDFLG_C"   // 贷记卡
	CRDFLG_OTH = "CRDFLG_OTH" // 不明确是借记卡还是贷记卡
)

// 交易商品类型
const (
	ProductType_1  = "1"  // 服饰箱包
	ProductType_2  = "2"  // 食品药品
	ProductType_3  = "3"  // 化妆品
	ProductType_4  = "4"  // 电子产品
	ProductType_5  = "5"  // 日用家居
	ProductType_7  = "7"  // 航空机票
	ProductType_8  = "8"  // 酒店住宿
	ProductType_9  = "9"  // 留学教育
	ProductType_10 = "10" // 旅游票务
	ProductType_11 = "11" // 国际物流
	ProductType_12 = "12" // 国际租车
	ProductType_13 = "13" // 国际会议
	ProductType_14 = "14" // 软件服务
	ProductType_15 = "15" // 医疗服务
	ProductType_16 = "16" // 通讯
	ProductType_17 = "17" // 休闲娱乐
)
