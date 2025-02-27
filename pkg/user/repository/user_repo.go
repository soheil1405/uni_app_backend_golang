package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error)
	GetByUserName(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(ID database.PID) error
	GetAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByUserName(username string) (user *models.User, err error) {
	if err := r.db.Preload("UserRoles").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, ID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.User{}, ID).Error
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
