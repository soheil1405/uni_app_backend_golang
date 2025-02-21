package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *models.Student) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Student, error)
	Update(student *models.Student) error
	Delete(ID database.PID) error
	GetAll() ([]models.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Student, error) {
	var student models.Student
	if err := r.db.First(&student, ID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Student{}, ID).Error
}

func (r *studentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
