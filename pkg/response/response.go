package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err string) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}

// WriteSuccess writes a success response to the fiber context
func WriteSuccess(c *fiber.Ctx, code int, message string, data any) error {
	return c.Status(code).JSON(SuccessResponse(message, data))
}

// WriteError writes an error response to the fiber context
func WriteError(c *fiber.Ctx, code int, message string, err string) error {
	return c.Status(code).JSON(ErrorResponse(message, err))
}
