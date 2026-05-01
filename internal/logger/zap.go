package logger

import (
	"strings"
	"time"

	"github.com/shivam-jainn/goldfiber/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	l *zap.Logger
}

func NewZap(cfg *config.Config) Logger {
	env := cfg.Env
	level := parseLevel(cfg.LogLevel)

	var zapCfg zap.Config

	if env == "prod" {
		zapCfg = zap.Config{
			Level:            zap.NewAtomicLevelAt(level),
			Encoding:         "json",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:      "ts",
				LevelKey:     "level",
				MessageKey:   "msg",
				CallerKey:    "caller",
				EncodeTime:   zapcore.ISO8601TimeEncoder,
				EncodeLevel:  zapcore.LowercaseLevelEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}
	} else {
		zapCfg = zap.Config{
			Level:            zap.NewAtomicLevelAt(level),
			Encoding:         "console",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:     "ts",
				LevelKey:    "level",
				MessageKey:  "msg",
				CallerKey:   "caller",
				EncodeLevel: zapcore.CapitalColorLevelEncoder,
				EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.Format("15:04:05"))
				},
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}
	}

	z, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}

	return &zapLogger{l: z}
}

func (z *zapLogger) Debug(msg string, fields ...Field) {
	z.l.Debug(msg, toZap(fields)...)
}

func (z *zapLogger) Info(msg string, fields ...Field) {
	z.l.Info(msg, toZap(fields)...)
}

func (z *zapLogger) Warn(msg string, fields ...Field) {
	z.l.Warn(msg, toZap(fields)...)
}

func (z *zapLogger) Error(msg string, fields ...Field) {
	z.l.Error(msg, toZap(fields)...)
}

func (z *zapLogger) With(fields ...Field) Logger {
	return &zapLogger{l: z.l.With(toZap(fields)...)}
}

func toZap(fields []Field) []zap.Field {
	out := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		out = append(out, zap.Any(f.Key, f.Value))
	}
	return out
}

func parseLevel(l string) zapcore.Level {
	switch strings.ToLower(l) {
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
