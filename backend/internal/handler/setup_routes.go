package handler

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo, h *handler) {
	e.POST("create_user", h.CreateUser)
}
