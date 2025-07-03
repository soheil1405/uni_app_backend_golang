package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/course_instance/usecase"

	"github.com/labstack/echo/v4"
)

type CourseInstanceHandler struct {
	usecase usecase.CourseInstanceUsecase
	group   echo.Group
}

func NewCourseInstanceHandler(usecase usecase.CourseInstanceUsecase, group echo.Group) {
	handler := &CourseInstanceHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *CourseInstanceHandler) initRoutes() {
	courseInstanceGroup := h.group.Group("/course-instances")

	courseInstanceGroup.POST("", h.Create)
	courseInstanceGroup.PUT("/:id", h.Update)
	courseInstanceGroup.DELETE("/:id", h.Delete)
	courseInstanceGroup.GET("/:id", h.GetByID)
	courseInstanceGroup.GET("", h.List)
	courseInstanceGroup.POST("/:id/students/:student_id", h.AddStudent)
	courseInstanceGroup.DELETE("/:id/students/:student_id", h.RemoveStudent)
	courseInstanceGroup.POST("/:id/schedules", h.AddSchedule)
	courseInstanceGroup.DELETE("/:id/schedules/:schedule_id", h.RemoveSchedule)
}

func (h *CourseInstanceHandler) Create(c echo.Context) error {
	var courseInstance models.CourseInstance
	if err := c.Bind(&courseInstance); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&courseInstance); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, courseInstance)
}

func (h *CourseInstanceHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var courseInstance models.CourseInstance
	if err := c.Bind(&courseInstance); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	courseInstance.ID = database.PID(id)
	if err := h.usecase.Update(&courseInstance); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, courseInstance)
}

func (h *CourseInstanceHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CourseInstanceHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	courseInstance, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, courseInstance)
}

func (h *CourseInstanceHandler) List(c echo.Context) error {
	var filters models.FetchCourseInstanceRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	courseInstances, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, courseInstances)
}

func (h *CourseInstanceHandler) AddStudent(c echo.Context) error {
	courseInstanceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid course instance ID")
	}

	studentID, err := strconv.ParseUint(c.Param("student_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid student ID")
	}

	if err := h.usecase.AddStudent(database.PID(courseInstanceID), database.PID(studentID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CourseInstanceHandler) RemoveStudent(c echo.Context) error {
	courseInstanceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid course instance ID")
	}

	studentID, err := strconv.ParseUint(c.Param("student_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid student ID")
	}

	if err := h.usecase.RemoveStudent(database.PID(courseInstanceID), database.PID(studentID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CourseInstanceHandler) AddSchedule(c echo.Context) error {
	courseInstanceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid course instance ID")
	}

	var schedule models.ClassSchedule
	if err := c.Bind(&schedule); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.AddSchedule(database.PID(courseInstanceID), &schedule); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, schedule)
}

func (h *CourseInstanceHandler) RemoveSchedule(c echo.Context) error {
	courseInstanceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid course instance ID")
	}

	scheduleID, err := strconv.ParseUint(c.Param("schedule_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid schedule ID")
	}

	if err := h.usecase.RemoveSchedule(database.PID(courseInstanceID), database.PID(scheduleID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
