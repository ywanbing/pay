package model

type OrderNotify struct {
	PayOrderNo         string            `json:"pay_order_no" validate:"required,max=64"`       // 支付订单号
	OutOrderNo         string            `json:"out_order_no" validate:"required,max=32"`       // 商户订单号
	ChannelID          string            `json:"channel_id" validate:"required,max=32"`         // 渠道号
	TransMerchantNo    string            `json:"trans_merchant_no" validate:"omitempty,max=32"` // 交易商户号
	TransTermNo        string            `json:"trans_term_no" validate:"omitempty,max=16"`     // 交易终端号
	MerchantNo         string            `json:"merchant_no" validate:"required,max=32"`        // 结算商户号
	TermNo             string            `json:"term_no" validate:"required,max=16"`            // 结算终端号
	OrderStatus        string            `json:"order_status" validate:"required,max=2"`        // 订单状态
	OrderInfo          string            `json:"order_info" validate:"omitempty,max=100"`       // 订单描述
	TotalAmount        int64             `json:"total_amount" validate:"required"`              // 订单金额，单位：分
	OrderCreateTime    string            `json:"order_create_time" validate:"required"`         // 订单创建时间(格式yyyyMMddHHmmss
	OrderEfficientTime string            `json:"order_efficient_time" validate:"required"`      // 订单有效时间(格式yyyyMMddHHmmss
	SplitMark          string            `json:"split_mark" validate:"omitempty,max=2"`         // 合单标识
	SplitInfo          []*OrderSplitInfo `json:"split_info" validate:"omitempty,dive"`          // 交易拆单信息(
	OrderTradeInfo     *OrderTradeInfo   `json:"order_trade_info" validate:"required"`          // 订单交易信息
}

// OrderTradeInfo represents the structure of trade information within an order.
type OrderTradeInfo struct {
	TradeNo                string `json:"trade_no" validate:"required,max=32"`            // 交易流水号
	LogNo                  string `json:"log_no" validate:"required,max=14"`              // 对账单流水号
	TradeRefNo             string `json:"trade_ref_no" validate:"omitempty,max=12"`       // 交易参考号
	TradeType              string `json:"trade_type" validate:"required,max=16"`          // 交易类型 (PAY-消费 REFUND-退款 CANCEL-撤销
	TradeStatus            string `json:"trade_status" validate:"required,max=2"`         // 支付状态
	TradeAmount            int64  `json:"trade_amount" validate:"required"`               // 交易金额，单位：分
	PayerAmount            int64  `json:"payer_amount" validate:"omitempty"`              // 付款人实际支付金额，单位：分
	UserID1                string `json:"user_id1" validate:"omitempty,max=64"`           // 用户标识1
	UserID2                string `json:"user_id2" validate:"omitempty,max=64"`           // 用户标识2
	BusiType               string `json:"busi_type" validate:"required,max=16"`           // 支付业务类型
	TradeTime              string `json:"trade_time" validate:"omitempty"`                // 交易完成时间
	AccTradeNo             string `json:"acc_trade_no" validate:"omitempty,max=32"`       // 付款受理交易流水号
	PayerAccountNo         string `json:"payer_account_no" validate:"omitempty,max=32"`   // 付款人账号
	PayerName              string `json:"payer_name" validate:"omitempty,max=32"`         // 付款人名称
	PayerAccountBank       string `json:"payer_account_bank" validate:"omitempty,max=32"` // 付款账号开户行
	AccType                string `json:"acc_type" validate:"omitempty,max=2"`            // 账户类型
	PayMode                string `json:"pay_mode" validate:"omitempty,max=2"`            // 付款方式
	ClientBatchNo          string `json:"client_batch_no" validate:"omitempty,max=6"`     // 终端批次号
	ClientSeqNo            string `json:"client_seq_no" validate:"omitempty,max=6"`       // 终端流水号
	SettleMerchantNo       string `json:"settle_merchant_no" validate:"omitempty,max=32"` // 结算商户号
	SettleTermNo           string `json:"settle_term_no" validate:"omitempty,max=16"`     // 结算终端号
	OriginTradeNo          string `json:"origin_trade_no" validate:"omitempty,max=32"`    // 原交易流水号
	TradeRemark            string `json:"trade_remark" validate:"omitempty,max=64"`       // 交易备注
	AuthCode               string `json:"auth_code" validate:"omitempty,max=64"`          // 快捷签约协议号
	BankType               string `json:"bank_type" validate:"omitempty,max=64"`          // 付款银行
	AccSettleAmount        string `json:"acc_settle_amount" validate:"omitempty"`         // 账户端结算金额
	AccMDiscountAmount     string `json:"acc_mdiscount_amount" validate:"omitempty"`      // 商户侧优惠金额(账户端)
	AccDiscountAmount      string `json:"acc_discount_amount" validate:"omitempty"`       // 账户端优惠金额
	AccOtherDiscountAmount string `json:"acc_other_discount_amount" validate:"omitempty"` // 账户端其它优惠金额
}

type OrderSplitInfo struct {
	SubTradeNo    string `json:"sub_trade_no" validate:"required,max=32"`             // 子单交易流水号
	SubLogNo      string `json:"sub_log_no" validate:"required,max=14"`               // 子单对账单流水号
	OutSubTradeNo string `json:"out_sub_trade_no" validate:"required,max=32"`         // 商户子交易流水号，商户号下唯一
	MerchantNo    string `json:"merchant_no" validate:"required,max=32"`              // 拉卡拉分配的银联商户号
	MerchantName  string `json:"merchant_name,omitempty" validate:"omitempty,max=64"` // 拉卡拉分配的银联商户名称
	TermNo        string `json:"term_no" validate:"required,max=32"`                  // 拉卡拉分配的业务终端号
	Amount        int64  `json:"amount" validate:"required"`                          // 单位分，整数型字符
}
