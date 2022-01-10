package httplient

import (
	"context"
	"io/ioutil"
	"net/http"
)

type origHTTPClient struct {
	client *http.Client
	host   string
}

func (h *origHTTPClient) doRes(req *http.Request) (*Response, error) {
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

func (h *origHTTPClient) Get(ctx context.Context, query QueryReq) (*Response, error) {
	req, err := newQuery(ctx, h.host, query)
	if err != nil {
		return nil, err
	}
	return h.doRes(req)
}

func (h *origHTTPClient) PostJSON(ctx context.Context, PostJSON PostJSONReq) (*Response, error) {
	req, err := newJSONPost(ctx, h.host, PostJSON)
	if err != nil {
		return nil, err
	}
	return h.doRes(req)
}
