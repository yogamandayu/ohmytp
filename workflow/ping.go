package workflow

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PingWorkflow struct {
	db *pgxpool.Pool
}

type StackStatus struct {
	Db    string
	Redis string
}

type PingStatus struct {
	Message     string
	Timestamp   string
	StackStatus StackStatus
}

func NewPingWorkflow(db *pgxpool.Pool) *PingWorkflow {
	return &PingWorkflow{
		db,
	}
}

func (p *PingWorkflow) Ping(ctx context.Context) PingStatus {
	status := PingStatus{
		Message:   "Pong!",
		Timestamp: time.Now().Format(time.RFC3339),
		StackStatus: StackStatus{
			Db:    "UNDEFINED",
			Redis: "UNDEFINED",
		},
	}
	err := p.db.Ping(ctx)
	status.StackStatus.Db = "OK"
	if err != nil {
		status.StackStatus.Db = "ERROR"
	}

	return status
}
