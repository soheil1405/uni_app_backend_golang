package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/uni/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type UniHandler struct {
	usecase usecases.UniUsecase
}

func NewUniHandler(usecase usecases.UniUsecase, e echo.Group) {
	uniHandler := &UniHandler{usecase}

	uniRouteGroup := e.Group("/unis")
	uniRouteGroup.POST("", uniHandler.CreateUni)
	uniRouteGroup.GET("/:id", uniHandler.GetUniByID)
	uniRouteGroup.PUT("/:id", uniHandler.UpdateUni)
	uniRouteGroup.DELETE("/:id", uniHandler.DeleteUni)
	uniRouteGroup.GET("", uniHandler.GetAllUnis)
}

func (h *UniHandler) CreateUni(c echo.Context) error {
	var uni models.Uni
	if err := c.Bind(&uni); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateUni(&uni); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, uni)
}

func (h *UniHandler) GetUniByID(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	uni, err := h.usecase.GetUniByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uni)
}

func (h *UniHandler) UpdateUni(c echo.Context) error {
	var (
		err error
		ID  database.PID
		uni models.Uni
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Bind(&uni); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	uni.ID = ID
	if err := h.usecase.UpdateUni(&uni); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uni)
}

func (h *UniHandler) DeleteUni(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteUni(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Uni deleted"})
}

func (h *UniHandler) GetAllUnis(c echo.Context) error {
	var request models.FetchUniRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	unis, paginate, err := h.usecase.GetAllUnis(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"unis": unis,
		"meta": paginate,
	})
}
