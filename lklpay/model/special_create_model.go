package model

// SpecialCreateReq 收银台订单创建请求
type SpecialCreateReq struct {
	OutOrderNo           string          `json:"out_order_no" validate:"required,max=32"`                          // 商户订单号，必填且长度不超过32
	MerchantNo           string          `json:"merchant_no" validate:"required,max=32"`                           // 银联商户号，必填且长度不超过32
	VposID               string          `json:"vpos_id,omitempty" validate:"omitempty,max=32"`                    // 交易设备标识，可选且长度不超过32
	ChannelID            string          `json:"channel_id,omitempty" validate:"omitempty,max=32"`                 // 渠道号，可选且长度不超过32
	TotalAmount          int64           `json:"total_amount" validate:"required,gte=0"`                           // 订单金额，单位：分，JPY和KRW的单位是元，即200日元，填“200"
	OrderEfficientTime   string          `json:"order_efficient_time" validate:"required,len=14"`                  // 订单有效期，必填且格式yyyyMMddHHmmss
	NotifyURL            string          `json:"notify_url,omitempty" validate:"omitempty,url"`                    // 订单支付成功后商户接收订单通知的地址，可选且需为有效URL
	SupportCancel        int             `json:"support_cancel" validate:"gte=0,lte=1"`                            // 是否支持撤销，默认0或1
	SupportRefund        int             `json:"support_refund" validate:"gte=0,lte=1"`                            // 是否支持退款，默认0或1
	SupportRepeatPay     int             `json:"support_repeat_pay" validate:"gte=0,lte=1"`                        // 是否支持“多次发起支付”，默认0或1
	OutUserID            string          `json:"out_user_id,omitempty" validate:"omitempty,max=64"`                // 发起订单方的userId，可选且长度不超过64
	CallbackURL          string          `json:"callback_url,omitempty" validate:"omitempty,url"`                  // 客户端下单完成支付后返回的商户网页跳转地址，可选且需为有效URL
	OrderInfo            string          `json:"order_info" validate:"required,max=64"`                            // 订单标题，必填且长度不超过64
	TermNo               string          `json:"term_no,omitempty" validate:"omitempty,max=32"`                    // 结算终端号，可选且长度不超过32
	SplitMark            string          `json:"split_mark,omitempty" validate:"omitempty,max=2"`                  // 合单标识，"1":为合单，不填默认是为非合单，可选且长度不超过2
	SettleType           string          `json:"settle_type,omitempty" validate:"omitempty,max=4"`                 // 结算类型（非合单），可选且长度不超过4
	OutSplitInfo         []*OutSplitInfo `json:"out_split_info,omitempty" validate:"omitempty,dive"`               // 拆单信息,合单标识为“1”时必传该字段。暂无特定验证规则
	CounterParam         string          `json:"counter_param,omitempty" validate:"omitempty,json"`                // 收银台参数，暂无特定验证规则( common.CounterParam
	CounterRemark        string          `json:"counter_remark,omitempty" validate:"omitempty,max=64"`             // 收银台备注，可选且长度不超过64
	BusiTypeParam        string          `json:"busi_type_param,omitempty" validate:"omitempty,json"`              // 业务类型控制参数（ []model.BusiTypeParam 类型），暂无特定验证规则
	SgnInfo              []string        `json:"sgn_info,omitempty"`                                               // 签约协议号列表，暂无特定验证规则
	ProductID            string          `json:"product_id,omitempty" validate:"omitempty,max=6"`                  // 指定产品编号，可选且长度不超过6
	GoodsMark            string          `json:"goods_mark,omitempty" validate:"omitempty,max=2"`                  // 商品信息标识（1:含商品信息，不填默认不含商品信息），暂无特定验证规则
	GoodsField           string          `json:"goods_field,omitempty"`                                            // 商品信息域，暂无特定验证规则
	OrderSceneField      string          `json:"order_scene_field,omitempty"`                                      // 订单场景域，暂无特定验证规则
	AgeLimit             string          `json:"age_limit,omitempty" validate:"omitempty,oneof=0 1"`               // 年龄限制，暂无特定验证规则
	RepeatPayAutoRefund  string          `json:"repeat_pay_auto_refund,omitempty" validate:"omitempty,oneof=0 1"`  // 重复支付后自动退货，暂无特定验证规则
	RepeatPayNotify      string          `json:"repeat_pay_notify,omitempty" validate:"omitempty,oneof=0 1"`       // 重复支付订单通知，暂无特定验证规则
	CloseOrderAutoRefund string          `json:"close_order_auto_refund,omitempty" validate:"omitempty,oneof=0 1"` // 关闭订单后支付成功触发自动退货，暂无特定验证规则
	ShopName             string          `json:"shop_name,omitempty" validate:"omitempty,max=64"`                  // 网点名称，可选且长度不超过64
	InteRouting          string          `json:"inte_routing,omitempty" validate:"omitempty,oneof=0 1"`            // 智能路由下单标识，可选且长度不超过2
}

