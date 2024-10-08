package lklpay

import (
	"context"
	"crypto"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/imroc/req/v3"
	"github.com/ywanbing/pay/lklpay/common"
	"github.com/ywanbing/pay/lklpay/model"
)

func creatClient() *Client {
	return New(Config{
		Appid:    "OP00000003",
		SerialNo: "00dfba8194c41b84cf",
		SyncPublicKey: `-----BEGIN CERTIFICATE-----
MIIEMTCCAxmgAwIBAgIGAXRTgcMnMA0GCSqGSIb3DQEBCwUAMHYxCzAJBgNVBAYT
AkNOMRAwDgYDVQQIDAdCZWlKaW5nMRAwDgYDVQQHDAdCZWlKaW5nMRcwFQYDVQQK
DA5MYWthbGEgQ28uLEx0ZDEqMCgGA1UEAwwhTGFrYWxhIE9yZ2FuaXphdGlvbiBW
YWxpZGF0aW9uIENBMB4XDTIwMTAxMDA1MjQxNFoXDTMwMTAwODA1MjQxNFowZTEL
MAkGA1UEBhMCQ04xEDAOBgNVBAgMB0JlaUppbmcxEDAOBgNVBAcMB0JlaUppbmcx
FzAVBgNVBAoMDkxha2FsYSBDby4sTHRkMRkwFwYDVQQDDBBBUElHVy5MQUtBTEEu
Q09NMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt1zHL54HiI8d2sLJ
lwoQji3/ln0nsvfZ/XVpOjuB+1YR6/0LdxEDMC/hxI6iH2Rm5MjwWz3dmN/6BZeI
gwGeTOWJUZFARo8UduKrlhC6gWMRpAiiGC8wA8stikc5gYB+UeFVZi/aJ0WN0cpP
JYCvPBhxhMvhVDnd4hNohnR1L7k0ypuWg0YwGjC25FaNAEFBYP9EYUyCJjE//9Z7
sMzHR9SJYCqqo6r9bOH9G6sWKuEp+osuAh+kJIxJMHfipw7w3tEcWG0hce9u/el4
cYJtg8/PPMVoccKmeCzMvarr7jdKP4lenJbtwlgyfs+JgNu60KMUJH8RS72wC9NY
uFz09wIDAQABo4HVMIHSMIGSBgNVHSMEgYowgYeAFCnH4DkZPR6CZxRn/kIqVsMo
dJHpoWekZTBjMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpSmluZzEQMA4GA1UE
BwwHQmVpSmluZzEXMBUGA1UECgwOTGFrYWxhIENvLixMdGQxFzAVBgNVBAMMDkxh
a2FsYSBSb290IENBggYBaiUALIowHQYDVR0OBBYEFJ2Kx9YZfmWpkKFnC33C0r5D
K3rFMAwGA1UdEwEB/wQCMAAwDgYDVR0PAQH/BAQDAgeAMA0GCSqGSIb3DQEBCwUA
A4IBAQBZoeU0XyH9O0LGF9R+JyGwfU/O5amoB97VeM+5n9v2z8OCiIJ8eXVGKN9L
tl9QkpTEanYwK30KkpHcJP1xfVkhPi/cCMgfTWQ5eKYC7Zm16zk7n4CP6IIgZIqm
TVGsIGKk8RzWseyWPB3lfqMDR52V1tdA1S8lJ7a2Xnpt5M2jkDXoArl3SVSwCb4D
AmThYhak48M++fUJNYII9JBGRdRGbfJ2GSFdPXgesUL2CwlReQwbW4GZkYGOg9LK
CNPK6XShlNdvgPv0CCR08KCYRwC3HZ0y1F0NjaKzYdGNPrvOq9lA495ONZCvzYDo
gmsu/kd6eqxTs/JwdaIYr4sCMg8Z
-----END CERTIFICATE-----
`,
		SignPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDvDBZyHUDndAGx
rIcsCV2njhNO3vCEZotTaWYSYwtDvkcAb1EjsBFabXZaKigpqFXk5XXNI3NIHP9M
8XKzIgGvc65NpLAfRjVql8JiTvLyYd1gIUcOXMInabu+oX7dQSI1mS8XzqaoVRhD
ZQWhXcJW9bxMulgnzvk0Ggw07AjGF7si+hP/Va8SJmN7EJwfQq6TpSxR+WdIHpbW
dhZ+NHwitnQwAJTLBFvfk28INM39G7XOsXdVLfsooFdglVTOHpNuRiQAj9gShCCN
rpGsNQxDiJIxE43qRsNsRwigyo6DPJk/klgDJa417E2wgP8VrwiXparO4FMzOGK1
5quuoD7DAgMBAAECggEBANhmWOt1EAx3OBFf3f4/fEjylQgRSiqRqg8Ymw6KGuh4
mE4Md6eW/B6geUOmZjVP7nIIR1wte28M0REWgn8nid8LGf+v1sB5DmIwgAf+8G/7
qCwd8/VMg3aqgQtRp0ckb5OV2Mv0h2pbnltkWHR8LDIMwymyh5uCApbn/aTrCAZK
NXcPOyAn9tM8Bu3FHk3Pf24Er3SN+bnGxgpzDrFjsDSHjDFT9UMIc2WdA3tuMv9X
3DDn0bRCsHnsIw3WrwY6HQ8mumdbURk+2Ey3eRFfMYxyS96kOgBC2hqZOlDwVPAK
TPtS4hoq+cQ0sRaJQ4T0UALJrBVHa+EESgRaTvrXqAECgYEA+WKmy9hcvp6IWZlk
9Q1JZ+dgIVxrO65zylK2FnD1/vcTx2JMn73WKtQb6vdvTuk+Ruv9hY9PEsf7S8gH
STTmzHOUgo5x0F8yCxXFnfji2juoUnDdpkjtQK5KySDcpQb5kcCJWEVi9v+zObM0
Zr1Nu5/NreE8EqUl3+7MtHOu1TMCgYEA9WM9P6m4frHPW7h4gs/GISA9LuOdtjLv
AtgCK4cW2mhtGNAMttD8zOBQrRuafcbFAyU9de6nhGwetOhkW9YSV+xRNa7HWTeI
RgXJuJBrluq5e1QGTIwZU/GujpNaR4Qiu0B8TodM/FME7htsyxjmCwEfT6SDYlke
MzTbMa9Q0DECgYBqsR/2+dvD2YMwAgZFKKgNAdoIq8dcwyfamUQ5mZ5EtGQL2yw4
8zibHh/LiIxgUD1Kjk/qQgNsX45NP4iOc0mCkrgomtRqdy+rumbPTNmQ0BEVJCBP
scd+8pIgNiTvnWpMRvj7gMP0NDTzLI3wnnCRIq8WAtR2jZ0Ejt+ZHBziLQKBgQDi
bEe/zqNmhDuJrpXEXmO7fTv3YB/OVwEj5p1Z/LSho2nHU3Hn3r7lbLYEhUvwctCn
Ll2fzC7Wic1rsGOqOcWDS5NDrZpUQGGF+yE/JEOiZcPwgH+vcjaMtp0TAfRzuQEz
NzV8YGwxB4mtC7E/ViIuVULHAk4ZGZI8PbFkDxjKgQKBgG8jEuLTI1tsP3kyaF3j
Aylnw7SkBc4gfe9knsYlw44YlrDSKr8AOp/zSgwvMYvqT+fygaJ3yf9uIBdrIilq
CHKXccZ9uA/bT5JfIi6jbg3EoE9YhB0+1aGAS1O2dBvUiD8tJ+BjAT4OB0UDpmM6
QsFLQgFyXgvDnzr/o+hQJelW
-----END PRIVATE KEY-----
`,
		SyncPubicPath:   "",
		SignPrivatePath: "",
	},
		WithVerifyResp(true))
}

func TestClient_OrderSpecialCreate(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.SpecialCreateReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *model.SpecialCreateRes
		wantErr  bool
	}{
		{
			name: "test1",
			fields: fields{
				cli: creatClient(),
			},
			args: args{
				ctx: context.Background(),
				reqData: model.SpecialCreateReq{
					OutOrderNo:           "0234456789",
					MerchantNo:           "822290059430BCY",
					TotalAmount:          1,
					OrderEfficientTime:   common.FormatTime(time.Now().Add(time.Minute * 5)),
					OrderInfo:            "测试订单",
					SupportRefund:        1,
					CloseOrderAutoRefund: "1",
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := c.OrderSpecialCreate(tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderSpecialCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("OrderSpecialCreate() gotResp = %+v", gotResp)
		})
	}
}

func TestClient_OrderClose(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.OrderCloseReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				cli: creatClient(),
			},
			args: args{
				ctx: context.Background(),
				reqData: model.OrderCloseReq{
					MerchantNo: "822290059430BCY",
					OutOrderNo: "234456789",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := c.OrderClose(tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderClose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("OrderClose() gotResp = %+v", gotResp)
		})
	}
}

func TestClient_OrderQuery(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.OrderQueryReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_query",
			fields: fields{
				cli: creatClient(),
			},
			args: args{
				ctx: context.Background(),
				reqData: model.OrderQueryReq{
					MerchantNo: "822290059430BCY",
					OutOrderNo: "0234456789",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := c.OrderQuery(tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bytes, _ := json.Marshal(gotResp)
			t.Logf("OrderQuery() gotResp = %s ", bytes)
		})
	}
}

func TestClient_Refund(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.RefundReq
	}
	cli := creatClient()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_refund",
			fields: fields{
				cli: cli,
			},
			args: args{
				ctx: context.Background(),
				reqData: model.RefundReq{
					MerchantNo:    "822290059430BCY",
					OutTradeNo:    "0234456788",
					TermNo:        "A9254710",
					RefundAmount:  "1",
					OriginTradeNo: "2024091066210316640558",
					OriginLogNo:   "66210316640558",
					LocationInfo: &model.LocationInfo{
						RequestIp: "127.0.0.1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := c.Refund(tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("OrderQuery() gotResp = %+v", gotResp)
		})
	}
}

func Test_Java(t *testing.T) {
	client := creatClient()
	body := fmt.Sprintf(`{"req_time":"20210907150256","version":"3.0","out_org_code":"OP00000003","req_data":{"merchant_no":"%s","term_no":"%s","out_trade_no":"FD660E1FAA3A4470933CDEDAE1EC1D8E","auth_code":"135178236713755038","total_amount":"123","location_info":{"request_ip":"10.176.1.192","location":"+37.123456789,-121.123456789"},"out_order_no":"08F4542EEC6A4497BC419161747A92FA"}}`,
		"822290059430BCY", "A9254710")

	cli := req.C().SetBaseURL(TestUrl).SetCommonHeaders(map[string]string{"Content-Type": "application/json"})

	now := time.Now().Unix()
	randomStr := common.RandomNumber(32)
	validStr := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", client.cfg.Appid, client.cfg.SerialNo, now, randomStr, body)

	sign, err := common.Sign([]byte(validStr), client.privateKey, crypto.SHA256, sha256.New())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sign: %s", sign)

	auth := fmt.Sprintf(common.AuthFormat, common.Algorism_SHA256, client.cfg.Appid, client.cfg.SerialNo, now, randomStr, sign)

	t.Logf("auth: %s", auth)

	response := cli.Post("/api/v3/labs/trans/micropay").SetHeader("Authorization", auth).SetBody(body).Do()

	t.Log(response)
}

func TestVerifySign(t *testing.T) {
	client := creatClient()
	/*
		Lklapi-Appid: OP00000003
		Lklapi-Traceid: eqn9QTc7qKmPCSCMm
		Lklapi-Serial: 1745381c327
		Lklapi-Signature: kVgF6/eH/05si6dM4WFqka7WmlFkCniBvGYnYcbFfnpL+nplLbY7swxfg4O4G0u8e9oh92+jrV+lCveakoE4K8wHqJUSJsW/g0X5qwDhThpKbL3bKSwUyjTXh+qrM6JkQk8hHoJt6eI+blTNqvJ27LfYJMjRigBYigNLWaZXFWd88mlkF9Zk/s60MPGXfUcJPPfGjzuT7x+zoRcRkVCWZO1KhsCeJphU6xM2VTK92eeFndsY+gikicRcJCKToV5Gc3fkOwk2bbFCw9ce2hlUba86NcQMiqbwD8ngy4QI/A5yzX+6BIAMTu2+x0hpBzyeoO0vHI0BuTvEDOF8sv38rg==
		Lklapi-Timestamp: 1720678638
		Lklapi-Nonce: ArgBcKYQtHLD
	*/
	var appid, serialNo, ts, nonce, body, sign string
	appid = "OP00000003"
	serialNo = "1745381c327"
	ts = "1720678638"
	nonce = "ArgBcKYQtHLD"
	body = `{"code":"000000","msg":"操作成功","resp_time":"20240711141715","resp_data":{"merchant_no":"822290059430BCY","channel_id":"95","out_order_no":"0234456789","order_create_time":"20240711141715","order_efficient_time":"20240711142216","pay_order_no":"24071111012001101011001304914","total_amount":"1","counter_url":"https://pay.wsmsd.cn/r/0000?pageStyle%3DV2%26token%3DCCSSIZ5wnqqemB40EXc7U40gGcb7rxoKuxVhhVI7XyulHEUboR1J21LJqZZPDO053tJ2vIjabliIY1f32w%3D%3D%26amount%3D1%26payOrderNo%3D24071111012001101011001304914%26mndf%3D1"}}`
	sign = "kVgF6/eH/05si6dM4WFqka7WmlFkCniBvGYnYcbFfnpL+nplLbY7swxfg4O4G0u8e9oh92+jrV+lCveakoE4K8wHqJUSJsW/g0X5qwDhThpKbL3bKSwUyjTXh+qrM6JkQk8hHoJt6eI+blTNqvJ27LfYJMjRigBYigNLWaZXFWd88mlkF9Zk/s60MPGXfUcJPPfGjzuT7x+zoRcRkVCWZO1KhsCeJphU6xM2VTK92eeFndsY+gikicRcJCKToV5Gc3fkOwk2bbFCw9ce2hlUba86NcQMiqbwD8ngy4QI/A5yzX+6BIAMTu2+x0hpBzyeoO0vHI0BuTvEDOF8sv38rg=="

	err := client.VerifySign(appid, serialNo, ts, nonce, body, sign)
	if err != nil {
		t.Log("verify sign failed")
	}
	t.Log("verify sign ok")
}

func TestClient_ConvergeActiveCreate(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.ConvergeActiveReq[model.AliPayAccBusiFields]
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *model.SpecialCreateRes
		wantErr  bool
	}{
		{
			name: "test1",
			fields: fields{
				cli: creatClient(),
			},
			args: args{
				ctx: context.Background(),
				reqData: model.ConvergeActiveReq[model.AliPayAccBusiFields]{
					MerchantNo:   "822290059430BCY",
					TermNo:       "A9254710",
					OutTradeNo:   "0234456777",
					AccountType:  common.AccountType_ALIPAY,
					TransType:    common.TransType_NATIVE,
					TotalAmount:  "1",
					Subject:      "测试订单",
					LocationInfo: &model.LocationInfo{RequestIp: "127.0.0.1"},
					AccBusiFields: &model.AliPayAccBusiFields{
						TimeoutExpress: "5",
						MinAge:         "18",
					},
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := ConvergeActive[model.AliPayAccBusiFields, model.AliPayNativeResp](c, tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderSpecialCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("OrderSpecialCreate() gotResp = %+v", gotResp)
			t.Logf("alipay resp: %+v", gotResp.AccRespFields)
		})
	}
}

func TestClient_ConvergeActiveQuery(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.ConvergeActiveQueryReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_query",
			fields: fields{
				cli: creatClient(),
			},
			args: args{
				ctx: context.Background(),
				reqData: model.ConvergeActiveQueryReq{
					MerchantNo: "822290059430BCY",
					TermNo:     "A9254710",
					OutTradeNo: "0234456777",
					// TradeNo:    "2024091066210316640558",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := c.ConvergeActiveQuery(tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bytes, _ := json.Marshal(gotResp)
			t.Logf("OrderQuery() gotResp = %s ", bytes)
		})
	}
}

func TestClient_ConvergeActiveClose(t *testing.T) {
	type fields struct {
		cli *Client
	}
	type args struct {
		ctx     context.Context
		reqData model.ConvergeActiveCloseReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				cli: creatClient(),
			},
			args: args{
				ctx: context.Background(),
				reqData: model.ConvergeActiveCloseReq{
					MerchantNo:       "822290059430BCY",
					TermNo:           "A9254710",
					OriginOutTradeNo: "0234456777",
					// OriginTradeNo:    "2024091066210316640583",
					LocationInfo: &model.LocationInfo{RequestIp: "127.0.0.1"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.cli
			gotResp, err := c.ConvergeActiveClose(tt.args.ctx, tt.args.reqData)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderClose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("OrderClose() gotResp = %+v", gotResp)
		})
	}
}
