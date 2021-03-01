package external

import (
	"github.com/KashEight/not/external/database"
	"gorm.io/gorm"
	"os"
)

func isDev() bool {
	mode := os.Getenv("MODE")

	if mode == "production" {
		return false
	}

	return true
}

func DBInit() *gorm.DB {
	dialector := database.NewDatabase(isDev())

	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
