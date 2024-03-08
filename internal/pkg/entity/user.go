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
	Password  string
	CreatedAt *time.Time
	DeletedAt *time.Time
}

func (User) TableName() string {
	return "app.users"
}

// Hooks
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}
