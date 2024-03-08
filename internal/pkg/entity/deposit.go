package entity

import (
	"time"
)

// Deposit represents a movement made by user of type 'expense'.
type Deposit struct {
	UserId      string    `firestore:"userID"`
	Category    string    `firestore:"category" binding:"required"`
	Name        string    `firestore:"name" binding:"required"`
	Description string    `firestore:"description" binding:"required"`
	Amount      float32   `firestore:"amount" binding:"required"`
	Date        time.Time `firestore:"date" binding:"required"`
}
