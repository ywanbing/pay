package lklpay

import (
	"context"
	"encoding/json"

	"github.com/ywanbing/pay/lklpay/model"
)

// doPost 发送POST请求
func doPost[D any, R any](ctx context.Context, c *Client, url string, reqData model.BaseRequest[D]) (resp *model.BaseResponse[R], err error) {
	resp = new(model.BaseResponse[R])
	bytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	// 获取签名信息
	auth, err := c.getRsaSign(bytes)
	if err != nil {
		return nil, err
	}

	err = c.cli.Post(url).SetHeader("Authorization", auth).SetBodyBytes(bytes).Do(ctx).Into(resp)
	return
}
