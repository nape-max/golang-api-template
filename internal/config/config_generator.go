package conf

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigGenerator struct {
	Generator generator
}

type generator struct {
	PathToGeneratedServer string
	PathToHandlers        string
}

func NewConfigGenerator(configPath string) (ConfigGenerator, error) {
	var cfg ConfigGenerator

	_, err := os.Stat(configPath)
	if err != nil {
		return cfg, fmt.Errorf("cannot receive stat of config file: %w", err)
	}

	_, err = toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("cannot decode config to struct: %w", err)
	}

	return cfg, nil
}
