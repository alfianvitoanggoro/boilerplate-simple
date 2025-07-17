package delivery

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/dto"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/usecase"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{Usecase: u}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	err := h.Usecase.Register(req.Username, req.Password, req.Role)
	if err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Register failed", err.Error())
	}
	resp := dto.RegisterResponse{
		Username: req.Username,
		Role:     req.Role,
	}
	return response.WriteSuccess(c, fiber.StatusOK, "Register success", resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	token, err := h.Usecase.Login(req.Username, req.Password)
	if err != nil {
		return response.WriteError(c, fiber.StatusUnauthorized, "Login failed", err.Error())
	}
	resp := dto.LoginResponse{Token: token}
	return response.WriteSuccess(c, fiber.StatusOK, "Login success", resp)
}
