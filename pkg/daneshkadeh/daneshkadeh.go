package faculty

import (
	"uni_app/models"
	handlers "uni_app/pkg/daneshkadeh/handler"
	repositories "uni_app/pkg/daneshkadeh/repository"
	usecases "uni_app/pkg/daneshkadeh/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *models.Config) {
	uniRepo := repositories.NewDaneshKadehRepository(db)
	uniUsecase := usecases.NewDaneshKadehUsecase(uniRepo)
	handlers.NewDaneshKadehHandler(uniUsecase, e)
}
