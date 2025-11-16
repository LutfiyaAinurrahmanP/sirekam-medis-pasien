package tests

import (
	"testing"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/stretchr/testify/assert"
)

// Test CreateTestTypeRequest
func TestCreateTestTypeRequest_Valid(t *testing.T) {
	price := 150000.0
	req := validators.CreateTestTypeRequest{
		Name:        "Complete Blood Count",
		Code:        "CBC001",
		Category:    "Hematologi",
		Description: "Complete blood count test for general health screening",
		Price:       &price,
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "Valid CreateTestTypeRequest should pass validation")
}

func TestCreateTestTypeRequest_MissingName(t *testing.T) {
	price := 150000.0
	req := validators.CreateTestTypeRequest{
		Name:        "", // Missing required field
		Code:        "CBC001",
		Category:    "Hematologi",
		Description: "Complete blood count test",
		Price:       &price,
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateTestTypeRequest without name should fail validation")
}

func TestCreateTestTypeRequest_MissingCode(t *testing.T) {
	price := 150000.0
	req := validators.CreateTestTypeRequest{
		Name:        "Complete Blood Count",
		Code:        "", // Missing required field
		Category:    "Hematologi",
		Description: "Complete blood count test",
		Price:       &price,
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateTestTypeRequest without code should fail validation")
}

func TestCreateTestTypeRequest_NameTooLong(t *testing.T) {
	price := 150000.0
	longName := ""
	for i := 0; i < 201; i++ {
		longName += "a"
	}

	req := validators.CreateTestTypeRequest{
		Name:        longName, // Exceeds max length of 200
		Code:        "CBC001",
		Category:    "Hematologi",
		Description: "Complete blood count test",
		Price:       &price,
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateTestTypeRequest with name exceeding max length should fail validation")
}

func TestCreateTestTypeRequest_CodeTooLong(t *testing.T) {
	price := 150000.0
	longCode := ""
	for i := 0; i < 51; i++ {
		longCode += "a"
	}

	req := validators.CreateTestTypeRequest{
		Name:        "Complete Blood Count",
		Code:        longCode, // Exceeds max length of 50
		Category:    "Hematologi",
		Description: "Complete blood count test",
		Price:       &price,
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateTestTypeRequest with code exceeding max length should fail validation")
}

func TestCreateTestTypeRequest_CategoryTooLong(t *testing.T) {
	price := 150000.0
	longCategory := ""
	for i := 0; i < 101; i++ {
		longCategory += "a"
	}

	req := validators.CreateTestTypeRequest{
		Name:        "Complete Blood Count",
		Code:        "CBC001",
		Category:    longCategory, // Exceeds max length of 100
		Description: "Complete blood count test",
		Price:       &price,
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.Error(t, err, "CreateTestTypeRequest with category exceeding max length should fail validation")
}

func TestCreateTestTypeRequest_MinimalValid(t *testing.T) {
	req := validators.CreateTestTypeRequest{
		Name:     "Basic Test",
		Code:     "BT001",
		IsActive: true,
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "CreateTestTypeRequest with minimal required fields should pass validation")
}

func TestCreateTestTypeRequest_WithNilPrice(t *testing.T) {
	req := validators.CreateTestTypeRequest{
		Name:        "Free Test",
		Code:        "FT001",
		Category:    "General",
		Description: "Free health screening test",
		Price:       nil, // Nil price is allowed
		IsActive:    true,
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "CreateTestTypeRequest with nil price should pass validation")
}

func TestCreateTestTypeRequest_WithoutOptionalFields(t *testing.T) {
	req := validators.CreateTestTypeRequest{
		Name:     "Simple Test",
		Code:     "ST001",
		IsActive: true,
		// Category, Description, Price are optional
	}

	err := validate.Struct(req)
	assert.NoError(t, err, "CreateTestTypeRequest without optional fields should pass validation")
}

func TestCreateTestTypeRequest_AllValidCategories(t *testing.T) {
	price := 100000.0
	validCategories := []string{"Hematologi", "Kimia Darah", "Imunologi", "Mikrobiologi", "Urinalisis"}

	for _, category := range validCategories {
		req := validators.CreateTestTypeRequest{
			Name:     "Test Type",
			Code:     "TT001",
			Category: category,
			Price:    &price,
			IsActive: true,
		}

		err := validate.Struct(req)
		assert.NoError(t, err, "CreateTestTypeRequest with category '%s' should pass validation", category)
	}
}

// // Test UpdateTestTypeRequest
// func TestUpdateTestTypeRequest_Valid(t *testing.T) {
// 	price := 200000.0
// 	req := validators.UpdateTestTypeRequest{
// 		Name:        "Updated Blood Test",
// 		Code:        "UBT001",
// 		Category:    "Kimia Darah",
// 		Description: "Updated description for blood test",
// 		Price:       &price,
// 		IsActive:    true,
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "Valid UpdateTestTypeRequest should pass validation")
// }

// func TestUpdateTestTypeRequest_Empty(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "Empty UpdateTestTypeRequest should pass validation (all fields are optional)")
// }

// func TestUpdateTestTypeRequest_OnlyName(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{
// 		Name: "Updated Name",
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with only name should pass validation")
// }

// func TestUpdateTestTypeRequest_OnlyCode(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{
// 		Code: "UPD001",
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with only code should pass validation")
// }

// func TestUpdateTestTypeRequest_OnlyCategory(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{
// 		Category: "Hematologi",
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with only category should pass validation")
// }

// func TestUpdateTestTypeRequest_OnlyDescription(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{
// 		Description: "Updated description",
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with only description should pass validation")
// }

// func TestUpdateTestTypeRequest_OnlyPrice(t *testing.T) {
// 	price := 250000.0
// 	req := validators.UpdateTestTypeRequest{
// 		Price: &price,
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with only price should pass validation")
// }

// func TestUpdateTestTypeRequest_OnlyIsActive(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{
// 		IsActive: false,
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with only is_active should pass validation")
// }

// func TestUpdateTestTypeRequest_NameTooLong(t *testing.T) {
// 	longName := ""
// 	for i := 0; i < 201; i++ {
// 		longName += "a"
// 	}

// 	req := validators.UpdateTestTypeRequest{
// 		Name: longName,
// 	}

// 	err := validate.Struct(req)
// 	assert.Error(t, err, "UpdateTestTypeRequest with name exceeding max length should fail validation")
// }

// func TestUpdateTestTypeRequest_CodeTooLong(t *testing.T) {
// 	longCode := ""
// 	for i := 0; i < 51; i++ {
// 		longCode += "a"
// 	}

// 	req := validators.UpdateTestTypeRequest{
// 		Code: longCode,
// 	}

// 	err := validate.Struct(req)
// 	assert.Error(t, err, "UpdateTestTypeRequest with code exceeding max length should fail validation")
// }

// func TestUpdateTestTypeRequest_CategoryTooLong(t *testing.T) {
// 	longCategory := ""
// 	for i := 0; i < 101; i++ {
// 		longCategory += "a"
// 	}

// 	req := validators.UpdateTestTypeRequest{
// 		Category: longCategory,
// 	}

// 	err := validate.Struct(req)
// 	assert.Error(t, err, "UpdateTestTypeRequest with category exceeding max length should fail validation")
// }

// func TestUpdateTestTypeRequest_WithNilPrice(t *testing.T) {
// 	req := validators.UpdateTestTypeRequest{
// 		Name:  "Updated Test",
// 		Price: nil, // Nil price is allowed
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with nil price should pass validation")
// }

// func TestUpdateTestTypeRequest_PartialUpdate(t *testing.T) {
// 	price := 175000.0
// 	req := validators.UpdateTestTypeRequest{
// 		Name:     "Partially Updated Test",
// 		Category: "Imunologi",
// 		Price:    &price,
// 		// Code, Description, IsActive not provided
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with partial fields should pass validation")
// }

// // Test ListTestTypeQuery
// func TestListTestTypeQuery_Valid(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page:   1,
// 		Limit:  10,
// 		Search: "blood",
// 		Sort:   "asc",
// 		SortBy: "id",
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "Valid ListTestTypeQuery should pass validation")
// }

// func TestListTestTypeQuery_Empty(t *testing.T) {
// 	query := validators.ListTestTypeQuery{}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "Empty ListTestTypeQuery should pass validation")
// }

// func TestListTestTypeQuery_PageZero(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page: 0,
// 	}

// 	err := validate.Struct(query)
// 	// Page 0 passes validation but should be set to 1 by SetTestTypeDefaults
// 	assert.NoError(t, err, "ListTestTypeQuery with page 0 passes validation (handled by SetTestTypeDefaults)")
// }

// func TestListTestTypeQuery_PageNegative(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page: -1,
// 	}

// 	err := validate.Struct(query)
// 	assert.Error(t, err, "ListTestTypeQuery with negative page should fail validation")
// }

// func TestListTestTypeQuery_LimitZero(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: 0,
// 	}

// 	err := validate.Struct(query)
// 	// Limit 0 passes validation but should be set to 10 by SetTestTypeDefaults
// 	assert.NoError(t, err, "ListTestTypeQuery with limit 0 passes validation (handled by SetTestTypeDefaults)")
// }

// func TestListTestTypeQuery_LimitExceedsMax(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: 101,
// 	}

// 	err := validate.Struct(query)
// 	assert.Error(t, err, "ListTestTypeQuery with limit exceeding max should fail validation")
// }

// func TestListTestTypeQuery_SearchTooLong(t *testing.T) {
// 	longSearch := ""
// 	for i := 0; i < 101; i++ {
// 		longSearch += "a"
// 	}

// 	query := validators.ListTestTypeQuery{
// 		Search: longSearch,
// 	}

// 	err := validate.Struct(query)
// 	assert.Error(t, err, "ListTestTypeQuery with search exceeding max length should fail validation")
// }

// func TestListTestTypeQuery_InvalidSort(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Sort: "invalid",
// 	}

// 	err := validate.Struct(query)
// 	assert.Error(t, err, "ListTestTypeQuery with invalid sort should fail validation")
// }

// func TestListTestTypeQuery_ValidSortValues(t *testing.T) {
// 	validSorts := []string{"asc", "desc"}

// 	for _, sort := range validSorts {
// 		query := validators.ListTestTypeQuery{
// 			Sort: sort,
// 		}

// 		err := validate.Struct(query)
// 		assert.NoError(t, err, "ListTestTypeQuery with sort '%s' should pass validation", sort)
// 	}
// }

// func TestListTestTypeQuery_InvalidSortBy(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		SortBy: "invalid_field",
// 	}

// 	err := validate.Struct(query)
// 	assert.Error(t, err, "ListTestTypeQuery with invalid sort_by should fail validation")
// }

// func TestListTestTypeQuery_ValidSortByValues(t *testing.T) {
// 	validSortBy := []string{"id", "category", "created_at"}

// 	for _, sortBy := range validSortBy {
// 		query := validators.ListTestTypeQuery{
// 			SortBy: sortBy,
// 		}

// 		err := validate.Struct(query)
// 		assert.NoError(t, err, "ListTestTypeQuery with sort_by '%s' should pass validation", sortBy)
// 	}
// }

// // Test SetTestTypeDefaults
// func TestListTestTypeQuery_SetDefaults(t *testing.T) {
// 	query := validators.ListTestTypeQuery{}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 1, query.Page, "Default page should be 1")
// 	assert.Equal(t, 10, query.Limit, "Default limit should be 10")
// 	assert.Equal(t, "desc", query.Sort, "Default sort should be 'desc'")
// 	assert.Equal(t, "id", query.SortBy, "Default sort_by should be 'id'")
// }

// func TestListTestTypeQuery_SetDefaultsWithInvalidPage(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page: 0,
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 1, query.Page, "Page should be set to 1 when invalid")
// }

// func TestListTestTypeQuery_SetDefaultsWithNegativePage(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page: -5,
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 1, query.Page, "Page should be set to 1 when negative")
// }

// func TestListTestTypeQuery_SetDefaultsWithInvalidLimit(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: 0,
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 10, query.Limit, "Limit should be set to 10 when 0")
// }

