package lklpay

import (
	"context"

	"github.com/ywanbing/pay/log"
)

type Option func(client *Client)

// WithLogger set custom logger
func WithLogger(logger log.Logger) Option {
	return func(client *Client) {
		client.log = logger
	}
}

// WithIsProd set isProd
func WithIsProd(isProd bool) Option {
	return func(client *Client) {
		client.isProd = isProd
	}
}

// WithContext set context
func WithContext(ctx context.Context) Option {
	return func(client *Client) {
		client.ctx = ctx
	}
}
