package model

import "github.com/ywanbing/pay/lklpay/common"

// https://o.lakala.com/#/home/document/detail?id=113

// ConvergeActiveReq 聚合主扫请求
//
//	现在只支持了 微信 和 支付宝
type ConvergeActiveReq[T AliPayAccBusiFields | WechatAccBusiFields] struct {
	MerchantNo    string             `json:"merchant_no" validate:"required,max=32"`             // 商户号(拉卡拉分配的商户号)
	TermNo        string             `json:"term_no" validate:"required,max=32"`                 // 终端号(拉卡拉分配的业务终端号)
	OutTradeNo    string             `json:"out_trade_no" validate:"required,max=32"`            // 商户交易流水号(商户系统唯一，对应数据库表中外部请求流水号)
	AccountType   common.AccountType `json:"account_type" validate:"required,max=32"`            // 钱包类型(微信：WECHAT 支付宝：ALIPAY 银联：UQRCODEPAY 翼支付: BESTPAY 苏宁易付宝: SUNING 拉卡拉支付账户：LKLACC 网联小钱包：NUCSPAY)
	TransType     common.TransType   `json:"trans_type" validate:"required,max=2"`               // 接入方式(41:NATIVE（（ALIPAY，云闪付支持）51:JSAPI（微信公众号支付，支付宝服务窗支付，银联JS支付，翼支付JS支付、拉卡拉钱包支付）71:微信小程序支付,61:APP支付（微信APP支付）
	TotalAmount   string             `json:"total_amount" validate:"required,max=12"`            // 金额(单位分，整数型字符)
	LocationInfo  *LocationInfo      `json:"location_info" validate:"required"`                  // 地址位置信息(地址位置信息，风控要求必送)
	BusiMode      string             `json:"busi_mode,omitempty" validate:"omitempty,max=8"`     // 业务模式(ACQ-收单 不填，默认为“ACQ-收单”
	Subject       string             `json:"subject,omitempty" validate:"omitempty,max=42"`      // 订单标题(标题，用于简单描述订单或商品主题，会传递给账户端 (账户端控制，实际最多42个字符)，微信支付必送。
	PayOrderNo    string             `json:"pay_order_no,omitempty" validate:"omitempty,max=64"` // 支付业务订单号(拉卡拉订单系统订单号，以拉卡拉支付业务订单号为驱动的支付行为，需上传该字段。
	NotifyUrl     string             `json:"notify_url,omitempty" validate:"omitempty,max=128"`  // 商户通知地址(商户通知地址，如果上传，且 pay_order_no 不存在情况下，则按此地址通知商户(详见“[交易通知]”接口)
	SettleType    string             `json:"settle_type,omitempty" validate:"omitempty,max=4"`   // 结算类型(0：普通结算 1：分账结算)
	Remark        string             `json:"remark,omitempty" validate:"omitempty,max=128"`      // 备注
	AccBusiFields *T                 `json:"acc_busi_fields,omitempty" validate:"omitempty"`     // 账户端业务信息域(不同的account_type和trans_type，需要传入的参数不一样
}

// AliPayAccBusiFields 支付宝主扫场景下acc_busi_fields域内容
type AliPayAccBusiFields struct {
	UserID             string        `json:"user_id,omitempty" validate:"omitempty,max=64"`              // 买家在支付宝的用户id
	TimeoutExpress     string        `json:"timeout_express,omitempty" validate:"omitempty,max=2"`       // 预下单有效时间（预下单的订单的有效时间，以分钟为单位。如果在有效时间内没有完成付款，则在账户端该订单失效。如果不上送，以账户端订单失效时间为准。 建议不超过15分钟。不传值则默认5分钟。
	ExtendParams       *ExtendParams `json:"extend_params,omitempty" validate:"omitempty"`               // 业务扩展参数
	GoodsDetail        string        `json:"goods_detail,omitempty" validate:"omitempty,max=6000"`       // 商品详情(订单包含的商品列表信息，Json GoodsDetail 数组。)
	StoreID            string        `json:"store_id,omitempty" validate:"omitempty,max=32"`             // 商户门店编号
	AlipayStoreID      string        `json:"alipay_store_id,omitempty" validate:"omitempty,max=32"`      // (不再使用)支付宝店铺编号
	DisablePayChannels string        `json:"disable_pay_channels,omitempty" validate:"omitempty,max=32"` // 支付宝禁用支付渠道( “credit_group”表示禁用信用支付类（包含信用卡,花呗，花呗分期）“pcredit”表示禁用花呗 “pcreditpayInstallment”表示禁用花呗分期 “creditCard“表示禁用信用卡 如果想禁用多个可在枚举间加,隔开
	BusinessParams     string        `json:"business_params,omitempty" validate:"omitempty,max=512"`     // 商户传入业务信息 (商户传入业务信息，应用于安全，营销等参数直传场景，格式为 json 格式。
	MinAge             string        `json:"min_age,omitempty" validate:"omitempty,max=2"`               // 允许的最小买家年龄 （买家年龄必须大于等于所传数值
}

