package logger

import (
	"time"
)

type LogEntry struct {
	Level     string // "debug", "info", "warn", "error"
	Message   string
	Service   string
	Method    string
	UserID    *int64
	Error     error
	Fields    map[string]any
	Timestamp time.Time
}
