package db

import (
	"fmt"
	"os"
	"uloom-api/models"

	"github.com/jinzhu/gorm"
)

// New create connection
func New() *gorm.DB {
	var dbHost string
	var dbName string
	var dbUser string
	var dbPassword string
	var exists bool
	if dbHost, exists = os.LookupEnv("DB_HOST"); exists == false {
		fmt.Println("db dbHost does not exist.")
	}
	if dbName, exists = os.LookupEnv("DB_NAME"); exists == false {
		fmt.Println("db dbName does not exist.")
	}
	if dbUser, exists = os.LookupEnv("DB_USER"); exists == false {
		fmt.Println("db dbUser does not exist.")
	}
	if dbPassword, exists = os.LookupEnv("DB_PASSWORD"); exists == false {
		fmt.Println("db dbPassword does not exist.")
	}
	db, err := gorm.Open("postgres", "host="+dbHost+" dbname="+dbName+" user="+dbUser+" password="+dbPassword+" sslmode=disable")
	if err != nil {
		fmt.Println("db error:", err)
	}
	db.SingularTable(true)
	return db
}

// AutoMigrate do migration
// TODO: handle error
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
