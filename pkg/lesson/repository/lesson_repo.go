package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LessonRepository interface {
	Create(lesson *models.Lesson) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Lesson, error)
	Update(lesson *models.Lesson) error
	Delete(ID database.PID) error
	GetAll() ([]models.Lesson, error)
}

type lessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &lessonRepository{db}
}

func (r *lessonRepository) Create(lesson *models.Lesson) error {
	return r.db.Create(lesson).Error
}

func (r *lessonRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Lesson, error) {
	var lesson models.Lesson
	if err := r.db.First(&lesson, ID).Error; err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *lessonRepository) Update(lesson *models.Lesson) error {
	return r.db.Save(lesson).Error
}

func (r *lessonRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Lesson{}, ID).Error
}

func (r *lessonRepository) GetAll() ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := r.db.Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
} 