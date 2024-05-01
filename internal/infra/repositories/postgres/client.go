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

// WARNING: Improve paramas with DTO's
// Transactions
type DestinationParams struct {
	destination interface{}
	column      string
	value       float64
}
type SourceParams struct {
	source interface{}
	column string
	value  float64
}
type TransactionParams struct {
	transactionInit interface{}
	source          SourceParams
	destination     DestinationParams
}

// WARNING: We need to improve Errors
func (c *client) SendTransaction(params TransactionParams) error {

	transaction := c.db.Transaction(func(tx *gorm.DB) error {
		// 1.- Update balance Card SOURCE
		if err := tx.Model(&params.source.source).Update(params.source.column, params.source.value).Error; err != nil {
			// return any error will rollback
			log.Println("Error updating Card source balance")
			return err
		}

		// 2.- Update balance Card DESTINATION
		if err := tx.Model(&params.destination.destination).Update(params.destination.column, params.destination.value).Error; err != nil {
			// return any error will rollback
			log.Println("Error updating Card destination balance")
			return err
		}

		// 3.- Create new transaction
		if err := tx.Create(params.transactionInit).Error; err != nil {
			// return any error will rollback
			log.Println("Error creating transaction")
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return transaction
}
