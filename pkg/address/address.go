package address

import (
	"uni_app/pkg/address/handler"
	"uni_app/pkg/address/repository"
	"uni_app/pkg/address/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, cfg *env.Config) {
	addressRepo := repository.NewAddressRepository(db)
	addressUsecase := usecase.NewAddressUsecase(addressRepo)
	handler.NewAddressHandler(addressUsecase, e)
}
