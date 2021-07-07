package common

import (
	"fmt"
	"log"
)

type Level int

const (
	Debug Level = 1
	Info  Level = 2
	Warn  Level = 3
	Error Level = 4
	Fatal Level = 5
)

func getLevleString(lvl Level) string {
	switch lvl {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO "
	case Warn:
		return "WARN "
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		panic("Unknown log level")
	}
}

type Logger struct {
	Name   string
	MinLvl Level
}

func CreateLogger(name string) *Logger {
	return &Logger{
		Name:   name,
		MinLvl: Info,
	}
}

func (l *Logger) logWithLevel(lvl Level, format string, v ...interface{}) {
	m := fmt.Sprintf(format, v...)
	m = fmt.Sprintf("[%s] %s", getLevleString(lvl), m)
	log.Print(m)
}

func (l *Logger) D(format string, v ...interface{}) {
	if Debug < l.MinLvl {
		return
	}

	l.logWithLevel(Debug, format, v...)
}

func (l *Logger) I(format string, v ...interface{}) {
	if Info < l.MinLvl {
		return
	}

	l.logWithLevel(Info, format, v...)
}

func (l *Logger) W(format string, v ...interface{}) {
	if Warn < l.MinLvl {
		return
	}

	l.logWithLevel(Warn, format, v...)
}

func (l *Logger) E(format string, v ...interface{}) {
	if Error < l.MinLvl {
		return
	}

	l.logWithLevel(Error, format, v...)
}

func (l *Logger) F(format string, v ...interface{}) {
	if Fatal < l.MinLvl {
		return
	}

	l.logWithLevel(Fatal, format, v...)
}

var rootLogger = CreateLogger("root-logger")

func LogD(format string, v ...interface{}) {
	rootLogger.D(format, v...)
}

func LogI(format string, v ...interface{}) {
	rootLogger.I(format, v...)
}

func LogW(format string, v ...interface{}) {
	rootLogger.W(format, v...)
}

func LogE(format string, v ...interface{}) {
	rootLogger.E(format, v...)
}

func LogF(format string, v ...interface{}) {
	rootLogger.F(format, v...)
}