// ExtendParams 支付宝扩展参数
type ExtendParams struct {
	SysServiceProviderID string `json:"sys_service_provider_id,omitempty" validate:"omitempty,max=64"` // 服务商的PID
	HbFqNum              string `json:"hb_fq_num" validate:",max=5"`                                   // 花呗分期期数(支付宝花呗分期必送字段: 花呗分期数 3：3期 6：6期 12：12期
	HbFqSellerPercent    string `json:"hb_fq_seller_percent" validate:"max=3"`                         // 卖家承担手续费比例(支付宝花呗分期必送字段: 卖家承担手续费比例，间连模式下只支持传0
	FoodOrderType        string `json:"food_order_type,omitempty" validate:"omitempty,max=32"`         // 点餐场景类型( qr_order（店内扫码点餐），pre_order（预点到店自提），home_delivery （外送到家），direct_payment（直接付款），other（其它）
}

type GoodsDetail struct {
	GoodsId        string  `json:"goods_id" validate:"max=32"`                             // 商品的编号
	AlipayGoodsId  string  `json:"alipay_goods_id,omitempty" validate:"omitempty,max=32"`  // 支付宝定义的统一商品编号
	GoodsName      string  `json:"goods_name" validate:"max=256"`                          // 商品名称
	Quantity       int64   `json:"quantity" validate:"gt=0"`                               // 商品数量
	Price          float64 `json:"price" validate:"gt=0"`                                  // 商品价格 单位元
	GoodsCategory  string  `json:"goods_category,omitempty" validate:"omitempty,max=24"`   // 商品类目
	CategoriesTree string  `json:"categories_tree,omitempty" validate:"omitempty,max=128"` // 商品类目树，从商品类目根节点到叶子节点的类目 id 组成，类目 id 值使用|分割
	Body           string  `json:"body,omitempty" validate:"omitempty,max=512"`            // 商品描述信息
	ShowUrl        string  `json:"show_url,omitempty" validate:"omitempty,max=256"`        // 商品的展示地址
}

// WechatAccBusiFields  微信主扫场景下acc_busi_fields域内容
type WechatAccBusiFields struct {
	TimeoutExpress string `json:"timeout_express,omitempty" validate:"omitempty,max=2"` // 预下单有效时间(下单的订单的有效时间，以分钟为单位。建议不超过15分钟。不传值则默认5分钟)
	SubAppid       string `json:"sub_appid,omitempty" validate:"omitempty,max=32"`      // 子商户公众账号ID(微信分配的子商户公众账号ID，sub_appid（即微信小程序支付-71、公众号支付-51、微信app支付-61）此参数必传，只对微信支付有效)
	UserId         string `json:"user_id,omitempty" validate:"omitempty,max=64"`        // 用户标识(用户在子商户sub_appid下的唯一标识，sub_openid，（即微信小程序支付-71、众号支付-51），此参数必传，只对微信支付有效
	Detail         string `json:"detail,omitempty" validate:"omitempty,max=1024"`       // 商品详情(单品优惠功能字段，详见下文说明)
	GoodsTag       string `json:"goods_tag,omitempty" validate:"omitempty,max=32"`      // 订单优惠标记(订单优惠标记，微信平台配置的商品标记，用于优惠券或者满减使用，accountType为WECHAT时，可选填此字段
	Attach         string `json:"attach,omitempty" validate:"omitempty,max=128"`        // 附加域(该字段主要用于商户携带订单的自定义数据。商户定制字段，直接送到账户端。
	DeviceInfo     string `json:"device_info,omitempty" validate:"omitempty,max=32"`    // 设备号(终端设备号(门店号或收银设备ID)，注意：PC网页或JSAPI支付请传”WEB”
	LimitPay       string `json:"limit_pay,omitempty" validate:"omitempty,max=8"`       // 指定支付方式(no_credit-指定不能使用信用卡支付
	SceneInfo      string `json:"scene_info,omitempty" validate:"omitempty,max=256"`    // 场景信息(该字段用于上报场景信息，目前支持上报实际门店信息。
	LimitPayer     string `json:"limit_payer,omitempty" validate:"omitempty,max=8"`     // 限定支付(ADULT-成年人
}

