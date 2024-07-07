package model

// RefundReq 退款
//
// 输入参数要么传out_order_no+merchant_no
// 要么传pay_order_no+channel_id
type RefundReq struct {
	MerchantNo       string        `json:"merchant_no" validate:"required,max=32"`                    // 拉卡拉分配的银联商户号
	TermNo           string        `json:"term_no" validate:"required,max=16"`                        // 拉卡拉分配的业务终端号
	OutTradeNo       string        `json:"out_trade_no" validate:"required,max=32"`                   // 商户订单号(商户系统唯一
	RefundAmount     string        `json:"refund_amount" validate:"required,gte=0"`                   // 退款金额
	RefundReason     string        `json:"refund_reason,omitempty" validate:"omitempty,max=32"`       // 退款原因
	OriginOutTradeNo string        `json:"origin_out_trade_no,omitempty" validate:"omitempty,max=32"` // 原商户交易流水号
	OriginTradeNo    string        `json:"origin_trade_no,omitempty" validate:"omitempty,max=32"`     // 原拉卡拉交易流水号
	OriginLogNo      string        `json:"origin_log_no,omitempty" validate:"omitempty,max=14"`       // 原始退款流水号
	LocationInfo     *LocationInfo `json:"location_info,omitempty" validate:"omitempty"`              // 地址位置信息
}

type LocationInfo struct {
	RequestIp   string `json:"request_ip" validate:"required,max=64"`              // 请求IP
	BaseStation string `json:"base_station,omitempty" validate:"omitempty,max=64"` // 基站信息
	Location    string `json:"location,omitempty" validate:"omitempty,max=64"`     // 位置信息
}

// 根据上面的注释生成

// RefundRes 退款
type RefundRes struct {
	MerchantNo       string `json:"merchant_no" validate:"required,max=32"`          // 拉卡拉分配的银联商户号（请求接口中商户号）
	OutOrderNo       string `json:"out_order_no" validate:"required,max=32"`         // 请求中的商户请求流水号
	TradeNo          string `json:"trade_no" validate:"required,max=32"`             // 拉卡拉交易流水号
	LogNo            string `json:"log_no" validate:"required,max=14"`               // 拉卡拉对账单流水号
	AccTradeNo       string `json:"acc_trade_no" validate:"required,max=32"`         // 账户端交易流水号
	AccountType      string `json:"account_type" validate:"required,max=32"`         // 钱包类型 WECHAT 支付宝:ALIPAY 银联:UQRCODEPAY 翼支付: BESTPAY 苏宁易付宝: SUNING
	TotalAmount      string `json:"total_amount" validate:"required"`                // 交易金额 单位分，整数数字型字符串
	RefundAmount     string `json:"refund_amount" validate:"required"`               // 申请退款金额 单位分，整数数字型字符串
	PayerAmount      string `json:"payer_amount" validate:"required"`                // 商户实际退还金额 单位分，整数数字型字符串
	TradeTime        string `json:"trade_time" validate:"required"`                  // 退款时间  实际退款时间。yyyyMMddHHmmss
	OriginTradeNo    string `json:"origin_trade_no" validate:"omitempty,max=32"`     // 如果请求中携带，则返回
	OriginOutTradeNo string `json:"origin_out_trade_no" validate:"omitempty,max=64"` // 如果请求中携带，则返回
	UpIssAddnData    string `json:"up_iss_addn_data" validate:"omitempty,max=8000"`  // 单品营销 附加数据
	UpCouponInfo     string `json:"up_coupon_info" validate:"omitempty,max=512"`     // 银联优惠信息、出资方信息
	TradeInfo        string `json:"trade_info" validate:"omitempty,max=512"`         // 出资方信息
}
