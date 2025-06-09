package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/role/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

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
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateRole(&role); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"role": role}, nil)
}

func (h *RoleHandler) GetRoleByID(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	role, err := h.usecase.GetRoleByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"role": role}, nil)
}

func (h *RoleHandler) UpdateRole(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	var role models.Role
	if err := c.Bind(&role); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	role.ID = ID
	if err := h.usecase.UpdateRole(&role); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"role": role}, nil)
}

func (h *RoleHandler) DeleteRole(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteRole(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Role deleted"}, nil)
}

func (h *RoleHandler) GetAllRoles(c echo.Context) error {
	roles, err := h.usecase.GetAllRoles()
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"roles": roles}, nil)
}
