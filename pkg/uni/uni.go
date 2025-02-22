package uni

import (
	handlers "uni_app/pkg/uni/handler"
	repositories "uni_app/pkg/uni/repository"
	usecases "uni_app/pkg/uni/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	uniRepo := repositories.NewUniRepository(db)
	uniUsecase := usecases.NewUniUsecase(uniRepo)
	handlers.NewUniHandler(uniUsecase, e)
}
