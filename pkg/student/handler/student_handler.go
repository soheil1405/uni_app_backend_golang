package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/student/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	usecase usecases.StudentUsecase
}

func NewStudentHandler(usecase usecases.StudentUsecase, e echo.Group) {
	studentHandler := &StudentHandler{usecase}
	studentRoutes := e.Group("/students")
	studentRoutes.GET("/", studentHandler.GetAllStudents)
	studentRoutes.GET("/:id", studentHandler.GetStudentByID)
	studentRoutes.POST("/", studentHandler.CreateStudent)
	studentRoutes.PUT("/:id", studentHandler.UpdateStudent)
	// studentRoutes.DELETE("/:id", userHandler.DeleteStudent)
}

func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var student models.Student

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateStudent(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) GetStudentByID(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	student, err := h.usecase.GetStudentByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	var (
		err     error
		ID      database.PID
		student models.Student
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	student.ID = ID
	if err := h.usecase.UpdateStudent(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteStudent(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}

func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	students, err := h.usecase.GetAllStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, students)
}
