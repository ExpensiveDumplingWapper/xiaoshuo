package logrus

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var BlueLogger *logrus.Logger
var LeavMessage *logrus.Logger
var AskBook *logrus.Logger

type Option func(*option)

func NewLogger() {
	BlueLogger = logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	BlueLogger.SetFormatter(customFormatter)
	BlueLogger.SetFormatter(&logrus.JSONFormatter{})
	BlueLogger.SetLevel(logrus.InfoLevel)
	logName := "/tmp/logs/" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file, os.Stdout}
	BlueLogger.SetOutput(io.MultiWriter(writers...))
	// BlueLogger.SetOutput(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func NewLeavMessage() {
	LeavMessage = logrus.New()
	LeavMessage.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	LeavMessage.SetFormatter(&logrus.JSONFormatter{})
	LeavMessage.SetLevel(logrus.InfoLevel)
	logName := "/tmp/logs/leavMessage.DB"
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file, os.Stdout}
	LeavMessage.SetOutput(io.MultiWriter(writers...))
	// BlueLogger.SetOutput(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func NewAskBook() {
	AskBook = logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	AskBook.SetFormatter(customFormatter)
	AskBook.SetFormatter(&logrus.JSONFormatter{})
	AskBook.SetLevel(logrus.InfoLevel)
	logName := "/tmp/logs/askBook.DB"
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file, os.Stdout}
	AskBook.SetOutput(io.MultiWriter(writers...))
	if err != nil {
		log.Fatal(err)
	}
}

func LeavMessDB(email, message, ip string) {
	LeavMessage.WithFields(logrus.Fields{
		"email": email,
		"ip":    ip,
	}).Info(message)
}
func AskBookDB(email, message, ip string) {
	AskBook.WithFields(logrus.Fields{
		"email": email,
		"ip":    ip,
	}).Info(message)
}

type option struct {
	level     uint32
	formatter *logrus.Formatter
	filePath  string
	service   string
	log_name  string
}

func Info(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

}

func Debug(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "xiaoshuo",
	}).Debug(msg)
}

func Error(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "xiaoshuo",
	}).Error(msg)
}

func Fatal(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "xiaoshuo",
	}).Fatal(msg)
}

func Panic(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "xiaoshuo",
	}).Panic(msg)
}
