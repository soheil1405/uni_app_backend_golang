package lesson

import (
	"uni_app/pkg/lesson/handler"
	"uni_app/pkg/lesson/repository"
	"uni_app/pkg/lesson/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, cfg *env.Config) {
	lessonRepo := repository.NewLessonRepository(db)
	lessonUsecase := usecase.NewLessonUsecase(lessonRepo)
	handler.NewLessonHandler(lessonUsecase, e)
}
