package conf

import (
	"context"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sethvargo/go-envconfig"
)

type ConfigServer struct {
	PostgresDatabase postgresDatabase
}

type postgresDatabase struct {
	Host     *string `env:"POSTGRES_HOST, noinit, overwrite"`
	Port     *string `env:"POSTGRES_PORT, noinit, overwrite"`
	Username *string `env:"POSTGRES_USERNAME, noinit, overwrite"`
	Password *string `env:"POSTGRES_PASSWORD, noinit, overwrite"`
	Database *string `env:"POSTGRES_DATABASE, noinit, overwrite"`
}

func NewConfigServer(ctx context.Context, configPath string) (ConfigServer, error) {
	var cfg ConfigServer

	_, err := os.Stat(configPath)
	if err != nil {
		return cfg, fmt.Errorf("cannot receive stat of config file: %w", err)
	}

	_, err = toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("cannot decode config from file to struct: %w", err)
	}

	if err := envconfig.Process(ctx, &cfg); err != nil {
		return cfg, fmt.Errorf("cannot receive config from env: %w", err)
	}

	return cfg, nil
}
