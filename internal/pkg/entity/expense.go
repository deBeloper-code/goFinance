package entity

import (
	"time"
)

// Expense represents a single movement of type 'expense'.
type Expense struct {
	UserId      string    `firestore:"userID"`
	Category    string    `firestore:"category"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Amount      float32   `firestore:"amount"`
	Date        time.Time `firestore:"date"`
}
