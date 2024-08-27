package pg_client

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewConfig(opts ...PostgresConfigOption) *PostgresConfig {
	const (
		defaultUsername = "postgres"
		defaultPassword = ""
		defaultHost     = "127.0.0.1"
		defaultPort     = "5432"
		defaultDatabase = "postgres"
	)

	c := &PostgresConfig{
		Username: defaultUsername,
		Password: defaultPassword,
		Host:     defaultHost,
		Port:     defaultPort,
		Database: defaultDatabase,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type PostgresConfigOption func(*PostgresConfig)

func WithUsername(username string) PostgresConfigOption {
	return func(c *PostgresConfig) {
		c.Username = username
	}
}

func WithPassword(password string) PostgresConfigOption {
	return func(c *PostgresConfig) {
		c.Password = password
	}
}

func WithHost(host string) PostgresConfigOption {
	return func(c *PostgresConfig) {
		c.Host = host
	}
}

func WithPort(port string) PostgresConfigOption {
	return func(c *PostgresConfig) {
		c.Port = port
	}
}

func WithDatabase(database string) PostgresConfigOption {
	return func(c *PostgresConfig) {
		c.Database = database
	}
}
