package config

import (
	"gopkg.in/yaml.v2"
	"mygo/infrastructure/database"
	"os"
)

type Config struct {
	DB database.Config
}

func DB(cfg *Config) *database.Config {
	return &cfg.DB
}

func New(configFilePath string) (*Config, error) {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	var cfg Config
	if err := yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

