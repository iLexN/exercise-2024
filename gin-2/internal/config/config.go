package config

type Config struct {
	AppEnv   *appEnv
	Server   *serverConfig
	Database *DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		AppEnv:   newAppEnv(),
		Server:   newServerConfig(),
		Database: NewMysql(),
	}
}
