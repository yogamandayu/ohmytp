package ping

// ResponseContract is ping response contract.
//
// @tag.name ResponseContract
// @tag.description Ping response API contract.
type ResponseContract struct {
	Message     string      `json:"message"`
	Timestamp   string      `json:"timestamp"`
	StackStatus StackStatus `json:"stack_status"`
}

// StackStatus is stack/dependency status.
type StackStatus struct {
	Db    DbStatus    `json:"db"`
	Redis RedisStatus `json:"redis"`
	Minio MinioStatus `json:"minio"`
}

type DbStatus struct {
	Status        string `json:"status"`
	TotalConns    uint32 `json:"total_conns"`
	IdleConns     uint32 `json:"idle_conns"`
	AcquiredConns uint32 `json:"acquired_conns"`
}

type RedisStatus struct {
	Status     string `json:"status"`
	TotalConns uint32 `json:"total_conns"`
	IdleConns  uint32 `json:"idle_conns"`
	StaleConns uint32 `json:"stale_conns"`
}

type MinioStatus struct {
	Status string `json:"status"`
}

// ErrResponseContract is ping error response contract.
type ErrResponseContract struct {
	Code    int    `json:"code"`
	Error   any    `json:"error,omitempty"`
	Message string `json:"message"`
}
