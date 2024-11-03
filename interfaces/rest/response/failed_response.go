package response

import (
	"encoding/json"
	"net/http"
)

type FailedResponse struct {
	Error          string `json:"error"`
	Code           string `json:"code"`
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"-"`
}

func NewHTTPFailedResponse(errCode string, err error, message string) *FailedResponse {
	return &FailedResponse{
		Code:    errCode,
		Error:   err.Error(),
		Message: message,
	}
}

func (f *FailedResponse) WithStatusCode(code int) *FailedResponse {
	f.HTTPStatusCode = code
	return f
}

func (f *FailedResponse) AsJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(f.HTTPStatusCode)
	_ = json.NewEncoder(w).Encode(f)
}
