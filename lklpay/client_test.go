package lklpay

import "testing"

func createClient() *Client {
	return New(Config{
		Appid:      "OP00000003",
		SerialNo:   "00dfba8194c41b84cf",
		MerchantNo: "822290059430BCY",
		TermNo:     "A9254710",
		SyncPublicKey: `-----BEGIN CERTIFICATE-----
MIIDYTCCAkmgAwIBAgIJAN+6gZTEG4TPMA0GCSqGSIb3DQEBCwUAMEkxCzAJBgNV
BAYTAlVTMREwDwYDVQQIEwhzaGFuZ2hhaTERMA8GA1UEBxMIc2hhbmdoYWkxFDAS
BgNVBAMUC2xha2FsYV8yMDIxMB4XDTIxMDYxODA3MjEzNFoXDTMxMDYxOTA3MjEz
NFowSTELMAkGA1UEBhMCVVMxETAPBgNVBAgTCHNoYW5naGFpMREwDwYDVQQHEwhz
aGFuZ2hhaTEUMBIGA1UEAxQLbGFrYWxhXzIwMjEwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDvDBZyHUDndAGxrIcsCV2njhNO3vCEZotTaWYSYwtDvkcA
b1EjsBFabXZaKigpqFXk5XXNI3NIHP9M8XKzIgGvc65NpLAfRjVql8JiTvLyYd1g
IUcOXMInabu+oX7dQSI1mS8XzqaoVRhDZQWhXcJW9bxMulgnzvk0Ggw07AjGF7si
+hP/Va8SJmN7EJwfQq6TpSxR+WdIHpbWdhZ+NHwitnQwAJTLBFvfk28INM39G7XO
sXdVLfsooFdglVTOHpNuRiQAj9gShCCNrpGsNQxDiJIxE43qRsNsRwigyo6DPJk/
klgDJa417E2wgP8VrwiXparO4FMzOGK15quuoD7DAgMBAAGjTDBKMAkGA1UdEwQC
MAAwEQYJYIZIAYb4QgEBBAQDAgTwMAsGA1UdDwQEAwIFoDAdBgNVHSUEFjAUBggr
BgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAI21YYAlH+Pc1ISv
nbQrGqL8suGL0Hh/8hGaFfrJEJEKr9OeC8jElUhck2MTmfu/Y1lB7r8RBrhGPXi4
kTXmB6ADs/9+ezNW3WXyFj7fhs3JcZ3mo33T9wyQySDKd//JrEtrTsc/s2PZ602y
qNmPomXSzjrlugaMyC7LI9sR44mc7sQnchjHoxrQFD5/usTFW72UQfYCORsQWYMt
0KKEyAcpRL51RE3xbX1WDtduFYGP62PbwLAn2nCL/j1wlF5hltWj7sditWqKgso5
F8BTffn2Bb0RdsNxqwMy1cTPrWLeXVOqMDu3ge7hvoav8lZKTjk5Kmqhs7wNAQXK
mg9qSwo=
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

func TestVerifySignForAuth(t *testing.T) {
	client := createClient()
	body := `{"code":"000000","msg":"操作成功","resp_time":"20240711141715","resp_data":{"merchant_no":"822290059430BCY","channel_id":"95","out_order_no":"0234456789","order_create_time":"20240711141715","order_efficient_time":"20240711142216","pay_order_no":"24071111012001101011001304914","total_amount":"1","counter_url":"https://pay.wsmsd.cn/r/0000?pageStyle%3DV2%26token%3DCCSSIZ5wnqqemB40EXc7U40gGcb7rxoKuxVhhVI7XyulHEUboR1J21LJqZZPDO053tJ2vIjabliIY1f32w%3D%3D%26amount%3D1%26payOrderNo%3D24071111012001101011001304914%26mndf%3D1"}}`

	// 获取签名信息
	auth, err := client.getRsaNotifySign([]byte(body))
	if err != nil {
		t.Log("get sign failed", err)
	}

	// 解析auth
	err = client.VerifySignForAuth(auth, []byte(body))
	if err != nil {
		t.Log("verify sign failed", err)
	}
	t.Log("verify sign ok")
}
