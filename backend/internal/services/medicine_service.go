package services

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
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
        panic("not implemented") // TODO: Implement
}

func (s *medicineService) UpdateMedicine(id uint, req *validators.UpdateMedicineRequest) (*models.Medicine, error) {
        panic("not implemented") // TODO: Implement
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
        panic("not implemented") // TODO: Implement
}

func (s *medicineService) GetAllMedicine(query *validators.ListMedicineQuery) ([]models.Medicine, *utils.PaginationMeta, error) {
        panic("not implemented") // TODO: Implement
}

func (s *medicineService) GetAllDeletedMedicine(query *validators.ListMedicineQuery) ([]models.Medicine, *utils.PaginationMeta, error) {
        panic("not implemented") // TODO: Implement
}