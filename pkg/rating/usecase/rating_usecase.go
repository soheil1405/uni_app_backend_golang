package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/rating/repository"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
)

type RatingUsecase interface {
	AddRating(rating *models.Rating) error
	GetRatingByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Rating, error)
	UpdateRating(rating *models.Rating) error
	DeleteRating(ID database.PID) error
	GetAllRatings(ctx echo.Context, request models.FetchRatingRequest) ([]models.Rating, error)
	GetRatableRatings(ratableID database.PID, ratableType string) ([]models.Rating, error)
	GetRatableAverageRating(ratableID database.PID, ratableType string) (float64, error)
	GetUserRating(userID database.PID, ratableID database.PID, ratableType string) (*models.Rating, error)
}

type ratingUsecase struct {
	repo   repositories.RatingRepository
	config *env.Config
}

func NewRatingUsecase(repo repositories.RatingRepository, config *env.Config) RatingUsecase {
	return &ratingUsecase{repo, config}
}

func (u *ratingUsecase) AddRating(rating *models.Rating) error {
	return u.repo.Create(rating)
}

func (u *ratingUsecase) GetRatingByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Rating, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *ratingUsecase) UpdateRating(rating *models.Rating) error {
	return u.repo.Update(rating)
}

func (u *ratingUsecase) DeleteRating(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *ratingUsecase) GetAllRatings(ctx echo.Context, request models.FetchRatingRequest) ([]models.Rating, error) {
	return u.repo.GetAll(ctx, request)
}

func (u *ratingUsecase) GetRatableRatings(ratableID database.PID, ratableType string) ([]models.Rating, error) {
	return u.repo.GetByRatable(ratableID, ratableType)
}

func (u *ratingUsecase) GetRatableAverageRating(ratableID database.PID, ratableType string) (float64, error) {
	return u.repo.GetAverageRating(ratableID, ratableType)
}

func (u *ratingUsecase) GetUserRating(userID database.PID, ratableID database.PID, ratableType string) (*models.Rating, error) {
	return u.repo.GetUserRating(userID, ratableID, ratableType)
}
