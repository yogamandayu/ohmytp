package ping

import (
	"net/http"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/workflow"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pingWorkflow := workflow.NewPingWorkflow(h.db)
	status := pingWorkflow.Ping(ctx)
	data := PingResponse{
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
