package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		UrlFetch
		AppPort
	}

	UrlFetch struct {
		Url string `env-required:"true" env:"COVID_DATA"`
	}
	AppPort struct {
		Port string `env-required:"true" env:"PORT"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	if _, err := os.Stat(".env"); err == nil {
		err = cleanenv.ReadConfig(".env", cfg)
		if err != nil {
			return nil, fmt.Errorf("config error: %w", err)
		}
	}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println("cfgxxx", cfg.UrlFetch)
	return cfg, nil
}
