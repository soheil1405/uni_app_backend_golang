package building

import (
	handlers "uni_app/pkg/building/handler"
	repositories "uni_app/pkg/building/repository"
	usecases "uni_app/pkg/building/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	buildingRepo := repositories.NewBuildingRepository(db)
	buildingUsecase := usecases.NewBuildingUsecase(buildingRepo, config)
	handlers.NewBuildingHandler(buildingUsecase, e)
}
