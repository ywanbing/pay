package lklpay

import (
	"context"
	"crypto"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/ywanbing/pay/log"
)

type Pool struct {
	pool sync.Pool
}

func NewPool() *Pool {
	return &Pool{
		pool: sync.Pool{
			New: func() interface{} {
				return &Client{
					ctx:      context.Background(),
					log:      log.DefLogger(),
					isProd:   false,
					valid:    validator.New(),
					hashType: crypto.SHA256,
				}
			},
		},
	}

}

func (p *Pool) GetClient(cfg Config, options ...Option) *Client {
	client := p.pool.Get().(*Client)
	client.ReSet(cfg, options...)
	return client
}

func (p *Pool) PutClient(client *Client) {
	p.pool.Put(client)
}
