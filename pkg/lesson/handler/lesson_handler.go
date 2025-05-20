package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/lesson/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type LessonHandler struct {
	usecase usecase.LessonUsecase
}

func NewLessonHandler(usecase usecase.LessonUsecase, e echo.Group) {
	lessonHandler := &LessonHandler{usecase}

	lessonsRouteGroup := e.Group("/lessons")
	lessonsRouteGroup.POST("", lessonHandler.CreateLesson)
	lessonsRouteGroup.GET("/:id", lessonHandler.GetLessonByID)
	lessonsRouteGroup.PUT("/:id", lessonHandler.UpdateLesson)
	lessonsRouteGroup.DELETE("/:id", lessonHandler.DeleteLesson)
	lessonsRouteGroup.GET("", lessonHandler.GetAllLessons)

}

func (h *LessonHandler) CreateLesson(c echo.Context) error {
	var lesson models.Lesson
	if err := c.Bind(&lesson); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateLesson(&lesson); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, lesson)
}

func (h *LessonHandler) GetLessonByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	lesson, err := h.usecase.GetLessonByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, lesson)
}

func (h *LessonHandler) UpdateLesson(c echo.Context) (err error) {
	var lesson models.Lesson
	if lesson.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateLesson(&lesson); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, lesson)
}

func (h *LessonHandler) DeleteLesson(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteLesson(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Lesson deleted"})
}

func (h *LessonHandler) GetAllLessons(c echo.Context) error {
	lessons, err := h.usecase.GetAllLessons()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, lessons)
} 