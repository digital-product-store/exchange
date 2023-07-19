package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	ListenAddr string `env:"LISTEN_ADDR" envDefault:":8000"`
}

func LoadConfig() (*Config, error) {
	conf := &Config{}
	err := env.Parse(conf)
	return conf, err
}
