package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Driver       string
	Hostname     string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

func NewDatabaseConnection(config DatabaseConfig) (*sqlx.DB, error) {
	dataSoruce := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Hostname, config.Port, config.Username, config.Password, config.DatabaseName)
	db, err := sqlx.Connect(config.Driver, dataSoruce)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
