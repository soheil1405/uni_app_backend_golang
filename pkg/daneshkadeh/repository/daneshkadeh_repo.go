package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type FacultyRepository interface {
	Create(faculty *models.DaneshKadeh) error
	GetByID(ID database.PID) (*models.DaneshKadeh, error)
	Update(faculty *models.DaneshKadeh) error
	Delete(ID database.PID) error
	GetAll() (*models.DaneshKadeha, error)
}

type facultyRepository struct {
	db *gorm.DB
}

func NewDaneshKadehRepository(db *gorm.DB) FacultyRepository {
	return &facultyRepository{db}
}

func (r *facultyRepository) Create(faculty *models.DaneshKadeh) error {
	return r.db.Create(faculty).Error
}

func (r *facultyRepository) GetByID(ID database.PID) (*models.DaneshKadeh, error) {
	var faculty models.DaneshKadeh
	if err := r.db.First(&faculty, ID).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (r *facultyRepository) Update(faculty *models.DaneshKadeh) error {
	return r.db.Save(faculty).Error
}

func (r *facultyRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.DaneshKadeha{}, ID).Error
}

func (r *facultyRepository) GetAll() (*models.DaneshKadeha, error) {
	var faculties *models.DaneshKadeha
	if err := r.db.Find(&faculties).Error; err != nil {
		return nil, err
	}
	return faculties, nil
}
