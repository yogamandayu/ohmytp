package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse is a struct to hold http success response data.
type SuccessResponse struct {
	Data           interface{} `json:"data"`
	Message        string      `json:"message"`
	HTTPStatusCode int         `json:"-"`
}

// NewHTTPSuccessResponse is a constructor.
func NewHTTPSuccessResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:    data,
		Message: message,
	}
}

// WithStatusCode is to set http status code to success response.
func (s *SuccessResponse) WithStatusCode(code int) *SuccessResponse {
	s.HTTPStatusCode = code
	return s
}

// AsJSON is to set success response as JSON.
func (s *SuccessResponse) AsJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.HTTPStatusCode)
	_ = json.NewEncoder(w).Encode(s)
}
