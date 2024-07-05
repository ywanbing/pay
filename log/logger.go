package log

import "github.com/imroc/req/v3"

// Logger is an interface for loggers
type Logger interface {
	req.Logger
	Logf(format string, v ...interface{})
}
