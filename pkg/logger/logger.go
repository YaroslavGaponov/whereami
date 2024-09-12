package logger

import (
	"context"
	"fmt"
	"time"
)

const (
	TRACE = 1 << 0
	DEBUG = 1 << 1
	INFO  = 1 << 2
	WARN  = 1 << 3
	ERROR = 1 << 4
	FATAL = 1 << 5
)

var (
	LOGGER = struct{}{}
)

type Logger struct {
	level uint8
}

func New() Logger {
	return Logger{
		level: INFO | WARN | ERROR | FATAL,
	}
}

func (l *Logger) AddToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, LOGGER, l)
}

func GetLogger(ctx context.Context) *Logger {
	return ctx.Value(LOGGER).(*Logger)
}

func (l *Logger) SetLogLevel(level string) {
	switch level {
	case "silent":
		l.level = 0
	case "trace":
		l.level = INFO | WARN | ERROR | FATAL | DEBUG | TRACE
	case "debug":
		l.level = INFO | WARN | ERROR | FATAL | DEBUG
	case "info":
		l.level = INFO | WARN | ERROR | FATAL
	case "all":
		l.level = INFO | WARN | ERROR | FATAL | TRACE | DEBUG
	default:
		l.level = INFO | WARN | ERROR | FATAL
	}
}

func (l *Logger) log(level uint8, format string, a ...any) {
	if (l.level & level) == level {
		fmt.Printf("%v [%s]\t", time.DateTime, level2Text(level))
		fmt.Printf(format, a...)
		fmt.Println()
	}
}

func (l *Logger) Trace(format string, a ...any) {
	l.log(TRACE, format, a...)
}

func (l *Logger) Debug(format string, a ...any) {
	l.log(DEBUG, format, a...)
}

func (l *Logger) Info(format string, a ...any) {
	l.log(INFO, format, a...)
}

func (l *Logger) Warn(format string, a ...any) {
	l.log(WARN, format, a...)
}

func (l *Logger) Error(format string, a ...any) {
	l.log(ERROR, format, a...)
}

func (l *Logger) Fatal(format string, a ...any) {
	l.log(FATAL, format, a...)
}

func level2Text(level uint8) string {
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "UNKNOWN"
}
