package config

import (
	"mygo/infrastructure/database"
	"mygo/infrastructure/trace"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB database.Config
	Trace trace.Config
}

func DB(cfg *Config) *database.Config {
	return &cfg.DB
}

func Trace(cfg *Config) *trace.Config {
	return &cfg.Trace
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
