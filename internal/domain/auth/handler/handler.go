package auth_handler

import (
	auth_dto "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/auth/dto"
	auth_usecase "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/auth/usecase"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Usecase auth_usecase.AuthUsecase
}

func NewAuthHandler(u auth_usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{Usecase: u}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req auth_dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	err := h.Usecase.Register(req.Username, req.Password, req.Role)
	if err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Register failed", err.Error())
	}
	resp := auth_dto.RegisterResponse{
		Username: req.Username,
		Role:     req.Role,
	}
	return response.WriteSuccess(c, fiber.StatusOK, "Register success", resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req auth_dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	token, err := h.Usecase.Login(req.Username, req.Password)
	if err != nil {
		return response.WriteError(c, fiber.StatusUnauthorized, "Login failed", err.Error())
	}
	resp := auth_dto.LoginResponse{Token: token}
	return response.WriteSuccess(c, fiber.StatusOK, "Login success", resp)
}
