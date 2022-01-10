package redisclient

import (
	"context"
	"fmt"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
	"github.com/go-redis/redis/v8"
)

type redisChain struct {
	address string
	redis   *redis.Client
	next    *redisChain
}

func (r *redisChain) heartbeat() error {
	ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
	defer cancel()
	if _, err := r.redis.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

type redisConfV2 struct {
	host              string
	redisPoolSize     int
	redisMinIdleConns int
	redisMaxConnAge   time.Duration
}

type redisStore struct {
	conf   redisConfV2
	curent *redisChain
	chain  *redisChain
}

func (s *redisStore) queryChain(address string) (*redisChain, bool) {
	if s.chain == nil {
		return nil, false
	}
	if s.chain.address == address {
		return s.chain, true
	}
	chain := s.chain.next
	for chain != s.chain {
		if chain.address == address {
			return chain, true
		}
		chain = chain.next
	}
	return nil, false
}

func (s *redisStore) queryLastChain() (*redisChain, bool) {
	if s.chain == nil {
		return nil, false
	}
	if s.chain.next == s.chain {
		return s.chain, true
	}
	chain := s.chain.next
	for {
		if chain.next == s.chain {
			break
		}
		chain = chain.next
	}
	return chain, true
}

func (s *redisStore) insert(redis *redis.Client, address string) {
	newchain := &redisChain{
		address: address,
		redis:   redis,
	}
	if s.chain == nil {
		newchain.next = newchain
		s.chain = newchain
		s.curent = s.chain
		return
	}
	newchain.next = s.chain.next
	s.chain.next = newchain
}

func (s *redisStore) release(chain *redisChain) {
	chain.next = nil
	time.Sleep(releaseTimeout)
	if err := chain.redis.Close(); err != nil {
		fmt.Printf("close expire client error, address = %v, error message is: %v\n", chain.address, err)
	}
}

func (s *redisStore) clean(addressList []string) {
	if s.chain == nil || s.chain.next == s.chain {
		return
	}
	chain := s.chain
	for {
		if chain.next == s.chain {
			break
		}
		if !inclued(chain.next.address, addressList) {
			out := chain.next
			chain.next = out.next
			s.curent = s.chain
			go s.release(out)
			continue
		}

		chain = chain.next
	}

	if s.chain.next == s.chain {
		return
	}

	if inclued(s.chain.address, addressList) {
		return
	}

	lastChain, ok := s.queryLastChain()
	if !ok {
		return
	}
	if lastChain == s.chain {
		return
	}

	out := s.chain
	s.chain = s.chain.next
	lastChain.next = s.chain
	s.curent = s.chain
	go s.release(out)
}

func (s *redisStore) queryAndNext() (*redis.Client, bool) {
	if s.curent == nil {
		return nil, false
	}
	client := s.curent.redis
	s.curent = s.curent.next
	return client, true
}

type redisClientV2 struct {
	store           map[string]*redisStore
	zconf           *zconf.ZConf
	refreshDuration time.Duration
	isPingCheck     bool
}

// ClientOptionsV2 redis 连接选项
type ClientOptionsV2 struct {
	ZookeeperPath   []string
	Conf            map[string]ClientConf
	QconfDuration   time.Duration
	RefreshDuration time.Duration
	IsPingCheck     bool
}

func (r *redisClientV2) heartbeat(c *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
	defer cancel()
	if err := c.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClientV2) connect(address string, store *redisStore) {
	// 判断连接是否已经存在
	if s, ok := store.queryChain(address); ok {
		if !r.isPingCheck {
			return
		}
		if err := s.heartbeat(); err != nil {
			fmt.Printf("check heartbeat field, address = %v, error message is: %v", address, err)
		}
		return
	}

	// 创建 -> 判断ping状态
	c := redis.NewClient(getOptions(redisOption{
		addr:          address,
		poolSize:      store.conf.redisPoolSize,
		minIdleConns:  store.conf.redisMinIdleConns,
		maxMaxConnAge: store.conf.redisMaxConnAge,
	}))
	if err := r.heartbeat(c); err != nil {
		fmt.Printf("connect heartbeat field, address = %v, error message is: %v", address, err)
	}
	store.insert(c, address)
}

func (r *redisClientV2) load(key string, store *redisStore, clean bool) {
	addressList, err := r.zconf.GetConfChildren(store.conf.host)
	if err != nil {
		fmt.Printf("get redis zconf error, key = %v, host = %v, error is %v\n", key, store.conf.host, err)
		return
	}
	if len(addressList) <= 0 {
		fmt.Printf("redis conf length is zero, key = %v, host = %v\n", key, store.conf.host)
		return
	}
	for _, address := range addressList {
		r.connect(address, store)
	}

	if !clean {
		return
	}
	store.clean(addressList)
}

func (r *redisClientV2) initClient(ctx context.Context) {
	idleTimeout := time.NewTimer(r.refreshDuration)
	defer idleTimeout.Stop()
	for {
		idleTimeout.Reset(r.refreshDuration)
		select {
		case <-idleTimeout.C:
			// 检测
			for key, store := range r.store {
				r.load(key, store, true)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (r *redisClientV2) GetClient(pathKey string) (*redis.Client, error) {
	store, ok := r.store[pathKey]
	if !ok {
		return nil, fmt.Errorf("the redis client does not exist, key = %v", pathKey)
	}
	client, ok := store.queryAndNext()
	if !ok {
		return nil, fmt.Errorf("the redis client is uninitialized, key = %v", pathKey)
	}
	return client, nil
}

// NewRedisClientV2 初始化连接
func NewRedisClientV2(ctx context.Context, option ClientOptionsV2) RedisClienter {
	r := &redisClientV2{
		store: make(map[string]*redisStore, len(option.Conf)),
		zconf: &zconf.ZConf{
			Path:     option.ZookeeperPath,
			Duration: option.QconfDuration,
		},
		refreshDuration: option.RefreshDuration,
		isPingCheck:     option.IsPingCheck,
	}

	for k, v := range option.Conf {
		r.store[k] = &redisStore{
			conf: redisConfV2{
				host:              v.Host,
				redisPoolSize:     v.RedisPoolSize,
				redisMinIdleConns: v.RedisMinIdleConns,
				redisMaxConnAge:   v.RedisMaxConnAge,
			},
		}
		if r.store[k].conf.redisMinIdleConns <= 0 {
			r.store[k].conf.redisMinIdleConns = r.store[k].conf.redisPoolSize / 2
		}
		r.load(k, r.store[k], false)
	}

	go r.initClient(ctx)
	return r
}
