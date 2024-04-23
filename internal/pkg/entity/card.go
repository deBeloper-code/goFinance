package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
//TODO: We need to be better.

type Card struct {
	CardID        string        `gorm:"primaryKey" json:"cardId"`
	UserID        string        `gorm:"index" json:"userId" binding:"required"`
	AccountNumber string        `gorm:"unique;not null" json:"accountNumber" binding:"required"`
	CardNumber    string        `gorm:"unique;not null" json:"cardNumber" binding:"required,len=16"`
	Balance       float64       `json:"balance"`
	ExpiryDate    time.Time     `json:"expiryDate"`
	CVV           string        `json:"cvv" binding:"required,len=3"`
	Issuer        string        `json:"issuer"`
	Type          string        `json:"type"`
	User          User          `gorm:"foreignKey:UserID"`
	Transactions  []Transaction `gorm:"foreignKey:CardID"`
}

func (Card) TableName() string {
	return "app.cards"
}

// Hooks
func (card *Card) BeforeCreate(tx *gorm.DB) (err error) {
	card.CardID = uuid.NewString()
	return
}
