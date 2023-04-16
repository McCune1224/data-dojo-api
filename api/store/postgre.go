package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect to given database with dns string as well as a flag for auto-migration
// of models used in the API
func Connect(dns string, autoMigrate bool) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// Can't really run the API without a database connection
	if err != nil {
		return err
	}

	// Auto-migrate models
	if autoMigrate {
	}

	if err != nil {
		panic(err)
	}
	return nil
}
