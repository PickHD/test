package helper

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   error  `json:"error,omitempty"`
}

func NewResponses[T any](ctx *fiber.Ctx, statusCode int, message string, data T, err error) error {
	if statusCode < 400 {
		return ctx.Status(statusCode).JSON(&BaseResponse{
			Message: message,
			Data:    data,
			Error:   nil,
		})
	}

	return ctx.Status(statusCode).JSON(&BaseResponse{
		Message: message,
		Data:    nil,
		Error:   err,
	})
}
