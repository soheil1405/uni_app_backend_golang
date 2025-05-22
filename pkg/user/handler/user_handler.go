package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/user/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase, e echo.Group) {
	userHandler := &UserHandler{usecase}

	// Public routes
	e.POST("/users/register", userHandler.RegisterUser)
	e.POST("/users/login", userHandler.LoginUser)

	// Protected routes
	users := e.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("/:id", userHandler.GetUserByID)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
	users.GET("", userHandler.GetAllUsers)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, err := h.usecase.GetUserByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var (
		user models.User
		err  error
	)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if user.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.DeleteUser(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	var request models.FetchRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	users, paginate, err := h.usecase.GetAllUsers(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"meta":  paginate,
	})
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.RegisterUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, err := h.usecase.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}
