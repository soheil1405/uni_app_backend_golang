package repositories

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type UniRepository interface {
	Create(uni *models.Uni) error
	GetByID(id uint) (*models.Uni, error)
	Update(uni *models.Uni) error
	Delete(id uint) error
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

func (r *uniRepository) GetByID(id uint) (*models.Uni, error) {
	var uni models.Uni
	if err := r.db.First(&uni, id).Error; err != nil {
		return nil, err
	}
	return &uni, nil
}

func (r *uniRepository) Update(uni *models.Uni) error {
	return r.db.Save(uni).Error
}

func (r *uniRepository) Delete(id uint) error {
	return r.db.Delete(&models.Uni{}, id).Error
}

func (r *uniRepository) GetAll() ([]models.Uni, error) {
	var unis []models.Uni
	if err := r.db.Find(&unis).Error; err != nil {
		return nil, err
	}
	return unis, nil
}
