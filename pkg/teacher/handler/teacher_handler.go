package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/teacher/usecase"

	"github.com/labstack/echo/v4"
)

type TeacherHandler struct {
	usecase usecase.TeacherUsecase
	group   echo.Group
}

func NewTeacherHandler(usecase usecase.TeacherUsecase, group echo.Group) {
	handler := &TeacherHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *TeacherHandler) initRoutes() {
	teacherGroup := h.group.Group("/teachers")

	teacherGroup.POST("", h.Create)
	teacherGroup.PUT("/:id", h.Update)
	teacherGroup.DELETE("/:id", h.Delete)
	teacherGroup.GET("/:id", h.GetByID)
	teacherGroup.GET("", h.List)
}

func (h *TeacherHandler) Create(c echo.Context) error {
	var teacher models.Teacher
	if err := c.Bind(&teacher); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&teacher); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, teacher)
}

func (h *TeacherHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var teacher models.Teacher
	if err := c.Bind(&teacher); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	teacher.ID = database.PID(id)
	if err := h.usecase.Update(&teacher); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TeacherHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	teacher, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) List(c echo.Context) error {
	var filters models.FetchTeacherRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	teachers, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teachers)
}