type WxGoodsDetail struct {
	GoodsId      string `json:"goods_id" validate:"max=32"`                           // 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。如“商品编码”
	WxpayGoodsId string `json:"wxpay_goods_id,omitempty" validate:"omitempty,max=32"` // 微信支付定义的统一商品编号
	GoodsName    string `json:"goods_name" validate:"max=256"`                        // 商品的实际名称
	Quantity     int64  `json:"quantity" validate:"gt=0"`                             // 用户购买的数量
	Price        int64  `json:"price" validate:"gt=0"`                                // 单位为：分。如果商户有优惠，需传输商户优惠后的单价
}

// ConvergeActiveRes 拉卡拉账户下单结果
type ConvergeActiveRes[T AliPayNativeResp | AliPayJsApiResp | WxResp | WxAppResp] struct {
	MerchantNo     string `json:"merchant_no" validate:"required,max=32"`                  // 商户号
	OutTradeNo     string `json:"out_trade_no" validate:"required,max=32"`                 // 商户请求流水号
	TradeNo        string `json:"trade_no" validate:"required,max=32"`                     // 拉卡拉交易流水号
	LogNo          string `json:"log_no" validate:"required,max=14"`                       // 拉卡拉对账单流水号
	SettleMerchant string `json:"settle_merchant" validate:"required,max=32"`              // 结算商户号
	SettleTerm     string `json:"settle_term" validate:"required,max=32"`                  // 结算终端号
	AccRespFields  *T     `json:"acc_resp_fields,omitempty" validate:"omitempty,required"` // 账户端返回信息域
}

type AliPayNativeResp struct {
	Code      string `json:"code" validate:"max=256"`       // 二维码信息
	CodeImage string `json:"code_image" validate:"max=256"` // 商户收款二维码图片。Base64编码，暂无
}

type AliPayJsApiResp struct {
	PrepayId string `json:"prepay_id" validate:"max=32"` // 预支付交易会话ID
}

// WxResp 微信(71-小程序)微信(51-JSAPI)场景下返回acc_resp_fields域
type WxResp struct {
	PrepayId  string `json:"prepay_id" validate:"max=32"`  // 预支付交易会话ID
	PaySign   string `json:"pay_sign" validate:"max=256"`  // 签名
	AppId     string `json:"app_id" validate:"max=32"`     // 商户注册具有支付权限的小程序成功后即可获得小程序id
	TimeStamp string `json:"time_stamp" validate:"max=32"` // 当前的时间
	NonceStr  string `json:"nonce_str" validate:"max=32"`  // 随机字符串
	Package   string `json:"package" validate:"max=128"`   // 订单详情扩展字符串
	SignType  string `json:"sign_type" validate:"max=32"`  // 签名类型，支持RSA
}

// WxAppResp 微信(61-APP)场景下返回acc_resp_fields域
type WxAppResp struct {
	PrepayId  string `json:"prepay_id" validate:"max=32"`  // 预支付交易会话ID
	PaySign   string `json:"pay_sign" validate:"max=256"`  // 签名
	AppId     string `json:"app_id" validate:"max=32"`     // 商户注册具有支付权限的安卓/IOSAPP成功后即可获得移动应用appid
	TimeStamp string `json:"time_stamp" validate:"max=32"` // 当前的时间
	NonceStr  string `json:"nonce_str" validate:"max=32"`  // 随机字符串
	Package   string `json:"package" validate:"max=128"`   // 订单详情扩展字符串
	SignType  string `json:"sign_type" validate:"max=32"`  // 签名类型，支持RSA
	PartnerId string `json:"partner_id" validate:"max=32"` // 从业机构号
}

