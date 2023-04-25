package router

import (
	"back-end/db"
	"back-end/handler"
	"back-end/repository/repo_impl"

	"github.com/labstack/echo/v4"
)

func LogRouter(e *echo.Echo, mong *db.MongoDB) {
	handler := handler.LogHandler{
		LogRepo: repo_impl.NewLogRepo(mong),
	}

	e.GET("/logs", handler.HandlerAllLogs)

}
