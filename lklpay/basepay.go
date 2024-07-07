package lklpay

import (
	"context"
	"encoding/json"

	"github.com/ywanbing/pay/lklpay/common"
	"github.com/ywanbing/pay/lklpay/model"
)

// OrderSpecialCreate 收银台订单创建
func (c *Client) OrderSpecialCreate(ctx context.Context, reqData model.SpecialCreateReq) (resp *model.SpecialCreateRes, err error) {
	// 验证请求参数
	err = c.valid.StructCtx(ctx, reqData)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}

	req := model.BaseRequest[model.SpecialCreateReq]{
		ReqData: &reqData,
		ReqTime: common.GetReqTime(),
		Version: "3.0",
	}

	baseResp, err := doPost[model.SpecialCreateReq, model.SpecialCreateRes](ctx, c, specialCreateUrl, req)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}
	if baseResp.Code != common.SuccessCode {
		return nil, common.NewErrMsg(baseResp.Code, baseResp.Msg)
	}

	return baseResp.RespData, nil
}

// ParseOrderNotify 解析订单通知
func (c *Client) ParseOrderNotify(body []byte) (notify *model.OrderNotify, err error) {
	notify = new(model.OrderNotify)
	err = json.Unmarshal(body, notify)
	if err != nil {
		return nil, err
	}
	return

}

// OrderQuery 收银台订单查询
func (c *Client) OrderQuery(ctx context.Context, reqData model.OrderQueryReq) (resp *model.OrderQueryRes, err error) {
	// 验证请求参数
	err = c.valid.StructCtx(ctx, reqData)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}

	req := model.BaseRequest[model.OrderQueryReq]{
		ReqData: &reqData,
		ReqTime: common.GetReqTime(),
		Version: "3.0",
	}

	baseResp, err := doPost[model.OrderQueryReq, model.OrderQueryRes](ctx, c, orderQueryUrl, req)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}
	if baseResp.Code != common.SuccessCode {
		return nil, common.NewErrMsg(baseResp.Code, baseResp.Msg)
	}

	return baseResp.RespData, nil
}

// OrderClose 收银台订单关闭
func (c *Client) OrderClose(ctx context.Context, reqData model.OrderCloseReq) (resp *model.OrderCloseRes, err error) {
	// 验证请求参数
	err = c.valid.StructCtx(ctx, reqData)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}

	req := model.BaseRequest[model.OrderCloseReq]{
		ReqData: &reqData,
		ReqTime: common.GetReqTime(),
		Version: "3.0",
	}

	baseResp, err := doPost[model.OrderCloseReq, model.OrderCloseRes](ctx, c, orderCloseUrl, req)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}
	if baseResp.Code != common.SuccessCode {
		return nil, common.NewErrMsg(baseResp.Code, baseResp.Msg)
	}

	return baseResp.RespData, nil
}

// Refund 收银台订单关闭
func (c *Client) Refund(ctx context.Context, reqData model.RefundReq) (resp *model.RefundRes, err error) {
	// 验证请求参数
	err = c.valid.StructCtx(ctx, reqData)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}

	req := model.BaseRequest[model.RefundReq]{
		ReqData: &reqData,
		ReqTime: common.GetReqTime(),
		Version: "3.0",
	}

	baseResp, err := doPost[model.RefundReq, model.RefundRes](ctx, c, refundUrl, req)
	if err != nil {
		return nil, common.NewErrMsg(common.InternalCode, err.Error())
	}
	if baseResp.Code != common.SuccessCode {
		return nil, common.NewErrMsg(baseResp.Code, baseResp.Msg)
	}

	return baseResp.RespData, nil
}
