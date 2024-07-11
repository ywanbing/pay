# pay
拉卡拉基础支付的 Golang SDK 

官方文档请看 [接口文档](http://open.lakala.com/#/home/document/detail?id=282)

## 功能介绍
1. 拉卡拉支付的接口SDK，目前只会增加聚合平台的支付接口
2. 采用泛型的方式进行接口封装

## 快速开始

该库采用泛型和一些新的特性，需要使用的 Go 版本需要不小于 1.22.4

### 安装

#### 1、使用 Go Modules 管理你的项目

如果你的项目还不是使用 Go Modules 做依赖管理，在项目根目录下执行：

```shell
go mod init
```

#### 2、无需 clone 仓库中的代码，直接在项目目录中执行：
```shell
go get -u github.com/ywanbing/pay
```
来添加依赖，完成 `go.mod` 修改与 SDK 下载。

### 发送支付订单创建

先初始化一个 `lklpay.Client` 实例，再向拉卡拉支付发送请求。

```go
package main

import (
	"context"
	"time"

	"github.com/ywanbing/pay/lklpay"
	"github.com/ywanbing/pay/lklpay/common"
	"github.com/ywanbing/pay/lklpay/model"
)

var (
	Appid           = "OP00000003"
	SerialNo        = "00dfba8194c41b84cf"
	MerchantNo      = "822290059430BCY"
	TermNo          = "A9254710"
	SyncPublicKey   = `-----BEGIN CERTIFICATE----- ******* END CERTIFICATE-----`
	SyncPubicPath   = "" // 读取证书文件（如果没有配置 SyncPublicKey 那么就会读取文件）
	SignPrivateKey  = `-----BEGIN RSA PRIVATE KEY----- ******* END RSA PRIVATE KEY-----`
	SignPrivatePath = "" // 读取私钥文件（如果没有配置 SignPrivateKey 那么就会读取文件）
)

func main() {
	client := lklpay.New(lklpay.Config{
		Appid:           Appid,
		SerialNo:        SerialNo,
		MerchantNo:      MerchantNo,
		TermNo:          TermNo,
		SyncPublicKey:   SyncPublicKey,
		SignPrivateKey:  SignPrivateKey,
		SyncPubicPath:   "",
		SignPrivatePath: "",
	},
		lklpay.WithIsProd(true),     // 是否生产环境
		lklpay.WithNonceStrLen(12),  // 随机字符串长度
		lklpay.WithVerifyResp(true)) // 验证响应的签名
	// ... 还有一些其他的配置

	// resp *model.SpecialCreateRes
	resp, err := client.OrderSpecialCreate(context.Background(), model.SpecialCreateReq{
		OutOrderNo:           "123456789",
		MerchantNo:           MerchantNo,
		TotalAmount:          1,
		OrderEfficientTime:   common.FormatTime(time.Now().Add(time.Minute * 5)), // 提供时间的格式化
		OrderInfo:            "测试订单",
		SupportRefund:        1,
		CloseOrderAutoRefund: "1",
	})
	if err != nil {
		panic(err)
	}
	println(resp)
}


```

## 实现接口

- 收银台订单创建 `client.OrderSpecialCreate` [doc地址](http://open.lakala.com/#/home/document/detail?id=283)
- 收银台订单查询 `client.OrderQuery` [doc地址](http://open.lakala.com/#/home/document/detail?id=284)
- 收银台订单通知 `client.ParseOrderNotify` [doc地址](http://open.lakala.com/#/home/document/detail?id=285)
- 收银台订单关单 `client.OrderClose` [doc地址](http://open.lakala.com/#/home/document/detail?id=722)
- 扫码-退款交易 `client.Refund` [doc地址](http://open.lakala.com/#/home/document/detail?id=113)

## 其他

在项目中的test中有完整的案列，如果想要详细了解可以去看一下。

也欢迎各位有需求可以 pr 或者 提问题哦