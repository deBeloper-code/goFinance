package postgres

import (
	"log"

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
func (c *client) AddCard(dest interface{}, value interface{}, conds ...interface{}) error {
	// First Looking for user exists
	if err := c.First(&value, conds...); err != nil {
		return err
	}
	// If the user exist you can create Card
	return c.db.Create(dest).Error
}

// Find gets records (Slice []) that match given conditions.
func (c *client) Find(dest interface{}, conds ...interface{}) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	// Find results
	if err := c.db.Model(dest).Find(&results, conds...).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// TODO: We need to investigate about Transaction (Best Practice)
func (c *client) SendTransaction(transactionInit interface{}, source interface{}, destination interface{}, condsSource []interface{}, condsDestination []interface{}) error {

	transaction := c.db.Transaction(func(tx *gorm.DB) error {
		// 1.- Find Source ID exists
		if err := tx.First(&source, condsSource...).Error; err != nil {
			// return any error will rollback
			log.Println("Error finding Source ID")
			return err
		}

		// 2.- Find Destination ID exists
		if err := tx.First(&destination, condsDestination...).Error; err != nil {
			// return any error will rollback
			log.Println("Error finding destination ID")
			return err
		}
		// 3.- Get balance source

		// return nil will commit the whole transaction
		return nil
	})

	return transaction
}
