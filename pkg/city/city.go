package city

import (
	handlers "uni_app/pkg/city/handler"
	repositories "uni_app/pkg/city/repository"
	usecases "uni_app/pkg/city/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group) {
	uniRepo := repositories.NewCityRepository(db)
	uniUsecase := usecases.NewCityUsecase(uniRepo)
	handlers.NewCityHandler(uniUsecase, e)
}
