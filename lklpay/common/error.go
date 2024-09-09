package common

import (
	"fmt"
)

const (
	SuccessCode   = "000000"
	InternalCode  = "-1" // 内部错误
	PramErrorCode = "-2" // 参数错误
)

type ErrMsg struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Error 组装错误信息
func (e ErrMsg) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}

func NewErrMsg(code, msg string) ErrMsg {
	return ErrMsg{
		Code: code,
		Msg:  msg,
	}
}
