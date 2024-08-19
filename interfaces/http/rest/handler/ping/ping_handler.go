package ping

import (
	"net/http"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/workflow"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	pingWorkflow := workflow.NewPingWorkflow()
	msg, ts := pingWorkflow.Ping()
	data := PingResponse{
		Message:   msg,
		Timestamp: ts,
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
