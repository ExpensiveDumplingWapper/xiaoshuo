### 下载
```bash
go get -u git.ourbluecity.com/overseas-server-tools/go-util-redis
```

### 使用方式
```go
// 代理模式 V1版
redisClient := redisclient.NewRedisClient(context.Background(), redisclient.ClientOptions{
  ZookeeperPath: []string{"10.9.158.210:2181", "10.9.114.167:2181"},
  Conf: map[string]redisclient.ClientConf{
    "live": {
      RedisPoolSize: 15,
      RedisMaxConnAge: time.Minute * 10, // redis连接最大存活时长，默认15分钟【传0取默认值】、-1 表示永久存活
      Host:          "/blued/backend/umem/live_oversea",
    },
  },
  QconfDuration:   time.Second * 3600,
  RefreshDuration: time.Second * 3600,
  IsPingCheck:     true,
})


// 代理模式 V2版「建议使用这版」
redisClient := redisclient.NewRedisClientV2(context.Background(), redisclient.ClientOptionsV2{
  ZookeeperPath: []string{"10.9.158.210:2181", "10.9.114.167:2181"},
  Conf: map[string]redisclient.ClientConf{
    "live": {
      RedisPoolSize: 15,
      RedisMaxConnAge: time.Minute * 10, // redis连接最大存活时长，默认15分钟【传0取默认值】、-1 表示永久存活
      Host:          "/blued/backend/umem/live_oversea",
    },
  },
  QconfDuration:   time.Second * 3600,
  RefreshDuration: time.Second * 3600,
  IsPingCheck:     true,
})

// cluster模式
redisClient := redisclient.NewClusterRedisClient(context.Background(), redisclient.ClusterClientOptions{
  ZookeeperPath: []string{"10.9.158.210:2181", "10.9.114.167:2181"},
  Conf: map[string]redisclient.ClusterClientConf{
    "live": {
      RedisPoolSize: 15,
      RedisMaxConnAge: time.Minute * 15, // redis连接最大存活时长，默认15分钟【传0取默认值】、-1 表示永久存活
      Host:          "/blued/backend/umem/live_oversea",
    },
  },
  QconfDuration:   time.Second * 3600,
  RefreshDuration: time.Second * 3600,
  IsPingCheck:     true,
})


// 使用
client, err := redisClient.GetClient("live")
if err != nil {
  return
}
data := client.Type(context.Background(), "test_live").Val()
fmt.Println(data)
```

>当使用***代理模式***时，建议使用***V2***版本「性能是V1版的8倍左右」

```shell
go test -test.run BenchmarkClient,BenchmarkClientV2 -bench=. -benchmem -benchtime=10s

BenchmarkRedisClient-8     	100000000	       219.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkRedisClientV2-8   	843702756	        27.80 ns/op	       0 B/op	       0 allocs/op
```