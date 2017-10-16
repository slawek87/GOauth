package storage

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"mysql",
		"slawek:k1k2k3k4k5k6@tcp(localhost:3306)/goauth?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}


func InitTestDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/test_gorm.db")

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}