// func TestListTestTypeQuery_SetDefaultsWithNegativeLimit(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: -10,
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 10, query.Limit, "Limit should be set to 10 when negative")
// }

// func TestListTestTypeQuery_SetDefaultsWithExcessiveLimit(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: 200,
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 100, query.Limit, "Limit should be capped at 100")
// }

// func TestListTestTypeQuery_SetDefaultsWithValidValues(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page:   5,
// 		Limit:  20,
// 		Sort:   "asc",
// 		SortBy: "category",
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, 5, query.Page, "Page should remain unchanged when valid")
// 	assert.Equal(t, 20, query.Limit, "Limit should remain unchanged when valid")
// 	assert.Equal(t, "asc", query.Sort, "Sort should remain unchanged when valid")
// 	assert.Equal(t, "category", query.SortBy, "SortBy should remain unchanged when valid")
// }

// func TestListTestTypeQuery_SetDefaultsWithEmptySort(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page:   1,
// 		Limit:  10,
// 		Sort:   "",
// 		SortBy: "",
// 	}
// 	query.SetTestTypeDefaults()

// 	assert.Equal(t, "desc", query.Sort, "Sort should be set to 'desc' when empty")
// 	assert.Equal(t, "id", query.SortBy, "SortBy should be set to 'id' when empty")
// }

