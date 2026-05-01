package logger

import "context"

type contextKey struct{}

var loggerKey = contextKey{}

// WithContext returns a new context with the given logger attached.
func WithContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// FromContext returns the logger attached to the context, or the global logger if none is found.
func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(loggerKey).(Logger); ok {
		return l
	}
	return get() // Fallback to global logger
}
