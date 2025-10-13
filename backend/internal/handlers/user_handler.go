package handlers

import (
	"strconv"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/middlewares"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req validators.CreateUserRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "username already exists" ||
			errorMessage == "email already exists" ||
			errorMessage == "phone already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}

		if errorMessage == "invalid role" {
			return utils.BadRequestResponse(c, errorMessage, nil)
		}

		return utils.InternalServerErrorResponse(c, "Failed to create user")
	}

	return utils.CreatedResponse(c, "User created successfully", fiber.Map{
		"user": user,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user Id", nil)
	}

	var req validators.UpdateUserRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "user not found" {
			return utils.NotFoundResponse(c, errorMessage)
		}

		if errorMessage == "username already exists" ||
			errorMessage == "email already exists" ||
			errorMessage == "phone already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}

		if errorMessage == "invalid role" {
			return utils.BadRequestResponse(c, errorMessage, nil)
		}
		return utils.InternalServerErrorResponse(c, "Failed to update user")
	}

	return utils.SuccessResponse(c, "User updated successfully", fiber.Map{
		"user": user,
	})
}
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID", nil)
	}

	currentUserID := middlewares.GetUserIDFromContext(c)
	if currentUserID == uint(id) {
		return utils.BadRequestResponse(c, "You cannot delete your own account", nil)
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		if err.Error() == "user not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, err.Error())
	}
	return utils.SuccessResponse(c, "User deleted successfully", nil)
}

func (h *UserHandler) HardDeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID", nil)
	}

	currentUserID := middlewares.GetUserIDFromContext(c)
	if currentUserID == uint(id) {
		return utils.BadRequestResponse(c, "You cannot delete your own account", nil)
	}

	if err := h.userService.HardDeleteUser(uint(id)); err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to permanetly delete user")
	}

	return utils.SuccessResponse(c, "User permanently deleted", nil)
}

func (h *UserHandler) RestoreUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID", nil)
	}

	if err := h.userService.RestoreUser(uint(id)); err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to restore user")
	}

	return utils.SuccessResponse(c, "User restored successfully", nil)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user Id", nil)
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to fetch user")
	}

	return utils.SuccessResponse(c, "User retrieved successfully", fiber.Map{
		"user": user,
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var query validators.ListUserQuery

	if err := c.QueryParser(&query); err != nil {
		return utils.BadRequestResponse(c, "Invalid query parameters", nil)
	}

	if err := validators.ValidateStruct(&query); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
	}

	users, meta, err := h.userService.GetAllUsers(&query)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch users")
	}

	return utils.PaginatedSeccessResponse(c, "User retrieved successfully", fiber.Map{
		"users": users,
	}, meta)
}

func (h *UserHandler) GetAllDeletedUsers(c *fiber.Ctx) error {
	var query validators.ListUserQuery

	if err := c.QueryParser(&query); err != nil {
		return utils.BadRequestResponse(c, "Invalid query parameters", nil)
	}

	if err := validators.ValidateStruct(&query); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
	}

	users, meta, err := h.userService.GetAllDeletedUsers(&query)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch deleted users")
	}

	return utils.PaginatedSeccessResponse(c, "Deleted users retrieved successfully", fiber.Map{
		"users": users,
	}, meta)
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userId := middlewares.GetUserIDFromContext(c)

	user, err := h.userService.GetProfile(userId)
	if err != nil {
		if err.Error() == "user not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to fetch profile")
	}

	return utils.SuccessResponse(c, "Profile retrieved successfully", fiber.Map{
		"profile": user,
	})
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userId := middlewares.GetUserIDFromContext(c)
	var req validators.UpdateProfileRequest

	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	user, err := h.userService.UpdateProfile(userId, &req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "user not found" {
			return utils.NotFoundResponse(c, errorMessage)
		}
		if errorMessage == "username already exists" ||
			errorMessage == "email already exists" ||
			errorMessage == "phone already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to update profile")
	}
	return utils.SuccessResponse(c, "Profile updated successfully", fiber.Map{
		"profile": user,
	})
}
