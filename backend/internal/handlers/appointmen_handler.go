package handlers

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentService services.AppointmentService
}

func NewAppointmentHandler(appointmentService services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

func (h *AppointmentHandler) CreateAppointment(c *fiber.Ctx) error {
	var req validators.CreateAppointmentRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	appointment, err := h.appointmentService.CreateAppointment(&req)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to create appointment")
	}

	return utils.SuccessResponse(c, "Appointment created successfully", fiber.Map{
		"appointment": appointment,
	})
}
