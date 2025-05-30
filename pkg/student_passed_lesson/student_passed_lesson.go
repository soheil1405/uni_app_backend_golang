package student_passed_lesson

import (
	handlers "uni_app/pkg/student_passed_lesson/handler"
	repositories "uni_app/pkg/student_passed_lesson/repository"
	usecases "uni_app/pkg/student_passed_lesson/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	passedLessonRepo := repositories.NewStudentPassedLessonRepository(db)
	passedLessonUsecase := usecases.NewStudentPassedLessonUsecase(passedLessonRepo, config)
	handlers.NewStudentPassedLessonHandler(passedLessonUsecase, e)
}
