package ping

import (
	"context"
	"net/http"
	"time"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/workflow"
)

// Ping is ping handler.
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)
	pingWorkflow := workflow.NewPingWorkflow(h.db, h.redis)
	status := pingWorkflow.Ping(ctx)
	data := PingResponseContract{
		Message:   status.Message,
		Timestamp: status.Timestamp,
		StackStatus: StackStatus{
			Db:    status.StackStatus.Db,
			Redis: status.StackStatus.Redis,
		},
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
