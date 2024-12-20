package workflow

import (
	"context"
	"time"

	"github.com/yogamandayu/ohmytp/internal/app"
)

type PingWorkflow struct {
	app *app.App
}

type StackStatus struct {
	Db    DbStatus
	Redis RedisStatus
	Minio MinioStatus
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

type MinioStatus struct {
	Status string
}

type PingStatus struct {
	Message     string
	Timestamp   string
	StackStatus StackStatus
}

func NewPingWorkflow(app *app.App) *PingWorkflow {
	return &PingWorkflow{
		app: app,
	}
}

func (p *PingWorkflow) Ping(ctx context.Context) PingStatus {
	status := PingStatus{
		Message:   "Pong!",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err := p.app.DB.Ping(ctx)
	if err != nil {
		p.app.Log.Error(err.Error())
		status.StackStatus.Db.Status = "ERROR"
	} else {
		status.StackStatus.Db.Status = "OK"
		status.StackStatus.Db.TotalConns = uint32(p.app.DB.Stat().TotalConns())
		status.StackStatus.Db.IdleConns = uint32(p.app.DB.Stat().IdleConns())
		status.StackStatus.Db.AcquiredConns = uint32(p.app.DB.Stat().AcquiredConns())
	}

	redisStatus := p.app.RedisAPI.Ping(ctx)
	if redisStatus.Err() != nil {
		p.app.Log.Error(redisStatus.Err().Error())
		status.StackStatus.Redis.Status = "ERROR"
	} else {
		status.StackStatus.Redis = RedisStatus{
			Status:     "OK",
			TotalConns: p.app.RedisAPI.PoolStats().TotalConns,
			IdleConns:  p.app.RedisAPI.PoolStats().IdleConns,
			StaleConns: p.app.RedisAPI.PoolStats().StaleConns,
		}
	}

	_, err = p.app.Minio.ListBuckets(ctx)
	if err != nil {
		p.app.Log.Error(err.Error())
		status.StackStatus.Minio.Status = "ERROR"
	} else {
		status.StackStatus.Minio = MinioStatus{
			Status: "OK",
		}
	}

	return status
}
