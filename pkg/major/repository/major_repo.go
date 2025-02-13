package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type MajorRepository interface {
	Create(major *models.Major) error
	GetByID(ID database.PID) (*models.Major, error)
	Update(major *models.Major) error
	Delete(ID database.PID) error
	GetAll() ([]models.Major, error)
}

type majorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) MajorRepository {
	return &majorRepository{db}
}

func (r *majorRepository) Create(major *models.Major) error {
	return r.db.Create(major).Error
}

func (r *majorRepository) GetByID(ID database.PID) (*models.Major, error) {
	var major models.Major
	if err := r.db.First(&major, ID).Error; err != nil {
		return nil, err
	}
	return &major, nil
}

func (r *majorRepository) Update(major *models.Major) error {
	return r.db.Save(major).Error
}

func (r *majorRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Major{}, ID).Error
}

func (r *majorRepository) GetAll() ([]models.Major, error) {
	var majors []models.Major
	if err := r.db.Find(&majors).Error; err != nil {
		return nil, err
	}
	return majors, nil
}
