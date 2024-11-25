package major_chart

import (
	handlers "uni_app/pkg/major_chart/handler"
	repositories "uni_app/pkg/major_chart/repository"
	usecases "uni_app/pkg/major_chart/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group) {

	uniRepo := repositories.NewChartRepository(db)

	uniUsecase := usecases.NewChartUsecase(uniRepo)

	handlers.NewChartHandler(uniUsecase, e)

}
