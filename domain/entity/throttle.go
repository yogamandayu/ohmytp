package entity

type Throttle struct {
	ID      string `json:"id"`
	Attempt uint8  `json:"attempt"`
	Timestamp
}
