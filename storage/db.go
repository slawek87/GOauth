package storage

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/slawek87/GOauth/settings"
)

func InitDB() (*gorm.DB, error) {
	mySettings := settings.Settings()

	db, err := gorm.Open(
		mySettings.Get("DATABASE_DRIVER"),
		mySettings.Get("DATABASE_ARGS"),
	)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}


//func InitTestDB() (*gorm.DB, error) {
//	db, err := gorm.Open("sqlite3", "/tmp/test_gorm.db")
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer db.Close()
//
//	return db, nil
//}