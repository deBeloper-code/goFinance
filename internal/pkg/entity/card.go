package entity

import "time"

type Card struct {
	CardID        string
	UserID        string
	AccountNumber int
	CardNumber    int
	Balance       float64
	ExpiryDate    time.Time
	CVV           int
	Issuer        string
	Type          string
}
