package config

import "github.com/uptrace/bun"

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
	Connect() (*bun.DB, error)
	Close() error
	Transaction() error
	NewConfig() DbConfig
}
