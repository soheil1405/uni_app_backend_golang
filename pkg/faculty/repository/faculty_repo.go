package repositories

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type FacultyRepository interface {
	Create(faculty *models.Faculty) error
	GetByID(id uint) (*models.Faculty, error)
	Update(faculty *models.Faculty) error
	Delete(id uint) error
	GetAll() ([]models.Faculty, error)
}

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) FacultyRepository {
	return &facultyRepository{db}
}

func (r *facultyRepository) Create(faculty *models.Faculty) error {
	return r.db.Create(faculty).Error
}

func (r *facultyRepository) GetByID(id uint) (*models.Faculty, error) {
	var faculty models.Faculty
	if err := r.db.First(&faculty, id).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (r *facultyRepository) Update(faculty *models.Faculty) error {
	return r.db.Save(faculty).Error
}

func (r *facultyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Faculty{}, id).Error
}

func (r *facultyRepository) GetAll() ([]models.Faculty, error) {
	var faculties []models.Faculty
	if err := r.db.Find(&faculties).Error; err != nil {
		return nil, err
	}
	return faculties, nil
}
