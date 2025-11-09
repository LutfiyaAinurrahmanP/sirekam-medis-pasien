package handlers

import (
	"strconv"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/gofiber/fiber/v2"
)

type PatientHandler interface {
	CreatePatient(c *fiber.Ctx) error
	UpdatePatient(c *fiber.Ctx) error
	DeletePatient(c *fiber.Ctx) error
	HardDeletePatient(c *fiber.Ctx) error
	RestorePatient(c *fiber.Ctx) error

	GetByIDPatient(c *fiber.Ctx) error
	GetAllPatient(c *fiber.Ctx) error
	GetAllDeletePatient(c *fiber.Ctx) error
}

type patientHandler struct {
	patientService services.PatientService
}

func NewPatientHandler(patientService services.PatientService) PatientHandler {
	return &patientHandler{
		patientService: patientService,
	}
}

func (h *patientHandler) CreatePatient(c *fiber.Ctx) error {
	var req validators.CreatePatientRequest

	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	patient, err := h.patientService.CreatePatient(&req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "patient code already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to create patient")
	}
	return utils.SuccessResponse(c, "Patient created successfully", fiber.Map{
		"patient": patient,
	})
}

func (h *patientHandler) UpdatePatient(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid patient id", nil)
	}

	var req validators.UpdatePatientRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	patient, err := h.patientService.UpdatePatient(uint(id), &req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "patient code already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to update pasien")
	}

	return utils.SuccessResponse(c, "Patient updated successfully", fiber.Map{
		"patient": patient,
	})
}

func (h *patientHandler) DeletePatient(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid patient ID", nil)
	}

	if err := h.patientService.DeletePatient(uint(id)); err != nil {
		if err.Error() == "patient not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to delete patient")
	}
	return utils.SuccessResponse(c, "Patient deleted successfully", nil)
}

func (h *patientHandler) HardDeletePatient(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid patient ID", nil)
	}

	if err := h.patientService.HardDeletePatient(uint(id)); err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to permanently delete patient")
	}

	return utils.SuccessResponse(c, "Patient permanently delete successfully", nil)
}

func (h *patientHandler) RestorePatient(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (h *patientHandler) GetByIDPatient(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (h *patientHandler) GetAllPatient(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (h *patientHandler) GetAllDeletePatient(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
