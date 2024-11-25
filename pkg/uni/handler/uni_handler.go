package handlers

import (
	"net/http"
	"strconv"
	"uni_app/models"
	usecases "uni_app/pkg/uni/usecase"

	"github.com/labstack/echo/v4"
)

type UniHandler struct {
	usecase usecases.UniUsecase
}

func NewUniHandler(usecase usecases.UniUsecase, e echo.Group) {
	uniHandler := &UniHandler{usecase}
	e.POST("/unis", uniHandler.CreateUni)
	e.GET("/unis/:id", uniHandler.GetUniByID)
	e.PUT("/unis/:id", uniHandler.UpdateUni)
	e.DELETE("/unis/:id", uniHandler.DeleteUni)
	e.GET("/unis", uniHandler.GetAllUnis)
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	uni, err := h.usecase.GetUniByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uni)
}

func (h *UniHandler) UpdateUni(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var uni models.Uni
	if err := c.Bind(&uni); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	uni.ID = uint(id)
	if err := h.usecase.UpdateUni(&uni); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uni)
}

func (h *UniHandler) DeleteUni(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.usecase.DeleteUni(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Uni deleted"})
}

func (h *UniHandler) GetAllUnis(c echo.Context) error {
	unis, err := h.usecase.GetAllUnis()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, unis)
}
