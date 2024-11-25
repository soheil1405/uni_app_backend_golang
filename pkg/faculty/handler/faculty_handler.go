package handlers

import (
	"net/http"
	"strconv"
	"uni_app/models"
	usecases "uni_app/pkg/faculty/usecase"

	"github.com/labstack/echo/v4"
)

type FacultyHandler struct {
	usecase usecases.FacultyUsecase
}

func NewFacultyHandler(usecase usecases.FacultyUsecase, e echo.Group) {
	facultyHandler := &FacultyHandler{usecase}
	e.POST("/faculties", facultyHandler.CreateFaculty)
	e.GET("/faculties/:id", facultyHandler.GetFacultyByID)
	e.PUT("/faculties/:id", facultyHandler.UpdateFaculty)
	e.DELETE("/faculties/:id", facultyHandler.DeleteFaculty)
	e.GET("/faculties", facultyHandler.GetAllFaculties)

}

func (h *FacultyHandler) CreateFaculty(c echo.Context) error {
	var faculty models.Faculty
	if err := c.Bind(&faculty); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateFaculty(&faculty); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, faculty)
}

func (h *FacultyHandler) GetFacultyByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	faculty, err := h.usecase.GetFacultyByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, faculty)
}

func (h *FacultyHandler) UpdateFaculty(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var faculty models.Faculty
	if err := c.Bind(&faculty); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	faculty.ID = uint(id)
	if err := h.usecase.UpdateFaculty(&faculty); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, faculty)
}

func (h *FacultyHandler) DeleteFaculty(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.usecase.DeleteFaculty(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Faculty deleted"})
}

func (h *FacultyHandler) GetAllFaculties(c echo.Context) error {
	faculties, err := h.usecase.GetAllFaculties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, faculties)
}
