package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func newSQLite() gorm.Dialector {
	fn := os.Getenv("SQLITE_FILENAME")

	return sqlite.Open(fn)
}
