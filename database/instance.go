package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func InitDatabase() error {
	var err error
	Instance, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	err = Instance.AutoMigrate(&SyncFolder{}, &User{})
	if err != nil {
		return err
	}
	return nil
}

func InitDefaultUser() {

}
