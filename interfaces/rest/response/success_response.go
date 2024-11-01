package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewHTTPSuccessResponse(w http.ResponseWriter, code int, data interface{}, message string) {
	res := &SuccessResponse{
		Data:    data,
		Message: message,
	}
	b, _ := json.Marshal(res)
	w.Write(b)
	w.WriteHeader(code)
	return
}
