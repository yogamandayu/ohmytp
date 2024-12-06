package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse is a struct to hold http success response data.
// @tag.name SuccessResponse
// @tag.description Response for success request
type SuccessResponse struct {
	Data           interface{} `json:"data,omitempty"`
	Message        string      `json:"message"`
	HTTPStatusCode int         `json:"-"`
}

// NewHTTPSuccessResponse is a constructor.
func NewHTTPSuccessResponse(data interface{}, message string) *SuccessResponse {
	if data == nil {
		data = ""
	}
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
