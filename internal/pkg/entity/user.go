package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents an application user.
type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	LastName  string
	Email     string `gorm:"unique"`
	Address   string
	Password  string
	CreatedAt *time.Time
	DeletedAt *time.Time
}

// Hooks
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}
