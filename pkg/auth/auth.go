package auth

import (
	handlers "uni_app/pkg/auth/handler"
	repositories "uni_app/pkg/auth/repository"
	usecases "uni_app/pkg/auth/usecase"
	routeRepo "uni_app/pkg/route/repository"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	uniRepo := repositories.NewAuthRepository(db)
	routeRepo := routeRepo.NewRouteRepository(db)
	uniUsecase := usecases.NewAuthUsecase(uniRepo, routeRepo)
	handlers.NewAuthHandler(uniUsecase, e)
}
