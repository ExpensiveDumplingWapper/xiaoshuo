package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	WithFields(fields logrus.Fields)
}

func SetLogger(newLogger *logrus.Logger) {
	Log = newLogger
}

func Info(description string) {
	requestEntry("i").Info(description)
}

func Error(description string) {
	requestEntry("e").Error(description)
}

func Debug(description string) {
	requestEntry("d").Debug(description)
}

func Warn(description string) {
	requestEntry("w").Warn(description)
}

func requestEntry(level string) *logrus.Entry {
	return Log.WithFields(logrus.Fields{
		"time":           time.Now().UnixNano() / 1e6,
		"service":        "adms-go",
		"log_name":       "adms-go",
		"level":          level,
		"uid":            "",
		"client_ip":      "",
		"server_ip":      "",
		"request_id":     "",
		"request_url":    "",
		"request_type":   "",
		"request_header": "",
	})
}
