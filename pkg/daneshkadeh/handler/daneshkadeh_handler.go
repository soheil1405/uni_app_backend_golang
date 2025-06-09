package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/daneshkadeh/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type DaneshKadehHandler struct {
	usecase usecase.DaneshKadehUsecase
}

func NewDaneshKadehHandler(usecase usecase.DaneshKadehUsecase, e echo.Group) {
	daneshKadehHandler := &DaneshKadehHandler{usecase}

	daneshKadehaRouteGroup := e.Group("/daneshkadeha")
	daneshKadehaRouteGroup.POST("", daneshKadehHandler.CreateDaneshKadeh)
	daneshKadehaRouteGroup.GET("/:id", daneshKadehHandler.GetDaneshKadehByID)
	daneshKadehaRouteGroup.PUT("/:id", daneshKadehHandler.UpdateDaneshKadeh)
	daneshKadehaRouteGroup.DELETE("/:id", daneshKadehHandler.DeleteDaneshKadeh)
	daneshKadehaRouteGroup.GET("", daneshKadehHandler.GetAllDaneshKadeha)
}

func (h *DaneshKadehHandler) CreateDaneshKadeh(c echo.Context) error {
	var daneshKadeh models.DaneshKadeh
	if err := c.Bind(&daneshKadeh); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateDaneshKadeh(&daneshKadeh); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"daneshkadeh": daneshKadeh}, nil)
}

func (h *DaneshKadehHandler) GetDaneshKadehByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	daneshKadeh, err := h.usecase.GetDaneshKadehByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"daneshkadeh": daneshKadeh}, nil)
}

func (h *DaneshKadehHandler) UpdateDaneshKadeh(c echo.Context) (err error) {
	var daneshKadeh models.DaneshKadeh
	if daneshKadeh.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateDaneshKadeh(&daneshKadeh); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"daneshkadeh": daneshKadeh}, nil)
}

func (h *DaneshKadehHandler) DeleteDaneshKadeh(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteDaneshKadeh(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "DaneshKadeh deleted"}, nil)
}

func (h *DaneshKadehHandler) GetAllDaneshKadeha(c echo.Context) error {
	var request models.FetchDaneshKadehRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	daneshKadeha, paginate, err := h.usecase.GetAllDaneshKadehs(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"daneshkadeha": daneshKadeha}, paginate)
}
