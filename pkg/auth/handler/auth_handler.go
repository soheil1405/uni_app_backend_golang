package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/auth/usecase"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
	group   echo.Group
}

func NewAuthHandler(usecase usecase.AuthUsecase, group echo.Group) {
	handler := &AuthHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *AuthHandler) initRoutes() {
	authGroup := h.group.Group("/auth")

	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
	authGroup.PUT("/:id", h.Update)
	authGroup.DELETE("/:id", h.Delete)
	authGroup.GET("/:id", h.GetByID)
	authGroup.GET("", h.List)
}

func (h *AuthHandler) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Register(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.usecase.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.ID = database.PID(id)
	if err := h.usecase.Update(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *AuthHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	user, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) List(c echo.Context) error {
	var filters models.FetchUserRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	users, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
