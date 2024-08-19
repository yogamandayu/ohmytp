package entity

type OTPBypass struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Code       string `json:"code"`
	Timestamp
}
