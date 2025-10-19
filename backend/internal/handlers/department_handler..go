package handlers

import (
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
