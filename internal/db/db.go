package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/yogamandayu/ohmytp/config"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func NewConnection(config config.Config) (*pgxpool.Pool, error) {
	if config.DB == nil {
		return nil, errors.New("missing config")
	}

	dbConfig := config.DB
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	if config.DB.Log {
		tracer := &CustomTracer{}
		pgxConfig.ConnConfig.Tracer = tracer
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

// CustomTracer implements the pgx.Trace interface
type CustomTracer struct{}

// TraceQueryStart logs the beginning of a query
func (ct *CustomTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	slog.Info(fmt.Sprintf("Executing query: %s, args: %v", data.SQL, data.Args))
	return ctx
}

// TraceQueryEnd logs the end of a query
func (ct *CustomTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		slog.Warn(fmt.Sprintf("Query failed: %v", data.Err))
	} else {
		slog.Info(fmt.Sprintf("Query successful, time: %v", data.CommandTag))
	}
}
