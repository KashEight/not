package database

import "gorm.io/gorm"

func NewDatabase(isDev bool) gorm.Dialector {
	if isDev {
		return newSQLite()
	} else {
		return newMySQL()
	}
}
