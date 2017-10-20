package utils

import (
	"github.com/slawek87/GOauth/settings"
	"os"
)

const (
	DATABSE_DRIVER = "sqlite3"
	DATABASE_ARGS = "/tmp/test_goauth.db"
)

func InitTestDB() {
	settings.Settings.Set("DATABASE_DRIVER", DATABSE_DRIVER)
	settings.Settings.Set("DATABASE_ARGS", DATABASE_ARGS)
}

func CleanTestDB() {
	err := os.Remove(DATABASE_ARGS)

	if err != nil {
		panic(err)
		return
	}
}