package consts

type RouteType string

const (
	EmailRouteType RouteType = "EMAIL"
	SMSRouteType   RouteType = "SMS"
)

func (r RouteType) ToString() string {
	return string(r)
}
