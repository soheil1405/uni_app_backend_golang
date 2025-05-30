package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RatingRepository interface {
	Create(rating *models.Rating) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Rating, error)
	Update(rating *models.Rating) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchRatingRequest) ([]models.Rating, error)
	GetByRatable(OwnerID database.PID, OwnerType string) ([]models.Rating, error)
	GetAverageRating(OwnerID database.PID, OwnerType string) (float64, error)
	GetUserRating(StudentID database.PID, OwnerID database.PID, OwnerType string) (*models.Rating, error)
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{db}
}

func (r *ratingRepository) Create(rating *models.Rating) error {
	return r.db.Create(rating).Error
}

func (r *ratingRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Rating, error) {
	var rating models.Rating
	if err := r.db.First(&rating, ID).Error; err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *ratingRepository) Update(rating *models.Rating) error {
	return r.db.Save(rating).Error
}

func (r *ratingRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Rating{}, ID).Error
}

func (r *ratingRepository) GetAll(ctx echo.Context, request models.FetchRatingRequest) ([]models.Rating, error) {
	var ratings []models.Rating
	query := r.db.Model(&models.Rating{})

	if request.StudentID > 0 {
		query = query.Where("student_id = ?", request.StudentID)
	}
	if request.OwnerID > 0 {
		query = query.Where("owner_id = ?", request.OwnerID)
	}
	if request.OwnerType != "" {
		query = query.Where("owner_type = ?", request.OwnerType)
	}
	if request.MinRating > 0 {
		query = query.Where("rating >= ?", request.MinRating)
	}
	if request.MaxRating > 0 {
		query = query.Where("rating <= ?", request.MaxRating)
	}

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&ratings).Error; err != nil {
		return nil, err
	}
	return ratings, nil
}

func (r *ratingRepository) GetByRatable(OwnerID database.PID, OwnerType string) ([]models.Rating, error) {
	var ratings []models.Rating
	if err := r.db.Where("owner_ = ? AND ratable_type = ?", OwnerID, OwnerType).
		Preload("User").
		Find(&ratings).Error; err != nil {
		return nil, err
	}
	return ratings, nil
}

func (r *ratingRepository) GetAverageRating(OwnerID database.PID, OwnerType string) (float64, error) {
	var avg float64
	if err := r.db.Model(&models.Rating{}).
		Where("owner_ = ? AND ratable_type = ?", OwnerID, OwnerType).
		Select("COALESCE(AVG(rating), 0)").
		Scan(&avg).Error; err != nil {
		return 0, err
	}
	return avg, nil
}

func (r *ratingRepository) GetUserRating(StudentID database.PID, OwnerID database.PID, OwnerType string) (*models.Rating, error) {
	var rating models.Rating
	if err := r.db.Where("student_id = ? AND owner_ = ? AND ratable_type = ?",
		StudentID, OwnerID, OwnerType).
		First(&rating).Error; err != nil {
		return nil, err
	}
	return &rating, nil
}
