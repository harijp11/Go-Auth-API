package services

import (
	"errors"
	"strings"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/models"
	"Go-Auth-API/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrInvalidInput       = errors.New("invalid registration input")
)

type UserServiceInterface interface {
	Register(req dtos.UserFormDTO) (*models.User, error)
}

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(req dtos.UserFormDTO) (*models.User, error) {
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if req.Username == "" || req.Email == "" || len(req.Password) < 8 {
		return nil, ErrInvalidInput
	}

	existingUser, err := us.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := us.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}