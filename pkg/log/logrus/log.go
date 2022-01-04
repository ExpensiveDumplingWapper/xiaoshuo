package logrus

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var BlueLogger *logrus.Logger

type Option func(*option)

func NewLogger() {
	BlueLogger = logrus.New()
	BlueLogger.SetFormatter(&logrus.JSONFormatter{})
	BlueLogger.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("/Users/jason/go/src/adms-go/stand_test1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	if err != nil {
		panic(err.Error())
	}

	writers := []io.Writer{file, os.Stdout}

	BlueLogger.SetOutput(io.MultiWriter(writers...))

	hook, err := NewKafkaLogrusHook(
		"klh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.DebugLevel},
		&logrus.JSONFormatter{},
		[]string{"192.168.1.103:9092"},
		"test",
		true,
	)

	if err != nil {
		log.Fatal(err)
	}

	BlueLogger.AddHook(hook)
}

//func NewLogger1(opts ...Option) {
//	opt := new(option)
//	for _, f := range opts{
//		f(opt)
//	}
//
//	Logger = &logger{
//		enableKafka: true,
//	}
//}

//var log logger

func withFormtter(formatter *logrus.Formatter) Option {
	return func(o *option) {
		o.formatter = formatter
	}
}

func withLevel(level logrus.Level) Option {
	return func(o *option) {
		o.level = uint32(level)
	}
}

func withFile(filePath string) Option {
	return func(o *option) {
		o.filePath = filePath
	}
}

//func withEnableKafka(enable bool) Option{
//	return func(o *option) {
//		o.
//	}
//}

type option struct {
	level     uint32
	formatter *logrus.Formatter
	filePath  string
	service   string
	log_name  string
}

func Info(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "service-test",
		"level":   "i",
	}).Info(msg)

	//Logger.WithFields(logrus.Fields{
	//		"service": "service-test",
	//	}).Info(msg)
}

func Debug(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "service-test",
	}).Debug(msg)
}

func Warn(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "service-test",
	}).Warn(msg)
}

func Error(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "service-test",
	}).Error(msg)
}

func Fatal(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "service-test",
	}).Fatal(msg)
}

func Panic(msg interface{}) {
	BlueLogger.WithFields(logrus.Fields{
		"service": "service-test",
	}).Panic(msg)
}
