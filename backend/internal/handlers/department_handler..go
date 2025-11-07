package handlers

import (
	"strconv"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/gofiber/fiber/v2"
)

type DepartmentHandler struct {
	departmentService services.DepartmentService
}

func NewDepartmentHandler(departmentService services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: departmentService,
	}
}

func (h *DepartmentHandler) CreateDepartment(c *fiber.Ctx) error {
	var req validators.CreateDepartmentRequest

	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	department, err := h.departmentService.CreateDepartment(&req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "code already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to create department")
	}
	return utils.CreatedResponse(c, "Department created successfully", fiber.Map{
		"department": department,
	})
}

func (h *DepartmentHandler) UpdateDepartment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid department id", nil)
	}

	var req validators.UpdateDepartmentRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	department, err := h.departmentService.UpdateDepartment(uint(id), &req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "code already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to update department")
	}
	return utils.SuccessResponse(c, "Department updated successfully", fiber.Map{
		"department": department,
	})
}

func (h *DepartmentHandler) GetAllDepartment(c *fiber.Ctx) error {
	var query validators.ListDepartmentQuery

	if err := c.QueryParser(&query); err != nil {
		return utils.BadRequestResponse(c, "Invalid query parameters", nil)
	}

	if err := validators.ValidateStruct(&query); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
	}

	departments, meta, err := h.departmentService.GetAllDepartments(&query)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch department")
	}

	return utils.PaginatedSeccessResponse(c, "Department retrieved successfully", fiber.Map{
		"department": departments,
	}, meta)
}

func (h *DepartmentHandler) DeleteDepartment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid department ID", nil)
	}

	if err := h.departmentService.DeleteDepartment(uint(id)); err != nil {
		if err.Error() == "department not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to permanetly delete department")
	}

	return utils.SuccessResponse(c, "Department deleted successfully", nil)
}

func (h *DepartmentHandler) HardDeleteDepartment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid department ID", nil)
	}

	if err := h.departmentService.HardDeleteDepartment(uint(id)); err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to permanently delete department")
	}

	return utils.SuccessResponse(c, "Department permanently deleted", nil)
}

func (h *DepartmentHandler) RestoreDepartment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid department ID", nil)
	}
	if err := h.departmentService.RestoreDepartment(uint(id)); err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to restore department")
	}
	return utils.SuccessResponse(c, "Department restore successfully", nil)
}

func (h *DepartmentHandler) GetDepartmentByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid department ID", nil)
	}
	department, err := h.departmentService.GetDepartmentByID(uint(id))
	if err != nil {
		if err.Error() == "department not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to fetch department")
	}

	return utils.SuccessResponse(c, "Department retrieved successfully", fiber.Map{
		"department": department,
	})
}

func (h *DepartmentHandler) GetAllDeletedDepartment(c *fiber.Ctx) error {
	var query validators.ListDepartmentQuery

	if err := c.QueryParser(&query); err != nil {
		return utils.BadRequestResponse(c, "Invalid query parameters", nil)
	}

	if err := validators.ValidateStruct(&query); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
	}

	departments, meta, err := h.departmentService.GetAllDeletedDepartments(&query)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch department")
	}

	return utils.PaginatedSeccessResponse(c, "Department retrieved successfully", fiber.Map{
		"department": departments,
	}, meta)
}
