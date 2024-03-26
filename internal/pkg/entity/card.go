package entity

import "time"

type Card struct {
	CardID     string    `firestore:"cardID" json:"cardId"`
	AccountID  string    `firestore:"accountID" binding:"required" json:"accountId"`
	CardNumber string    `firestore:"cardNumber" binding:"required" json:"cardNumber" validate:"required,len=16"`
	ExpiryDate time.Time `firestore:"expiryDate" binding:"required" json:"expiryDate"`
	CVV        string    `firestore:"cvv" binding:"required" json:"cvv" validate:"required,len=3"`
	Issuer     string    `firestore:"issuer" binding:"required" json:"issuer"`
	Type       string    `firestore:"type" binding:"required" json:"type"`
}