// // Test GetTestTypeOffSet
// func TestListTestTypeQuery_GetOffset(t *testing.T) {
// 	tests := []struct {
// 		page     int
// 		limit    int
// 		expected int
// 	}{
// 		{1, 10, 0},
// 		{2, 10, 10},
// 		{3, 10, 20},
// 		{1, 20, 0},
// 		{2, 20, 20},
// 		{5, 25, 100},
// 		{10, 15, 135},
// 	}

// 	for _, tt := range tests {
// 		query := validators.ListTestTypeQuery{
// 			Page:  tt.page,
// 			Limit: tt.limit,
// 		}

// 		offset := query.GetTestTypeOffSet()
// 		assert.Equal(t, tt.expected, offset, "Offset calculation should be correct for page %d and limit %d", tt.page, tt.limit)
// 	}
// }

// func TestListTestTypeQuery_GetOffsetWithDefaults(t *testing.T) {
// 	query := validators.ListTestTypeQuery{}
// 	query.SetTestTypeDefaults()

// 	offset := query.GetTestTypeOffSet()
// 	assert.Equal(t, 0, offset, "Offset should be 0 for first page with default limit")
// }

// func TestListTestTypeQuery_GetOffsetSecondPage(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page:  2,
// 		Limit: 10,
// 	}

