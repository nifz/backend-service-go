package main

import (
	"backend-service-go/configs"
	"backend-service-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	routes.New(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
