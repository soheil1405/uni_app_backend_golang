package faculty

import (
	handlers "uni_app/pkg/faculty/handler"
	repositories "uni_app/pkg/faculty/repository"
	usecases "uni_app/pkg/faculty/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	facultyRepo := repositories.NewFacultyRepository(db)
	facultyUsecase := usecases.NewFacultyUsecase(facultyRepo, config)
	handlers.NewFacultyHandler(facultyUsecase, e)
}
