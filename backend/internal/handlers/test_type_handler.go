package handlers

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/gofiber/fiber/v2"
)

type TestTypeHandler interface {
	CreateTestType(c *fiber.Ctx) error
	UpdateTestType(c *fiber.Ctx) error
	DeleteTestType(c *fiber.Ctx) error
	HardDeleteTestType(c *fiber.Ctx) error
	RestoreTestType(c *fiber.Ctx) error

	GetTestTypeByID(c *fiber.Ctx) error
	GetAllTestType(c *fiber.Ctx) error
	GetAllDeleteTestType(c *fiber.Ctx) error
}

type testTypeHandler struct {
	testTypeService services.TestTypeService
}

func NewTestTypeHandler(testTypeService services.TestTypeService) TestTypeHandler {
	return &testTypeHandler{
		testTypeService: testTypeService,
	}
}

func (h *testTypeHandler) CreateTestType(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) UpdateTestType(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) DeleteTestType(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) HardDeleteTestType(c *fiber.Ctx) error {      
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) RestoreTestType(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) GetTestTypeByID(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) GetAllTestType(c *fiber.Ctx) error {
        panic("not implemented") // TODO: Implement
}

func (h *testTypeHandler) GetAllDeleteTestType(c *fiber.Ctx) error {    
        panic("not implemented") // TODO: Implement
}