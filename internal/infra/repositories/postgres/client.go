package postgres

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"gorm.io/gorm"
)

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

// Create card a new record in our DB.
func (c *client) AddCard(dest interface{}) error {
	card := dest.(*entity.Card)
	// First Looking for if userId exists
	var user entity.User
	if err := c.First(&user, "id = ?", card.UserID); err != nil {
		return err
	}
	// Assign the UserID to the card.
	card.UserID = user.ID
	// If the user exist you can create Card
	return c.db.Create(card).Error
}

// First find first record that match given conditions.
func (c *client) GetUserCard(cardID, accountID string, conds ...interface{}) ([]*entity.Card, error) {
	card := []*entity.Card{}
	return card, c.db.First(cardID, conds...).Error
}
