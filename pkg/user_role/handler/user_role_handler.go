package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/user_role/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type UserRoleHandler struct {
	usecase usecase.UserRoleUsecase
}

func NewUserRoleHandler(usecase usecase.UserRoleUsecase, e echo.Group) {
	userRoleHandler := &UserRoleHandler{usecase}

	userRolesRouteGroup := e.Group("/user-roles")
	userRolesRouteGroup.POST("", userRoleHandler.CreateUserRole)
	userRolesRouteGroup.GET("/:id", userRoleHandler.GetUserRoleByID)
	userRolesRouteGroup.PUT("/:id", userRoleHandler.UpdateUserRole)
	userRolesRouteGroup.DELETE("/:id", userRoleHandler.DeleteUserRole)
	userRolesRouteGroup.GET("", userRoleHandler.GetAllUserRoles)

}

func (h *UserRoleHandler) CreateUserRole(c echo.Context) error {
	var userRole models.UserRole
	if err := c.Bind(&userRole); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateUserRole(&userRole); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, userRole)
}

func (h *UserRoleHandler) GetUserRoleByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	userRole, err := h.usecase.GetUserRoleByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, userRole)
}

func (h *UserRoleHandler) UpdateUserRole(c echo.Context) (err error) {
	var userRole models.UserRole
	if userRole.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateUserRole(&userRole); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, userRole)
}

func (h *UserRoleHandler) DeleteUserRole(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteUserRole(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "UserRole deleted"})
}

func (h *UserRoleHandler) GetAllUserRoles(c echo.Context) error {
	userRoles, err := h.usecase.GetAllUserRoles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, userRoles)
} 