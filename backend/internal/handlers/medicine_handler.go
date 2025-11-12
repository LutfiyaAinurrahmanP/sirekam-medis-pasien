package handlers

import (
	"strconv"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/gofiber/fiber/v2"
)

type MedicineHandler interface {
	CreateMedicine(c *fiber.Ctx) error
	UpdateMedicine(c *fiber.Ctx) error
	DeleteMedicine(c *fiber.Ctx) error
	HardDeleteMedicine(c *fiber.Ctx) error
	RestoreMedicine(c *fiber.Ctx) error

	GetByIDMedicine(c *fiber.Ctx) error
	GetAllMedicine(c *fiber.Ctx) error
	GetAllDeleteMedicine(c *fiber.Ctx) error
}

type medicineHandler struct {
	medicineService services.MedicineService
}

func NewMedicineHandler(medicineService services.MedicineService) MedicineHandler {
	return &medicineHandler{
		medicineService: medicineService,
	}
}

func (h *medicineHandler) CreateMedicine(c *fiber.Ctx) error {
	var req validators.CreateMedicineRequest

	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	medicine, err := h.medicineService.CreateMedicine(&req)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed create medicine")
	}
	return utils.SuccessResponse(c, "Medicine created successfully", fiber.Map{
		"medicine": medicine,
	})
}

func (h *medicineHandler) UpdateMedicine(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid medicine ID: %w", err)
	}

	var req validators.UpdateMedicineRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	medicine, err := h.medicineService.UpdateMedicine(uint(id), &req)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to updated medicine")
	}
	return utils.SuccessResponse(c, "Medicine update successfully", fiber.Map{
		"medicine": medicine,
	})
}

func (h *medicineHandler) DeleteMedicine(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid medicine ID", nil)
	}

	if err := h.medicineService.DeleteMedicine(uint(id)); err != nil {
		if err.Error() == "medicine not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to delete medicine")
	}

	return utils.SuccessResponse(c, "Medicine delete successfully", nil)
}

func (h *medicineHandler) HardDeleteMedicine(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) RestoreMedicine(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) GetByIDMedicine(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32);
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid medicine ID: %w", err)
	}

	medicine, err := h.medicineService.GetMedicineByID(uint(id))
	if err != nil {
		if err.Error() == "medicine not found"{
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to fetch medicine")
	}

	return utils.SuccessResponse(c, "Medicine retrieved successfully", fiber.Map{
		"medicine": medicine,
	})
}

func (h *medicineHandler) GetAllMedicine(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) GetAllDeleteMedicine(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
