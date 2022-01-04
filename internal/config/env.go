package config

import (
	"os"
)

// Env 环境变量
type Env string

// 环境变量：Pro->正是环境，Dev->测试环境
const (
	Pro  Env = "Pro"
	Dev  Env = "Dev"
	Test Env = "Test"
)

// SGetENV 获取环境变量
func GetENV() Env {
	env := os.Getenv("RUN_ENV")
	switch env {
	case "pro":
		return Pro
	case "dev":
		return Dev
	default:
		return Test
	}
}

func GetZK() string {
	zk := os.Getenv("ZOOKEEPER_PATH")
	if zk == "" {
		zk = string(Setting.Zookeeper)
	}

	return zk
}
