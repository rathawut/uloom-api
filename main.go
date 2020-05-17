package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"

	"uloom-api/config"
	v1 "uloom-api/controllers/v1"
	"uloom-api/db"
	"uloom-api/providers"
)

func main() {
	e := echo.New()

	c := config.New()
	d, err := db.New(c.DbHost, c.DbName, c.DbUser, c.DbPassword, 3*time.Minute)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(d)

	facebookProvider := providers.NewFacebookProvider(c.FacebookClientID, c.FacebookClientSecret, c.FacebookRedirectURL)

	v1Handler := v1.NewHandler(facebookProvider)
	v1Handler.Register(e.Group("/v1"))

	e.Logger.Fatal(e.Start(":80"))
}
