package ping

type PingResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type PingErrResponse struct {
	Code    int
	Error   any
	Message string
}
