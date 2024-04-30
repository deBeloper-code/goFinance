package entity

import "time"

type Transaction struct {
	TransactionID     string    `gorm:"primaryKey"`
	SourceCardID      string    `gorm:"source_card_id"`
	DestinationCardID string    `gorm:"destination_card_id"`
	TypeOfTransaction string    `gorm:"type_of_transaction"`
	Category          string    `gorm:"category"`
	Name              string    `gorm:"name"`
	Description       string    `gorm:"description"`
	Amount            float64   `gorm:"type:numeric(15,2);"`
	Date              time.Time `gorm:"autoCreateTime"`
	SourceCard        Card      `gorm:"foreignKey:SourceCardID"`
	DestinationCard   Card      `gorm:"foreignKey:DestinationCardID"`
}
