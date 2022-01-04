package bootstrap

import (
	"context"
	"strings"

	"xiaoshuo/internal/config"
	"xiaoshuo/internal/models"

	redisclient "git.ourbluecity.com/overseas-server-tools/go-util-redis"

	"time"
)

func loadCache() {
	models.RedisCli = redisclient.NewRedisClient(context.Background(), redisclient.ClientOptions{
		ZookeeperPath: strings.Split(config.GetZK(), ","),
		Conf: map[string]redisclient.ClientConf{
			"adms": {
				RedisPoolSize: 15,
				Host:          "/blued/backend/umem/adms",
			},
			"adms_read": {
				RedisPoolSize: 15,
				Host:          "/blued/backend/umem/adms_read",
			},
		},
		QconfDuration:   time.Second * 4,
		RefreshDuration: time.Second * 2,
		IsPingCheck:     true,
	})
}
