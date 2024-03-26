package entity

type Account struct {
	AccountID     string `gorm:"primaryKey"`
	UserID        string
	AccountNumber string
	Balance       float32
}
