package bootstrap

import (
	"xiaoshuo/internal/config"

	"github.com/pkg/errors"
)

func InitApp() error {
	var err error
	var file string

	env := config.GetENV()
	if env == config.Pro {
		file = config.ProdConfig
	} else {
		file = config.DevConfig
	}
	file = config.ProdConfig
	config.Setting, err = loadConfig(file)
	if err != nil {
		return errors.Wrap(err, "loadConfig")
	}

	LoadZConf()
	loadLogger()
	loadCache()
	loadDb()
	loadIPDb()

	return nil
}

func loadConfig(filename string) (*config.App, error) {
	ac, err := config.ReadConfig(filename)

	if err != nil {
		return nil, errors.Wrap(err, "loadConfig")
	}
	return ac, nil
}
