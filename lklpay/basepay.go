package lklpay

import (
	"context"

	"github.com/ywanbing/pay/lklpay/basepay/ccss"
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

	baseResp, err := ccss.OrderSpecialCreate[model.SpecialCreateReq, model.SpecialCreateRes](ctx, c.cli, req)
	if err != nil {
		return nil, err
	}
	if baseResp.Code != common.SuccessCode {
		return nil, common.NewErrMsg(baseResp.Code, baseResp.Msg)
	}

	return baseResp.RespData, nil
}
