package consts

type CircuitBreakerState string

const (
	CircuitBreakerStateOpen     CircuitBreakerState = "OPEN"
	CircuitBreakerStateClose    CircuitBreakerState = "CLOSE"
	CircuitBreakerStateHalfOpen CircuitBreakerState = "HALF_OPEN"
)

func (c CircuitBreakerState) String() string {
	return string(c)
}
