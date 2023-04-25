package handler

import (
	"back-end/common/response"
	"back-end/repository"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LangHandler struct {
	LangRepo repository.LangRepo
}

func (l *LangHandler) HandlerAllLangs(c echo.Context) error {
	defer c.Request().Body.Close()
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	limit := c.Param("limit")
	defer cancel()
	if len(limit) == 0 {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Thiếu limit",
			Data:       nil,
		})
	}
	posts, err := l.LangRepo.GetAllLangs(ctx, limit)
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
		Data:       posts,
	})
}
