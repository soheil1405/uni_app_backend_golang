package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/faculty/usecase"

	"github.com/labstack/echo/v4"
)

type FacultyHandler struct {
	usecase usecase.FacultyUsecase
	group   echo.Group
}

func NewFacultyHandler(usecase usecase.FacultyUsecase, group echo.Group) {
	handler := &FacultyHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *FacultyHandler) initRoutes() {
	facultyGroup := h.group.Group("/faculties")

	facultyGroup.POST("", h.Create)
	facultyGroup.PUT("/:id", h.Update)
	facultyGroup.DELETE("/:id", h.Delete)
	facultyGroup.GET("/:id", h.GetByID)
	facultyGroup.GET("", h.List)
	facultyGroup.POST("/:id/teachers/:teacher_id", h.AddTeacher)
	facultyGroup.DELETE("/:id/teachers/:teacher_id", h.RemoveTeacher)
	facultyGroup.POST("/:id/staff/:user_id", h.AddStaff)
	facultyGroup.DELETE("/:id/staff/:user_id", h.RemoveStaff)
}

func (h *FacultyHandler) Create(c echo.Context) error {
	var faculty models.Faculty
	if err := c.Bind(&faculty); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&faculty); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, faculty)
}

func (h *FacultyHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var faculty models.Faculty
	if err := c.Bind(&faculty); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	faculty.ID = database.PID(id)
	if err := h.usecase.Update(&faculty); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, faculty)
}

func (h *FacultyHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FacultyHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	faculty, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, faculty)
}

func (h *FacultyHandler) List(c echo.Context) error {
	var filters models.FetchFacultyRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	faculties, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, faculties)
}

func (h *FacultyHandler) AddTeacher(c echo.Context) error {
	facultyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid faculty ID")
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid teacher ID")
	}

	if err := h.usecase.AddTeacher(database.PID(facultyID), database.PID(teacherID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FacultyHandler) RemoveTeacher(c echo.Context) error {
	facultyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid faculty ID")
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid teacher ID")
	}

	if err := h.usecase.RemoveTeacher(database.PID(facultyID), database.PID(teacherID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FacultyHandler) AddStaff(c echo.Context) error {
	facultyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid faculty ID")
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.usecase.AddStaff(database.PID(facultyID), database.PID(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FacultyHandler) RemoveStaff(c echo.Context) error {
	facultyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid faculty ID")
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.usecase.RemoveStaff(database.PID(facultyID), database.PID(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
