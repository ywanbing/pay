package ccss

import (
	"context"
	"encoding/json"

	"github.com/imroc/req/v3"
	"github.com/ywanbing/pay/lklpay/model"
)

// OrderSpecialCreate 收银台订单创建
func OrderSpecialCreate[D any, R any](ctx context.Context, cli *req.Client, reqData model.BaseRequest[D]) (resp *model.BaseResponse[R], err error) {
	resp = new(model.BaseResponse[R])
	bytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	// TODO 增加签名

	err = cli.Post(specialCreateUrl).SetBodyBytes(bytes).Do(ctx).Into(resp)
	return
}
