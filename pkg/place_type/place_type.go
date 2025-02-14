package place_type

import (
	handlers "uni_app/pkg/place_type/handler"
	repositories "uni_app/pkg/place_type/repository"
	usecases "uni_app/pkg/place_type/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group) {
	uniRepo := repositories.NewPlaceTypeRepository(db)
	uniUsecase := usecases.NewPlaceTypeUsecase(uniRepo)
	handlers.NewPlaceTypeHandler(uniUsecase, e)
}
