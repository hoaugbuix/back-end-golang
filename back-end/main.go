package main

import (
	"back-end/db"
	"back-end/router"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "@hoangbb98",
		DbName:   "test",
	}
	sql.Connect()
	defer sql.Close()

	mongo := &db.MongoDB{
		Url: "mongodb+srv://hoangbui:hoangbb98@atlascluster.6rkxpwx.mongodb.net/?retryWrites=true&w=majority",
	}
	mongo.ConnectMG()
	defer mongo.CloseMG()

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	api := router.API{
		Echo:   e,
		Sql:    sql,
		MongDB: mongo,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
