package conf

type ConfigServer struct {
	Database map[string]database
}

type database struct {
	Host     string
	Port     string
	Username string
	Password string
}
