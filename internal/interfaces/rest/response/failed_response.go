package response

import (
	"encoding/json"
	"net/http"
)

// FailedResponse is a struct to hold http failed response data.
type FailedResponse struct {
	Error          string `json:"error"`
	Code           string `json:"code"`
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"-"`
}

// NewHTTPFailedResponse is a constructor.
func NewHTTPFailedResponse(errCode string, err error, message string) *FailedResponse {
	return &FailedResponse{
		Code: errCode,
		Error: func() string {
			if err != nil {
				return err.Error()
			}
			return ""
		}(),
		Message: message,
	}
}

// WithStatusCode is to set http status code to failed response.
func (f *FailedResponse) WithStatusCode(code int) *FailedResponse {
	f.HTTPStatusCode = code
	return f
}

// AsJSON is to set success response as JSON.
func (f *FailedResponse) AsJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(f.HTTPStatusCode)
	_ = json.NewEncoder(w).Encode(f)
}
