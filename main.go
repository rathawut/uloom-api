package main

import (
	"uloom-api/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"

	v1 "uloom-api/controllers/v1"
)

func main() {
	e := echo.New()

	d := db.New()
	db.AutoMigrate(d)

	v1Handler := v1.NewHandler()
	v1Handler.Register(e.Group("/v1"))

	e.Logger.Fatal(e.Start(":80"))
}
