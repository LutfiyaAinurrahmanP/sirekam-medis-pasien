package handlers

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
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
	return &medicineHandler {
		medicineService: medicineService,
	}
}

func (h *medicineHandler) CreateMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) UpdateMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) DeleteMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) HardDeleteMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) RestoreMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) GetByIDMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) GetAllMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *medicineHandler) GetAllDeleteMedicine(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}