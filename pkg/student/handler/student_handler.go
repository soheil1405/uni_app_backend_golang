package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/student/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	usecase usecase.StudentUsecase
}

func NewStudentHandler(usecase usecase.StudentUsecase, e echo.Group) {
	studentHandler := &StudentHandler{usecase}

	studentsRouteGroup := e.Group("/students")
	studentsRouteGroup.POST("", studentHandler.CreateStudent)
	studentsRouteGroup.GET("/:id", studentHandler.GetStudentByID)
	studentsRouteGroup.PUT("/:id", studentHandler.UpdateStudent)
	studentsRouteGroup.DELETE("/:id", studentHandler.DeleteStudent)
	studentsRouteGroup.GET("", studentHandler.GetAllStudents)
	studentsRouteGroup.POST("/register", studentHandler.RegisterStudent)
	studentsRouteGroup.POST("/login", studentHandler.LoginStudent)
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
		ID  database.PID
		err error
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

func (h *StudentHandler) UpdateStudent(c echo.Context) (err error) {
	var student models.Student
	if student.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateStudent(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteStudent(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Student deleted"})
}

func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	var request models.FetchStudentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	students, paginate, err := h.usecase.GetAllStudents(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"students": students,
		"meta":     paginate,
	})
}

func (h *StudentHandler) RegisterStudent(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.RegisterStudent(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) LoginStudent(c echo.Context) error {
	var loginRequest struct {
		StudentCode database.PID `json:"student_code"`
		Password    string       `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	student, err := h.usecase.LoginStudent(loginRequest.StudentCode, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}
