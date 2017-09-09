package storage

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"mysql",
		"slawek:k1k2k3k4k5k6@tcp(localhost:6603)/goauth?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	return db, nil
}
