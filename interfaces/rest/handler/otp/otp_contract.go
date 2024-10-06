package otp

import "github.com/yogamandayu/ohmytp/domain/entity"

type RequestOtpRequestContract struct {
	RouteType string `json:"route_type"`
}

func (r RequestOtpRequestContract) TransformToOtpEntity() entity.Otp {
	return entity.Otp{
		RouteType: r.RouteType,
	}
}

type RequestOtpResponse struct {
	Message string `json:"message"`
}
