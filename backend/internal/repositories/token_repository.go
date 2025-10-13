package repositories

import (
	"time"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"gorm.io/gorm"
)

type TokenRepository interface {
	AddToBlacklist(token *models.TokenBlacklist) error
	IsBlacklisted(token string) (bool, error)
	CleanupExpiredTokens() error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (r *tokenRepository) AddToBlacklist(token *models.TokenBlacklist) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) IsBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&models.TokenBlacklist{}).
		Where("token = ? AND expires_at > ?", token, time.Now()).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *tokenRepository) CleanupExpiredTokens() error {
	return r.db.Where("expires_at <= ?", time.Now()).
		Delete(&models.TokenBlacklist{}).Error
}
