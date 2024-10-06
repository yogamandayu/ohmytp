package otp

import (
	"context"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/yogamandayu/ohmytp/workflow/otp"
)

func (h *Handler) Request(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)

	var payload RequestOtpRequestContract
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, "X-Request-ID", r.Header.Get("X-Request-ID"))

	workflow := otp.NewRequestOtpWorkflow(h.app.DB, h.app.Log)

	data := RequestOtpResponse{
		Message: "OK",
	}
	err = workflow.Request(ctx, payload.TransformToOtpEntity())
	if err != nil {
		data.Message = "ERROR"
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
