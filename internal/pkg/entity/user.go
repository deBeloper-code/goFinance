package entity

import (
	"time"
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
