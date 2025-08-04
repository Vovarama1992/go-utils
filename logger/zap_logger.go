package logger

import (
	"time"

	"go.uber.org/zap"
)

type ZapLogger struct {
	z *zap.SugaredLogger
}

func NewZapLogger(z *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{z}
}

func (l *ZapLogger) Log(entry LogEntry) {

	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	fields := []any{
		"service", entry.Service,
		"method", entry.Method,
		"timestamp", entry.Timestamp.Format(time.RFC3339),
	}

	if entry.UserID != nil {
		fields = append(fields, "user_id", *entry.UserID)
	}

	if entry.Error != nil {
		fields = append(fields, "error", entry.Error.Error())
	}

	for k, v := range entry.Fields {
		fields = append(fields, k, v)
	}

	switch entry.Level {
	case "debug":
		l.z.Debugw(entry.Message, fields...)
	case "info":
		l.z.Infow(entry.Message, fields...)
	case "warn":
		l.z.Warnw(entry.Message, fields...)
	case "error":
		l.z.Errorw(entry.Message, fields...)
	default:
		l.z.Infow(entry.Message, fields...)
	}
}
