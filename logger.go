package eyes

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var KiteLogger LoggerInterface

type LoggerInterface interface {
	SetConfig(int, bool, string, ...ConfigOption)
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

type ParamSetter func(logger *Logger)

func (l *Logger)SetConfig(level int, isFile bool, timeZone string, opts ...ConfigOption) {
	var err error
	if l.TimeZone, err = time.LoadLocation(timeZone); err != nil {
		panic(err)
	}
	for _, opt := range opts {
		opt(l)
	}
	l.Level = level
	if !isFile {
		l.OutConsole = os.Stdout
	} else {
		l.IsFile = true
		l.MakeDir()
		l.OpenFile()
	}
}

type ConfigOption func(loggerType *Logger)

func (l *Logger)Debug(messages ...interface{}) {
	l.OrganizeText(DEBUG, "DEBUG", messages...)
}

func (l *Logger)Info(messages ...interface{}) {
	l.OrganizeText(INFO, "INFO", messages...)
}

func (l *Logger)Warn(messages ...interface{}) {
	l.OrganizeText(WARN, "WARN", messages...)
}

func (l *Logger)Error(messages ...interface{}) {
	l.OrganizeText(ERROR, "ERROR", messages...)
}

func (l *Logger)Fatal(messages ...interface{}) {
	l.OrganizeText(FATAL, "FATAL", messages...)
}

func (l *Logger)OrganizeText(level int, leverStr string, messages ...interface{}) {
	if level >= l.Level {
		text :=  l.GetNowTime() + " " + GetCaller(3) + strings.TrimRight(fmt.Sprint(leverStr, " ", "IP=",
			ServerIP, " " + Format(messages...)), " ") + "\n"
		if l.IsFile {
			l.Write(text)
		}
	}
}

func GetLogger() LoggerInterface {
	if KiteLogger == nil {
		timeZone, _ := time.LoadLocation("Asia/Shanghai")
		KiteLogger = &Logger{Level: INFO, IsFile: true, TimeZone: timeZone}
	}
	return KiteLogger
}
