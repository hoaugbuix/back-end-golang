package router

import (
	"back-end/db"
	"back-end/handler"
	"back-end/repository/repo_impl"

	"github.com/labstack/echo/v4"
)

func LangRouter(e *echo.Echo, sql *db.Sql) {
	handler := handler.LangHandler{
		LangRepo: repo_impl.NewLangRepo(sql),
	}

	e.GET("/langs", handler.HandlerAllLangs)

}
