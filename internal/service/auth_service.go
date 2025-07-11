package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
	"github.com/briandang59/be_scada/internal/repository"
)

type AuthService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type authService struct {
	accountRepo repository.AccountRepository
}

func NewAuthService(accountRepo repository.AccountRepository) AuthService {
	return &authService{accountRepo: accountRepo}
}

// ✅ Đăng ký tài khoản
func (s *authService) Register(username, password string) error {
	// Kiểm tra trùng username
	existing, _ := s.accountRepo.FindByUsername(username)
	if existing != nil {
		return errors.New("username already exists")
	}

	// Băm password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	account := &model.Account{
		Username: username,
		Password: string(hashedPassword),
	}

	return s.accountRepo.Create(account)
}

// ✅ Đăng nhập và trả về JWT
func (s *authService) Login(username, password string) (string, error) {
	account, err := s.accountRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Thêm id + username vào claims
	claims := jwt.MapClaims{
		"sub":      account.ID,       // subject
		"id":       account.ID,       // thêm ID riêng
		"username": account.Username, // thêm username
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
