package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentPassedCourseRepository interface {
	Create(passedCourse *models.StudentPassedCourse) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedCourse, error)
	Update(passedCourse *models.StudentPassedCourse) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchStudentPassedCourseRequest) ([]models.StudentPassedCourse, error)
	GetByStudentID(studentID database.PID) ([]models.StudentPassedCourse, error)
}

type studentPassedCourseRepository struct {
	db *gorm.DB
}

func NewStudentPassedCourseRepository(db *gorm.DB) StudentPassedCourseRepository {
	return &studentPassedCourseRepository{db}
}

func (r *studentPassedCourseRepository) Create(passedCourse *models.StudentPassedCourse) error {
	return r.db.Create(passedCourse).Error
}

func (r *studentPassedCourseRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedCourse, error) {
	var passedCourse models.StudentPassedCourse
	if err := r.db.First(&passedCourse, ID).Error; err != nil {
		return nil, err
	}
	return &passedCourse, nil
}

func (r *studentPassedCourseRepository) Update(passedCourse *models.StudentPassedCourse) error {
	return r.db.Save(passedCourse).Error
}

func (r *studentPassedCourseRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.StudentPassedCourse{}, ID).Error
}

func (r *studentPassedCourseRepository) GetAll(ctx echo.Context, request models.FetchStudentPassedCourseRequest) ([]models.StudentPassedCourse, error) {
	var passedCourses []models.StudentPassedCourse
	query := r.db.Model(&models.StudentPassedCourse{})

	if request.StudentID > 0 {
		query = query.Where("student_id = ?", request.StudentID)
	}
	if request.CourseID > 0 {
		query = query.Where("lesson_id = ?", request.CourseID)
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

	if err := query.Find(&passedCourses).Error; err != nil {
		return nil, err
	}
	return passedCourses, nil
}

func (r *studentPassedCourseRepository) GetByStudentID(studentID database.PID) ([]models.StudentPassedCourse, error) {
	var passedCourses []models.StudentPassedCourse
	if err := r.db.Where("student_id = ?", studentID).
		Preload("Lesson").
		Find(&passedCourses).Error; err != nil {
		return nil, err
	}
	return passedCourses, nil
}
