package themealsdb_client

import (
	"context"
	"time"

	app_config "github.com/Vyach12/meal-api/internal/config"
	"github.com/gomeal/config/pkg/config"
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

func NewConfig(ctx context.Context, provider config.Provider) (*configImpl, error) {
	c := provider.GetConfigClient().GetValue(app_config.TheMealsDbUrl).String()
	t := provider.GetConfigClient().GetValue(app_config.TheMealsDbTimeout).Duration()

	return &configImpl{c, t}, nil
}
