package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Data           interface{} `json:"data"`
	Message        string      `json:"message"`
	HTTPStatusCode int         `json:"-"`
}

func NewHTTPSuccessResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:    data,
		Message: message,
	}
}

func (s *SuccessResponse) WithStatusCode(code int) *SuccessResponse {
	s.HTTPStatusCode = code
	return s
}

func (s *SuccessResponse) AsJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.HTTPStatusCode)
	_ = json.NewEncoder(w).Encode(s)
}
