package otp

import (
	"github.com/yogamandayu/ohmytp/domain/entity"
	"strings"
)

// RequestOtpRequestContract is request otp request contract.
type RequestOtpRequestContract struct {
	RouteType  string `json:"route_type"`
	RouteValue string `json:"route_value"`
	Purpose    string `json:"purpose"`
	Length     int    `json:"length"`
	Expiration int    `json:"expiration"`
}

// TransformToOtpEntity is to transform to otp entity.
func (r RequestOtpRequestContract) TransformToOtpEntity() entity.Otp {
	return entity.Otp{
		RouteType: strings.ToUpper(r.RouteType),
	}
}

// RequestOtpResponseContract is request otp response contract.
type RequestOtpResponseContract struct {
	Message string `json:"message"`
}
