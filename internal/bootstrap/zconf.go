package bootstrap

import (
	"strings"
	"time"

	"xiaoshuo/internal/config"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
)

func LoadZConf() {
	config.ZConf = &zconf.ZConf{
		Path:     strings.Split(config.GetZK(), ","),
		Duration: time.Second * 8,
	}
}
