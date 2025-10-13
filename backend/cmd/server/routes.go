package main

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/handlers"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/middlewares"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	AuthHandler *handlers.AuthHandler
	UserHandler *handlers.UserHandler
	JWTSecret   string
	TokenRepo   repositories.TokenRepository
}

// SetupRoutes mendaftarkan semua routes ke Fiber app
func SetupRoutes(app *fiber.App, config *RouteConfig) {
	// ============================================
	// ROOT & HEALTH CHECK ENDPOINTS
	// ============================================
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Welcome to User Management API",
			"version": "1.0.0",
			"status":  "running",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"status":  "healthy",
			"message": "Server is running properly",
		})
	})

	// ============================================
	// PUBLIC ROUTES - No Authentication Required
	// ============================================

	// Authentication routes (public)
	auth := app.Group("/auth")
	{
		// POST /auth/register - Public registration (default role: user)
		auth.Post("/register", config.AuthHandler.Register)

		// POST /auth/login - Login and get JWT token
		auth.Post("/login", config.AuthHandler.Login)

		// POST /auth/logout - Logout (client-side operation)
		auth.Post("/logout",
			middlewares.JWTAuthMiddleware(config.JWTSecret, config.TokenRepo),
			config.AuthHandler.Logout,
		)
	}

	// ============================================
	// ADMIN ROUTES - Admin Role Required
	// ============================================
	// Prefix: /admin
	// Middleware: JWT Authentication + Admin Role

	admin := app.Group("/admin")
	admin.Use(middlewares.JWTAuthMiddleware(config.JWTSecret, config.TokenRepo)) // Require authentication
	admin.Use(middlewares.RequireAdmin())                                        // Require admin role
	{
		// GET /admin/dashboard - Admin dashboard
		// TODO: Implement admin dashboard handler
		admin.Get("/dashboard", func(c *fiber.Ctx) error {
			// Get admin info from context
			username := middlewares.GetUsernameFromContext(c)
			userID := middlewares.GetUserIDFromContext(c)

			return c.JSON(fiber.Map{
				"success": true,
				"message": "Welcome to Admin Dashboard",
				"data": fiber.Map{
					"admin_id":       userID,
					"admin_username": username,
					"role":           "admin",
				},
			})
		})

		// User Management Routes (Admin)
		// Prefix: /admin/user
		user := admin.Group("/user")
		{
			// GET /admin/user - List all users with pagination, search, filter
			// Query params: page, limit, search, role, sort, sort_by
			user.Get("/", config.UserHandler.GetAllUsers)

			// POST /admin/user/create - Create new user (admin can choose role)
			user.Post("/create", config.UserHandler.CreateUser)

			// GET /admin/user/deleted - List all soft deleted users
			user.Get("/deleted", config.UserHandler.GetAllDeletedUsers)

			// GET /admin/user/:id - Get specific user by ID
			user.Get("/:id", config.UserHandler.GetUserByID)

			// PUT /admin/user/update/:id - Update user by ID
			user.Put("/update/:id", config.UserHandler.UpdateUser)

			// DELETE /admin/user/:id - Soft delete user by ID
			user.Delete("/:id", config.UserHandler.DeleteUser)

			// DELETE /admin/user/permanent/:id - Hard delete user (permanent)
			user.Delete("/permanent/:id", config.UserHandler.HardDeleteUser)

			// POST /admin/user/restore/:id - Restore soft deleted user
			user.Post("/restore/:id", config.UserHandler.RestoreUser)
		}

		// Future admin routes bisa ditambahkan di sini
		// admin.Get("/reports", config.ReportHandler.GetReports)
		// admin.Get("/settings", config.SettingHandler.GetSettings)
		// admin.Post("/settings", config.SettingHandler.UpdateSettings)
	}

	// ============================================
	// USER ROUTES - User Role Required
	// ============================================
	// Prefix: /user
	// Middleware: JWT Authentication + User Role

	userRoute := app.Group("/user")
	userRoute.Use(middlewares.JWTAuthMiddleware(config.JWTSecret, config.TokenRepo)) // Require authentication
	userRoute.Use(middlewares.RequireUser())                                         // Require user role
	{
		// GET /user/dashboard - User dashboard
		// TODO: Implement user dashboard handler
		userRoute.Get("/dashboard", func(c *fiber.Ctx) error {
			// Get user info from context
			username := middlewares.GetUsernameFromContext(c)
			userID := middlewares.GetUserIDFromContext(c)

			return c.JSON(fiber.Map{
				"success": true,
				"message": "Welcome to User Dashboard",
				"data": fiber.Map{
					"user_id":       userID,
					"user_username": username,
					"role":          "user",
				},
			})
		})

		// Profile routes
		profile := userRoute.Group("/profile")
		{
			// GET /user/profile - Get own profile
			profile.Get("/", config.UserHandler.GetProfile)

			// PUT /user/profile/update - Update own profile
			profile.Put("/update", config.UserHandler.UpdateProfile)

			// Future profile routes
			// profile.Put("/change-password", config.UserHandler.ChangePassword)
			// profile.Post("/avatar", config.UserHandler.UploadAvatar)
		}

		// Future user routes
		// userRoute.Get("/notifications", config.NotificationHandler.GetNotifications)

		// Future user routes bisa ditambahkan di sini
		// userRoute.Get("/notifications", config.NotificationHandler.GetNotifications)
		// userRoute.Get("/settings", config.UserSettingHandler.GetSettings)
	}

	// ============================================
	// 404 NOT FOUND HANDLER
	// ============================================
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Route not found",
			"path":    c.Path(),
			"method":  c.Method(),
		})
	})
}
