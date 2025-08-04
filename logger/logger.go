package logger

type Logger interface {
	Log(entry LogEntry)
}
