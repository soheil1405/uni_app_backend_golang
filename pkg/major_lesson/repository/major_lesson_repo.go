package repository

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MajorLessonRepository interface {
	Create(majorLesson *models.MajorLesson) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error)
	Update(majorLesson *models.MajorLesson) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchMajorLessonRequest) ([]models.MajorLesson, *helpers.PaginateTemplate, error)
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

func (r *majorLessonRepository) GetAll(ctx echo.Context, request models.FetchMajorLessonRequest) ([]models.MajorLesson, *helpers.PaginateTemplate, error) {
	var majorLessons []models.MajorLesson
	query := r.db.Model(&models.MajorLesson{})

	// Apply filters
	if request.MajorID != 0 {
		query = query.Where("major_id = ?", request.MajorID)
	}
	if request.MajorChartID != 0 {
		query = query.Where("major_chart_id = ?", request.MajorChartID)
	}
	if request.LessonID != 0 {
		query = query.Where("lesson_id = ?", request.LessonID)
	}
	if request.IsOptional {
		query = query.Where("is_optional = ?", request.IsOptional)
	}
	if request.IsTechnical {
		query = query.Where("is_technical = ?", request.IsTechnical)
	}
	if request.RecommendedTerm != 0 {
		query = query.Where("recommended_term = ?", request.RecommendedTerm)
	}

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&majorLessons).Error; err != nil {
		return nil, nil, err
	}

	return majorLessons, paginate, nil
}
