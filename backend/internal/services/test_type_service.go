package services

import (
	"errors"
	"fmt"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
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
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
}

func (s *testTypeService) GetAllTestType(query *validators.ListTestTypeQuery) ([]models.TestType, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}

func (s *testTypeService) GetAllDeleteTestType(query *validators.ListTestTypeQuery) ([]models.TestType, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}
