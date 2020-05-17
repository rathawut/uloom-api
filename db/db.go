package db

import (
	"fmt"
	"uloom-api/models"

	"github.com/jinzhu/gorm"
)

// New create connection
func New(dbHost string, dbName string, dbUser string, dbPassword string) *gorm.DB {
	db, err := gorm.Open("postgres", "host="+dbHost+" dbname="+dbName+" user="+dbUser+" password="+dbPassword+" sslmode=disable")
	if err != nil {
		fmt.Println("db error:", err)
	}
	db.SingularTable(true)
	return db
}

// AutoMigrate do auto migrate for all models
// TODO: handle error
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
