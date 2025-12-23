package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ApplicationName string                 `yaml:"application_name"`
	Env             string                 `yaml:"env" env-default:"LOCAL"`
	Integration     map[string]Integration `yaml:"integrations"`
}

type Integration struct {
	Url     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
}

func NewConfig(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s\n", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s\n", err)
	}

	return &cfg
}

func (c *Config) GetIntegration(name string) (Integration, error) {
	i, exists := c.Integration[name]
	if !exists {
		return Integration{}, fmt.Errorf("integration '%s' not found in config", name)
	}

	return i, nil
}
