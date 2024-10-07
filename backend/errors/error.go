package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) AppError {
	return AppError{StatusCode: code, Message: message}
}

type ErrorCategory struct {
	StatusCode int
}

func (ec ErrorCategory) Error(message string) AppError {
	return NewAppError(ec.StatusCode, message)
}

var (
	NotFound       = ErrorCategory{StatusCode: http.StatusNotFound}
	AlreadyExists  = ErrorCategory{StatusCode: http.StatusConflict}
	BadRequest     = ErrorCategory{StatusCode: http.StatusBadRequest}
	InternalServer = ErrorCategory{StatusCode: http.StatusInternalServerError}
	Unauthorized   = ErrorCategory{StatusCode: http.StatusUnauthorized}
	Forbidden      = ErrorCategory{StatusCode: http.StatusForbidden}
)

var (
	ErrInvalidID          = BadRequest.Error("Invalid ID")
	ErrInvalidRequestBody = BadRequest.Error("Invalid request body")
	ErrInternalServer     = InternalServer.Error("Internal server error")
)

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func NotFoundError(resourceName string) AppError {
	return NotFound.Error(fmt.Sprintf("%s not found", capitalizeFirst(resourceName)))
}

func AlreadyExistsError(resourceName string) AppError {
	return AlreadyExists.Error(fmt.Sprintf("%s already exists", capitalizeFirst(resourceName)))
}

func CreateFailedError(resourceName string) AppError {
	return InternalServer.Error(fmt.Sprintf("Failed to create %s", resourceName))
}

func UpdateFailedError(resourceName string) AppError {
	return InternalServer.Error(fmt.Sprintf("Failed to update %s", resourceName))
}

func DeleteFailedError(resourceName string) AppError {
	return InternalServer.Error(fmt.Sprintf("Failed to delete %s", resourceName))
}

func ParseFailedError(resourceName string) AppError {
	return BadRequest.Error(fmt.Sprintf("Failed to parse %s request body", resourceName))
}

func ValidateIdFailedError(resourceName string) AppError {
	return BadRequest.Error(fmt.Sprintf("Failed to validate %s id", resourceName))
}

func InternalServerError(err string) AppError {
	return InternalServer.Error("Internal server error" + err)
}

func HandleError(c *fiber.Ctx, err error) error {
	var appErr AppError
	if errors.As(err, &appErr) {
		return c.Status(appErr.StatusCode).JSON(fiber.Map{"error": appErr.Message})
	}
	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
}
