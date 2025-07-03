package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/course/usecase"

	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	usecase usecase.CourseUsecase
	group   echo.Group
}

func NewCourseHandler(usecase usecase.CourseUsecase, group echo.Group) {
	handler := &CourseHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *CourseHandler) initRoutes() {
	courseGroup := h.group.Group("/courses")

	courseGroup.POST("", h.Create)
	courseGroup.PUT("/:id", h.Update)
	courseGroup.DELETE("/:id", h.Delete)
	courseGroup.GET("/:id", h.GetByID)
	courseGroup.GET("", h.List)
}

func (h *CourseHandler) Create(c echo.Context) error {
	var course models.Course
	if err := c.Bind(&course); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&course); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, course)
}

func (h *CourseHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var course models.Course
	if err := c.Bind(&course); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	course.ID = database.PID(id)
	if err := h.usecase.Update(&course); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CourseHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	course, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) List(c echo.Context) error {
	var filters models.FetchCourseRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	courses, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, courses)
}
