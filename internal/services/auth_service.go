package services

import (
	"fmt"
	"net/mail"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(req dtos.LoginFormDTO) (string, error)
}

type AuthService struct {
	userService UserServiceInterface
}

func NewAuthService(userService UserServiceInterface) AuthServiceInterface {
	return &AuthService{
		userService: userService,
	}
}

func (as *AuthService) Login(req dtos.LoginFormDTO) (string, error) {
	if err := validateLoginForm(&req); err != nil {
		return "", err
	}

	user, err := as.userService.FindByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("find user by email: %w", err)
	}

	if user == nil {
		return "", utils.NewUnauthorizedError("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		return "", utils.NewUnauthorizedError("Invalid email or password")
	}

	secret := os.Getenv("JWT_SECRET")
	if strings.TrimSpace(secret) == "" {
		return "", fmt.Errorf("JWT_SECRET is not configured")
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
		return "", fmt.Errorf("sign JWT: %w", err)
	}

	return signedToken, nil
}

func validateLoginForm(dto *dtos.LoginFormDTO) error {
	dto.Email = strings.ToLower(strings.TrimSpace(dto.Email))

	if dto.Email == "" {
		return utils.NewValidationError(
			"Email is required",
			"INVALID_INPUT",
			"email",
		)
	}

	email, err := mail.ParseAddress(dto.Email)
	if err != nil || email.Address != dto.Email {
		return utils.NewValidationError(
			"Invalid email address",
			"INVALID_INPUT",
			"email",
		)
	}

	if dto.Password == "" {
		return utils.NewValidationError(
			"Password is required",
			"INVALID_INPUT",
			"password",
		)
	}

	if utf8.RuneCountInString(dto.Password) > 72 {
		return utils.NewValidationError(
			"Invalid email or password",
			"INVALID_INPUT",
			"password",
		)
	}

	return nil
}