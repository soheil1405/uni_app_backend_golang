package handlers

import (
	"net/http"
	"strconv"
	"uni_app/models"
	usecases "uni_app/pkg/major/usecase"

	"github.com/labstack/echo/v4"
)

type MajorHandler struct {
	usecase usecases.MajorUsecase
}

func NewMajorHandler(usecase usecases.MajorUsecase, e echo.Group) {
	majorHandler := &MajorHandler{usecase}

	e.POST("/majors", majorHandler.CreateMajor)
	e.GET("/majors/:id", majorHandler.GetMajorByID)
	e.PUT("/majors/:id", majorHandler.UpdateMajor)
	e.DELETE("/majors/:id", majorHandler.DeleteMajor)
	e.GET("/majors", majorHandler.GetAllMajors)

}

func (h *MajorHandler) CreateMajor(c echo.Context) error {
	var major models.Major
	if err := c.Bind(&major); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateMajor(&major); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, major)
}
func (h *MajorHandler) GetMajorByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	major, err := h.usecase.GetMajorByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, major)
}

func (h *MajorHandler) UpdateMajor(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var major models.Major
	if err := c.Bind(&major); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	major.ID = uint(id)
	if err := h.usecase.UpdateMajor(&major); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, major)
}

func (h *MajorHandler) DeleteMajor(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.usecase.DeleteMajor(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Major deleted"})
}

func (h *MajorHandler) GetAllMajors(c echo.Context) error {
	majors, err := h.usecase.GetAllMajors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, majors)
}
