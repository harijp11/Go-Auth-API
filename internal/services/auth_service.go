package services

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrAuthConfig         = errors.New("authentication configuration error")
)

type AuthServiceInterface interface {
	Login(req dtos.LoginFormDTO) (string, error)
}

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewAuthService(userRepo repositories.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (as *AuthService) Login(req dtos.LoginFormDTO) (string, error) {
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if req.Email == "" || req.Password == "" {
		return "", ErrInvalidCredentials
	}

	user, err := as.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(req.Password)); err != nil {
		return "", ErrInvalidCredentials
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", ErrAuthConfig
	}

	now := time.Now()

	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(user.ID),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}