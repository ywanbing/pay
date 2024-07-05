package lklpay

import (
	"context"

	"github.com/imroc/req/v3"
	"github.com/ywanbing/pay/log"
)

// Client pay client
type Client struct {
	ctx context.Context // 上下文
	cfg Config          // 配置
	log log.Logger      // logger

	isProd bool // 是否生产环境
	cli    *req.Client
}

// New a pay client
func New(cfg Config, options ...Option) *Client {
	client := &Client{
		ctx:    context.Background(),
		cfg:    cfg,
		log:    log.DefLogger(),
		isProd: false,
	}

	for _, option := range options {
		option(client)
	}

	// 初始化http客户端
	client.initHttpClient()

	return client
}

func (c *Client) initHttpClient() {
	cli := req.NewClient()
	if c.isProd {
		cli.SetBaseURL(ProdUrl)
	} else {
		cli.SetBaseURL(TestUrl)
	}

	cli.SetLogger(c.log)
	cli.SetCommonContentType("application/json")
	c.cli = cli
}
