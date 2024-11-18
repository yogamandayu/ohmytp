package ping

import (
	"context"
	"net/http"
	"time"

	"github.com/yogamandayu/ohmytp/internal/workflow"

	"encoding/json"
)

// Ping is ping handler.
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)
	pingWorkflow := workflow.NewPingWorkflow(h.db, h.redis)
	status := pingWorkflow.Ping(ctx)
	data := ResponseContract{
		Message:   status.Message,
		Timestamp: status.Timestamp,
		StackStatus: StackStatus{
			Db: DbStatus{
				Status:        status.StackStatus.Db.Status,
				TotalConns:    status.StackStatus.Db.TotalConns,
				IdleConns:     status.StackStatus.Db.IdleConns,
				AcquiredConns: status.StackStatus.Db.AcquiredConns,
			},
			Redis: RedisStatus{
				Status:     status.StackStatus.Redis.Status,
				TotalConns: status.StackStatus.Redis.TotalConns,
				IdleConns:  status.StackStatus.Redis.IdleConns,
				StaleConns: status.StackStatus.Redis.StaleConns,
			},
		},
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
