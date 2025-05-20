package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/uni_type/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type UniTypeHandler struct {
	usecase usecase.UniTypeUsecase
}

func NewUniTypeHandler(usecase usecase.UniTypeUsecase) *UniTypeHandler {
	return &UniTypeHandler{usecase}
}

func (h *UniTypeHandler) CreateUniType(c echo.Context) error {
	var uniType models.UniType
	if err := c.Bind(&uniType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.CreateUniType(&uniType); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, uniType)
}

func (h *UniTypeHandler) GetUniTypeByID(c echo.Context) error {
	ID := ctxHelper.GetIDFromContext(c)
	useCache := ctxHelper.GetUseCacheFromContext(c)

	uniType, err := h.usecase.GetUniTypeByID(c, ID, useCache)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, uniType)
}

func (h *UniTypeHandler) UpdateUniType(c echo.Context) error {
	ID := ctxHelper.GetIDFromContext(c)

	var uniType models.UniType
	if err := c.Bind(&uniType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	uniType.ID = ID
	if err := h.usecase.UpdateUniType(&uniType); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, uniType)
}

func (h *UniTypeHandler) DeleteUniType(c echo.Context) error {
	ID := ctxHelper.GetIDFromContext(c)

	if err := h.usecase.DeleteUniType(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UniTypeHandler) GetAllUniTypes(c echo.Context) error {
	uniTypes, err := h.usecase.GetAllUniTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, uniTypes)
}

func (h *UniTypeHandler) RegisterRoutes(e *echo.Echo) {
	uniTypes := e.Group("/uni-types")
	uniTypes.POST("", h.CreateUniType)
	uniTypes.GET("/:id", h.GetUniTypeByID)
	uniTypes.PUT("/:id", h.UpdateUniType)
	uniTypes.DELETE("/:id", h.DeleteUniType)
	uniTypes.GET("", h.GetAllUniTypes)
} 