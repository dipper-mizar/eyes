package eyes

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func WithFile(filePrefix string, path string) ConfigOption {
	return func(l *Logger) {
		l.FilePrefix = filePrefix
		l.Path = path
		l.AbsoluteFileName = l.GetAbsoluteFileName()
	}
}

func (l *Logger) FormatFileName() string {
	return l.FilePrefix + "_" + l.GetNowTimeWithoutHMS() + "." + DefaultFileSuffix
}

func (l *Logger) GetAbsoluteFileName() string {
	return strings.TrimRight(l.Path, "/") + "/" + l.FormatFileName()
}

// Support multilevel directory creating like project/name/log.
func (l *Logger) MakeDir() {
	_, err := os.Stat(l.Path)
	if !os.IsExist(err) {
		err := os.MkdirAll(l.Path, 0766)
		if os.IsPermission(err) {
			panic(err)
		}
	}
}

func (l *Logger) OpenFile() {
	var err error
	l.LogFile, err = os.OpenFile(l.AbsoluteFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
	if err != nil {
		panic("Can not create or open the file. Please check permissions. " + err.Error())
	}
	l.FileWriter = l.LogFile
}

func (l *Logger) Write(messages string) {
	// TODO: Goroutine should be here.
	_, err := l.FileWriter.Write([]byte(messages))
	if err != nil {
		panic(err)
	}
}

func (l *Logger) Close() {
	panic(l.LogFile.Close())
}

func GetCaller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		panic("Caller depth for log file prefix error.")
	} else {
		str := path.Base(file) + ":" + strconv.Itoa(line) + ":"
		return str
	}
}
