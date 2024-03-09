package entity

import (
	"time"
)

// Expense represents a single movement of type 'expense'.
type Expense struct {
	UserId      string    `firestore:"userID"`
	Category    string    `firestore:"category" binding:"required"`
	Name        string    `firestore:"name" binding:"required"`
	Description string    `firestore:"description" binding:"required"`
	Amount      float32   `firestore:"amount" binding:"required"`
	Date        time.Time `firestore:"date" binding:"required"`
}
