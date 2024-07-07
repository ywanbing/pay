package model

// OrderCloseReq 关闭订单
//
// 输入参数要么传out_order_no，要么传pay_order_no
type OrderCloseReq struct {
	MerchantNo string `json:"merchant_no" validate:"required,max=32"`             // 商户号
	OutOrderNo string `json:"out_order_no,omitempty" validate:"omitempty,max=32"` // 商户订单号
	PayOrderNo string `json:"pay_order_no,omitempty" validate:"omitempty,max=64"` // 支付订单号
	ChannelID  string `json:"channel_id,omitempty" validate:"omitempty,max=32"`   // 渠道号
}

// OrderCloseRes 关闭订单
type OrderCloseRes struct {
	MerchantNo  string `json:"merchant_no" validate:"required,max=32"`  // 商户号
	OutOrderNo  string `json:"out_order_no" validate:"required,max=32"` // 商户订单号
	PayOrderNo  string `json:"pay_order_no" validate:"required,max=64"` // 支付订单号
	ChannelID   string `json:"channel_id" validate:"required,max=32"`   // 渠道号
	OrderStatus string `json:"order_status" validate:"required,max=2"`  // 订单状态 (0:待支付 1:支付中 2:支付成功 3:支付失败 4:已过期 5:已取消 6：部分退款或者全部退款 7:已关单)
}
