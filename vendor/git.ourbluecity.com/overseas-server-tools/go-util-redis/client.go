package redisclient

import (
	"context"
	"fmt"
	"sync"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
	"github.com/go-redis/redis/v8"
)

const (
	releaseTimeout   = time.Second * 2
	heartbeatTimeout = time.Second * 2
)

const (
	redisDialTimeout        = 6 * time.Second
	redisReadTimeout        = 4 * time.Second
	redisWriteTimeout       = 4 * time.Second
	redisPoolTimeout        = 6 * time.Second
	redisIdleTimeout        = 3 * time.Minute
	redisIdleCheckFrequency = 1 * time.Minute
	redisMaxRetries         = 5
	redisMaxConnAge         = 15 * time.Minute
)

func compairAddress(old []string, new []string) bool {
	if len(old) <= 0 {
		return true
	}
	if len(old) != len(new) {
		return true
	}

	// 顺序一般不会做改动，无需比较
	for i := range old {
		if old[i] != new[i] {
			return true
		}
	}

	return false
}

func inclued(t string, l []string) bool {
	for _, v := range l {
		if t == v {
			return true
		}
	}
	return false
}

type redisOption struct {
	addr          string
	poolSize      int
	minIdleConns  int
	maxMaxConnAge time.Duration
}

func getOptions(option redisOption) *redis.Options {
	o := &redis.Options{
		Addr:               option.addr,
		DialTimeout:        redisDialTimeout,
		ReadTimeout:        redisReadTimeout,
		WriteTimeout:       redisWriteTimeout,
		PoolSize:           option.poolSize,
		MinIdleConns:       option.minIdleConns,
		MaxRetries:         redisMaxRetries,
		PoolTimeout:        redisPoolTimeout,
		IdleTimeout:        redisIdleTimeout,
		IdleCheckFrequency: redisIdleCheckFrequency,
	}

	// 判断连接最大时长
	switch {
	case option.maxMaxConnAge == 0:
		o.MaxConnAge = redisMaxConnAge
	case option.maxMaxConnAge < 0:
		o.MaxConnAge = 0 * time.Second
	default:
		o.MaxConnAge = option.maxMaxConnAge
	}
	return o
}

// RedisClienter redis client
type RedisClienter interface {
	// GetClient 获取连接对象
	GetClient(pathKey string) (*redis.Client, error)
}

type redisConf struct {
	currentIndex      int
	addressList       []string
	host              string
	redisPoolSize     int
	redisMinIdleConns int
	redisMaxConnAge   time.Duration
}

type redisClient struct {
	redisClientConf sync.Map
	zconf           *zconf.ZConf
	cacheClient     sync.Map
	refreshDuration time.Duration
	isPingCheck     bool
}

// ClientConf redis conf
type ClientConf struct {
	Host              string
	RedisPoolSize     int
	RedisMinIdleConns int
	RedisMaxConnAge   time.Duration
}

// ClientOptions redis 连接选项
type ClientOptions struct {
	ZookeeperPath   []string
	Conf            map[string]ClientConf
	QconfDuration   time.Duration
	RefreshDuration time.Duration
	IsPingCheck     bool
}

func (r *redisClient) getAddress(pathKey string) (string, error) {
	basePath, ok := r.redisClientConf.Load(pathKey)
	if !ok {
		return "", fmt.Errorf("the key does not exist, key = %v", pathKey)
	}
	path, ok := basePath.(*redisConf)
	if !ok {
		return "", fmt.Errorf("value type error, the value type must be a redisConf object, key = %v", pathKey)
	}
	addressLength := len(path.addressList)
	if addressLength <= 0 {
		return "", fmt.Errorf("address list is empty, key = %v", pathKey)
	}
	path.currentIndex = (path.currentIndex + 1) % addressLength
	return path.addressList[path.currentIndex], nil
}

func (r *redisClient) heartbeat(c *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
	defer cancel()
	if _, err := c.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) release(address string, client *redis.Client, remove bool) {
	time.Sleep(releaseTimeout)
	if remove {
		r.cacheClient.Delete(address)
	}
	if err := client.Close(); err != nil {
		fmt.Printf("close expire client error, address = %v, error message is: %v\n", address, err)
	}
}

func (r *redisClient) connect(address string, poolSize int, minIdleConns int, maxConnAge time.Duration) {
	out, ok := r.getCacheClient(address)
	c := redis.NewClient(getOptions(redisOption{
		addr:          address,
		poolSize:      poolSize,
		minIdleConns:  minIdleConns,
		maxMaxConnAge: maxConnAge,
	}))

	// 开启ping检测 -> 判断ping状态
	if r.isPingCheck {
		if err := r.heartbeat(c); err != nil {
			fmt.Printf("connect heartbeat field, address = %v, error message is: %v", address, err)
		}
	}
	r.cacheClient.Store(address, c)

	// 没有过期的链接 -> 直接返回
	if !ok || out == nil {
		return
	}

	// 过期的链接 -> 释放掉
	go r.release(address, out, false)
}

