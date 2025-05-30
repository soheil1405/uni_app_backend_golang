package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/auth/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase, e echo.Group) {
	authHandler := &AuthHandler{usecase}

	authRouteGroup := e.Group("/auth-rules")
	authRouteGroup.POST("", authHandler.CreateAuth)
	authRouteGroup.GET("/:id", authHandler.GetAuthByID)
	authRouteGroup.PUT("/:id", authHandler.UpdateAuth)
	authRouteGroup.DELETE("/:id", authHandler.DeleteAuth)
	authRouteGroup.GET("", authHandler.GetAllAuths)
}

func (h *AuthHandler) CreateAuth(c echo.Context) error {
	var auth models.AuthRules
	if err := c.Bind(&auth); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateAuth(&auth); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"auth": auth}, nil)
}

func (h *AuthHandler) GetAuthByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	auth, err := h.usecase.GetAuthByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"auth": auth}, nil)
}

func (h *AuthHandler) UpdateAuth(c echo.Context) (err error) {
	var auth models.AuthRules
	if auth.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateAuth(&auth); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"auth": auth}, nil)
}

func (h *AuthHandler) DeleteAuth(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteAuth(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Auth rule deleted"}, nil)
}

func (h *AuthHandler) GetAllAuths(c echo.Context) error {
	auths, err := h.usecase.GetAllAuths()
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"auths": auths}, nil)
}
