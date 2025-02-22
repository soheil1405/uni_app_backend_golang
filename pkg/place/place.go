package place

import (
	handlers "uni_app/pkg/place/handler"
	repositories "uni_app/pkg/place/repository"
	usecases "uni_app/pkg/place/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	uniRepo := repositories.NewPlaceRepository(db)
	uniUsecase := usecases.NewPlaceUsecase(uniRepo)
	handlers.NewPlaceHandler(uniUsecase, e)
}