// ConvergeActiveQueryReq 拉卡拉主扫结果查询
type ConvergeActiveQueryReq struct {
	MerchantNo string `json:"merchant_no" validate:"required,max=32"`             // 商户号 拉卡拉分配的商户号
	TermNo     string `json:"term_no" validate:"required,max=32"`                 // 终端号 拉卡拉分配的终端号
	OutTradeNo string `json:"out_trade_no,omitempty" validate:"omitempty,max=32"` // 商户交易流水号 下单时的商户请求流水号 说明：out_trade_no、trade_no、必有其一。如果存在多个字段上送，优先级顺序如下： trade_no、 out_trade_no
	TradeNo    string `json:"trade_no,omitempty" validate:"omitempty,max=32"`     // 拉卡拉交易流水号
}

// ConvergeActiveQueryRes 拉卡拉主扫结果查询响应
type ConvergeActiveQueryRes struct {
	MerchantNo             string             `json:"merchant_no" validate:"required,max=32"`                    // 商户号 (拉卡拉分配的商户号（请求接口中商户号）)
	OutTradeNo             string             `json:"out_trade_no" validate:"required,max=32"`                   // 商户请求流水号(请求接口中out_trade_no)
	TradeNo                string             `json:"trade_no" validate:"required,max=32"`                       // 拉卡拉生成的交易流水
	LogNo                  string             `json:"log_no" validate:"required,max=14"`                         // 拉卡拉对账单流水号
	TradeMainType          string             `json:"trade_main_type,omitempty" validate:"omitempty,max=32"`     // 交易大类 (PREORDER-主扫，MICROPAY-被扫，REFUND-退款，CANCEL-撤销，无-其它类型)
	SplitAttr              string             `json:"split_attr,omitempty" validate:"omitempty,max=2"`           // 拆单属性 (只有涉及合单交易时会出现：M-主单，S-子单)
	SplitInfo              []*OrderSplitInfo  `json:"split_info,omitempty" validate:"omitempty,max=32"`          // 拆单信息
	AccTradeNo             string             `json:"acc_trade_no" validate:"required,max=32"`                   // 账户端交易订单号
	AccountType            string             `json:"account_type" validate:"required,max=32"`                   // 钱包类型 (微信：WECHAT 支付宝：ALIPAY 银联：UQRCODEPAY 翼支付: BESTPAY 苏宁易付宝: SUNING)
	TradeState             string             `json:"trade_state" validate:"required,max=16"`                    // 交易状态(INIT-初始化 CREATE-下单成功 SUCCESS-交易成功 FAIL-交易失败 DEAL-交易处理中 UNKNOWN-未知状态 CLOSE-订单关闭 PART_REFUND-部分退款 REFUND-全部退款(或订单被撤销))
	TradeStateDesc         string             `json:"trade_state_desc,omitempty" validate:"omitempty,max=256"`   // 交易状态描述
	TotalAmount            string             `json:"total_amount" validate:"required"`                          // 订单金额 (单位分，整数数字型字符)
	PayerAmount            string             `json:"payer_amount,omitempty" validate:"omitempty"`               // 付款人实付金额(单位分)
	AccSettleAmount        string             `json:"acc_settle_amount,omitempty" validate:"omitempty"`          // 账户端结算金额
	AccMDiscountAmount     string             `json:"acc_mdiscount_amount,omitempty" validate:"omitempty"`       // 商户侧优惠金额（账户端）
	AccDiscountAmount      string             `json:"acc_discount_amount,omitempty" validate:"omitempty"`        // 账户端优惠金额
	AccOtherDiscountAmount string             `json:"acc_other_discount_amount,omitempty" validate:"omitempty"`  // 账户端其它优惠金额
	TradeTime              string             `json:"trade_time,omitempty" validate:"omitempty"`                 // 交易完成时间
	UserId1                string             `json:"user_id1,omitempty" validate:"omitempty,max=128"`           // 用户标识1
	UserId2                string             `json:"user_id2,omitempty" validate:"omitempty,max=128"`           // 用户标识2
	BankType               string             `json:"bank_type,omitempty" validate:"omitempty,max=128"`          // 付款银行
	CardType               string             `json:"card_type,omitempty" validate:"omitempty,max=16"`           // 银行卡类型 (00：借记 01：贷记 02：微信零钱 03：支付宝花呗 04：支付宝其他 05：数字货币 06：拉卡拉支付账户 99：未知)
	AccActivityId          string             `json:"acc_activity_id,omitempty" validate:"omitempty,max=32"`     // 活动 ID(在账户端商户后台配置的批次 ID)
	AccRespFields          map[string]any     `json:"acc_resp_fields,omitempty" validate:"omitempty,max=1024"`   // 账户端返回信息域
	RefundSplitInfo        []*RefundSplitInfo `json:"refund_split_info,omitempty" validate:"omitempty,max=1024"` // 合单退款拆单信息
}

