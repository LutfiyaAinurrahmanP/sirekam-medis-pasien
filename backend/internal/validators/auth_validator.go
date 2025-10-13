package validators

import (
	"fmt"
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,min=10,max=15"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Role            string `json:"role" validate:"omitempty,oneof=user admin"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			var message string
			field := strings.ToLower(e.Field())

			switch e.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", field)
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", field)
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters", field, e.Param())
			case "max":
				message = fmt.Sprintf("%s must not exceed %s characters", field, e.Param())
			case "eqfield":
				message = fmt.Sprintf("%s must match %s", field, strings.ToLower(e.Param()))
			case "oneof":
				message = fmt.Sprintf("%s must be one of: %s", field, e.Param())
			default:
				message = fmt.Sprintf("%s is invalid", field)
			}

			errors[field] = message
		}
	}
	return errors
}

func ParseAndValidate(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := ValidateStruct(data); err != nil {
		return err
	}

	return nil
}

func (r *RegisterRequest) SetDefaultRole() {
	if r.Role == "" {
		r.Role = models.RoleUser
	}
}
