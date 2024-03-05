package entity

import (
	"time"
)

// Deposit represents a movement made by user of type 'expense'.
type Deposit struct {
	UserId      string    `firestore:"userID"`
	Category    string    `firestore:"category"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Amount      float32   `firestore:"amount"`
	Date        time.Time `firestore:"date"`
}
