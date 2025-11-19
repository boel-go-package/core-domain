package config

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Schema   string
	JDBCURL  string
}

type DbConnection interface {
	Connect()
	NewConfig() DbConfig
}
