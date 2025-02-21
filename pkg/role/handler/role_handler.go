package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/role/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	usecase usecases.RoleUsecase
}

func NewRoleHandler(usecase usecases.RoleUsecase, e echo.Group) {
	roleHandler := &RoleHandler{usecase}
	e.POST("/roles", roleHandler.CreateRole)
	e.GET("/roles/:id", roleHandler.GetRoleByID)
	e.PUT("/roles/:id", roleHandler.UpdateRole)
	e.DELETE("/roles/:id", roleHandler.DeleteRole)
	e.GET("/roles", roleHandler.GetAllRoles)

}

func (h *RoleHandler) CreateRole(c echo.Context) error {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateRole(&role); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetRoleByID(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	role, err := h.usecase.GetRoleByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) UpdateRole(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var role models.Role
	if err := c.Bind(&role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	role.ID = ID
	if err := h.usecase.UpdateRole(&role); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) DeleteRole(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteRole(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Role deleted"})
}

func (h *RoleHandler) GetAllRoles(c echo.Context) error {
	roles, err := h.usecase.GetAllRoles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, roles)
}
