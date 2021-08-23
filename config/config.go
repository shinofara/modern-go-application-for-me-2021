package config

import (
	"mygo/infrastructure/database"
	"mygo/infrastructure/logger"
	"mygo/infrastructure/trace"
	"os"

	"go.uber.org/dig"

	"gopkg.in/yaml.v2"
)

type Config struct {
	dig.Out

	DB     *database.Config
	Trace  *trace.Config
	Logger *logger.Config
}

func (cfg Config) Clone() Config {
	return cfg
}

func New(configFilePath string) (Config, error) {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()
	var cfg Config
	if err := yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
