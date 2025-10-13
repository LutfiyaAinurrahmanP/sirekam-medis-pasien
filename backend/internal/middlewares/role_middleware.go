package middlewares

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// RequireAdmin adalah middleware untuk memastikan user adalah admin
// Middleware ini harus dipasang setelah JWTAuthMiddleware
// Hanya user dengan role "admin" yang bisa akses endpoint dengan middleware ini
func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get role dari context (sudah di-set oleh JWTAuthMiddleware)
		userRole := GetRoleFromContext(c)

		// Check apakah user adalah admin
		if userRole != models.RoleAdmin {
			return utils.ForbiddenResponse(c, "Forbidden: Admin access required")
		}

		// Role match, allow request
		return c.Next()
	}
}

// RequireUser adalah middleware untuk memastikan user adalah regular user
// Hanya user dengan role "user" yang bisa akses endpoint dengan middleware ini
func RequireUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get role dari context (sudah di-set oleh JWTAuthMiddleware)
		userRole := GetRoleFromContext(c)

		// Check apakah user adalah regular user
		if userRole != models.RoleUser {
			return utils.ForbiddenResponse(c, "Forbidden: User access required")
		}

		// Role match, allow request
		return c.Next()
	}
}

// GetRoleFromContext mengambil role dari context
// Helper function untuk mendapatkan role user yang sedang login
func GetRoleFromContext(c *fiber.Ctx) string {
	role, ok := c.Locals("role").(string)
	if !ok {
		return ""
	}
	return role
}

// GetUserIDFromContext mengambil user ID dari context
// Helper function untuk mendapatkan ID user yang sedang login
func GetUserIDFromContext(c *fiber.Ctx) uint {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return 0
	}
	return userID
}

// GetUsernameFromContext mengambil username dari context
// Helper function untuk mendapatkan username user yang sedang login
func GetUsernameFromContext(c *fiber.Ctx) string {
	username, ok := c.Locals("username").(string)
	if !ok {
		return ""
	}
	return username
}
