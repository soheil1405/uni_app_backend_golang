package faculty

import (
	handlers "uni_app/pkg/faculty/handler"
	repositories "uni_app/pkg/faculty/repository"
	usecases "uni_app/pkg/faculty/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group) {
	uniRepo := repositories.NewFacultyRepository(db)
	uniUsecase := usecases.NewFacultyUsecase(uniRepo)
	handlers.NewFacultyHandler(uniUsecase, e)
}
