package main

import (
	"net/http"
	"uloom-api/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	d := db.New()
	db.AutoMigrate(d)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":80"))
}
