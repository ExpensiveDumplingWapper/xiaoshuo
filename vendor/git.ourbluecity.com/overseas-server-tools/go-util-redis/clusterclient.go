package redisclient

import (
	"context"
	"fmt"
	"sync"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
	"github.com/go-redis/redis/v8"
)

// ClusterRedisClienter cluster redis client
type ClusterRedisClienter interface {
	// GetClient 获取连接对象
	GetClient(pathKey string) (*redis.ClusterClient, error)
}

type clusterRedis struct {
	addressList       []string
	host              string
	redisPoolSize     int
	redisMinIdleConns int
	redisMaxConnAge   time.Duration
	clusterRedis      *redis.ClusterClient
}

type clusterRedisClient struct {
	zconf           *zconf.ZConf
	cacheClient     sync.Map
	refreshDuration time.Duration
	isPingCheck     bool
}

// create redis client options
type (
	// ClusterClientConf cluster redis conf
	ClusterClientConf struct {
		Host              string
		RedisPoolSize     int
		RedisMinIdleConns int
		RedisMaxConnAge   time.Duration
	}

	// ClusterClientOptions cluster redis options
	ClusterClientOptions struct {
		ZookeeperPath   []string
		Conf            map[string]ClusterClientConf
		QconfDuration   time.Duration
		RefreshDuration time.Duration
		IsPingCheck     bool
	}
)

func (c *clusterRedisClient) getOptions(client *clusterRedis) *redis.ClusterOptions {
	o := &redis.ClusterOptions{
		Addrs:              client.addressList,
		DialTimeout:        redisDialTimeout,
		ReadTimeout:        redisReadTimeout,
		WriteTimeout:       redisWriteTimeout,
		PoolSize:           client.redisPoolSize,
		MinIdleConns:       client.redisMinIdleConns,
		PoolTimeout:        redisPoolTimeout,
		IdleTimeout:        redisIdleTimeout,
		IdleCheckFrequency: redisIdleCheckFrequency,
	}

	// 判断连接最大时长
	switch {
	case client.redisMaxConnAge == 0:
		o.MaxConnAge = redisMaxConnAge
	case client.redisMaxConnAge < 0:
		o.MaxConnAge = 0 * time.Second
	default:
		o.MaxConnAge = client.redisMaxConnAge
	}
	return o
}

func (c *clusterRedisClient) heartbeat(r *redis.ClusterClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
	defer cancel()
	if err := r.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (c *clusterRedisClient) release(key string, client *redis.ClusterClient) {
	time.Sleep(releaseTimeout)
	if err := client.Close(); err != nil {
		fmt.Printf("close expire client error, key = %v, error message is: %v\n", key, err)
	}
}

func (c *clusterRedisClient) connect(key string, client *clusterRedis) {
	out := client.clusterRedis
	r := redis.NewClusterClient(c.getOptions(client))

	// 开启ping检测 -> 判断ping状态
	if c.isPingCheck {
		if err := c.heartbeat(r); err != nil {
			fmt.Printf("connect heartbeat field, key = %v, error message is: %v", key, err)
		}
	}
	client.clusterRedis = r
	c.cacheClient.Store(key, client)

	// 没有过期的链接 -> 直接返回
	if out == nil {
		return
	}

	// 过期的链接 -> 释放掉
	go c.release(key, out)
}

func (c *clusterRedisClient) getCacheClient(key string) (*redis.ClusterClient, bool) {
	originalCacheClient, ok := c.cacheClient.Load(key)
	if !ok {
		return nil, false
	}
	cacheClient, ok := originalCacheClient.(*clusterRedis)
	if !ok {
		return nil, false
	}
	if cacheClient.clusterRedis == nil {
		return nil, false
	}
	return cacheClient.clusterRedis, true
}

func (c *clusterRedisClient) getCluster(client *clusterRedis, addressList []string) *clusterRedis {
	return &clusterRedis{
		host:              client.host,
		redisPoolSize:     client.redisPoolSize,
		redisMinIdleConns: client.redisMinIdleConns,
		redisMaxConnAge:   client.redisMaxConnAge,
		addressList:       addressList,
		clusterRedis:      client.clusterRedis,
	}
}

func (c *clusterRedisClient) load(originKey, originValue interface{}) bool {
	key, ok := originKey.(string)
	if !ok {
		fmt.Printf("the type of redis-conf-key is not string, %v\n", key)
		return false
	}
	client, ok := originValue.(*clusterRedis)
	if !ok || client.host == "" {
		fmt.Printf("redis type of conf-value is not redisConf or host is empty, key = %v\n", key)
		return false
	}
	addressList, err := c.zconf.GetConfChildren(client.host)
	if err != nil {
		fmt.Printf("get redis zconf error, key = %v, host = %v, error is %v\n", key, client.host, err)
		return false
	}
	if len(addressList) <= 0 {
		fmt.Printf("redis conf length is zero, key = %v, host = %v\n", key, client.host)
		return false
	}

	// 判断新旧地址是否满足新建连接要求
	isCreateNew := compairAddress(client.addressList, addressList)

	// 首次创建 or 地址发生变化 -> 创建cluster链接
	if client.clusterRedis == nil || isCreateNew {
		c.connect(key, c.getCluster(client, addressList))
		return true
	}

	// 不需要ping检测
	if !c.isPingCheck {
		return true
	}

	// 心跳检测
	if err = c.heartbeat(client.clusterRedis); err != nil {
		fmt.Printf("check heartbeat field, key = %v, error message is: %v", key, err)
	}
	return true
}

func (c *clusterRedisClient) initClient(ctx context.Context) {
	idleTimeout := time.NewTimer(c.refreshDuration)
	defer idleTimeout.Stop()
	for {
		idleTimeout.Reset(c.refreshDuration)
		select {
		case <-idleTimeout.C:
			c.cacheClient.Range(c.load)
		case <-ctx.Done():
			return
		}
	}
}

func (c *clusterRedisClient) GetClient(pathKey string) (*redis.ClusterClient, error) {
	r, ok := c.getCacheClient(pathKey)
	if !ok {
		return nil, fmt.Errorf("the redis client does not exist, key = %v", pathKey)
	}
	return r, nil
}

// NewClusterRedisClient 初始化连接
func NewClusterRedisClient(ctx context.Context, option ClusterClientOptions) ClusterRedisClienter {
	c := &clusterRedisClient{
		zconf: &zconf.ZConf{
			Path:     option.ZookeeperPath,
			Duration: option.QconfDuration,
		},
		cacheClient:     sync.Map{},
		refreshDuration: option.RefreshDuration,
		isPingCheck:     option.IsPingCheck,
	}
	for k, v := range option.Conf {
		if v.Host == "" {
			panic("redis option host is empty")
		}
		client := &clusterRedis{
			host:              v.Host,
			redisPoolSize:     v.RedisPoolSize,
			redisMinIdleConns: v.RedisMinIdleConns,
			redisMaxConnAge:   v.RedisMaxConnAge,
		}
		if client.redisMinIdleConns <= 0 {
			client.redisMinIdleConns = client.redisPoolSize / 2
		}
		c.load(k, client)
	}
	go c.initClient(ctx)
	return c
}
