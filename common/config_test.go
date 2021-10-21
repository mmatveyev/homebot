package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvConfig(t *testing.T)  {
	os.Setenv("TELEGRAM_TOKEN", "token1")
	os.Setenv("TELEGRAM_DEBUG", "1")
	os.Setenv("TELEGRAM_CLIENT_TIMEZONE", "tz123")
	config, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, Config{Telegram: Telegram{
		Token:          "token1",
		Debug:          true,
		ClientTimezone: "tz123",
	}}, *config)
}
