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

type DepartmentService interface {
	CreateDepartment(req *validators.CreateDepartmentRequest) (*models.Department, error)
	UpdateDepartment(id uint, req *validators.UpdateDepartmentRequest) (*models.Department, error)
	DeleteDepartment(id uint) error
	HardDeleteDepartment(id uint) error
	RestoreDepartment(id uint) error

	GetDepartmentByID(id uint) (*models.Department, error)
	GetAllDepartments(query *validators.ListDepartmentQuery) ([]models.Department, *utils.PaginationMeta, error)
	GetAllDeletedDepartments(query *validators.ListDepartmentQuery) ([]models.Department, *utils.PaginationMeta, error)
}

type departmentService struct {
	departmentRepo repositories.DepartmentRepository
}

func NewDepartmentService(departmentRepo repositories.DepartmentRepository) DepartmentService {
	return &departmentService{
		departmentRepo: departmentRepo,
	}
}

func (s *departmentService) CreateDepartment(req *validators.CreateDepartmentRequest) (*models.Department, error) {
	exists, err := s.departmentRepo.ExistsByCode(req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to check code: %w", err)
	}
	if exists {
		return nil, errors.New("code already exists")
	}

	department := &models.Department{
		Name:          req.Name,
		Code:          req.Code,
		Description:   req.Description,
		FloorLocation: req.FloorLocation,
	}

	if err := s.departmentRepo.Create(department); err != nil {
		return nil, fmt.Errorf("failed to created department: %w", err)
	}

	return department, nil
}

func (s *departmentService) UpdateDepartment(id uint, req *validators.UpdateDepartmentRequest) (*models.Department, error) {
	department, err := s.departmentRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("department not found")
		}
		return nil, fmt.Errorf("failed to find department: %w", err)
	}

	if req.Code != "" && req.Code != department.Code {
		exists, err := s.departmentRepo.ExistsByCode(req.Code)
		if err != nil {
			return nil, fmt.Errorf("failed to check code: %w", err)
		}
		if exists {
			return nil, errors.New("code already exists")
		}
		department.Code = req.Code
	}

	if req.Description != "" && req.Description != department.Description {
		department.Description = req.Description
	}

	if req.FloorLocation != "" && req.FloorLocation != department.FloorLocation {
		department.FloorLocation = req.FloorLocation
	}

	if req.Name != "" && req.Name != department.Name {
		department.Name = req.Name
	}

	if err := s.departmentRepo.Update(department); err != nil {
		return nil, fmt.Errorf("failed to updated department: %w", err)
	}
	return department, nil
}

func (s *departmentService) DeleteDepartment(id uint) error {
	_, err := s.departmentRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return errors.New("department not found")
		}
		return fmt.Errorf("failed to find department: %w", err)
	}

	if err := s.departmentRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete department: %w", err)
	}
	return nil
}

func (s *departmentService) HardDeleteDepartment(id uint) error {
	if err := s.departmentRepo.HardDelete(id); err != nil {
		return fmt.Errorf("failed to permanently delete department: %w", err)
	}
	return nil
}

func (s *departmentService) RestoreDepartment(id uint) error {
	if err := s.departmentRepo.Restore(id); err != nil {
		return fmt.Errorf("failed to restore department: %w", err)
	}
	return nil
}

func (s *departmentService) GetDepartmentByID(id uint) (*models.Department, error) {
	department, err := s.departmentRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("department not found")
		}
		return nil, fmt.Errorf("failed to get department: %w", err)
	}
	return department, nil
}

func (s *departmentService) GetAllDepartments(query *validators.ListDepartmentQuery) ([]models.Department, *utils.PaginationMeta, error) {
	query.SetDepartmentDefaults()

	departments, total, err := s.departmentRepo.FindAll(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch department: %w", err)
	}

	meta := &utils.PaginationMeta{
		CurrentPage: query.Page,
		PerPage: query.Limit,
		Total: total,
		TotalPages: (total + int64(query.Limit) - 1) /int64(query.Limit),
	}
	return departments, meta, nil
}

func (s *departmentService) GetAllDeletedDepartments(query *validators.ListDepartmentQuery) ([]models.Department, *utils.PaginationMeta, error) {
	query.SetDepartmentDefaults()

	departments, total, err := s.departmentRepo.FindAllDelete(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch department: %w", err)
	}

	meta := &utils.PaginationMeta{
		CurrentPage: query.Page,
		PerPage: query.Limit,
		Total: total,
		TotalPages: (total + int64(query.Limit) - 1) /int64(query.Limit),
	}
	return departments, meta, nil
}
