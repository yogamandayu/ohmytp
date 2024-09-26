package workflow

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type PingWorkflow struct {
	db *pgx.Conn
}

func NewPingWorkflow(db *pgx.Conn) *PingWorkflow {
	return &PingWorkflow{
		db,
	}
}

func (p *PingWorkflow) Ping(ctx context.Context) (message, timestamp string, err error) {
	err = p.db.Ping(ctx)
	if err != nil {
		return "", time.Now().Format(time.RFC3339), err
	}

	return "Pong!", time.Now().Format(time.RFC3339), nil
}
