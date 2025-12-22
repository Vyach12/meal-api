package themealsdb_client

import (
	"context"
	"time"
)

type configImpl struct {
	url     string
	timeout time.Duration
}

func (c *configImpl) Url() string {
	return c.url
}

func (c *configImpl) Timeout() time.Duration {
	return c.timeout
}

func NewConfig(ctx context.Context) (*configImpl, error) {
	c := &configImpl{
		url:     "https://www.themealdb.com/api/json/v1/1/random.php",
		timeout: 10000,
	}

	return c, nil
}