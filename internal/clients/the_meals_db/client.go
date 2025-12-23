package themealsdb_client

import "github.com/Vyach12/meal-api/internal/clients"

type clientImpl struct {
	cl     clients.HTTPClient
	config clients.TheMealsDbClientConfig
}

func New(config clients.TheMealsDbClientConfig, httpClient clients.HTTPClient) *clientImpl {
	return &clientImpl{
		config: config,
		cl:     httpClient,
	}
}
