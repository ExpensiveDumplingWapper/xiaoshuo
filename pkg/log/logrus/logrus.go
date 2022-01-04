package logrus

import (
	"io/ioutil"

	"xiaoshuo/internal/config"
	"xiaoshuo/pkg/log/logger"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func RegisterLog(lc config.LogConfig) error {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)
	err := customizeLogFromConfig(log, lc)
	if err != nil {
		return errors.Wrap(err, "")
	}

	logger.SetLogger(log)
	return nil
}

func customizeLogFromConfig(log *logrus.Logger, lc config.LogConfig) error {
	log.SetReportCaller(lc.EnableCaller)

	// if lc.FileName != "" {
	// 	_ = os.MkdirAll(filepath.Dir(lc.FileName), 0666)
	// 	f, _ := os.OpenFile(lc.FileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	// 	log.SetOutput(f)
	// } else {
	// 	log.SetOutput(os.Stdout)
	// }

	log.SetOutput(ioutil.Discard)

	l := &log.Level
	err := l.UnmarshalText([]byte(lc.Level))
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.SetLevel(*l)

	var format logrus.Formatter
	if lc.Format == "json" {
		format = &logrus.JSONFormatter{
			DisableTimestamp: true,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@time",
				logrus.FieldKeyLevel: "@level",
				logrus.FieldKeyMsg:   "description",
			}}
	} else {
		format = &logrus.TextFormatter{}
	}
	log.SetFormatter(format)

	hook, err := NewKafkaLogrusHook(
		"klh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.DebugLevel},
		format,
		lc.KafkaBootstrap,
		lc.Topic,
		true,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.AddHook(hook)

	return nil
}
