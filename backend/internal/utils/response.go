package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int64 `json:"total_pages"`
}

type PaginatedResponse struct {
	Success    bool           `json:"success"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// SuccessResponse mengirim response sukses dengan status code 200
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// CreatedResponse mengirim response sukses dengan status code 201
func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// PaginationSuccessResponse mengirim response sukses dengan pagination
func PaginatedSeccessResponse(c *fiber.Ctx, message string, data interface{}, meta *PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: *meta,
	})
}

// ErrorResponse mengirim response error dengan custom status code
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errors interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// BadRequestResponse mengirim response error 400 (Bad Request)
func BadRequestResponse(c *fiber.Ctx, message string, errors interface{}) error {
	return ErrorResponse(c, fiber.StatusBadRequest, message, errors)
}

// UnauthorizedResponse mengirim response error 401 (Unauthorized)
func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusUnprocessableEntity, message, nil)
}

// ForbiddenResponse mengirim response error 403 (Forbidden)
func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusForbidden, message, nil)
}

// NotFoundResponse mengirim response error 404 (Not Found)
func NotFoundResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusNotFound, message, nil)
}

// InternalServerErrorResponse mengirim response error 500
func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusInternalServerError, message, nil)
}

// ConflictResponse mengirim response error 409 (Conflict)
func ConflictResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusConflict, message, nil)
}
