package handlers

import (
	"net/http"
	"uni_app/models"
	usecases "uni_app/pkg/student_passed_lesson/usecase"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type StudentPassedLessonHandler struct {
	usecase usecases.StudentPassedLessonUsecase
}

func NewStudentPassedLessonHandler(usecase usecases.StudentPassedLessonUsecase, e echo.Group) {
	handler := &StudentPassedLessonHandler{
		usecase: usecase,
	}

	e.POST("/student-passed-lessons", handler.AddPassedLesson)
	e.GET("/student-passed-lessons/:id", handler.GetPassedLessonByID)
	e.PUT("/student-passed-lessons/:id", handler.UpdatePassedLesson)
	e.DELETE("/student-passed-lessons/:id", handler.DeletePassedLesson)
	e.GET("/student-passed-lessons", handler.GetAllPassedLessons)
	e.GET("/students/:student_id/passed-lessons", handler.GetStudentPassedLessons)
}

func (h *StudentPassedLessonHandler) AddPassedLesson(c echo.Context) error {
	var passedLesson models.StudentPassedLesson
	if err := c.Bind(&passedLesson); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.AddPassedLesson(&passedLesson); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, passedLesson)
}

func (h *StudentPassedLessonHandler) GetPassedLessonByID(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedLesson, err := h.usecase.GetPassedLessonByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedLesson)
}

func (h *StudentPassedLessonHandler) UpdatePassedLesson(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var passedLesson models.StudentPassedLesson
	if err := c.Bind(&passedLesson); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedLesson.ID = ID
	if err := h.usecase.UpdatePassedLesson(&passedLesson); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedLesson)
}

func (h *StudentPassedLessonHandler) DeletePassedLesson(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeletePassedLesson(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *StudentPassedLessonHandler) GetAllPassedLessons(c echo.Context) error {
	var request models.FetchStudentPassedLessonRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedLessons, err := h.usecase.GetAllPassedLessons(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedLessons)
}

func (h *StudentPassedLessonHandler) GetStudentPassedLessons(c echo.Context) error {
	studentID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedLessons, err := h.usecase.GetStudentPassedLessons(studentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedLessons)
}
