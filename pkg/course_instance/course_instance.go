package course_instance

import (
	handlers "uni_app/pkg/course_instance/handler"
	repositories "uni_app/pkg/course_instance/repository"
	usecases "uni_app/pkg/course_instance/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	courseInstanceRepo := repositories.NewCourseInstanceRepository(db)
	courseInstanceUsecase := usecases.NewCourseInstanceUsecase(courseInstanceRepo, config)
	handlers.NewCourseInstanceHandler(courseInstanceUsecase, e)
}
