package handlers

import (
	"net/http"
	"uni_app/models"
	usecases "uni_app/pkg/student_passed_lesson/usecase"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type StudentPassedCourseHandler struct {
	usecase usecases.StudentPassedCourseUsecase
}

func NewStudentPassedCourseHandler(usecase usecases.StudentPassedCourseUsecase, e echo.Group) {
	handler := &StudentPassedCourseHandler{
		usecase: usecase,
	}

	e.POST("/student-passed-course", handler.AddPassedCourse)
	e.GET("/student-passed-course/:id", handler.GetPassedCourseByID)
	e.PUT("/student-passed-course/:id", handler.UpdatePassedCourse)
	e.DELETE("/student-passed-course/:id", handler.DeletePassedCourse)
	e.GET("/student-passed-course", handler.GetAllPassedCourses)
	e.GET("/students/:student_id/passed-course", handler.GetStudentPassedCourses)
}

func (h *StudentPassedCourseHandler) AddPassedCourse(c echo.Context) error {
	var passedCourse models.StudentPassedCourse
	if err := c.Bind(&passedCourse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.AddPassedCourse(&passedCourse); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, passedCourse)
}

func (h *StudentPassedCourseHandler) GetPassedCourseByID(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedCourse, err := h.usecase.GetPassedCourseByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedCourse)
}

func (h *StudentPassedCourseHandler) UpdatePassedCourse(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var passedLCourse models.StudentPassedCourse
	if err := c.Bind(&passedLCourse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedLCourse.ID = ID
	if err := h.usecase.UpdatePassedCourse(&passedLCourse); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedLCourse)
}

func (h *StudentPassedCourseHandler) DeletePassedCourse(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeletePassedCourse(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *StudentPassedCourseHandler) GetAllPassedCourses(c echo.Context) error {
	var request models.FetchStudentPassedCourseRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedCourses, err := h.usecase.GetAllPassedCourses(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedCourses)
}

func (h *StudentPassedCourseHandler) GetStudentPassedCourses(c echo.Context) error {
	studentID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	passedCourses, err := h.usecase.GetStudentPassedCourses(studentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, passedCourses)
}
