package otp

import "github.com/yogamandayu/ohmytp/domain/entity"

// RequestOtpRequestContract is request otp request contract.
type RequestOtpRequestContract struct {
	RouteType  string `json:"route_type"`
	RouteValue string `json:"route_value"`
}

// TransformToOtpEntity is to transform to otp entity.
func (r RequestOtpRequestContract) TransformToOtpEntity() entity.Otp {
	return entity.Otp{
		RouteType: r.RouteType,
	}
}

// RequestOtpResponseContract is request otp response contract.
type RequestOtpResponseContract struct {
	Message string `json:"message"`
}
