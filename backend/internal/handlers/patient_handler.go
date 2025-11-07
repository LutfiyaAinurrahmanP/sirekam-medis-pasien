package handlers

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
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

func NewPatientHandler(patientService services.PatientService) PatientHandler{
	return &patientHandler{
		patientService: patientService,
	}
}

func (h *patientHandler) CreatePatient(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *patientHandler) UpdatePatient(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *patientHandler) DeletePatient(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *patientHandler) HardDeletePatient(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
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