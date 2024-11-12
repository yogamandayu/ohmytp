package ping

// PingResponseContract is ping response contract.
type PingResponseContract struct {
	Message     string      `json:"message"`
	Timestamp   string      `json:"timestamp"`
	StackStatus StackStatus `json:"stack_status"`
}

// StackStatus is stack/dependency status.
type StackStatus struct {
	Db    string `json:"db"`
	Redis string `json:"redis"`
}

// PingErrResponseContract is ping error response contract.
type PingErrResponseContract struct {
	Code    int    `json:"code"`
	Error   any    `json:"error,omitempty"`
	Message string `json:"message"`
}
