package faculty

import (
	handlers "uni_app/pkg/daneshkadeh/handler"
	repositories "uni_app/pkg/daneshkadeh/repository"
	usecases "uni_app/pkg/daneshkadeh/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	uniRepo := repositories.NewDaneshKadehRepository(db)
	uniUsecase := usecases.NewDaneshKadehUsecase(uniRepo)
	handlers.NewDaneshKadehHandler(uniUsecase, e)
}
