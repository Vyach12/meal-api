package themealsdb_client

import (
	"net/http"
	"time"
)

type clientImpl struct {
	cl     HTTPClient
	config TheMealsDbClientConfig
}

type TheMealsDbClientConfig interface {
	Url() string
	Timeout() time.Duration
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(config TheMealsDbClientConfig, httpClient HTTPClient) *clientImpl {
	return &clientImpl{
		config: config,
		cl:     httpClient,
	}
}
