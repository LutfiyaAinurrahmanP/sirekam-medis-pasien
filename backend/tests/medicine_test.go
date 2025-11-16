package tests

import (
	"testing"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Test CreateMedicineRequest
func TestCreateMedicineRequest_Valid(t *testing.T) {
	price := 50000.0
	req := validators.CreateMedicineRequest{
		Name:          "Paracetamol",
		GenericName:   "Acetaminophen",
		BrandName:     "Panadol",
		Type:          "tablet",
		Strength:      "500mg",
		Manufacturer:  "PT. Pharma",
		Unit:          "tablet",
		StockQuantity: 100,
		Price:         &price,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "Valid CreateMedicineRequest should pass validation")
}

func TestCreateMedicineRequest_MissingName(t *testing.T) {
	price := 50000.0
	req := validators.CreateMedicineRequest{
		Name:          "", // Missing required field
		GenericName:   "Acetaminophen",
		BrandName:     "Panadol",
		Type:          "tablet",
		Strength:      "500mg",
		Manufacturer:  "PT. Pharma",
		Unit:          "tablet",
		StockQuantity: 100,
		Price:         &price,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateMedicineRequest without name should fail validation")
}

func TestCreateMedicineRequest_MissingType(t *testing.T) {
	price := 50000.0
	req := validators.CreateMedicineRequest{
		Name:          "Paracetamol",
		GenericName:   "Acetaminophen",
		BrandName:     "Panadol",
		Type:          "", // Missing required field
		Strength:      "500mg",
		Manufacturer:  "PT. Pharma",
		Unit:          "tablet",
		StockQuantity: 100,
		Price:         &price,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateMedicineRequest without type should fail validation")
}

func TestCreateMedicineRequest_InvalidType(t *testing.T) {
	price := 50000.0
	req := validators.CreateMedicineRequest{
		Name:          "Paracetamol",
		GenericName:   "Acetaminophen",
		BrandName:     "Panadol",
		Type:          "invalid_type", // Invalid type
		Strength:      "500mg",
		Manufacturer:  "PT. Pharma",
		Unit:          "tablet",
		StockQuantity: 100,
		Price:         &price,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateMedicineRequest with invalid type should fail validation")
}

func TestCreateMedicineRequest_NameTooLong(t *testing.T) {
	price := 50000.0
	longName := ""
	for i := 0; i < 201; i++ {
		longName += "a"
	}

	req := validators.CreateMedicineRequest{
		Name:          longName, // Exceeds max length of 200
		GenericName:   "Acetaminophen",
		BrandName:     "Panadol",
		Type:          "tablet",
		Strength:      "500mg",
		Manufacturer:  "PT. Pharma",
		Unit:          "tablet",
		StockQuantity: 100,
		Price:         &price,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateMedicineRequest with name exceeding max length should fail validation")
}

func TestCreateMedicineRequest_AllValidTypes(t *testing.T) {
	price := 50000.0
	validTypes := []string{"tablet", "capsule", "syrup", "injection", "ointment", "other"}

	for _, medicineType := range validTypes {
		req := validators.CreateMedicineRequest{
			Name:          "Test Medicine",
			Type:          medicineType,
			StockQuantity: 100,
			IsActive:      true,
			Price:         &price,
		}

		err := validate.Struct(req)
		assert.NoError(t, err, "CreateMedicineRequest with type '%s' should pass validation", medicineType)
	}
}

func TestCreateMedicineRequest_MinimalValid(t *testing.T) {
	req := validators.CreateMedicineRequest{
		Name:          "Aspirin",
		Type:          "tablet",
		StockQuantity: 100,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "CreateMedicineRequest with minimal required fields should pass validation")
}

// Test UpdateMedicineRequest
func TestUpdateMedicineRequest_Valid(t *testing.T) {
	price := 60000.0
	req := validators.UpdateMedicineRequest{
		Name:          "Paracetamol Updated",
		GenericName:   "Acetaminophen",
		BrandName:     "Panadol",
		Type:          "capsule",
		Strength:      "250mg",
		Manufacturer:  "PT. Pharma Indonesia",
		Unit:          "capsule",
		StockQuantity: 150,
		Price:         &price,
		IsActive:      true,
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "Valid UpdateMedicineRequest should pass validation")
}

func TestUpdateMedicineRequest_Empty(t *testing.T) {
	req := validators.UpdateMedicineRequest{}

	err := validate.Struct(req)
	assert.NoError(t, err, "Empty UpdateMedicineRequest should pass validation (all fields are optional)")
}

func TestUpdateMedicineRequest_OnlyName(t *testing.T) {
	req := validators.UpdateMedicineRequest{
		Name: "Updated Name",
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "UpdateMedicineRequest with only name should pass validation")
}

func TestUpdateMedicineRequest_InvalidType(t *testing.T) {
	req := validators.UpdateMedicineRequest{
		Type: "invalid_type",
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with invalid type should fail validation")
}

func TestUpdateMedicineRequest_NameTooLong(t *testing.T) {
	longName := ""
	for i := 0; i < 201; i++ {
		longName += "a"
	}

	req := validators.UpdateMedicineRequest{
		Name: longName,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with name exceeding max length should fail validation")
}

func TestUpdateMedicineRequest_GenericNameTooLong(t *testing.T) {
	longName := ""
	for i := 0; i < 201; i++ {
		longName += "a"
	}

	req := validators.UpdateMedicineRequest{
		GenericName: longName,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with generic name exceeding max length should fail validation")
}

func TestUpdateMedicineRequest_BrandNameTooLong(t *testing.T) {
	longName := ""
	for i := 0; i < 201; i++ {
		longName += "a"
	}

	req := validators.UpdateMedicineRequest{
		BrandName: longName,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with brand name exceeding max length should fail validation")
}

func TestUpdateMedicineRequest_StrengthTooLong(t *testing.T) {
	longStrength := ""
	for i := 0; i < 51; i++ {
		longStrength += "a"
	}

	req := validators.UpdateMedicineRequest{
		Strength: longStrength,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with strength exceeding max length should fail validation")
}

func TestUpdateMedicineRequest_ManufacturerTooLong(t *testing.T) {
	longManufacturer := ""
	for i := 0; i < 101; i++ {
		longManufacturer += "a"
	}

	req := validators.UpdateMedicineRequest{
		Manufacturer: longManufacturer,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with manufacturer exceeding max length should fail validation")
}

func TestUpdateMedicineRequest_UnitTooLong(t *testing.T) {
	longUnit := ""
	for i := 0; i < 21; i++ {
		longUnit += "a"
	}

	req := validators.UpdateMedicineRequest{
		Unit: longUnit,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "UpdateMedicineRequest with unit exceeding max length should fail validation")
}

// Test ListMedicineQuery
func TestListMedicineQuery_Valid(t *testing.T) {
	query := validators.ListMedicineQuery{
		Page:   1,
		Limit:  10,
		Search: "paracetamol",
		Sort:   "asc",
		SortBy: "id",
	}

	err := validate.Struct(query)
	assert.NoError(t, err, "Valid ListMedicineQuery should pass validation")
}

func TestListMedicineQuery_Empty(t *testing.T) {
	query := validators.ListMedicineQuery{}

	err := validate.Struct(query)
	assert.NoError(t, err, "Empty ListMedicineQuery should pass validation")
}

func TestListMedicineQuery_PageZero(t *testing.T) {
	query := validators.ListMedicineQuery{
		Page: 0,
	}

	err := validate.Struct(query)
	// Page 0 passes validation but should be set to 1 by SetMedicineDefaults
	assert.NoError(t, err, "ListMedicineQuery with page 0 passes validation (handled by SetMedicineDefaults)")
}

func TestListMedicineQuery_PageNegative(t *testing.T) {
	query := validators.ListMedicineQuery{
		Page: -1,
	}

	err := validate.Struct(query)
	assert.Error(t, err, "ListMedicineQuery with negative page should fail validation")
}

func TestListMedicineQuery_LimitZero(t *testing.T) {
	query := validators.ListMedicineQuery{
		Limit: 0,
	}

	err := validate.Struct(query)
	// Limit 0 passes validation but should be set to 10 by SetMedicineDefaults
	assert.NoError(t, err, "ListMedicineQuery with limit 0 passes validation (handled by SetMedicineDefaults)")
}

func TestListMedicineQuery_LimitExceedsMax(t *testing.T) {
	query := validators.ListMedicineQuery{
		Limit: 101,
	}

	err := validate.Struct(query)
	assert.Error(t, err, "ListMedicineQuery with limit exceeding max should fail validation")
}

func TestListMedicineQuery_SearchTooLong(t *testing.T) {
	longSearch := ""
	for i := 0; i < 101; i++ {
		longSearch += "a"
	}

	query := validators.ListMedicineQuery{
		Search: longSearch,
	}

	err := validate.Struct(query)
	assert.Error(t, err, "ListMedicineQuery with search exceeding max length should fail validation")
}

func TestListMedicineQuery_InvalidSort(t *testing.T) {
	query := validators.ListMedicineQuery{
		Sort: "invalid",
	}

	err := validate.Struct(query)
	assert.Error(t, err, "ListMedicineQuery with invalid sort should fail validation")
}

func TestListMedicineQuery_ValidSortValues(t *testing.T) {
	validSorts := []string{"asc", "desc"}

	for _, sort := range validSorts {
		query := validators.ListMedicineQuery{
			Sort: sort,
		}

		err := validate.Struct(query)
		assert.NoError(t, err, "ListMedicineQuery with sort '%s' should pass validation", sort)
	}
}

func TestListMedicineQuery_InvalidSortBy(t *testing.T) {
	query := validators.ListMedicineQuery{
		SortBy: "invalid_field",
	}

	err := validate.Struct(query)
	assert.Error(t, err, "ListMedicineQuery with invalid sort_by should fail validation")
}

func TestListMedicineQuery_ValidSortByValues(t *testing.T) {
	validSortBy := []string{"id", "type", "stock_quantity", "created_at"}

	for _, sortBy := range validSortBy {
		query := validators.ListMedicineQuery{
			SortBy: sortBy,
		}

		err := validate.Struct(query)
		assert.NoError(t, err, "ListMedicineQuery with sort_by '%s' should pass validation", sortBy)
	}
}

// Test SetMedicineDefaults
func TestListMedicineQuery_SetDefaults(t *testing.T) {
	query := validators.ListMedicineQuery{}
	query.SetMedicineDefaults()

	assert.Equal(t, 1, query.Page, "Default page should be 1")
	assert.Equal(t, 10, query.Limit, "Default limit should be 10")
	assert.Equal(t, "desc", query.Sort, "Default sort should be 'desc'")
	assert.Equal(t, "id", query.SortBy, "Default sort_by should be 'id'")
}

func TestListMedicineQuery_SetDefaultsWithInvalidPage(t *testing.T) {
	query := validators.ListMedicineQuery{
		Page: 0,
	}
	query.SetMedicineDefaults()

	assert.Equal(t, 1, query.Page, "Page should be set to 1 when invalid")
}

func TestListMedicineQuery_SetDefaultsWithInvalidLimit(t *testing.T) {
	query := validators.ListMedicineQuery{
		Limit: 0,
	}
	query.SetMedicineDefaults()

	assert.Equal(t, 10, query.Limit, "Limit should be set to 10 when 0")
}

func TestListMedicineQuery_SetDefaultsWithExcessiveLimit(t *testing.T) {
	query := validators.ListMedicineQuery{
		Limit: 200,
	}
	query.SetMedicineDefaults()

	assert.Equal(t, 100, query.Limit, "Limit should be capped at 100")
}

func TestListMedicineQuery_SetDefaultsWithValidValues(t *testing.T) {
	query := validators.ListMedicineQuery{
		Page:   5,
		Limit:  20,
		Sort:   "asc",
		SortBy: "type",
	}
	query.SetMedicineDefaults()

	assert.Equal(t, 5, query.Page, "Page should remain unchanged when valid")
	assert.Equal(t, 20, query.Limit, "Limit should remain unchanged when valid")
	assert.Equal(t, "asc", query.Sort, "Sort should remain unchanged when valid")
	assert.Equal(t, "type", query.SortBy, "SortBy should remain unchanged when valid")
}

// Test GetMedicineOffSet
func TestListMedicineQuery_GetOffset(t *testing.T) {
	tests := []struct {
		page     int
		limit    int
		expected int
	}{
		{1, 10, 0},
		{2, 10, 10},
		{3, 10, 20},
		{1, 20, 0},
		{2, 20, 20},
		{5, 25, 100},
	}

	for _, tt := range tests {
		query := validators.ListMedicineQuery{
			Page:  tt.page,
			Limit: tt.limit,
		}

		offset := query.GetMedicineOffSet()
		assert.Equal(t, tt.expected, offset, "Offset calculation should be correct for page %d and limit %d", tt.page, tt.limit)
	}
}

// Edge case tests
func TestCreateMedicineRequest_WithNilPrice(t *testing.T) {
	req := validators.CreateMedicineRequest{
		Name:          "Free Medicine",
		Type:          "tablet",
		StockQuantity: 50,
		IsActive:      true,
		Price:         nil, // Nil price is allowed
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "CreateMedicineRequest with nil price should pass validation")
}

func TestUpdateMedicineRequest_WithNilPrice(t *testing.T) {
	req := validators.UpdateMedicineRequest{
		Name:  "Updated Medicine",
		Price: nil, // Nil price is allowed
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "UpdateMedicineRequest with nil price should pass validation")
}

func TestCreateMedicineRequest_AllOptionalFields(t *testing.T) {
	req := validators.CreateMedicineRequest{
		Name:          "Basic Medicine",
		Type:          "tablet",
		StockQuantity: 100,
		IsActive:      true,
		// All optional fields omitted
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "CreateMedicineRequest without optional fields should pass validation")
}
