package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/department/usecase"

	"github.com/labstack/echo/v4"
)

type DepartmentHandler struct {
	usecase usecase.DepartmentUsecase
	group   echo.Group
}

func NewDepartmentHandler(usecase usecase.DepartmentUsecase, group echo.Group) {
	handler := &DepartmentHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *DepartmentHandler) initRoutes() {
	departmentGroup := h.group.Group("/departments")

	departmentGroup.POST("", h.Create)
	departmentGroup.PUT("/:id", h.Update)
	departmentGroup.DELETE("/:id", h.Delete)
	departmentGroup.GET("/:id", h.GetByID)
	departmentGroup.GET("", h.List)
	departmentGroup.POST("/:id/head/:teacher_id", h.SetHead)
	departmentGroup.POST("/:id/teachers/:teacher_id", h.AddTeacher)
	departmentGroup.DELETE("/:id/teachers/:teacher_id", h.RemoveTeacher)
	departmentGroup.POST("/:id/majors/:major_id", h.AddMajor)
	departmentGroup.DELETE("/:id/majors/:major_id", h.RemoveMajor)
}

func (h *DepartmentHandler) Create(c echo.Context) error {
	var department models.Department
	if err := c.Bind(&department); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&department); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, department)
}

func (h *DepartmentHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var department models.Department
	if err := c.Bind(&department); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	department.ID = database.PID(id)
	if err := h.usecase.Update(&department); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, department)
}

func (h *DepartmentHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *DepartmentHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	department, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, department)
}

func (h *DepartmentHandler) List(c echo.Context) error {
	var filters models.FetchDepartmentRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	departments, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, departments)
}

func (h *DepartmentHandler) SetHead(c echo.Context) error {
	departmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid department ID")
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid teacher ID")
	}

	if err := h.usecase.SetHead(database.PID(departmentID), database.PID(teacherID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *DepartmentHandler) AddTeacher(c echo.Context) error {
	departmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid department ID")
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid teacher ID")
	}

	if err := h.usecase.AddTeacher(database.PID(departmentID), database.PID(teacherID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *DepartmentHandler) RemoveTeacher(c echo.Context) error {
	departmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid department ID")
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid teacher ID")
	}

	if err := h.usecase.RemoveTeacher(database.PID(departmentID), database.PID(teacherID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *DepartmentHandler) AddMajor(c echo.Context) error {
	departmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid department ID")
	}

	majorID, err := strconv.ParseUint(c.Param("major_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid major ID")
	}

	if err := h.usecase.AddMajor(database.PID(departmentID), database.PID(majorID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *DepartmentHandler) RemoveMajor(c echo.Context) error {
	departmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid department ID")
	}

	majorID, err := strconv.ParseUint(c.Param("major_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid major ID")
	}

	if err := h.usecase.RemoveMajor(database.PID(departmentID), database.PID(majorID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
