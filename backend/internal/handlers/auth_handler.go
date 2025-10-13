package handlers

import (
	"errors"
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/middlewares"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req validators.RegisterRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		errorMessage := err.Error()
		if errorMessage == "username already exists" ||
			errorMessage == "email already exists" ||
			errorMessage == "phone already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to register user")
	}
	return utils.CreatedResponse(c, "User registered successfully", fiber.Map{
		"user": user,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req validators.LoginRequest

	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	token, user, err := h.authService.Login(&req)
	if err != nil {
		errorMessage := err.Error()
		if errorMessage == "invalid username or password" {
			return utils.UnauthorizedResponse(c, errorMessage)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.UnauthorizedResponse(c, "Invalid username or password")
		}
		return utils.InternalServerErrorResponse(c, "Failed to login")
	}
	return utils.SuccessResponse(c, "Login succesful", fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// 1. Get token dari Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.UnauthorizedResponse(c, "Missing authorization header")
	}

	// 2. Extract token dari "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return utils.UnauthorizedResponse(c, "Invalid authorization header format")
	}
	token := parts[1]

	// 3. Get user ID dari context (di-set oleh JWT middleware)
	userID := middlewares.GetUserIDFromContext(c)

	// Debug: Log untuk debugging
	// Uncomment baris di bawah untuk debugging
	// fmt.Printf("DEBUG Logout - UserID from context: %d\n", userID)

	if userID == 0 {
		// Jika userID masih 0, berarti middleware tidak jalan atau ada masalah
		return utils.UnauthorizedResponse(c, "Invalid user session. Please ensure you are logged in.")
	}

	// 4. Call service untuk logout (blacklist token)
	if err := h.authService.Logout(token, userID); err != nil {
		// Log error untuk debugging
		// fmt.Printf("DEBUG Logout Error: %v\n", err)
		return utils.InternalServerErrorResponse(c, "Failed to logout: "+err.Error())
	}

	// 5. Response sukses
	return utils.SuccessResponse(c, "Logout successful. Token has been revoked.", nil)
}
