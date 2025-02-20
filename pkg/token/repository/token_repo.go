package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(token *models.Token) error
	GetByID(ID database.PID) (*models.Token, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *userRepository) GetByID(ID database.PID) (*models.Token, error) {
	var token models.Token
	if err := r.db.First(&token, ID).Error; err != nil {
		return nil, err
	}
	return &token, nil
}
