package major

import (
	"uni_app/models"
	handlers "uni_app/pkg/major/handler"
	repositories "uni_app/pkg/major/repository"
	usecases "uni_app/pkg/major/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *models.Config) {
	uniRepo := repositories.NewMajorRepository(db)
	uniUsecase := usecases.NewMajorUsecase(uniRepo)
	handlers.NewMajorHandler(uniUsecase, e)
}
