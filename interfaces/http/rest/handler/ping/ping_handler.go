package ping

import (
	"context"
	"net/http"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/workflow"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	pingWorkflow := workflow.NewPingWorkflow(h.db)
	msg, ts, err := pingWorkflow.Ping(context.Background())
	var data any
	if err != nil {
		data = PingErrResponse{
			Code:    http.StatusInternalServerError,
			Error:   nil,
			Message: err.Error(),
		}
	} else {
		data = PingResponse{
			Message:   msg,
			Timestamp: ts,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
