package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var serverdb *gorm.DB

func OpenDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/labcontrol-server.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database! " + err.Error())
	}

	db.AutoMigrate(&Agent{})

	serverdb = db

	return db
}

func GetServerDB() *gorm.DB {
	if serverdb == nil {
		return nil
	}
	return serverdb
}
