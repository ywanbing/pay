package model

// BaseRequest 基础请求
type BaseRequest[T any] struct {
	ReqTime string `json:"req_time" validate:"required"`
	Version string `json:"version" validate:"required"`
	ReqData *T     `json:"req_data" validate:"required"`
}

// BaseResponse 基础响应
type BaseResponse[T any] struct {
	Code     string `json:"code"`                // 返回业务代码(000000为成功，其余按照错误信息来定)
	Msg      string `json:"msg"`                 // 返回业务代码描述
	RespTime string `json:"resp_time"`           // 响应时间，格式yyyyMMddHHmmss
	RespData *T     `json:"resp_data,omitempty"` // 返回数据.下文定义的响应均为该属性中的内容
}
