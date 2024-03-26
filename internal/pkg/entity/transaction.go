package entity

import "time"

type Transaction struct {
	TransactionID     string    `firestore:"transactionID" json:"transactionId"`
	AccountID         string    `firestore:"accountID" json:"accountId"`
	CardID            string    `firestore:"cardID" json:"cardId"`
	TypeOfTransaction string    `firestore:"typeOfTransaction" json:"typeOfTransaction"`
	Category          string    `firestore:"category" json:"category"`
	Name              string    `firestore:"name" json:"name"`
	Description       string    `firestore:"description" json:"description"`
	Amount            float64   `firestore:"amount" json:"amount"`
	Date              time.Time `firestore:"date" json:"date"`
}