// 	offset := query.GetTestTypeOffSet()
// 	assert.Equal(t, 10, offset, "Offset should be 10 for second page with limit 10")
// }

// // Edge case tests
// func TestCreateTestTypeRequest_IsActiveFalse(t *testing.T) {
// 	req := validators.CreateTestTypeRequest{
// 		Name:     "Inactive Test",
// 		Code:     "IT001",
// 		IsActive: false, // IsActive can be false
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "CreateTestTypeRequest with IsActive=false should pass validation")
// }

// func TestUpdateTestTypeRequest_AllFields(t *testing.T) {
// 	price := 300000.0
// 	req := validators.UpdateTestTypeRequest{
// 		Name:        "Complete Update",
// 		Code:        "CU001",
// 		Category:    "Mikrobiologi",
// 		Description: "Fully updated test type",
// 		Price:       &price,
// 		IsActive:    true,
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with all fields should pass validation")
// }

// func TestCreateTestTypeRequest_LongDescription(t *testing.T) {
// 	price := 100000.0
// 	longDescription := ""
// 	for i := 0; i < 1000; i++ {
// 		longDescription += "a"
// 	}

// 	req := validators.CreateTestTypeRequest{
// 		Name:        "Test with Long Description",
// 		Code:        "TLD001",
// 		Category:    "General",
// 		Description: longDescription, // Long description is allowed (no max length validation)
// 		Price:       &price,
// 		IsActive:    true,
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "CreateTestTypeRequest with long description should pass validation")
// }

// func TestUpdateTestTypeRequest_LongDescription(t *testing.T) {
// 	longDescription := ""
// 	for i := 0; i < 1000; i++ {
// 		longDescription += "a"
// 	}

// 	req := validators.UpdateTestTypeRequest{
// 		Description: longDescription, // Long description is allowed
// 	}

// 	err := validate.Struct(req)
// 	assert.NoError(t, err, "UpdateTestTypeRequest with long description should pass validation")
// }

// func TestListTestTypeQuery_MinimumValidPage(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page: 1,
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with minimum valid page should pass validation")
// }

// func TestListTestTypeQuery_MinimumValidLimit(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: 1,
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with minimum valid limit should pass validation")
// }

// func TestListTestTypeQuery_MaximumValidLimit(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Limit: 100,
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with maximum valid limit should pass validation")
// }

// func TestListTestTypeQuery_WithSearchOnly(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Search: "hematologi",
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with only search should pass validation")
// }

// func TestListTestTypeQuery_WithSortOnly(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Sort: "asc",
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with only sort should pass validation")
// }

// func TestListTestTypeQuery_WithSortByOnly(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		SortBy: "category",
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with only sort_by should pass validation")
// }

// func TestListTestTypeQuery_CompleteQuery(t *testing.T) {
// 	query := validators.ListTestTypeQuery{
// 		Page:   2,
// 		Limit:  25,
// 		Search: "blood test",
// 		Sort:   "desc",
// 		SortBy: "created_at",
// 	}

// 	err := validate.Struct(query)
// 	assert.NoError(t, err, "ListTestTypeQuery with all fields should pass validation")
// }
