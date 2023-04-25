package router

import (
	"back-end/db"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo   *echo.Echo
	Sql    *db.Sql
	MongDB *db.MongoDB
}

func (api *API) SetupRouter() {
	LangRouter(api.Echo, api.Sql)
	LogRouter(api.Echo, api.MongDB)
}
