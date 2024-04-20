package handler

import (
	"context"
	"example.com/m/v2/backend/internal/models"
	"example.com/m/v2/backend/internal/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func New(s service.Service) *handler {
	return &handler{s: s}
}

type handler struct {
	s service.Service
}

// CreateUser godoc
// @Summary      CreateUser
// @Description  CreateUser
// @Accept       json
// @Produce      json
// @Param        CreateUser   body      models.CreateUserRequest  true "CreateUser"
// @Success      200  {object}  models.CreateUserResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.BadRequestResponse
// @Router       /create_user [post]
func (h *handler) CreateUser(ctx echo.Context) error {
	var request models.CreateUserRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	token, err := h.s.CreateUser(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, models.CreateUserResponse{Token: token})
}

func (h *handler) AddKeys(ctx echo.Context) error {
	// Проверить токен

	type in struct {
		Login   string `json:"login"`
		OpenKey string `json:"open_key"`
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *handler) GetChallenge(ctx echo.Context) error {
	type in struct {
		Login   string
		OpenKey string
	}

	type out struct {
		Challenge   string
		OperationID string
	}

	return nil
}

func (h *handler) SolveChallenge(ctx echo.Context) error {
	type in struct {
		OperationID string `json:"operation_id"`
		Solution    string
	}

	token := "228"

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *handler) GetAccess(ctx echo.Context) error {
	type in struct {
		Login string
	}

	type out struct {
		Number string `json:"number"`
	}

	q := out{}

	return ctx.JSON(http.StatusOK, q)
}

func (h *handler) OkAccess(ctx echo.Context) error {
	// Проверяем токен
	type in struct {
		Number string
	}

	return ctx.JSON(http.StatusCreated, nil)
}
