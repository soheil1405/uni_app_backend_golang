package news

import (
	handlers "uni_app/pkg/news/handler"
	repositories "uni_app/pkg/news/repository"
	usecases "uni_app/pkg/news/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	newsRepo := repositories.NewNewsRepository(db)
	newsUsecase := usecases.NewNewsUsecase(newsRepo)
	handlers.NewNewsHandler(newsUsecase, e)
}
