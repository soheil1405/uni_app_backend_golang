package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/uni_major/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type UniMajorHandler struct {
	usecase usecase.UniMajorUsecase
}

func NewUniMajorHandler(usecase usecase.UniMajorUsecase, e echo.Group) {
	uniMajorHandler := &UniMajorHandler{usecase}

	uniMajorsRouteGroup := e.Group("/uni-majors")
	uniMajorsRouteGroup.POST("", uniMajorHandler.CreateUniMajor)
	uniMajorsRouteGroup.GET("/:id", uniMajorHandler.GetUniMajorByID)
	uniMajorsRouteGroup.PUT("/:id", uniMajorHandler.UpdateUniMajor)
	uniMajorsRouteGroup.DELETE("/:id", uniMajorHandler.DeleteUniMajor)
	uniMajorsRouteGroup.GET("", uniMajorHandler.GetAllUniMajors)

}

func (h *UniMajorHandler) CreateUniMajor(c echo.Context) error {
	var uniMajor models.UniMajor
	if err := c.Bind(&uniMajor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateUniMajor(&uniMajor); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, uniMajor)
}

func (h *UniMajorHandler) GetUniMajorByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	uniMajor, err := h.usecase.GetUniMajorByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uniMajor)
}

func (h *UniMajorHandler) UpdateUniMajor(c echo.Context) (err error) {
	var uniMajor models.UniMajor
	if uniMajor.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateUniMajor(&uniMajor); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uniMajor)
}

func (h *UniMajorHandler) DeleteUniMajor(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteUniMajor(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "UniMajor deleted"})
}

func (h *UniMajorHandler) GetAllUniMajors(c echo.Context) error {
	uniMajors, err := h.usecase.GetAllUniMajors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, uniMajors)
} 