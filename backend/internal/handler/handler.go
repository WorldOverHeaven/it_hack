package handler

import (
	"context"
	"example.com/m/v2/backend/internal/models"
	"example.com/m/v2/backend/internal/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
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

// GetChallenge godoc
// @Summary      GetChallenge
// @Description  GetChallenge
// @Accept       json
// @Produce      json
// @Param        GetChallenge   body      models.GetChallengeRequest  true "GetChallenge"
// @Success      200  {object}  models.GetChallengeResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.BadRequestResponse
// @Router       /get_challenge [post]
func (h *handler) GetChallenge(ctx echo.Context) error {
	var request models.GetChallengeRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	resp, err := h.s.GetChallenge(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, resp)
}

// SolveChallenge godoc
// @Summary      SolveChallenge
// @Description  SolveChallenge
// @Accept       json
// @Produce      json
// @Param        SolveChallenge   body      models.SolveChallengeRequest  true "SolveChallenge"
// @Success      200  {object}  models.SolveChallengeResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.BadRequestResponse
// @Router       /solve_challenge [post]
func (h *handler) SolveChallenge(ctx echo.Context) error {
	var request models.SolveChallengeRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	resp, err := h.s.SolveChallenge(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, resp)
}

func (h *handler) Verify(ctx echo.Context) error {
	// Нужен токен авторизации сервиса
	return nil
}

func (h *handler) AddKeys(ctx echo.Context) error {
	// Проверить токен

	type in struct {
		Login   string `json:"login"`
		OpenKey string `json:"open_key"`
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *handler) RegisterCloud(ctx echo.Context) error {
	return nil
}

func (h *handler) AuthCloud(ctx echo.Context) error {
	return nil
}

func (h *handler) GetContainer(ctx echo.Context) error {
	// Нужен токен авторизации облака
	return nil
}

func (h *handler) PutContainer(ctx echo.Context) error {
	// Нужен токен авторизации облака
	return nil
}
