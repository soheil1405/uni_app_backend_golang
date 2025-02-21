package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/daneshkadeh/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type FacultyHandler struct {
	usecase usecases.FacultyUsecase
}

func NewDaneshKadehHandler(usecase usecases.FacultyUsecase, e echo.Group) {
	facultyHandler := &FacultyHandler{usecase}
	e.POST("/daneshkadeha", facultyHandler.CreateDaneshKadeh)
	e.GET("/daneshkadeha/:id", facultyHandler.GetDaneshKadehByID)
	e.PUT("/daneshkadeha/:id", facultyHandler.UpdateDaneshKadeh)
	e.DELETE("/daneshkadeha/:id", facultyHandler.DeleteDaneshKadeh)
	e.GET("/daneshkadeha", facultyHandler.GetAlldaneshkadeha)

}

func (h *FacultyHandler) CreateDaneshKadeh(c echo.Context) error {
	var faculty models.DaneshKadeh
	if err := c.Bind(&faculty); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateDaneshKadeh(&faculty); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, faculty)
}

func (h *FacultyHandler) GetDaneshKadehByID(c echo.Context) error {
	var (
		err         error
		daneshkadeh *models.DaneshKadeh
	)

	if daneshkadeh.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	daneshkadeh, err = h.usecase.GetDaneshKadehByID(c, daneshkadeh.ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, daneshkadeh)
}

func (h *FacultyHandler) UpdateDaneshKadeh(c echo.Context) error {
	var (
		err         error
		daneshkadeh models.DaneshKadeh
	)

	if err := c.Bind(&daneshkadeh); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if daneshkadeh.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.UpdateDaneshKadeh(&daneshkadeh); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, daneshkadeh)
}

func (h *FacultyHandler) DeleteDaneshKadeh(c echo.Context) error {
	var (
		err error
		id  database.PID
	)
	if id, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteDaneshKadeh(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Faculty deleted"})
}

func (h *FacultyHandler) GetAlldaneshkadeha(c echo.Context) error {
	daneshkadeha, err := h.usecase.GetAllDaneshKadeha()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, daneshkadeha)
}
