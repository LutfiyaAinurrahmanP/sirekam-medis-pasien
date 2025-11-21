package services

import (
	"errors"
	"fmt"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type TestTypeService interface {
	CreateTestType(req *validators.CreateTestTypeRequest) (*models.TestType, error)
	UpdateTestType(id uint, req *validators.UpdateTestTypeRequest) (*models.TestType, error)
	DeleteTestType(id uint) error
	HardDeleteTestType(id uint) error
	RestoreTestType(id uint) error

	GetTestTypeByID(id uint) (*models.TestType, error)
	GetAllTestType(query *validators.ListTestTypeQuery) ([]models.TestType, *utils.PaginationMeta, error)
	GetAllDeleteTestType(query *validators.ListTestTypeQuery) ([]models.TestType, *utils.PaginationMeta, error)
}

type testTypeService struct {
	testTypeRepo repositories.TestTypeRepository
}

func NewTestTypeService(testTypeRepo repositories.TestTypeRepository) TestTypeService {
	return &testTypeService{
		testTypeRepo: testTypeRepo,
	}
}

func (s *testTypeService) CreateTestType(req *validators.CreateTestTypeRequest) (*models.TestType, error) {
	exists, err := s.testTypeRepo.ExistsByCode(req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to check code: %w", err)
	}

	if exists {
		return nil, errors.New("code already exists")
	}

	testType := &models.TestType{
		Name:        req.Name,
		Code:        req.Code,
		Category:    req.Category,
		Description: req.Description,
		Price:       req.Price,
		IsActive:    req.IsActive,
	}

	if err := s.testTypeRepo.Crete(testType); err != nil {
		return nil, fmt.Errorf("failed to create test type: %w", err)
	}

	return testType, nil
}

func (s *testTypeService) UpdateTestType(id uint, req *validators.UpdateTestTypeRequest) (*models.TestType, error) {
	testType, err := s.testTypeRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, fmt.Errorf("test type not found")
		}
		return nil, fmt.Errorf("failed to find test type: %w", err)
	}

	if req.Code != "" && req.Code != testType.Code {
		exists, err := s.testTypeRepo.ExistsByCode(req.Code)
		if err != nil {
			return nil, fmt.Errorf("failed to check test type code: %w", err)
		}

		if exists {
			return nil, errors.New("test type code already exists")
		}
		testType.Code = req.Code
	}
	
	if req.Name != "" && req.Name != testType.Name{
		testType.Name = req.Name
	}

	if req.Category != "" && req.Category != testType.Category {
		testType.Category = req.Category
	}

	if req.Description != "" && req.Description != testType.Description {
		testType.Description = req.Description
	}

	if req.Price != nil && req.Price != testType.Price {
		testType.Price = req.Price
	}

	if req.IsActive != testType.IsActive {
		testType.IsActive = req.IsActive
	}

	if err := s.testTypeRepo.Update(testType); err != nil {
		return nil, fmt.Errorf("failed to updated test type: %w", err)
	}
	return testType, nil
}

func (s *testTypeService) DeleteTestType(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *testTypeService) HardDeleteTestType(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *testTypeService) RestoreTestType(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *testTypeService) GetTestTypeByID(id uint) (*models.TestType, error) {
	testType, err := s.testTypeRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("test type not found")
		}
		return nil, fmt.Errorf("failed to get test type: %w", err)
	}
	return testType, nil
}

func (s *testTypeService) GetAllTestType(query *validators.ListTestTypeQuery) ([]models.TestType, *utils.PaginationMeta, error) {
	query.SetTestTypeDefaults()

	testTypes, total, err := s.testTypeRepo.FindAll(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch test type: %w", err)
	}
	meta := &utils.PaginationMeta{
		CurrentPage: query.Page,
		PerPage: query.Limit,
		Total: total,
		TotalPages: (total + int64(query.Limit) - 1) / int64(query.Limit),
	}

	return testTypes, meta, nil
}

func (s *testTypeService) GetAllDeleteTestType(query *validators.ListTestTypeQuery) ([]models.TestType, *utils.PaginationMeta, error) {
	query.SetTestTypeDefaults()

	testTypes, total, err := s.testTypeRepo.FindAllDelete(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch test types: %w", err)
	}
	meta := &utils.PaginationMeta{
		CurrentPage: query.Page,
		PerPage: query.Limit,
		Total: total,
		TotalPages: (total + int64(query.Limit) - 1) / int64(query.Limit),
	}

	return testTypes, meta, nil	
}
