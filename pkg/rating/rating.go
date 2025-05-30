package rating

import (
	handlers "uni_app/pkg/rating/handler"
	repositories "uni_app/pkg/rating/repository"
	usecases "uni_app/pkg/rating/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	ratingRepo := repositories.NewRatingRepository(db)
	ratingUsecase := usecases.NewRatingUsecase(ratingRepo, config)
	handlers.NewRatingHandler(ratingUsecase, e)
}
