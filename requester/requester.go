package requester

import (
	"net/http"

	"github.com/google/uuid"
)

// Requester is a requester of all action in this project. Mostly will hold all of REST and gRPC metadata.
type Requester struct {
	Metadata Metadata
}

// Metadata is a metadata of requester.
type Metadata struct {
	RequestID string
	IPAddress string
	UserAgent string
}

// NewRequester is a constructor.
func NewRequester() *Requester {
	return &Requester{}
}

// SetMetadataFromREST is to set metadata from REST API.
func (req *Requester) SetMetadataFromREST(r *http.Request) *Requester {
	req.Metadata = Metadata{
		RequestID: func() string {
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				return uuid.NewString()
			}
			return requestID
		}(),
		IPAddress: r.RemoteAddr,
		UserAgent: r.UserAgent(),
	}
	return req
}
