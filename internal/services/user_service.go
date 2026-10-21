package services

import (
	"net/mail"
	"strings"
	"unicode/utf8"

	"Go-Auth-API/internal/dtos"
	"Go-Auth-API/internal/models"
	"Go-Auth-API/internal/repositories"
	"Go-Auth-API/internal/utils"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	FindDetailsByID(id uint) (*dtos.UserDTO, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Register(req dtos.UserFormDTO) (uint, error)
}

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) FindDetailsByID(id uint) (*dtos.UserDTO, error) {
	var userDetails dtos.UserDTO

	user, err := us.repo.FindDetailsByID(id)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(user, &userDetails); err != nil {
		return nil, err
	}

	return &userDetails, nil
}

func (us *UserService) FindByID(id uint) (*models.User, error) {
	return us.repo.FindByID(id)
}

func (us *UserService) FindByEmail(email string) (*models.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	return us.repo.FindByEmail(email)
}

func (us *UserService) Register(req dtos.UserFormDTO) (uint, error) {
	if err := validateUserForm(&req); err != nil {
		return 0, err
	}

	existingUser, err := us.FindByEmail(req.Email)
	if err != nil {
		return 0, err
	}

	if existingUser != nil {
		return 0, utils.NewConflictError("Email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return 0, err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		PhoneNumber:  req.PhoneNumber,
	}

	if err := us.repo.Create(user); err != nil {
		return 0, err
	}

	return user.ID, nil
}

func validateUserForm(dto *dtos.UserFormDTO) error {
	dto.Username = strings.TrimSpace(dto.Username)
	dto.Email = strings.ToLower(strings.TrimSpace(dto.Email))
	dto.PhoneNumber = strings.TrimSpace(dto.PhoneNumber)

	if dto.Username == "" {
		return utils.NewValidationError(
			"Username is required",
			"INVALID_INPUT",
			"username",
		)
	}

	if utf8.RuneCountInString(dto.Username) < 3 {
		return utils.NewValidationError(
			"Username must be at least 3 characters",
			"INVALID_INPUT",
			"username",
		)
	}

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

	if dto.PhoneNumber == "" {
		return utils.NewValidationError(
			"Phone number is required",
			"INVALID_INPUT",
			"phone_number",
		)
	}

	if utf8.RuneCountInString(dto.PhoneNumber) > 20 {
		return utils.NewValidationError(
			"Phone number must not exceed 20 characters",
			"INVALID_INPUT",
			"phone_number",
		)
	}

	if dto.Password == "" {
		return utils.NewValidationError(
			"Password is required",
			"INVALID_INPUT",
			"password",
		)
	}

	if utf8.RuneCountInString(dto.Password) < 8 {
		return utils.NewValidationError(
			"Password must be at least 8 characters",
			"INVALID_INPUT",
			"password",
		)
	}

	return nil
}
