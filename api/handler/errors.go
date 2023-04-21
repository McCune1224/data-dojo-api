package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error             string `json:"error"`
	Error_Description string `json:"error_description"`
}

func NewErrorResponse(err string, desc string) ErrorResponse {
	return ErrorResponse{
		Error:             err,
		Error_Description: desc,
	}
}

var (
	// Used for returning a 400 status JSON response with error details
	error400 = ErrorResponse{
		Error:             "bad_request",
		Error_Description: "The request was malformed",
	}

	// Used for returning a 404 status JSON response with error details
	Error404 = ErrorResponse{
		Error:             "not_found",
		Error_Description: "The requested resource was not found",
	}

	// Used for returning a 500 status JSON response with error details
	Error500 = ErrorResponse{
		Error:             "internal_server_error",
		Error_Description: "An internal server error occurred",
	}
)

func handleNotFound(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(Error404)
	} else {
		return c.Status(500).JSON(Error500)
	}
}
