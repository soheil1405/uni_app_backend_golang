package uni_major

import (
	"uni_app/pkg/uni_major/handler"
	"uni_app/pkg/uni_major/repository"
	"uni_app/pkg/uni_major/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, cfg *env.Config) {
	uniMajorRepo := repository.NewUniMajorRepository(db)
	uniMajorUsecase := usecase.NewUniMajorUsecase(uniMajorRepo)
	handler.NewUniMajorHandler(uniMajorUsecase, e)
}
