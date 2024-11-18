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
	Db    DbStatus
	Redis RedisStatus
}

type DbStatus struct {
	Status        string
	TotalConns    uint32
	IdleConns     uint32
	AcquiredConns uint32
}

type RedisStatus struct {
	Status     string
	TotalConns uint32
	IdleConns  uint32
	StaleConns uint32
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
			Db: DbStatus{
				Status: "ERROR",
			},
			Redis: RedisStatus{
				Status: "ERROR",
			},
		},
	}

	err := p.db.Ping(ctx)
	if err == nil {
		status.StackStatus.Db.Status = "OK"
		status.StackStatus.Db.TotalConns = uint32(p.db.Stat().TotalConns())
		status.StackStatus.Db.IdleConns = uint32(p.db.Stat().IdleConns())
		status.StackStatus.Db.AcquiredConns = uint32(p.db.Stat().AcquiredConns())
	}

	redisStatus := p.redis.Ping(ctx)
	if redisStatus.Err() == nil {
		status.StackStatus.Redis = RedisStatus{
			Status:     "OK",
			TotalConns: p.redis.PoolStats().TotalConns,
			IdleConns:  p.redis.PoolStats().IdleConns,
			StaleConns: p.redis.PoolStats().StaleConns,
		}
	}

	return status
}
