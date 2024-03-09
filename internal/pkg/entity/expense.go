package entity

import (
	"time"
)

// Expense represents a single movement of type 'expense'.
type Expense struct {
	UserId      string    `firestore:"userID" json:"userId"`
	Category    string    `firestore:"category" binding:"required" json:"category"`
	Name        string    `firestore:"name" binding:"required" json:"name"`
	Description string    `firestore:"description" binding:"required" json:"description"`
	Amount      float32   `firestore:"amount" binding:"required" json:"amount"`
	Date        time.Time `firestore:"date" binding:"required" json:"date"`
}
