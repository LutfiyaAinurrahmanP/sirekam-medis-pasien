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

type MedicineService interface {
	CreateMedicine(req *validators.CreateMedicineRequest) (*models.Medicine, error)
	UpdateMedicine(id uint, req *validators.UpdateMedicineRequest) (*models.Medicine, error)
	DeleteMedicine(id uint) error
	HardDeleteMedicine(id uint) error
	RestoreMedicine(id uint) error

	GetMedicineByID(id uint) (*models.Medicine, error)
	GetAllMedicine(query *validators.ListMedicineQuery) ([]models.Medicine, *utils.PaginationMeta, error)
	GetAllDeletedMedicine(query *validators.ListMedicineQuery) ([]models.Medicine, *utils.PaginationMeta, error)
}

type medicineService struct {
	medicineRepo repositories.MedicineRepository
}

func NewMedicineService(medicineRepo repositories.MedicineRepository) MedicineService {
	return &medicineService{
		medicineRepo: medicineRepo,
	}
}

func (s *medicineService) CreateMedicine(req *validators.CreateMedicineRequest) (*models.Medicine, error) {
	medicine := &models.Medicine{
		Name:          req.Name,
		GenericName:   req.GenericName,
		BrandName:     req.BrandName,
		Type:          req.Type,
		Strength:      req.Strength,
		Manufacturer:  req.Manufacturer,
		Unit:          req.Unit,
		StockQuantity: req.StockQuantity,
		Price:         req.Price,
		IsActive:      req.IsActive,
	}

	if err := s.medicineRepo.Create(medicine); err != nil {
		return nil, fmt.Errorf("failed to created medicine: %w", err)
	}

	return medicine, nil
}

func (s *medicineService) UpdateMedicine(id uint, req *validators.UpdateMedicineRequest) (*models.Medicine, error) {
	medicine, err := s.medicineRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, fmt.Errorf("medicine not found")
		}
		return nil, fmt.Errorf("failed to find medicine: %w", err)
	}

	if req.Name != "" && req.Name != medicine.Name {
		medicine.Name = req.Name
	}

	if req.GenericName != "" && req.GenericName != medicine.GenericName {
		medicine.GenericName = req.GenericName
	}

	if req.BrandName != "" && req.BrandName != medicine.BrandName {
		medicine.BrandName = req.BrandName
	}

	if req.Type != "" && req.Type != medicine.Type {
		medicine.Type = req.Type
	}

	if req.Strength != "" && req.Strength != medicine.Strength {
		medicine.Strength = req.Strength
	}

	if req.Manufacturer != "" && req.Manufacturer != medicine.Manufacturer {
		medicine.Manufacturer = req.Manufacturer
	}

	if req.Unit != "" && req.Unit != medicine.Unit {
		medicine.Unit = req.Unit
	}

	if req.StockQuantity != 0 && req.StockQuantity != medicine.StockQuantity {
		medicine.StockQuantity = req.StockQuantity
	}

	if req.Price != nil && req.Price != medicine.Price {
		medicine.Price = req.Price
	}

	if req.IsActive != medicine.IsActive {
		medicine.IsActive = req.IsActive
	}

	if err := s.medicineRepo.Update(medicine); err != nil {
		return nil, fmt.Errorf("failed to updated medicine: %w", err)
	}

	return medicine, nil
}

func (s *medicineService) DeleteMedicine(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *medicineService) HardDeleteMedicine(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *medicineService) RestoreMedicine(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *medicineService) GetMedicineByID(id uint) (*models.Medicine, error) {
	medicine, err := s.medicineRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, fmt.Errorf("medicine not found")
		}
		return nil, fmt.Errorf("failed to get medicine: %w", err)
	}
	return medicine, nil
}

func (s *medicineService) GetAllMedicine(query *validators.ListMedicineQuery) ([]models.Medicine, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}

func (s *medicineService) GetAllDeletedMedicine(query *validators.ListMedicineQuery) ([]models.Medicine, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}
