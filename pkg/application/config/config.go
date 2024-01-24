package config

import (
	"fmt"
)

type Config struct{}

func New() (*Config, error) {
	cfg := Default()

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate a config: %w", err)
	}

	return &cfg, nil
}
