package otp

import (
	"strings"

	"github.com/yogamandayu/ohmytp/internal/domain/entity"
)

// RequestOtpRequestContract is request otp request contract.
type RequestOtpRequestContract struct {
	Identifier string `json:"identifier"`
	RouteType  string `json:"route_type"`
	RouteValue string `json:"route_value"`
	Purpose    string `json:"purpose"`
	Length     uint16 `json:"length"`
	Expiration int    `json:"expiration"`
}

// TransformToOtpEntity is to transform to otp entity.
func (r RequestOtpRequestContract) TransformToOtpEntity() entity.Otp {
	return entity.Otp{
		Purpose:    r.Purpose,
		Identifier: r.Identifier,
		RouteType:  strings.ToUpper(r.RouteType),
	}
}

// RequestOtpResponseContract is request otp response contract.
type RequestOtpResponseContract struct {
	ExpiredAt string `json:"expired_at"`
}

// ConfirmOtpRequestContract is request otp request contract.
//
// @tag.name ConfirmOtpRequestContract
// @tag.description Confirm OTP response API contract.
type ConfirmOtpRequestContract struct {
	Identifier string `json:"identifier"`
	Code       string `json:"code"`
	Purpose    string `json:"purpose"`
}

// TransformToOtpEntity is to transform to otp entity.
func (r ConfirmOtpRequestContract) TransformToOtpEntity() entity.Otp {
	return entity.Otp{
		Identifier: r.Identifier,
		Code:       r.Code,
		Purpose:    r.Purpose,
	}
}
