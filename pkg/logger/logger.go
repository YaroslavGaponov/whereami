package logger

import (
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

type Logger struct {
	level uint8
}

func New() Logger {
	return Logger{
		level: INFO | WARN | ERROR | FATAL,
	}
}

func (l *Logger) SetLogLevel(level string) {
	switch level {
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
	fmt.Printf("%v ",time.DateTime)
	if (l.level & level) == level {
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
