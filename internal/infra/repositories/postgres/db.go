package postgres

import (
	"fmt"
	"sync"
	"time"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var onceDBLoad sync.Once

var tables = []interface{}{
	&entity.User{},
	&entity.Card{},
}

func connect() *gorm.DB {
	onceDBLoad.Do(func() {
		source := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			"0.0.0.0",
			"root",
			"root",
			"financeDb",
			"5432",
		)
		var i int
		for {
			var err error
			if i >= 30 {
				panic("Failed to connect: " + source)
			}
			time.Sleep(3 * time.Second)
			db, err = gorm.Open(postgres.Open(source), &gorm.Config{})
			if err != nil {
				log.Info("Retrying connection...", err)
				i++
				continue
			}
			break
		}
		// Init migration. Create Tables
		migrate()
		// Success DB connection.
		log.Info("Connected DB!")
	})
	return db
}

func migrate() {
	dbName := db.Migrator().CurrentDatabase()
	if dbName == "financeDb" {
		for _, table := range tables {
			error := db.AutoMigrate(table)
			if error != nil {
				log.Panic("Auto migration field")
			}
		}
	}
}
