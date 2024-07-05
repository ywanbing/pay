package log

import logx "log"

type log struct{}

func (l *log) Errorf(format string, v ...interface{}) {
	logx.Printf(format, v...)
}

func (l *log) Warnf(format string, v ...interface{}) {
	logx.Printf(format, v...)
}

func (l *log) Debugf(format string, v ...interface{}) {
	logx.Printf(format, v...)
}

func (l *log) Logf(format string, v ...interface{}) {
	logx.Printf(format, v...)
}

// DefLogger returns default logger instance
func DefLogger() Logger {
	return &log{}
}
