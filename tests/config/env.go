package config

import "github.com/kelseyhightower/envconfig"

const envPrefix = "QA"

type Config struct {
	Host string `split_words:"true" default:"localhost:8081"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
