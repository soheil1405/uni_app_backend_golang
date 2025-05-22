package repositories

import (
	"errors"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *models.Student) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Student, error)
	Update(student *models.Student) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchRequest) ([]models.Student, *helpers.PaginateTemplate, error)
	GetByStudentCode(studentCode database.PID) (*models.Student, error)
	GetByNationalCode(nationalCode database.PID) (*models.Student, error)
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

func (r *studentRepository) GetAll(ctx echo.Context, request models.FetchRequest) ([]models.Student, *helpers.PaginateTemplate, error) {
	var students []models.Student
	query := r.db.Model(&models.Student{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&students).Error; err != nil {
		return nil, nil, err
	}

	return students, paginate, nil
}

func (r *studentRepository) GetByStudentCode(studentCode database.PID) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("student_code = ?", studentCode).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) GetByNationalCode(nationalCode database.PID) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("national_code = ?", nationalCode).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}
