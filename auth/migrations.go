package auth

import "github.com/slawek87/GOauth/storage"


func InitMigrations() {
	db, _ := storage.InitDB()
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&TokenHistory{})
	db.AutoMigrate(&Service{})
}

