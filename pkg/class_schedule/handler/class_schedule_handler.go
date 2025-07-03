package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/class_schedule/usecase"

	"github.com/labstack/echo/v4"
)

type ClassScheduleHandler struct {
	usecase usecase.ClassScheduleUsecase
	group   echo.Group
}

func NewClassScheduleHandler(usecase usecase.ClassScheduleUsecase, group echo.Group) {
	handler := &ClassScheduleHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *ClassScheduleHandler) initRoutes() {
	classScheduleGroup := h.group.Group("/class-schedules")

	classScheduleGroup.POST("", h.Create)
	classScheduleGroup.PUT("/:id", h.Update)
	classScheduleGroup.DELETE("/:id", h.Delete)
	classScheduleGroup.GET("/:id", h.GetByID)
	classScheduleGroup.GET("", h.List)
}

func (h *ClassScheduleHandler) Create(c echo.Context) error {
	var classSchedule models.ClassSchedule
	if err := c.Bind(&classSchedule); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&classSchedule); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, classSchedule)
}

func (h *ClassScheduleHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var classSchedule models.ClassSchedule
	if err := c.Bind(&classSchedule); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	classSchedule.ID = database.PID(id)
	if err := h.usecase.Update(&classSchedule); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, classSchedule)
}

func (h *ClassScheduleHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ClassScheduleHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	classSchedule, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, classSchedule)
}

func (h *ClassScheduleHandler) List(c echo.Context) error {
	var filters models.FetchClassScheduleRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	classSchedules, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, classSchedules)
}
