package handler

import (
	"back-end/common/response"
	"back-end/repository"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LogHandler struct {
	LogRepo repository.LogRepo
}

func (l *LogHandler) HandlerAllLogs(c echo.Context) error {
	defer c.Request().Body.Close()
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	logs, err := l.LogRepo.GetAllLog(ctx)
	if err != nil {
		return c.JSON(http.StatusConflict, response.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       logs,
	})

}
