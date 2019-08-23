package common

import (
	"errors"
	"github.com/BurntSushi/toml"
	"os"
)

type Telegram struct {
	Token          string `toml:"token"`
	Debug          bool   `toml:"debug"`
	ClientTimezone string `toml:"client_timezone"`
}

type AuthorizedClient struct {
	TelegramId   string `toml:"telegram_id"`
	DropboxToken string `toml:"dropbox_token"`
	EcovacsToken string `toml:"ecovacs_token"`
}

type Config struct {
	Telegram Telegram           `toml:"telegram"`
	Clients  []AuthorizedClient `toml:"authorized_clients"`
}

func NewConfigFromFile(path string) (*Config, error) {
	var result = new(Config)
	if _, err := toml.DecodeFile(path, &result); err != nil {
		return nil, err
	}
	err := envConfig(result)
	return result, err
}

func envConfig(config *Config) error {
	if v := os.Getenv("HOMEBOT_TELEGRAM_TOKEN"); v != "" {
		config.Telegram.Token = v
	} else if config.Telegram.Token == "" {
		return errors.New("Telegram token is undefined")
	}
	return nil
}
