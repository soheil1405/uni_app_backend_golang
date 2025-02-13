package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type FacultyRepository interface {
	Create(faculty *models.Faculty) error
	GetByID(ID database.PID) (*models.Faculty, error)
	Update(faculty *models.Faculty) error
	Delete(ID database.PID) error
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

func (r *facultyRepository) GetByID(ID database.PID) (*models.Faculty, error) {
	var faculty models.Faculty
	if err := r.db.First(&faculty, ID).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (r *facultyRepository) Update(faculty *models.Faculty) error {
	return r.db.Save(faculty).Error
}

func (r *facultyRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Faculty{}, ID).Error
}

func (r *facultyRepository) GetAll() ([]models.Faculty, error) {
	var faculties []models.Faculty
	if err := r.db.Find(&faculties).Error; err != nil {
		return nil, err
	}
	return faculties, nil
}
