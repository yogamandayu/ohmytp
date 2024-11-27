package consts

type RollbarSeverityLevel string

const (
	RollbarSeverityLevelDebug    RollbarSeverityLevel = "DEBUG"
	RollbarSeverityLevelInfo     RollbarSeverityLevel = "INFO"
	RollbarSeverityLevelWarning  RollbarSeverityLevel = "WARNING"
	RollbarSeverityLevelError    RollbarSeverityLevel = "ERROR"
	RollbarSeverityLevelCritical RollbarSeverityLevel = "CRITICAL"
)

func (r RollbarSeverityLevel) String() string {
	return string(r)
}

func (r RollbarSeverityLevel) ToCode() int {
	switch r {
	case RollbarSeverityLevelDebug:
		return 10
	case RollbarSeverityLevelInfo:
		return 20
	case RollbarSeverityLevelWarning:
		return 30
	case RollbarSeverityLevelError:
		return 40
	case RollbarSeverityLevelCritical:
		return 50
	}
	return 0
}
