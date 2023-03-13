package log

import (
	"context"
	"go.uber.org/zap"
)

const (
	logContextKey = iota
)

// WithContext returns a copy of context in which the log value is set.
func WithContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

// FromContext returns the value of the log key on the ctx.
func FromContext(ctx context.Context) *zap.Logger {
	if ctx != nil {
		logger := ctx.Value(logContextKey)
		if logger != nil {
			return logger.(*zap.Logger)
		}
	}

	return WithName("Unknown-Context")
}
