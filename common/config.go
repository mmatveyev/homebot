package common

import (
	"github.com/caarlos0/env/v6"
)

type Telegram struct {
	Token          string `env:"TELEGRAM_TOKEN"`
	Debug          bool   `env:"TELEGRAM_DEBUG"`
	ClientTimezone string `env:"TELEGRAM_CLIENT_TIMEZONE"`
}

type Config struct {
	Telegram Telegram
}

func GetConfig() (*Config, error) {
	config := new(Config)
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
