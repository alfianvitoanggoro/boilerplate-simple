package user_handler

import (
	user_dto "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/dto"
	user_usecase "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/usecase"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Usecase user_usecase.UserUsecase
}

func NewUserHandler(u user_usecase.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: u}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	// Example: get pagination params from query
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	users, meta, err := h.Usecase.GetAll(page, limit)
	if err != nil {
		return response.WriteError(c, fiber.StatusInternalServerError, "Failed to get users", err.Error())
	}
	return response.WriteSuccessWithMeta(c, fiber.StatusOK, "Get all users success", users, meta)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.Usecase.GetByID(id)
	if err != nil {
		return response.WriteError(c, fiber.StatusNotFound, "User not found", err.Error())
	}
	return response.WriteSuccess(c, fiber.StatusOK, "Get user success", user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req user_dto.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	user, err := h.Usecase.Create(req)
	if err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Create user failed", err.Error())
	}
	return response.WriteSuccess(c, fiber.StatusCreated, "Create user success", user)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req user_dto.UserUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}
	user, err := h.Usecase.Update(id, req)
	if err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Update user failed", err.Error())
	}
	return response.WriteSuccess(c, fiber.StatusOK, "Update user success", user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.Usecase.Delete(id)
	if err != nil {
		return response.WriteError(c, fiber.StatusBadRequest, "Delete user failed", err.Error())
	}
	return response.WriteSuccess(c, fiber.StatusOK, "Delete user success", nil)
}
