package database

import (
	"database/sql"
	"os"
	"sync"

	"github.com/boel-go-package/core-domain/cmd/domain/config"
	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type PostgresConnection struct {
}

var (
	dbInstance *bun.DB
	once       sync.Once
	err        error
)

func PGGetPConnection() *bun.DB {

	once.Do(func() {
		var dbPostgres config.DbConnection = PostgresConnection{}
		dbPostgres.NewConfig()

		dbInstance, err = dbPostgres.Connect()
		if err != nil {
			panic(err)
		}
	})

	return dbInstance
}

func PGCloseConnection(db *bun.DB) error {
	if db == nil {
		return sql.ErrConnDone
	}
	return db.Close()
}

func PGStartTransaction(db *bun.DB) (*bun.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func PGCommitTransaction(tx *bun.Tx) error {
	if tx == nil {
		return sql.ErrTxDone
	}
	return tx.Commit()
}

func PGRollbackTransaction(tx *bun.Tx) error {
	if tx == nil {
		return sql.ErrTxDone
	}
	return tx.Rollback()
}

// NewConfig implements config.DbConnection.
func (p PostgresConnection) NewConfig() config.DbConfig {
	return config.DbConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DATABASE"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		Schema:   os.Getenv("POSTGRES_SCHEMA"),
		JDBCURL:  os.Getenv("POSTGRES_JDBCURL"),
	}
}

// Connect implements config.DbConnection.
func (p PostgresConnection) Connect() (*bun.DB, error) {

	var sqldb sql.DB

	if p.NewConfig().JDBCURL != "" {
		sqldb = *sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithDSN(p.NewConfig().JDBCURL),
		))
	} else if p.NewConfig().Host != "" {
		strUrl := "postgres://" + p.NewConfig().User + ":" + p.NewConfig().Password + "@" + p.NewConfig().Host + ":" + p.NewConfig().Port + "/" + p.NewConfig().DBName

		if p.NewConfig().SSLMode != "" {
			strUrl = strUrl + "?sslmode=" + p.NewConfig().SSLMode
		}

		if p.NewConfig().Schema != "" && p.NewConfig().SSLMode != "" {
			strUrl = strUrl + "&search_path=" + p.NewConfig().Schema
		} else if p.NewConfig().Schema != "" {
			strUrl = strUrl + "?search_path=" + p.NewConfig().Schema
		}

		sqldb = *sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithDSN(strUrl),
		))
	} else {
		return nil, sql.ErrConnDone
	}

	db := bun.NewDB(&sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),

		// BUNDEBUG=1 logs failed queries
		// BUNDEBUG=2 logs all queries
		bundebug.FromEnv(),
	))

	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbInstance = db
	return db, nil
}
