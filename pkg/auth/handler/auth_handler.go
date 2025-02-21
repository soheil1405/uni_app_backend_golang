package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/auth/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase usecases.AuthUsecase
}

func NewAuthHandler(usecase usecases.AuthUsecase, e echo.Group) {
	authHandler := &AuthHandler{usecase}
	authRoutes := e.Group("/authorization")
	authRoutes.GET("/:id", authHandler.GetAuthByID)
	authRoutes.GET("", authHandler.GetAllAuthes)
	authRoutes.POST("", authHandler.CreateAuth)
	authRoutes.PUT("/:id", authHandler.UpdateAuth)
	authRoutes.DELETE("/:id", authHandler.DeleteAuth)
}

func (h *AuthHandler) CreateAuth(c echo.Context) error {
	var auth models.AuthRules
	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateAuth(&auth); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, auth)
}

func (h *AuthHandler) GetAuthByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	auth, err := h.usecase.GetAuthByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, auth)
}

func (h *AuthHandler) UpdateAuth(c echo.Context) (err error) {
	var auth models.AuthRules
	if auth.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateAuth(&auth); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, auth)
}

func (h *AuthHandler) DeleteAuth(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteAuth(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "auth deleted"})
}

func (h *AuthHandler) GetAllAuthes(c echo.Context) error {
	authes, err := h.usecase.GetAllAuths()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, authes)
}
