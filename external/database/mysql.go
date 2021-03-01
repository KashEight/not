package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMySQL() gorm.Dialector {
	// TODO: add mysql driver settings
	return mysql.Open("")
}
