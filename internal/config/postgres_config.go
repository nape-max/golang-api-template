package conf

import "maxnap/platform/internal/pkg/pg_client"

func NewPostgresConfig(cfg postgresDatabase) *pg_client.PostgresConfig {
	options := make([]pg_client.PostgresConfigOption, 0, 5)

	if cfg.Username != nil {
		options = append(options, pg_client.WithUsername(*cfg.Username))
	}
	if cfg.Password != nil {
		options = append(options, pg_client.WithPassword(*cfg.Password))
	}
	if cfg.Host != nil {
		options = append(options, pg_client.WithHost(*cfg.Host))
	}
	if cfg.Port != nil {
		options = append(options, pg_client.WithPort(*cfg.Port))
	}
	if cfg.Database != nil {
		options = append(options, pg_client.WithDatabase(*cfg.Database))
	}

	return pg_client.NewConfig(options...)
}
