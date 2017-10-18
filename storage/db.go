package storage

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/slawek87/GOauth/settings"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		settings.Settings.Get("DATABASE_DRIVER"),
		settings.Settings.Get("DATABASE_ARGS"),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
