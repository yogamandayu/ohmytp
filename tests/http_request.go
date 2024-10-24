package tests

import (
	"net/http"

	"github.com/google/uuid"
)

func FakeHTTPRequest() *http.Request {
	r := &http.Request{
		Header: http.Header{},
	}
	r.Header.Set("X-Request-ID", uuid.NewString())
	r.RemoteAddr = "127.0.0.1"
	r.Header.Set("User-Agent", "local-testing")
	return r
}
