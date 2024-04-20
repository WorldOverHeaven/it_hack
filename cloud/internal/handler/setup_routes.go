package handler

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo, h *handler) {
	e.POST("create_user", h.CreateUser)
	e.POST("auth_user", h.AuthUser)
	e.POST("put_payload", h.PutPayload)
	e.POST("get_payload", h.GetPayload)
}