type RefundSplitInfo struct {
	OutSubTradeNo string `json:"out_sub_trade_no" validate:"required,max=32"`       // 外部子退款交易流水号(商户子交易流水号,商户号下唯一)
	MerchantNo    string `json:"merchant_no" validate:"required,max=32"`            // 商户号(拉卡拉分配的商户号)
	TermNo        string `json:"term_no" validate:"required,max=32"`                // 终端号(拉卡拉分配的业务终端号)
	RefundAmount  string `json:"refund_amount" validate:"required"`                 // 申请退款金额(单位分，整数型字符)
	SubTradeNo    string `json:"sub_trade_no" validate:"required,max=32"`           // 拉卡拉子交易流水号(商户子交易流水号，商户号下唯一)
	SubLogNo      string `json:"sub_log_no,omitempty" validate:"omitempty,max=14"`  // 对账单子流水号
	TradeState    string `json:"trade_state,omitempty" validate:"omitempty,max=16"` // 子交易状态
	ResultCode    string `json:"result_code,omitempty" validate:"omitempty,max=32"` // 处理结果码
	ResultMsg     string `json:"result_msg,omitempty" validate:"omitempty,max=128"` // 处理描述
}

// ConvergeActiveCloseReq 关闭订单
//
// 输入参数要么传out_order_no，要么传pay_order_no
type ConvergeActiveCloseReq struct {
	MerchantNo       string        `json:"merchant_no" validate:"required,max=32"`                    // 商户号
	TermNo           string        `json:"term_no" validate:"required,max=32"`                        // 终端号(拉卡拉分配的业务终端号)
	OriginOutTradeNo string        `json:"origin_out_trade_no,omitempty" validate:"omitempty,max=32"` // 原商户交易流水号(下单时商户请求流水号，origin_out_trade_no、origin_trade_no、origin_out_order_source+origin_out_order_no必有其一，前两者为交易关单，后者为商户订单关单)
	OriginTradeNo    string        `json:"origin_trade_no,omitempty" validate:"omitempty,max=32"`     // 原交易拉卡拉交易流水号(下单成功时，返回的SAAS生成的交易流水)
	LocationInfo     *LocationInfo `json:"location_info" validate:"required"`                         // 地址位置信息(地址位置信息，风控要求必送)
}

// ConvergeActiveCloseRes 关闭订单
type ConvergeActiveCloseRes struct {
	OriginTradeNo        string `json:"origin_trade_no" validate:"required,max=32"`                    // 原交易拉卡拉交易流水号
	OriginOutTradeNo     string `json:"origin_out_trade_no" validate:"required,max=32"`                // 原商户请求流水号
	OriginOutOrderSource string `json:"origin_out_order_source,omitempty" validate:"omitempty,max=16"` // 原订单外部订单来源
	OriginOutOrderNo     string `json:"origin_out_order_no,omitempty" validate:"omitempty,max=32"`     // 原订单外部商户订单号
	TradeTime            string `json:"trade_time" validate:"required,max=14"`                         // 交易时间
}
