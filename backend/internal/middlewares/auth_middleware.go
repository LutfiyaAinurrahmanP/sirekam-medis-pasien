package middlewares

import (
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(jwtSecret string, tokenRepo repositories.TokenRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Get token dari Authorization header
		// Format: "Bearer <token>"
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.UnauthorizedResponse(c, "Missing authorization header")
		}

		// 2. Extract token dari "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.UnauthorizedResponse(c, "Invalid authorization header format")
		}

		tokenString := parts[1]

		// Check token blacklist (sudah logout)
		isBlacklisted, err := tokenRepo.IsBlacklisted(tokenString)
		if err != nil {
			return utils.InternalServerErrorResponse(c, "Failed to verify token")
		}

		if isBlacklisted {
			return utils.UnauthorizedResponse(c, "Token has been revoked. Please login again.")
		}

		// 4. Parse dan validate token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return []byte(jwtSecret), nil
		})

		// 5. Check parsing error
		if err != nil {
			return utils.UnauthorizedResponse(c, "Invalid or expired token")
		}

		// 6. Validate token
		if !token.Valid {
			return utils.UnauthorizedResponse(c, "Invalid token")
		}

		// 7. Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.UnauthorizedResponse(c, "Invalid tokenn claims")
		}

		// 8. Set user info ke context untuk digunakan di handler
		c.Locals("userID", uint(claims["sub"].(float64)))
		c.Locals("username", claims["username"].(string))
		c.Locals("role", claims["role"].(string))

		// 9. Continue ke handler berikutnya
		return c.Next()
	}
}
