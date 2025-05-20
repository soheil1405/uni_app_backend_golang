package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MajorLessonRepository interface {
	Create(majorLesson *models.MajorLesson) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error)
	Update(majorLesson *models.MajorLesson) error
	Delete(ID database.PID) error
	GetAll() ([]models.MajorLesson, error)
}

type majorLessonRepository struct {
	db *gorm.DB
}

func NewMajorLessonRepository(db *gorm.DB) MajorLessonRepository {
	return &majorLessonRepository{db}
}

func (r *majorLessonRepository) Create(majorLesson *models.MajorLesson) error {
	return r.db.Create(majorLesson).Error
}

func (r *majorLessonRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error) {
	var majorLesson models.MajorLesson
	if err := r.db.First(&majorLesson, ID).Error; err != nil {
		return nil, err
	}
	return &majorLesson, nil
}

func (r *majorLessonRepository) Update(majorLesson *models.MajorLesson) error {
	return r.db.Save(majorLesson).Error
}

func (r *majorLessonRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.MajorLesson{}, ID).Error
}

func (r *majorLessonRepository) GetAll() ([]models.MajorLesson, error) {
	var majorLessons []models.MajorLesson
	if err := r.db.Find(&majorLessons).Error; err != nil {
		return nil, err
	}
	return majorLessons, nil
} 