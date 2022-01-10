package httplient

import (
	"bytes"
	"context"
	"net/http"
	"reflect"
	"strings"
	"unsafe"
)

// QueryParam 查询字段
type QueryParam map[string]string

// HeaderParam header头部字段
type HeaderParam map[string]string

// QueryReq query对象
type QueryReq struct {
	URL         string
	QueryParam  QueryParam
	HeaderParam HeaderParam
}

// PostJSONReq post对象
type PostJSONReq struct {
	URL         string
	HeaderParam HeaderParam
	Body        []byte
}

func parseHost(host string, url string) string {
	if strings.Contains(host, "http://") || strings.Contains(host, "https://") {
		return strings.Join([]string{host, url}, "")
	}
	return strings.Join([]string{"http://", host, url}, "")
}

func newQuery(ctx context.Context, host string, queryReq QueryReq) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parseHost(host, queryReq.URL), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range queryReq.HeaderParam {
		req.Header.Set(k, v)
	}
	q := req.URL.Query()
	for k, v := range queryReq.QueryParam {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return req, err
}

func newJSONPost(ctx context.Context, host string, PostJSONReq PostJSONReq) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, parseHost(host, PostJSONReq.URL), bytes.NewBuffer(PostJSONReq.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range PostJSONReq.HeaderParam {
		req.Header.Set(k, v)
	}
	return req, err
}

// StringToBytes String to []byte。
// 需要注意的是：该转换是直接替换指针的指向，从而使得string和[]byte指向同一个底层数组，性能更好。
// 但是该转换会有一定的安全隐患：即修改string和[]byte的底层数组会对原值产生影响，所以：
// 1、在你不确定安全隐患的条件下，尽量采用标准方式进行数据转换「[]byte(data)」.
// 2、当程序对运行性能有高要求，同时满足对数据仅仅只有读操作的条件，且存在频繁转换，可以使用该转换。
func StringToBytes(data string) []byte {
	sth := (*reflect.StringHeader)(unsafe.Pointer(&data))
	slh := reflect.SliceHeader{
		Data: sth.Data,
		Len:  sth.Len,
		Cap:  sth.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&slh))
}

// BytesToString []byte to String。
// 需要注意的是：该转换是直接替换指针的指向，从而使得string和[]byte指向同一个底层数组，性能更好。
// 但是该转换会有一定的安全隐患：即修改string和[]byte的底层数组会对原值产生影响，所以：
// 1、在你不确定安全隐患的条件下，尽量采用标准方式进行数据转换「String(data)」.
// 2、当程序对运行性能有高要求，同时满足对数据仅仅只有读操作的条件，且存在频繁转换，可以使用该转换。
func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
