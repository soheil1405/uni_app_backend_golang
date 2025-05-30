package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentPassedLessonRepository interface {
	Create(passedLesson *models.StudentPassedLesson) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedLesson, error)
	Update(passedLesson *models.StudentPassedLesson) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchStudentPassedLessonRequest) ([]models.StudentPassedLesson, error)
	GetByStudentID(studentID database.PID) ([]models.StudentPassedLesson, error)
}

type studentPassedLessonRepository struct {
	db *gorm.DB
}

func NewStudentPassedLessonRepository(db *gorm.DB) StudentPassedLessonRepository {
	return &studentPassedLessonRepository{db}
}

func (r *studentPassedLessonRepository) Create(passedLesson *models.StudentPassedLesson) error {
	return r.db.Create(passedLesson).Error
}

func (r *studentPassedLessonRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedLesson, error) {
	var passedLesson models.StudentPassedLesson
	if err := r.db.First(&passedLesson, ID).Error; err != nil {
		return nil, err
	}
	return &passedLesson, nil
}

func (r *studentPassedLessonRepository) Update(passedLesson *models.StudentPassedLesson) error {
	return r.db.Save(passedLesson).Error
}

func (r *studentPassedLessonRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.StudentPassedLesson{}, ID).Error
}

func (r *studentPassedLessonRepository) GetAll(ctx echo.Context, request models.FetchStudentPassedLessonRequest) ([]models.StudentPassedLesson, error) {
	var passedLessons []models.StudentPassedLesson
	query := r.db.Model(&models.StudentPassedLesson{})

	if request.StudentID > 0 {
		query = query.Where("student_id = ?", request.StudentID)
	}
	if request.LessonID > 0 {
		query = query.Where("lesson_id = ?", request.LessonID)
	}
	if request.Term > 0 {
		query = query.Where("term = ?", request.Term)
	}

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&passedLessons).Error; err != nil {
		return nil, err
	}
	return passedLessons, nil
}

func (r *studentPassedLessonRepository) GetByStudentID(studentID database.PID) ([]models.StudentPassedLesson, error) {
	var passedLessons []models.StudentPassedLesson
	if err := r.db.Where("student_id = ?", studentID).
		Preload("Lesson").
		Find(&passedLessons).Error; err != nil {
		return nil, err
	}
	return passedLessons, nil
}
