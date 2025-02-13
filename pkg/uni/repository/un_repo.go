package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type UniRepository interface {
	Create(uni *models.Uni) error
	GetByID(ID database.PID) (*models.Uni, error)
	Update(uni *models.Uni) error
	Delete(ID database.PID) error
	GetAll() ([]models.Uni, error)
}

type uniRepository struct {
	db *gorm.DB
}

func NewUniRepository(db *gorm.DB) UniRepository {
	return &uniRepository{db}
}

func (r *uniRepository) Create(uni *models.Uni) error {
	return r.db.Create(uni).Error
}

func (r *uniRepository) GetByID(ID database.PID) (*models.Uni, error) {
	var uni models.Uni
	if err := r.db.First(&uni, ID).Error; err != nil {
		return nil, err
	}
	return &uni, nil
}

func (r *uniRepository) Update(uni *models.Uni) error {
	return r.db.Save(uni).Error
}

func (r *uniRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Uni{}, ID).Error
}

func (r *uniRepository) GetAll() ([]models.Uni, error) {
	var unis []models.Uni
	if err := r.db.Find(&unis).Error; err != nil {
		return nil, err
	}
	return unis, nil
}
