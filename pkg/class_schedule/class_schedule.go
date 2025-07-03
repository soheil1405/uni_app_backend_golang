package class_schedule

import (
	handlers "uni_app/pkg/class_schedule/handler"
	repositories "uni_app/pkg/class_schedule/repository"
	usecases "uni_app/pkg/class_schedule/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	classScheduleRepo := repositories.NewClassScheduleRepository(db)
	classScheduleUsecase := usecases.NewClassScheduleUsecase(classScheduleRepo, config)
	handlers.NewClassScheduleHandler(classScheduleUsecase, e)
}
