### 下载
```bash
go get -u git.ourbluecity.com/overseas-server-tools/go-util-http
```

### 使用方式
```go
// V1版本「HTTPOptions.Client可以设置为nil」
client := httpClient.NewHTTPClient(context.Background(), &httpClient.HTTPOptions{
    ZookeeperPath: []string{"10.9.158.210:2181", "10.9.114.167:2181"},
    Conf: map[string]string{
    "tools": "/blued/service/live/oversea_tools_host",
  },
  QconfDuration:   time.Second * 4,
  RefreshDuration: time.Second * 2,
  Client: &http.Client{
    Transport: &http.Transport{
      TLSClientConfig: &tls.Config{
        InsecureSkipVerify: true,
      },
      },
    Timeout: 30 * time.Second,
  },
})


// V2版「HTTPOptions.Client可以设置为nil」
client := httpClient.NewHTTPClientV2(context.Background(), &httpClient.HTTPOptions{
    ZookeeperPath: []string{"10.9.158.210:2181", "10.9.114.167:2181"},
    Conf: map[string]string{
    "tools": "/blued/service/live/oversea_tools_host",
  },
  QconfDuration:   time.Second * 4,
  RefreshDuration: time.Second * 2,
  Client: &http.Client{
    Transport: &http.Transport{
      TLSClientConfig: &tls.Config{
        InsecureSkipVerify: true,
      },
    },
    Timeout: 30 * time.Second,
  },
})

// 基本版「client 可以自定义」
client := httpClient.NewOrigHTTPClient("http://127.0.0.1:8801", nil)
```

#### 使用
```go
_, err = client.PostJSON(context.Background(), "tools", httpClient.PostJSONReq{
  URL:  "/test",
  Body: bodyData,
  HeaderParam: map[string]string{
    "x-iris-uid": "12",
  },
})
if err != nil {
  b.Error(err)
  return
}
```

>两个版本效率差不多，使用哪个都行

```shell
go test -test.run BenchmarkClient,BenchmarkClientV2 -bench=. -benchmem -benchtime=10s

BenchmarkClient-8                 	    1024	  18789704 ns/op	    5788 B/op	      69 allocs/op
BenchmarkClientV2-8               	    1107	  14611280 ns/op	    5771 B/op	      69 allocs/op
```