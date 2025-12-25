package app_config

import "github.com/gomeal/config/pkg/config"

const (
	ApplicationName = config.Key("application_name")
	Env             = config.Key("env")

	TheMealsDbUrl     = config.Key("the_meals_db_http_url")
	TheMealsDbTimeout = config.Key("the_meals_db_http_timeout")
)
