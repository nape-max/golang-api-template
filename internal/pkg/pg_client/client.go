package pg_client

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(cfg PostgresConfig) (*sqlx.DB, error) {
	connectionString := buildConnectionString(cfg)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to databae: %w", err)
	}

	return db, nil
}

func buildConnectionString(cfg PostgresConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}
