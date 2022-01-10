package httplient

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	zconf "git.ourbluecity.com/overseas-server-tools/go-util-zconf"
)

type chain struct {
	address string
	next    *chain
}

type store struct {
	host   string
	curent *chain
	chain  *chain
}

func (s *store) queryChain(address string) (*chain, bool) {
	if s.chain == nil {
		return nil, false
	}
	if s.chain.address == address {
		return s.chain, true
	}
	c := s.chain.next
	for c != s.chain {
		if c.address == address {
			return c, true
		}
		c = c.next
	}
	return nil, false
}

func (s *store) queryLastChain() (*chain, bool) {
	if s.chain == nil {
		return nil, false
	}
	if s.chain.next == s.chain {
		return s.chain, true
	}
	c := s.chain.next
	for {
		if c.next == s.chain {
			break
		}
		c = c.next
	}
	return c, true
}

func (s *store) insert(address string) {
	newchain := &chain{
		address: address,
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

func (s *store) release(c *chain) {
	c.next = nil
}

func (s *store) clean(addressList []string) {
	if s.chain == nil || s.chain.next == s.chain {
		return
	}
	c := s.chain
	for {
		if c.next == s.chain {
			break
		}
		if !inclued(c.next.address, addressList) {
			out := c.next
			c.next = out.next
			s.curent = s.chain
			go s.release(out)
			continue
		}

		c = c.next
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

func (s *store) queryAndNext() (string, bool) {
	if s.curent == nil {
		return "", false
	}
	address := s.curent.address
	s.curent = s.curent.next
	return address, true
}

type httpClientV2 struct {
	store           map[string]*store
	zconf           *zconf.ZConf
	refreshDuration time.Duration
	client          *http.Client
}

func (h *httpClientV2) load(key string, s *store, clean bool) {
	addressList, err := h.zconf.GetConfChildren(s.host)
	if err != nil {
		fmt.Printf("get http zconf error, key = %v, host = %v, error is %v\n", key, s.host, err)
		return
	}
	if len(addressList) <= 0 {
		fmt.Printf("http conf length is zero, key = %v, host = %v\n", key, s.host)
		return
	}

	for _, address := range addressList {
		if _, ok := s.queryChain(address); ok {
			continue
		}
		s.insert(address)
	}
	if !clean {
		return
	}
	s.clean(addressList)
}

func (h *httpClientV2) initClient(ctx context.Context) {
	idleTimeout := time.NewTimer(h.refreshDuration)
	defer idleTimeout.Stop()
	for {
		idleTimeout.Reset(h.refreshDuration)
		select {
		case <-idleTimeout.C:
			for key, s := range h.store {
				h.load(key, s, true)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (h *httpClientV2) doRes(req *http.Request) (*Response, error) {
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

func (h *httpClientV2) Get(ctx context.Context, confKey string, query QueryReq) (*Response, error) {
	s, ok := h.store[confKey]
	if !ok {
		return nil, fmt.Errorf("the http client does not exist, key = %v", confKey)
	}
	host, ok := s.queryAndNext()
	if !ok {
		return nil, fmt.Errorf("the http client is uninitialized, key = %v", confKey)
	}
	req, err := newQuery(ctx, host, query)
	if err != nil {
		return nil, err
	}
	return h.doRes(req)
}

func (h *httpClientV2) PostJSON(ctx context.Context, confKey string, PostJSON PostJSONReq) (*Response, error) {
	s, ok := h.store[confKey]
	if !ok {
		return nil, fmt.Errorf("the http client does not exist, key = %v", confKey)
	}
	host, ok := s.queryAndNext()
	if !ok {
		return nil, fmt.Errorf("the http client is uninitialized, key = %v", confKey)
	}
	req, err := newJSONPost(ctx, host, PostJSON)
	if err != nil {
		return nil, err
	}
	return h.doRes(req)
}
