package lklpay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imroc/req/v3"
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

	c.log.Debugf("auth: %s", auth)

	response := c.cli.Post(url).SetHeader("Authorization", auth).SetBodyBytes(bytes).Do(ctx)

	// 是否验证签名
	if c.verifyResp && !verifyResp(c, response) {
		return nil, fmt.Errorf("verify response error")
	}

	err = response.Into(resp)
	return
}

func verifyResp(c *Client, resp *req.Response) bool {
	body, _ := resp.ToBytes()
	return c.VerifyResponseSign(resp.Header, body) == nil
}
