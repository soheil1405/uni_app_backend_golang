package room

import (
	handlers "uni_app/pkg/room/handler"
	repositories "uni_app/pkg/room/repository"
	usecases "uni_app/pkg/room/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, config *env.Config) {
	roomRepo := repositories.NewRoomRepository(db)
	roomUsecase := usecases.NewRoomUsecase(roomRepo, config)
	handlers.NewRoomHandler(roomUsecase, e)
}
