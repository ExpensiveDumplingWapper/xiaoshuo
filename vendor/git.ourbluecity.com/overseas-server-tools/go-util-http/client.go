package httplient

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
)

type httpConf struct {
	currentIndex int
	addressList  []string
	host         string
}

type httpClient struct {
	clientConf      sync.Map
	zconf           *zconf.ZConf
	cacheHost       sync.Map
	refreshDuration time.Duration
	client          *http.Client
}

func (h *httpClient) load(originKey, originValue interface{}) bool {
	key, ok := originKey.(string)
	if !ok {
		fmt.Printf("the type of http-conf-key is not string, %v\n", key)
		return false
	}
	oldPath, ok := originValue.(*httpConf)
	if !ok || oldPath.host == "" {
		fmt.Printf("http type of conf-value is not httpConf or host is empty, key = %v\n", key)
		return false
	}
	conf, err := h.zconf.GetConfChildren(oldPath.host)
	if err != nil {
		fmt.Printf("get http zconf error, key = %v, host = %v, error is %v\n", key, oldPath.host, err)
		return false
	}
	confLength := len(conf)
	if confLength <= 0 {
		fmt.Printf("http conf length is zero, key = %v, host = %v\n", key, oldPath.host)
		return false
	}
	path := &httpConf{
		host:         oldPath.host,
		currentIndex: oldPath.currentIndex,
		addressList:  conf,
	}
	if path.currentIndex >= confLength {
		path.currentIndex = 0
	}
	h.clientConf.Store(key, path)
	return true
}

func (h *httpClient) initClient(ctx context.Context) {
	idleTimeout := time.NewTimer(h.refreshDuration)
	defer idleTimeout.Stop()
	for {
		idleTimeout.Reset(h.refreshDuration)
		select {
		case <-idleTimeout.C:
			h.clientConf.Range(h.load)
		case <-ctx.Done():
			return
		}
	}
}

func (h *httpClient) getHost(pathKey string) (string, error) {
	basePath, ok := h.clientConf.Load(pathKey)
	if !ok {
		return "", fmt.Errorf("the key does not exist, key = %v", pathKey)
	}
	path, ok := basePath.(*httpConf)
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

func (h *httpClient) doRes(req *http.Request) (*Response, error) {
	res, err := h.client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		ResBody:    resBody,
	}, nil
}

func (h *httpClient) Get(ctx context.Context, confKey string, query QueryReq) (*Response, error) {
	host, err := h.getHost(confKey)
	if err != nil {
		return nil, err
	}
	req, err := newQuery(ctx, host, query)
	if err != nil {
		return nil, err
	}
	return h.doRes(req)
}

func (h *httpClient) PostJSON(ctx context.Context, confKey string, PostJSON PostJSONReq) (*Response, error) {
	host, err := h.getHost(confKey)
	if err != nil {
		return nil, err
	}
	req, err := newJSONPost(ctx, host, PostJSON)
	if err != nil {
		return nil, err
	}
	return h.doRes(req)
}
