package log

import "xiaoshuo/internal/config"

type logFbInterface interface {
	Build(*config.LogConfig) error
}

var logfactoryBuilderMap = map[string]logFbInterface{
	//config.ZAP:    &ZapFactory{},
	config.LOGRUS: &LogrusFactory{},
}

func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
