package themealsdb_client

import "github.com/Vyach12/meal-api/internal/clients"

type Client struct {
	cl     clients.HTTPClient
	config clients.TheMealsDbClientConfig
}

func New(config clients.TheMealsDbClientConfig, httpClient clients.HTTPClient) *Client {
	return &Client{
		config: config,
		cl:     httpClient,
	}
}
