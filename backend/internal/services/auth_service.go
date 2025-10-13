package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req *validators.RegisterRequest) (*models.User, error)
	Login(req *validators.LoginRequest) (string, *models.User, error)
	Logout(token string, userID uint) error
	ValidateToken(token string) error
}

type authService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
	jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(req *validators.RegisterRequest) (*models.User, error) {
	req.SetDefaultRole()
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already exists")
	}
	existingUser, err = s.userRepo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}
	existingUser, err = s.userRepo.FindByPhone(req.Phone)
	if err == nil && existingUser != nil {
		return nil, errors.New("phone already exists")
	}

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
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *authService) Login(req *validators.LoginRequest) (string, *models.User, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("invalid username or password")
		}
		return "", nil, fmt.Errorf("failed to find user: %w", err)
	}

	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	token, err := s.generateJWTToken(user)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return token, user, nil
}

func (s *authService) Logout(token string, userID uint) error {
	claims, err := s.parseToken(token)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid token expiration")
	}
	expiresAt := time.Unix(int64(exp), 0)

	blacklistedToken := &models.TokenBlacklist{
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}

	if err := s.tokenRepo.AddToBlacklist(blacklistedToken); err != nil {
		return fmt.Errorf("failed to blacklist token: %w", err)
	}

	return nil
}

func (s *authService) ValidateToken(token string) error {
	isBlacklisted, err := s.tokenRepo.IsBlacklisted(token)
	if err != nil {
		return fmt.Errorf("failed to check token blacklist: %w", err)
	}

	if isBlacklisted {
		return errors.New("token has been revoked")
	}

	_, err = s.parseToken(token)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	return nil
}

func (s *authService) generateJWTToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *authService) parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ExtractTokenFromHeader(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}
