package logger

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)

	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value any
}

var log Logger

func SetLogger(l Logger) {
	if l == nil {
		panic("logger cannot be nil")
	}
	log = l
}

func get() Logger {
	if log == nil {
		panic("logger not initialized")
	}
	return log
}

func Debug(msg string, fields ...Field) {
	get().Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	get().Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	get().Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	get().Error(msg, fields...)
}
