package db

import (
	"fmt"
	"time"
	"uloom-api/models"

	"github.com/jinzhu/gorm"
)

// New db connection
func New(dbHost string, dbName string, dbUser string, dbPassword string, timeout time.Duration) (*gorm.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)
		case <-ticker.C:
			db, err := gorm.Open("postgres", "host="+dbHost+" dbname="+dbName+" user="+dbUser+" password="+dbPassword+" sslmode=disable")
			if err == nil {
				db.SingularTable(true)
				return db, nil
			}
			fmt.Println("db connection failed, try again, error:", err)
		}
	}
}

// AutoMigrate migrate for all models
// TODO: handle error
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
