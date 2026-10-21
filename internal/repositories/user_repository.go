package repositories

import (
	"errors"

	"Go-Auth-API/internal/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindDetailsByID(id uint) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (repo *UserRepository) FindDetailsByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	err := repo.db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}
