package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/yogamandayu/ohmytp/config"
)

func NewConnection(config *config.Config) (*sql.DB, error) {
	if config.DB == nil {
		return nil, errors.New("missing config")
	}

	dbConfig := config.DB
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := sql.Open(dbConfig.Driver, connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
