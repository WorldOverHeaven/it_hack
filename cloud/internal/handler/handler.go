package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"mephi_hack/cloud/internal/modelscloud"
	"mephi_hack/cloud/internal/service"
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
// @Param        CreateUser   body      modelscloud.CreateUserRequest  true "CreateUser"
// @Success      200  {object}  modelscloud.CreateUserResponse
// @Failure      400  {object}  modelscloud.BadRequestResponse
// @Failure      500  {object}  modelscloud.BadRequestResponse
// @Router       /create_user [post]
func (h *handler) CreateUser(ctx echo.Context) error {
	var request modelscloud.CreateUserRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	token, err := h.s.CreateUser(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, modelscloud.CreateUserResponse{Token: token})
}

// AuthUser godoc
// @Summary      AuthUser
// @Description  AuthUser
// @Accept       json
// @Produce      json
// @Param        AuthUser   body      modelscloud.AuthUserRequest  true "AuthUser"
// @Success      200  {object}  modelscloud.AuthUserResponse
// @Failure      400  {object}  modelscloud.BadRequestResponse
// @Failure      500  {object}  modelscloud.BadRequestResponse
// @Router       /auth_user [post]
func (h *handler) AuthUser(ctx echo.Context) error {
	var request modelscloud.AuthUserRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	token, err := h.s.AuthUser(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, modelscloud.AuthUserResponse{Token: token})
}

// PutPayload godoc
// @Summary      PutPayload
// @Description  PutPayload
// @Accept       json
// @Produce      json
// @Param        PutPayload   body      modelscloud.PutPayloadRequest  true "PutPayload"
// @Success      200  {object}  modelscloud.PutPayloadResponse
// @Failure      400  {object}  modelscloud.BadRequestResponse
// @Failure      500  {object}  modelscloud.BadRequestResponse
// @Router       /put_payload [post]
func (h *handler) PutPayload(ctx echo.Context) error {
	var request modelscloud.PutPayloadRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	resp, err := h.s.PutPayload(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// GetPayload godoc
// @Summary      GetPayload
// @Description  GetPayload
// @Accept       json
// @Produce      json
// @Param        GetPayload   body      modelscloud.GetPayloadRequest  true "GetPayload"
// @Success      200  {object}  modelscloud.GetPayloadResponse
// @Failure      400  {object}  modelscloud.BadRequestResponse
// @Failure      500  {object}  modelscloud.BadRequestResponse
// @Router       /get_payload [post]
func (h *handler) GetPayload(ctx echo.Context) error {
	var request modelscloud.GetPayloadRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	resp, err := h.s.GetPayload(context.Background(), request)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, modelscloud.BadRequestResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, resp)
}
