package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/yogamandayu/ohmytp/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func NewConnection(config *config.Config) (*sql.DB, error) {
	if config.DB == nil {
		return nil, errors.New("missing config")
	}

	dbConfig := config.DB
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	dbConn, err := sql.Open(dbConfig.Driver, connectionString)
	if err != nil {
		return nil, err
	}
	dbConn.Ping()
	return dbConn, nil
}
