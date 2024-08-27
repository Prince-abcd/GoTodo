package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)
var DB *gorm.DB
func ConnectDatabase()  {
	var err error
    DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database!", err)
    }
	err = DB.AutoMigrate(&Todo{})
    if err != nil {
        log.Fatal("Failed to migrate the database!", err)
    }

}