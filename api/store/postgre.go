package store

import (
	"github.com/mccune1224/data-dojo/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect to given database with dns string as well as a flag for auto-migration
// of model used in the API
func Connect(dns string, autoMigrate bool) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// Can't really run the API without a database connection
	if err != nil {
		return err
	}

	// Auto-migrate model
	if autoMigrate {
		err = DB.AutoMigrate(
			&model.Game{},
			&model.Character{},
			&model.Move{})
	}

	if err != nil {
		panic(err)
	}
	return nil
}
