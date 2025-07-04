package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/student/usecase"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	usecase usecase.StudentUsecase
	group   echo.Group
}

func NewStudentHandler(usecase usecase.StudentUsecase, group echo.Group) {
	handler := &StudentHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *StudentHandler) initRoutes() {
	studentGroup := h.group.Group("/students")

	studentGroup.POST("", h.Create)
	studentGroup.PUT("/:id", h.Update)
	studentGroup.DELETE("/:id", h.Delete)
	studentGroup.GET("/:id", h.GetByID)
	studentGroup.GET("", h.List)
}

func (h *StudentHandler) Create(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var student models.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	student.ID = database.PID(id)
	if err := h.usecase.Update(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *StudentHandler) GetByID(c echo.Context) error {
	id := database.Parse(c.Param("id"))
	if !id.IsValid() {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	student, err := h.usecase.GetByID(c, id, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) List(c echo.Context) error {
	var filters models.FetchStudentRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	students, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) RegisterStudent(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.RegisterStudent(&student); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"student": student}, nil)
}

func (h *StudentHandler) LoginStudent(c echo.Context) error {
	var loginRequest struct {
		StudentCode database.PID `json:"student_code"`
		Password    string       `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	student, err := h.usecase.LoginStudent(loginRequest.StudentCode, loginRequest.Password)
	if err != nil {
		return helpers.Reply(c, http.StatusUnauthorized, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"student": student}, nil)
}
