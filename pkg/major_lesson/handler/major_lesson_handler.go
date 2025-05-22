package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/major_lesson/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type MajorLessonHandler struct {
	usecase usecase.MajorLessonUsecase
}

func NewMajorLessonHandler(usecase usecase.MajorLessonUsecase, e echo.Group) {
	majorLessonHandler := &MajorLessonHandler{usecase}

	majorLessonsRouteGroup := e.Group("/major-lessons")
	majorLessonsRouteGroup.POST("", majorLessonHandler.CreateMajorLesson)
	majorLessonsRouteGroup.GET("/:id", majorLessonHandler.GetMajorLessonByID)
	majorLessonsRouteGroup.PUT("/:id", majorLessonHandler.UpdateMajorLesson)
	majorLessonsRouteGroup.DELETE("/:id", majorLessonHandler.DeleteMajorLesson)
	majorLessonsRouteGroup.GET("", majorLessonHandler.GetAllMajorLessons)
}

func (h *MajorLessonHandler) CreateMajorLesson(c echo.Context) error {
	var majorLesson models.MajorLesson
	if err := c.Bind(&majorLesson); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateMajorLesson(&majorLesson); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, majorLesson)
}

func (h *MajorLessonHandler) GetMajorLessonByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	majorLesson, err := h.usecase.GetMajorLessonByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, majorLesson)
}

func (h *MajorLessonHandler) UpdateMajorLesson(c echo.Context) (err error) {
	var majorLesson models.MajorLesson
	if majorLesson.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateMajorLesson(&majorLesson); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, majorLesson)
}

func (h *MajorLessonHandler) DeleteMajorLesson(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteMajorLesson(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "MajorLesson deleted"})
}

func (h *MajorLessonHandler) GetAllMajorLessons(c echo.Context) error {
	var request models.FetchMajorLessonRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	majorLessons, paginate, err := h.usecase.GetAllMajorLessons(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"major_lessons": majorLessons,
		"meta":          paginate,
	})
}
