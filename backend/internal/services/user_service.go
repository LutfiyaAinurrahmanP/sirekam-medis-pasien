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

type UserService interface {
	// Admin
	CreateUser(req *validators.CreateUserRequest) (*models.User, error)
	UpdateUser(id uint, req *validators.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uint) error
	HardDeleteUser(id uint) error
	RestoreUser(id uint) error

	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(query *validators.ListUserQuery) ([]models.User, *utils.PaginationMeta, error)
	GetAllDeletedUsers(query *validators.ListUserQuery) ([]models.User, *utils.PaginationMeta, error)

	// User
	GetProfile(userID uint) (*models.User, error)
	UpdateProfile(userID uint, req *validators.UpdateProfileRequest) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(req *validators.CreateUserRequest) (*models.User, error) {
	// 1. Validasi role
	if !models.ValidateRole(req.Role) {
		return nil, errors.New("invalid role")
	}

	// 2. Cek uniqueness
	exists, err := s.userRepo.ExistsByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username: %w", err)
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	exists, err = s.userRepo.ExistsByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	exists, err = s.userRepo.ExistsByPhone(req.Phone)
	if err != nil {
		return nil, fmt.Errorf("failed to check phone: %w", err)
	}
	if exists {
		return nil, errors.New("phone already exists")
	}

	// 3. Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to created user: %w", err)
	}

	return user, nil
}

func (s *userService) UpdateUser(id uint, req *validators.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if req.Username != "" && req.Username != user.Username {
		exists, err := s.userRepo.ExistsByUsername(req.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to check username: %w", err)
		}
		if exists {
			return nil, errors.New("username already exists")
		}
		user.Username = req.Username
	}

	if req.Email != "" && req.Email != user.Email {
		exists, err := s.userRepo.ExistsByEmail(req.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to check email: %w", err)
		}
		if exists {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.Phone != "" && req.Phone != user.Phone {
		exists, err := s.userRepo.ExistsByPhone(req.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed to check phone: %w", err)
		}
		if exists {
			return nil, errors.New("phone already exists")
		}
		user.Phone = req.Phone
	}

	if req.Role != "" {
		if !models.ValidateRole(req.Role) {
			return nil, errors.New("invalid role")
		}
	user.Role = req.Role
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to updated user: %w", err)
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	if err := s.userRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete userL %w", err)
	}
	return nil
}

func (s *userService) HardDeleteUser(id uint) error {
	if err := s.userRepo.HardDelete(id); err != nil {
		return fmt.Errorf("failed to permanently delete user: %w", err)
	}
	return nil
}

func (s *userService) RestoreUser(id uint) error {
	if err := s.userRepo.Restore(id); err != nil {
		return fmt.Errorf("failed to respore user: %w", err)
	}
	return nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (s *userService) GetAllUsers(query *validators.ListUserQuery) ([]models.User, *utils.PaginationMeta, error) {
	query.SetDefaults()

	users, total, err := s.userRepo.FindAll(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	meta := &utils.PaginationMeta{
		CurrentPage: query.Page,
		PerPage:     query.Limit,
		Total:       total,
		TotalPages:  (total + int64(query.Limit) - 1) / int64(query.Limit),
	}
	return users, meta, nil
}

func (s *userService) GetAllDeletedUsers(query *validators.ListUserQuery) ([]models.User, *utils.PaginationMeta, error) {
	query.SetDefaults()

	users, total, err := s.userRepo.FindAllDelete(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch deleted users: %w", err)
	}

	meta := &utils.PaginationMeta{
		CurrentPage: query.Page,
		PerPage:     query.Limit,
		Total:       total,
		TotalPages:  (total + int64(query.Limit) - 1) / int64(query.Limit),
	}
	return users, meta, nil
}

func (s *userService) GetProfile(userID uint) (*models.User, error) {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}
	return user, nil
}

func (s *userService) UpdateProfile(userID uint, req *validators.UpdateProfileRequest) (*models.User, error) {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if req.Username != "" && req.Username != user.Username {
		exists, err := s.userRepo.ExistsByUsername(req.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to check username: %w", err)
		}
		if exists {
			return nil, errors.New("username already exists")
		}
		user.Username = req.Username
	}

	if req.Email != "" && req.Email != user.Email {
		exists, err := s.userRepo.ExistsByEmail(req.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to check email: %w", err)
		}
		if exists {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.Phone != "" && req.Phone != user.Phone {
		exists, err := s.userRepo.ExistsByPhone(req.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed to check phone: %w", err)
		}
		if exists {
			return nil, errors.New("phone already exists")
		}
		user.Phone = req.Phone
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return user, nil
}
