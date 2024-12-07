package config

import (
	"errors"
	"github.com/joho/godotenv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DSN string
}

type AppConfig struct {
	Port      string
	UrlPrefix string
	Secret    string
}

func New() (*Config, error) {
	_ = godotenv.Load()

	v := viper.New()

	v.SetDefault("App.Port", "8000")
	v.SetDefault("App.UrlPrefix", "")
	v.SetDefault("Database.DSN", "")
	v.SetDefault("App.Secret", "")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.App.Secret == "" {
		return nil, errors.New("secret is required")
	}

	return &cfg, nil
}
