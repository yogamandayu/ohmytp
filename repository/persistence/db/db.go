package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/yogamandayu/ohmytp/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func NewConnection(config *config.Config) (*pgxpool.Pool, error) {
	if config.DB == nil {
		return nil, errors.New("missing config")
	}

	dbConfig := config.DB
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	dbConn, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
