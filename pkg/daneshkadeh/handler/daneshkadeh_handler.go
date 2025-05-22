package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/daneshkadeh/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type DaneshKadehHandler struct {
	usecase usecase.DaneshKadehUsecase
}

func NewDaneshKadehHandler(usecase usecase.DaneshKadehUsecase, e echo.Group) {
	daneshKadehHandler := &DaneshKadehHandler{usecase}

	daneshKadehsRouteGroup := e.Group("/daneshkadehs")
	daneshKadehsRouteGroup.POST("", daneshKadehHandler.CreateDaneshKadeh)
	daneshKadehsRouteGroup.GET("/:id", daneshKadehHandler.GetDaneshKadehByID)
	daneshKadehsRouteGroup.PUT("/:id", daneshKadehHandler.UpdateDaneshKadeh)
	daneshKadehsRouteGroup.DELETE("/:id", daneshKadehHandler.DeleteDaneshKadeh)
	daneshKadehsRouteGroup.GET("", daneshKadehHandler.GetAllDaneshKadehs)
}

func (h *DaneshKadehHandler) CreateDaneshKadeh(c echo.Context) error {
	var daneshKadeh models.DaneshKadeh
	if err := c.Bind(&daneshKadeh); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateDaneshKadeh(&daneshKadeh); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, daneshKadeh)
}

func (h *DaneshKadehHandler) GetDaneshKadehByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	daneshKadeh, err := h.usecase.GetDaneshKadehByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, daneshKadeh)
}

func (h *DaneshKadehHandler) UpdateDaneshKadeh(c echo.Context) (err error) {
	var daneshKadeh models.DaneshKadeh
	if daneshKadeh.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateDaneshKadeh(&daneshKadeh); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, daneshKadeh)
}

func (h *DaneshKadehHandler) DeleteDaneshKadeh(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteDaneshKadeh(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "DaneshKadeh deleted"})
}

func (h *DaneshKadehHandler) GetAllDaneshKadehs(c echo.Context) error {
	var request models.FetchDaneshKadehRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	daneshKadehs, paginate, err := h.usecase.GetAllDaneshKadehs(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"daneshkadehs": daneshKadehs,
		"meta":         paginate,
	})
}
