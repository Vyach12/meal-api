package themealsdb_client

import (
	"context"
	"fmt"
	"time"

	"github.com/Vyach12/meal-api/internal/config"
)

var (
	integrationName = "the_meal_db"
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

func NewConfig(ctx context.Context, cfg config.Config) (*configImpl, error) {
	c, err := cfg.GetIntegration(integrationName)

	if err != nil {
		return nil, fmt.Errorf("failed to get %s integration config: %w",
			integrationName, err)
	}

	return &configImpl{c.Url, c.Timeout}, nil
}
