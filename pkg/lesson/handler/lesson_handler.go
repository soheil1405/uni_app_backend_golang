package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/lesson/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

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
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateLesson(&lesson); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"lesson": lesson}, nil)
}

func (h *LessonHandler) GetLessonByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	lesson, err := h.usecase.GetLessonByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"lesson": lesson}, nil)
}

func (h *LessonHandler) UpdateLesson(c echo.Context) (err error) {
	var lesson models.Lesson
	if lesson.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateLesson(&lesson); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"lesson": lesson}, nil)
}

func (h *LessonHandler) DeleteLesson(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteLesson(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Lesson deleted"}, nil)
}

func (h *LessonHandler) GetAllLessons(c echo.Context) error {
	lessons, err := h.usecase.GetAllLessons()
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"lessons": lessons}, nil)
}
