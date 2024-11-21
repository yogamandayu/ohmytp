package requester

import (
	"net"
	"net/http"
	"strings"

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
		IPAddress: getIPAddress(r),
		UserAgent: r.UserAgent(),
	}
	return req
}

func getIPAddress(r *http.Request) string {
	// Check X-Forwarded-For header
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // Return the first IP in the list
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Return raw value if parsing fails
	}

	return ip
}
