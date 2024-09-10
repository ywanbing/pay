package model

// OrderQueryReq ...
//
//	Note: 输入参数要么传out_order_no+merchant_no
//	要么传pay_order_no+channel_id(如果订单创建时传channel_id，查询时也要传channel_id)
type OrderQueryReq struct {
	OutOrderNo string `json:"out_order_no,omitempty" validate:"omitempty,max=32"` // 商户订单号
	MerchantNo string `json:"merchant_no,omitempty" validate:"omitempty,max=32"`  // 银联商户号
	PayOrderNo string `json:"pay_order_no,omitempty" validate:"omitempty,max=64"` // 支付订单号
	ChannelID  string `json:"channel_id,omitempty" validate:"omitempty,max=32"`   // 渠道号
}

// OrderQueryRes ...
type OrderQueryRes struct {
	PayOrderNo         string            `json:"pay_order_no" validate:"required,max=64"`         // 支付订单号
	OutOrderNo         string            `json:"out_order_no" validate:"required,max=32"`         // 商户订单号
	ChannelID          string            `json:"channel_id" validate:"required,max=32"`           // 渠道号
	TransMerchantNo    string            `json:"trans_merchant_no" validate:"omitempty,max=32"`   // 交易商户号
	TransTermNo        string            `json:"trans_term_no" validate:"omitempty,max=16"`       // 交易终端号
	MerchantNo         string            `json:"merchant_no" validate:"required,max=32"`          // 结算商户号
	TermNo             string            `json:"term_no" validate:"required,max=16"`              // 结算终端号
	OrderStatus        string            `json:"order_status" validate:"required,max=2"`          // 订单状态
	OrderInfo          string            `json:"order_info" validate:"omitempty,max=100"`         // 订单描述
	TotalAmount        int64             `json:"total_amount" validate:"required"`                // 订单金额，单位：分
	OrderCreateTime    string            `json:"order_create_time" validate:"required"`           // 订单创建时间(格式yyyyMMddHHmmss)
	OrderEfficientTime string            `json:"order_efficient_time" validate:"required"`        // 订单有效时间(格式yyyyMMddHHmmss)
	SettleType         string            `json:"settle_type" validate:"omitempty,max=4"`          // 结算类型（非合单）
	SplitMark          string            `json:"split_mark" validate:"omitempty,max=2"`           // 合单标识
	CounterParam       string            `json:"counter_param" validate:"omitempty,max=1024"`     // 收银台参数(json字符串,类型： common.CounterParam)
	CounterRemark      string            `json:"counter_remark" validate:"omitempty,max=64"`      // 收银台备注
	BusiTypeParam      string            `json:"busi_type_param" validate:"omitempty,max=256"`    // 业务类型控制参数(json字符串：[]model.BusiTypeParam)
	OutSplitInfo       []*OutSplitInfo   `json:"out_split_info" validate:"omitempty,dive"`        // 商户拆单信息
	SplitInfo          []*OrderSplitInfo `json:"split_info" validate:"omitempty,dive"`            // 交易拆单信息
	SgnInfo            []string          `json:"sgn_info" validate:"omitempty"`                   // 签约协议号列表
	GoodsMark          string            `json:"goods_mark" validate:"omitempty,max=2"`           // 商品标识
	GoodsField         string            `json:"goods_field" validate:"omitempty,max=2048"`       // 商品信息
	OrderTradeInfoList []*OrderTradeInfo `json:"order_trade_info_list" validate:"omitempty,dive"` // 订单交易信息列表
}
