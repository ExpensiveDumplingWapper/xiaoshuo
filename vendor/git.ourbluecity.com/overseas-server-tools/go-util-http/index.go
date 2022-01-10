package httplient

import (
	"context"
	"crypto/tls"
	"net/http"
	"sync"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
)

func inclued(t string, l []string) bool {
	for _, v := range l {
		if t == v {
			return true
		}
	}
	return false
}

// HTTPOptions http client连接选项
type HTTPOptions struct {
	ZookeeperPath   []string
	Conf            map[string]string
	QconfDuration   time.Duration
	RefreshDuration time.Duration
	Client          *http.Client
}

// Response response 返回对象
type Response struct {
	Status     string
	StatusCode int
	ResBody    []byte
}

// HTTPClienter http client object
type HTTPClienter interface {
	// Get get查询
	Get(ctx context.Context, confKey string, query QueryReq) (*Response, error)

	// PostJSON post json查询
	PostJSON(ctx context.Context, confKey string, PostJSON PostJSONReq) (*Response, error)
}

// OrigHTTPClienter http client object
type OrigHTTPClienter interface {
	// Get get查询
	Get(ctx context.Context, query QueryReq) (*Response, error)

	// PostJSON post json查询
	PostJSON(ctx context.Context, PostJSON PostJSONReq) (*Response, error)
}

// NewHTTPClient 初始化http client对象 基于zconf
func NewHTTPClient(ctx context.Context, option *HTTPOptions) HTTPClienter {
	client := option.Client
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: 30 * time.Second,
		}
	}

	h := &httpClient{
		refreshDuration: option.RefreshDuration,
		clientConf:      sync.Map{},
		zconf: &zconf.ZConf{
			Path:     option.ZookeeperPath,
			Duration: option.QconfDuration,
		},
		client: client,
	}
	for k, v := range option.Conf {
		h.clientConf.Store(k, &httpConf{
			currentIndex: 0,
			host:         v,
		})
	}
	h.clientConf.Range(h.load)
	go h.initClient(ctx)
	return h
}

// NewHTTPClientV2 初始化http client对象 基于zconf
func NewHTTPClientV2(ctx context.Context, option *HTTPOptions) HTTPClienter {
	client := option.Client
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: 30 * time.Second,
		}
	}

	h := &httpClientV2{
		refreshDuration: option.RefreshDuration,
		zconf: &zconf.ZConf{
			Path:     option.ZookeeperPath,
			Duration: option.QconfDuration,
		},
		client: client,
		store:  make(map[string]*store, len(option.Conf)),
	}
	for k, v := range option.Conf {
		h.store[k] = &store{
			host: v,
		}
		h.load(k, h.store[k], false)
	}
	go h.initClient(ctx)
	return h
}

// NewOrigHTTPClient 初始化http client对象 不基于zconf
func NewOrigHTTPClient(host string, client *http.Client) OrigHTTPClienter {
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: 30 * time.Second,
		}
	}
	return &origHTTPClient{
		client: client,
		host:   host,
	}
}
