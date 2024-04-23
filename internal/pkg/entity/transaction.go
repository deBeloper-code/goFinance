package entity

import "time"

type Transaction struct {
	TransactionID            string    `gorm:"primaryKey" json:"transactionId"`
	CardID                   string    `json:"cardId"`
	UserID                   string    `json:"userId"`
	TypeOfTransaction        string    `json:"typeOfTransaction"`
	Category                 string    `json:"category"`
	Name                     string    `json:"name"`
	Description              string    `json:"description"`
	Amount                   float64   `json:"amount"`
	Date                     time.Time `json:"date"`
	DestinationAccountNumber string    `json:"destinationAccountNumber"`
}
