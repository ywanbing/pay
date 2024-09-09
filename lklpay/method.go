package lklpay

const (
	// specialCreateUrl 收银台订单创建
	specialCreateUrl = "/api/v3/ccss/counter/order/special_create"
	orderQueryUrl    = "/api/v3/ccss/counter/order/query"
	orderCloseUrl    = "/api/v3/ccss/counter/order/close"
	refundUrl        = "/api/v3/labs/relation/refund"

	// 聚合扫码
	convergeActive      = "/api/v3/labs/trans/preorder"   // 聚合主扫
	convergeActiveQuery = "/api/v3/labs/query/tradequery" // 聚合扫码交易查询
	convergeActiveClose = "/api/v3/labs/relation/close"   // 聚合主扫交易订单关闭
)
