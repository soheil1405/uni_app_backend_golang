package route

import (
	"uni_app/pkg/route/handler"
	"uni_app/pkg/route/repository"
	"uni_app/pkg/route/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, cfg *env.Config) {
	routeRepo := repository.NewRouteRepository(db)
	routeUsecase := usecase.NewRouteUsecase(routeRepo)
	handler.NewRouteHandler(routeUsecase, e)
}