type OutSplitInfo struct {
	OutSubOrderNo string `json:"out_sub_order_no" validate:"required,max=32"`           // 外部子订单号，必填且长度不超过32
	MerchantNo    string `json:"merchant_no" validate:"required,max=32"`                // 商户号，必填且长度不超过32
	TermNo        string `json:"term_no" validate:"required,max=32"`                    // 终端号，必填且长度不超过32
	Amount        int64  `json:"amount" validate:"required"`                            // 金额，单位分，整数型字符，必填
	SettleType    string `json:"settle_type,omitempty" validate:"omitempty,oneof=0 ''"` // 结算类型（合单），可选，只能是"0"或空字符串
}

// BusiTypeParam 业务类型控制参数
type BusiTypeParam struct {
	BusiType string `json:"busi_type" validate:"required,max=32"` // 业务类型，必填且长度不超过32
	Params   struct {
		PayMode string `json:"pay_mode,omitempty" validate:"omitempty"`
		CrdFlg  string `json:"crd_flg"`
	} `json:"params" validate:"required"`
}

// GoodsField
type GoodsField struct {
	GoodsAmt         int64  `json:"goods_amt" validate:"required"`                // 商品单价，单位：分，必填
	GoodsNum         int    `json:"goods_num" validate:"required,gt=0"`           // 商品数量，必填
	GoodsPricingUnit string `json:"goods_pricing_unit" validate:"required,max=8"` // 商品计价单位，1-箱 2-件 3-瓶 4-个，必填且长度不超过8
	GoodsName        string `json:"goods_name" validate:"required,max=128"`       // 商品名称，必填且长度不超过128
	TePlatformType   string `json:"te_platform_type" validate:"required,max=2"`   // 交易电商平台类型，1-境内平台 2-境外平台，必填且长度不超过2
	TePlatformName   string `json:"te_platform_name" validate:"required,max=256"` // 交易电商平台名称，必填且长度不超过256
	GoodsType        string `json:"goods_type" validate:"required,max=8"`         // 交易商品类型，必填且长度不超过8
}

// OrderSceneField
type OrderSceneField struct {
	OrderSceneType string `json:"order_scene_type" validate:"required,max=16"` // 订单场景类型，HB_FQ：花呗分期场景;KL_FQ：考拉分期场景。必填且长度不超过16
	SceneInfo      string `json:"scene_info,omitempty"`                        // 订单场景信息，（HBFQSceneInfo）可选
}

// HBFQSceneInfo
type HBFQSceneInfo struct {
	HbFqNum           string `json:"hbFqNum" validate:"required"`           // 花呗分期期数，必填
	HbFqSellerPercent string `json:"hbFqSellerPercent" validate:"required"` // 卖家承担手续费比例，必填
}

// SpecialCreateRes 收银台订单创建响应
type SpecialCreateRes struct {
	MerchantNo         string `json:"merchant_no"`          // 银联商户号，必填且长度不超过32
	ChannelID          string `json:"channel_id"`           // 字段，必填且长度不超过32
	OutOrderNo         string `json:"out_order_no"`         // 商户订单号，必填且长度不超过32
	OrderCreateTime    string `json:"order_create_time"`    // 创建订单时间，必填且格式yyyyMMddHHmmss
	OrderEfficientTime string `json:"order_efficient_time"` // 订单有效截至时间，必填且格式yyyyMMddHHmmss
	PayOrderNo         string `json:"pay_order_no"`         // 平台订单号，必填且长度不超过64
	TotalAmount        string `json:"total_amount"`         // 订单金额，单位：分，必填
	CounterURL         string `json:"counter_url"`          // 收银台地址信息，必填且长度不超过256
}
