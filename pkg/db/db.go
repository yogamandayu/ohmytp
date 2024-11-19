package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config is db config.
type Config struct {
	Host              string
	Port              string
	Username          string
	Password          string
	Database          string
	TimeZone          string
	Log               bool
	MaxConns          int
	MinConns          int
	MaxConnIdleTime   time.Duration
	MaxConnLifeTime   time.Duration
	HealthCheckPeriod time.Duration
}

// NewConnection is to set new db connection.
func NewConnection(config *Config) (*pgxpool.Pool, error) {
	if config == nil {
		return nil, errors.New("db.error.missing_config")
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	if config.MaxConns > 0 {
		pgxConfig.MaxConns = int32(config.MaxConns)
	}
	if config.MinConns > 0 {
		pgxConfig.MinConns = int32(config.MinConns)
	}
	if config.MaxConnIdleTime > 0 {
		pgxConfig.MaxConnIdleTime = config.MaxConnIdleTime
	}
	if config.MaxConnLifeTime > 0 {
		pgxConfig.MaxConnLifetime = config.MaxConnLifeTime
	}
	if config.HealthCheckPeriod > 0 {
		pgxConfig.HealthCheckPeriod = config.HealthCheckPeriod
	}

	if config.Log {
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
