package common

const (
	// Authorization auth
	Authorization = "Authorization"

	Algorism_SHA256 = "LKLAPI-SHA256withRSA"

	// AuthFormat 签名格式(${Algorism}+空格+appid=“${appid}“,serial_no=“${serialNo}“,timestamp=“${timeStamp}“,nonce_str=“${nonceStr}“,signature=“${signature}“)
	AuthFormat    = `%s appid="%s",serial_no="%s",timestamp="%d",nonce_str="%s",signature="%s"`
	AuthAppid     = "appid"
	AuthSerialNo  = "serial_no"
	AuthTimestamp = "timestamp"
	AuthNonceStr  = "nonce_str"
	AuthSignature = "signature"

	// ReplyOrderNotifySuccess 收银台回调通知成功
	ReplyOrderNotifySuccess = `{"code":"SUCCESS","message":"执行成功"}`

	Lklapi_Appid     = "Lklapi-Appid"     // 原始送上来的Lklapi-Appid，异步交易结果通知无此字段
	Lklapi_Serial    = "Lklapi-Serial"    // Lklapi-Serial：拉卡拉证书序列号
	Lklapi_Timestamp = "Lklapi-Timestamp" // Lklapi-Timestamp：时间戳
	Lklapi_Nonce     = "Lklapi-Nonce"     // Lklapi-Nonce：随机字符串
	Lklapi_Sign      = "Lklapi-Signature" // Lklapi-Signature：签名值
	Lklapi_Trace     = "Lklapi-Trace"     // Lklapi-Traceid：供研发查询定位问题使用，建议打印出来 异步交易结果通知无此字段

)

// PayMode 收银台的支付方式
type PayMode string

const (
	PayMode_ALIPAY    PayMode = "ALIPAY"    // 支付宝
	PayMode_WECHAT    PayMode = "WECHAT"    // 微信
	PayMode_UNION     PayMode = "UNION"     // 银联云闪付
	PayMode_CARD      PayMode = "CARD"      // POS刷卡交易
	PayMode_LKLAT     PayMode = "LKLAT"     // 线上转账
	PayMode_QUICK_PAY PayMode = "QUICK_PAY" // 快捷支付
	PayMode_EBANK     PayMode = "EBANK"     // 网银支付
	PayMode_UNION_CC  PayMode = "UNION_CC"  // 银联支付
	PayMode_BESTPAY   PayMode = "BESTPAY"   // 翼支付

	PayMode_UNION_FQ PayMode = "UNION_FQ" // 银联聚分期
)

// CounterParam 收银台参数格式化 + PayMode
const CounterParam = `{"pay_mode":"%s"}`

// 业务类型控制参数

type BusinessType string

const (
	BusinessType_UPCARD     BusinessType = "UPCARD"     // 刷卡
	BusinessType_SCPAY      BusinessType = "SCPAY"      // 扫码
	BusinessType_CRDFLG_D   BusinessType = "CRDFLG_D"   // 借记卡
	BusinessType_CRDFLG_C   BusinessType = "CRDFLG_C"   // 贷记卡
	BusinessType_CRDFLG_OTH BusinessType = "CRDFLG_OTH" // 不明确是借记卡还是贷记卡
)

// ProductType 交易商品类型
type ProductType string

const (
	ProductType_1  ProductType = "1"  // 服饰箱包
	ProductType_2  ProductType = "2"  // 食品药品
	ProductType_3  ProductType = "3"  // 化妆品
	ProductType_4  ProductType = "4"  // 电子产品
	ProductType_5  ProductType = "5"  // 日用家居
	ProductType_7  ProductType = "7"  // 航空机票
	ProductType_8  ProductType = "8"  // 酒店住宿
	ProductType_9  ProductType = "9"  // 留学教育
	ProductType_10 ProductType = "10" // 旅游票务
	ProductType_11 ProductType = "11" // 国际物流
	ProductType_12 ProductType = "12" // 国际租车
	ProductType_13 ProductType = "13" // 国际会议
	ProductType_14 ProductType = "14" // 软件服务
	ProductType_15 ProductType = "15" // 医疗服务
	ProductType_16 ProductType = "16" // 通讯
	ProductType_17 ProductType = "17" // 休闲娱乐
)

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderStatus_0 OrderStatus = "0" // 待支付
	OrderStatus_1 OrderStatus = "1" // 支付中
	OrderStatus_2 OrderStatus = "2" // 支付成功
	OrderStatus_3 OrderStatus = "3" // 支付失败
	OrderStatus_4 OrderStatus = "4" // 已过期
	OrderStatus_5 OrderStatus = "5" // 已取消
	OrderStatus_6 OrderStatus = "6" // 部分退款或者全部退款
	OrderStatus_7 OrderStatus = "7" // 订单已关闭
)
