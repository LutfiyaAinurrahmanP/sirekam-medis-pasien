package handlers

import (
	"strconv"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/services"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
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
	var req validators.CreateTestTypeRequest

	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	testType, err := h.testTypeService.CreateTestType(&req)
	if err != nil {
		errorMessage := err.Error()

		if errorMessage == "code already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to create test type")
	}

	return utils.SuccessResponse(c, "Test Type successfully created", fiber.Map{
		"test_type": testType,
	})
}

func (h *testTypeHandler) UpdateTestType(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid test type id", nil)
	}

	var req validators.UpdateTestTypeRequest
	if err := validators.ParseAndValidate(c, &req); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
		return utils.BadRequestResponse(c, err.Error(), nil)
	}

	testType, err := h.testTypeService.UpdateTestType(uint(id), &req)
	if err != nil {
		errorMessage := err.Error()
		if errorMessage == "test type code already exists" {
			return utils.ConflictResponse(c, errorMessage)
		}
		return utils.InternalServerErrorResponse(c, "Failed to update test type")
	}

	return utils.SuccessResponse(c, "Test type updated successfully", fiber.Map{
		"test_type": testType,
	})
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
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid test type ID", nil)
	}

	testType, err := h.testTypeService.GetTestTypeByID(uint(id))
	if err != nil {
		if err.Error() == "test type not found" {
			return utils.NotFoundResponse(c, err.Error())
		}
		return utils.InternalServerErrorResponse(c, "Failed to fetch test type")
	}
	return utils.SuccessResponse(c, "Test type retrieved successfully", fiber.Map{
		"test_type": testType,
	})
}

func (h *testTypeHandler) GetAllTestType(c *fiber.Ctx) error {
	var query validators.ListTestTypeQuery

	if err := c.QueryParser(&query); err != nil{
		return utils.BadRequestResponse(c, "Invalid query parameters", nil)
	}

	if err := validators.ValidateStruct(&query); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
	}

	testTypes, meta, err := h.testTypeService.GetAllTestType(&query)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch test types")
	}

	return utils.PaginatedSeccessResponse(c, "Test types retrieved successfully", fiber.Map{
		"test_type": testTypes,
	}, meta)
}

func (h *testTypeHandler) GetAllDeleteTestType(c *fiber.Ctx) error {
	var query validators.ListTestTypeQuery

	if err := c.QueryParser(&query); err != nil {
		return utils.BadRequestResponse(c, "Invalid query parameter", nil)
	}

	if err := validators.ValidateStruct(&query); err != nil {
		if validationErrors := validators.FormatValidationError(err); len(validationErrors) > 0 {
			return utils.BadRequestResponse(c, "Validation failed", validationErrors)
		}
	}

	testTypes, meta, err := h.testTypeService.GetAllDeleteTestType(&query)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch test types")
	}

	return utils.PaginatedSeccessResponse(c, "Test types retrieved successfully", fiber.Map{
		"test_type": testTypes,
	},meta)
}
