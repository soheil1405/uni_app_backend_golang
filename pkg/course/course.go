package course

import (
	handlers "uni_app/pkg/course/handler"
	repositories "uni_app/pkg/course/repository"
	usecases "uni_app/pkg/course/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	courseRepo := repositories.NewCourseRepository(db)
	courseUsecase := usecases.NewCourseUsecase(courseRepo, config)
	handlers.NewCourseHandler(courseUsecase, e)
}