func (r *redisClient) load(originKey, originValue interface{}) bool {
	key, ok := originKey.(string)
	if !ok {
		fmt.Printf("the type of redis-conf-key is not string, %v\n", key)
		return false
	}
	oldConf, ok := originValue.(*redisConf)
	if !ok || oldConf.host == "" {
		fmt.Printf("redis type of conf-value is not redisConf or host is empty, key = %v\n", key)
		return false
	}
	addressList, err := r.zconf.GetConfChildren(oldConf.host)
	if err != nil {
		fmt.Printf("get redis zconf error, key = %v, host = %v, error is %v\n", key, oldConf.host, err)
		return false
	}
	addressListLength := len(addressList)
	if addressListLength <= 0 {
		fmt.Printf("redis conf length is zero, key = %v, host = %v\n", key, oldConf.host)
		return false
	}
	conf := &redisConf{
		host:              oldConf.host,
		currentIndex:      oldConf.currentIndex,
		redisPoolSize:     oldConf.redisPoolSize,
		redisMinIdleConns: oldConf.redisMinIdleConns,
		redisMaxConnAge:   oldConf.redisMaxConnAge,
		addressList:       addressList,
	}
	if conf.currentIndex >= addressListLength {
		conf.currentIndex = 0
	}
	for _, address := range conf.addressList {
		cacheClient, ok := r.getCacheClient(address)

		// 链接不存在 -> 直接创建
		if !ok {
			r.connect(address, conf.redisPoolSize, conf.redisMinIdleConns, conf.redisMaxConnAge)
			continue
		}

		// 不需要ping检测
		if !r.isPingCheck {
			continue
		}

		// 心跳检测
		if err = r.heartbeat(cacheClient); err != nil {
			fmt.Printf("check heartbeat field, address = %v, error message is: %v", address, err)
		}
	}
	r.redisClientConf.Store(key, conf)
	return true
}

func (r *redisClient) getAddressList() []string {
	addressList := make([]string, 0, 50)
	r.redisClientConf.Range(func(key, value interface{}) bool {
		if conf, ok := value.(*redisConf); ok {
			for _, v := range conf.addressList {
				if inclued(v, addressList) {
					continue
				}
				addressList = append(addressList, v)
			}
		}
		return true
	})
	return addressList
}

func (r *redisClient) clean() {
	addressList := r.getAddressList()
	if len(addressList) <= 0 {
		return
	}

	// 遍历所有client
	r.cacheClient.Range(func(key, value interface{}) bool {
		address, ok := key.(string)
		if !ok {
			r.cacheClient.Delete(address)
			return true
		}
		db, ok := value.(*redis.Client)
		if !ok {
			r.cacheClient.Delete(address)
			return true
		}

		// 判断链接是否过期
		if inclued(address, addressList) {
			return true
		}

		// 过期的链接 -> 释放掉
		go r.release(address, db, true)
		return true
	})
}

func (r *redisClient) initClient(ctx context.Context) {
	idleTimeout := time.NewTimer(r.refreshDuration)
	defer idleTimeout.Stop()
	for {
		idleTimeout.Reset(r.refreshDuration)
		select {
		case <-idleTimeout.C:
			// 检测
			r.redisClientConf.Range(r.load)

			// 清理资源
			r.clean()
		case <-ctx.Done():
			return
		}
	}
}

func (r *redisClient) getCacheClient(address string) (*redis.Client, bool) {
	originalCacheClient, ok := r.cacheClient.Load(address)
	if !ok {
		return nil, false
	}
	cacheClient, ok := originalCacheClient.(*redis.Client)
	if !ok {
		return nil, false
	}
	return cacheClient, true
}

func (r *redisClient) GetClient(pathKey string) (*redis.Client, error) {
	address, err := r.getAddress(pathKey)
	if err != nil {
		return nil, err
	}
	c, ok := r.getCacheClient(address)
	if !ok {
		return nil, fmt.Errorf("the redis client does not exist, address = %v", address)
	}
	return c, nil
}

// NewRedisClient 初始化连接
func NewRedisClient(ctx context.Context, option ClientOptions) RedisClienter {
	r := &redisClient{
		redisClientConf: sync.Map{},
		zconf: &zconf.ZConf{
			Path:     option.ZookeeperPath,
			Duration: option.QconfDuration,
		},
		cacheClient:     sync.Map{},
		refreshDuration: option.RefreshDuration,
		isPingCheck:     option.IsPingCheck,
	}
	for k, v := range option.Conf {
		conf := &redisConf{
			currentIndex:      0,
			host:              v.Host,
			redisPoolSize:     v.RedisPoolSize,
			redisMinIdleConns: v.RedisMinIdleConns,
			redisMaxConnAge:   v.RedisMaxConnAge,
		}
		if conf.redisMinIdleConns <= 0 {
			conf.redisMinIdleConns = conf.redisPoolSize / 2
		}
		r.load(k, conf)
	}
	go r.initClient(ctx)
	return r
}
