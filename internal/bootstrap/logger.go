package bootstrap

import (
	"log"
	"xiaoshuo/internal/config"
	newlog "xiaoshuo/pkg/log"

	"github.com/pkg/errors"
)

func loadLogger() error {
	var err error
	setting := config.Setting.Log
	setting.KafkaBootstrap, err = config.ZConf.GetConfChildren("/blued/bigdata/das/backendlog/overseas/bootstraps")
	if err != nil {
		log.Fatal(err.Error())
	}

	setting.Topic, err = config.ZConf.GetConf("/blued/bigdata/das/backendlog/overseas/topic")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = newlog.GetLogFactoryBuilder(setting.Code).Build(&setting)
	if err != nil {
		log.Fatal(err.Error())

		return errors.Wrap(err, "")
	}
	return nil
}
