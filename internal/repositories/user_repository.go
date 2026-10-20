
package repositories

import (
	"errors"

	"Go-Auth-API/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := ur.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
