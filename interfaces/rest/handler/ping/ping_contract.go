package ping

type PingResponse struct {
	Message     string      `json:"message"`
	Timestamp   string      `json:"timestamp"`
	StackStatus StackStatus `json:"stack_status"`
}

type StackStatus struct {
	Db    string `json:"db"`
	Redis string `json:"redis"`
}

type PingErrResponse struct {
	Code    int    `json:"code"`
	Error   any    `json:"error,omitempty"`
	Message string `json:"message"`
}
