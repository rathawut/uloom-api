package db

import (
	"fmt"
	"uloom-api/models"

	"github.com/jinzhu/gorm"
)

// New create connection
func New() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db user=uloom password=uloompassword dbname=uloom sslmode=disable")
	if err != nil {
		fmt.Println("db error: ", err)
	}
	return db
}

// AutoMigrate do migration
// TODO: handle error
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
