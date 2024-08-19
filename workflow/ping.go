package workflow

import "time"

type PingWorkflow struct {
}

func NewPingWorkflow() *PingWorkflow {
	return &PingWorkflow{}
}

func (p *PingWorkflow) Ping() (message, timestamp string) {
	return "Pong!", time.Now().Format(time.RFC3339)
}
