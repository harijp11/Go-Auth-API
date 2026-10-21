package services

import (
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(req dtos.LoginFormDTO) (string, error)
}

type AuthService struct {
	userService UserServiceInterface
	jwtService  JWTServiceInterface
}

func NewAuthService(userService UserServiceInterface, jwtService  JWTServiceInterface,) AuthServiceInterface {
	return &AuthService{
		userService: userService,
		jwtService : jwtService,
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

	token, err := as.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	return token, nil
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
