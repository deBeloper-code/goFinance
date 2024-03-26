package postgres

import "gorm.io/gorm"

type client struct {
	db *gorm.DB
}

// NewClient returns a new instance to use postgres.
func NewClient() *client {
	return &client{
		db: connect(),
	}
}

// Create stores a new record in our DB.
func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}

// First find first record that match given conditions.
func (c *client) First(dest interface{}, conds ...interface{}) error {
	return c.db.First(dest, conds...).Error
}

// Create Bank Account
func (c *client) Add(value interface{}) error {
	return c.db.Create(value).Error
}

// Get Bank Account
func (c *client) GetUserAccount(value interface{}, conds ...interface{}) error {
	return c.db.First(value, conds...).Error
}
