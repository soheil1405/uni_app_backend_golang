package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/uni/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type UniHandler struct {
	usecase usecase.UniUsecase
}

func NewUniHandler(usecase usecase.UniUsecase, e echo.Group) {
	uniHandler := &UniHandler{usecase}

	unisRouteGroup := e.Group("/unis")
	unisRouteGroup.POST("", uniHandler.CreateUni)
	unisRouteGroup.GET("/:id", uniHandler.GetUniByID)
	unisRouteGroup.PUT("/:id", uniHandler.UpdateUni)
	unisRouteGroup.DELETE("/:id", uniHandler.DeleteUni)
	unisRouteGroup.GET("", uniHandler.GetAllUnis)
}

func (h *UniHandler) CreateUni(c echo.Context) error {
	var uni models.Uni
	if err := c.Bind(&uni); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateUni(&uni); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"uni": uni}, nil)
}

func (h *UniHandler) GetUniByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	uni, err := h.usecase.GetUniByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"uni": uni}, nil)
}

func (h *UniHandler) UpdateUni(c echo.Context) (err error) {
	var uni models.Uni
	if uni.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateUni(&uni); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"uni": uni}, nil)
}

func (h *UniHandler) DeleteUni(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteUni(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Uni deleted"}, nil)
}

func (h *UniHandler) GetAllUnis(c echo.Context) error {
	var request models.FetchUniRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	unis, paginate, err := h.usecase.GetAllUnis(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"unis": unis}, paginate)
}
