package comment

import (
	"uni_app/pkg/comment/handler"
	"uni_app/pkg/comment/repository"
	"uni_app/pkg/comment/usecase"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group, cfg *env.Config) {
	commentRepository := repository.NewCommentRepository(db)
	commentUseCase := usecase.NewCommentUseCase(commentRepository)
	commentHandler := handler.NewCommentHandler(commentUseCase)
	commentHandler.RegisterRoutes(e)
}
