package workflow

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PingWorkflow struct {
	db    *pgxpool.Pool
	redis *redis.Client
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

func NewPingWorkflow(db *pgxpool.Pool, redis *redis.Client) *PingWorkflow {
	return &PingWorkflow{
		db:    db,
		redis: redis,
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

	redisStatus := p.redis.Ping(ctx)
	status.StackStatus.Redis = "OK"
	if redisStatus.Err() != nil {
		status.StackStatus.Redis = "ERROR"
	}

	return status
}
