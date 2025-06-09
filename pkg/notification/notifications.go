package notification

import (
	"uni_app/pkg/notification/handler"
	"uni_app/pkg/notification/repository"
	"uni_app/pkg/notification/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitNotification(e echo.Group, db *gorm.DB) {
	notificationRepository := repository.NewNotificationRepository(db)
	notificationUseCase := usecase.NewNotificationUseCase(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(notificationUseCase)
	notificationHandler.RegisterRoutes(e)
}
