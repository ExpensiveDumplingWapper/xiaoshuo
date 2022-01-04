package user

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	HTTPClient "git.ourbluecity.com/overseas-server-tools/go-util-http"
)

var (
	baseClient HTTPClient.HTTPClienter
)

const (
	HTTPTools = "tools"
)

func InitApp() {
	baseClient = HTTPClient.NewHTTPClientV2(context.Background(), &HTTPClient.HTTPOptions{
		ZookeeperPath: []string{"10.9.158.210:2181", "10.9.114.167:2181"},
		Conf: map[string]string{
			"tools": "/blued/service/users/service/profiles",
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
}
