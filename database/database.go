package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DBInstance struct {
    DB *gorm.DB
}

var DB DBInstance


func ConnectToDB()  {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
	log.Panicf("Failed to connect to database: %v", err)
    }

    DB = DBInstance{
	DB: db,
    }
}
