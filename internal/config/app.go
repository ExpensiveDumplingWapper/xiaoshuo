package config

import (
	"io/ioutil"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	LOGRUS     string = "logrus"
	ZAP        string = "zap"
	DevConfig  string = "internal/config/yaml/dev.yaml"
	ProdConfig string = "internal/config/yaml/prod.yaml"
)

var Setting *App
var ZConf *zconf.ZConf

type App struct {
	ZapConfig    LogConfig `yaml:"zap         -config"`
	LogrusConfig LogConfig `yaml:"logrus-config"`
	Log          LogConfig `yaml:"log-config"`
	Server       Server    `yaml:"server"`
	Zookeeper    Zookeeper `yaml:"zookeeper"`
}

type Server struct {
	RunMode      string        `yaml:"run-mode"`
	HttpPort     string        `yaml:"http-port"`
	IpdbPath     string        `yaml:"ipdb-path"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

type Zookeeper string

type LogConfig struct {
	Code           string `yaml:"code"`
	Level          string `yaml:"level"`
	EnableCaller   bool   `yaml:"enable-caller"`
	FileName       string `yaml:"file-name"`
	Format         string `yaml:"format"`
	EnableKafka    string `yaml:"enable-kafka"`
	KafkaBootstrap []string
	Topic          string
}

func ReadConfig(filename string) (*App, error) {
	var appConfig App
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read config error.")
	}
	err = yaml.Unmarshal(file, &appConfig)

	if err != nil {
		return nil, errors.Wrap(err, "parse config error.")
	}
	err = validateConfig(appConfig)
	if err != nil {
		return nil, errors.Wrap(err, "validate config error.")
	}
	return &appConfig, nil
}
