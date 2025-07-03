package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/building/usecase"

	"github.com/labstack/echo/v4"
)

type BuildingHandler struct {
	usecase usecase.BuildingUsecase
	group   echo.Group
}

func NewBuildingHandler(usecase usecase.BuildingUsecase, group echo.Group) {
	handler := &BuildingHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *BuildingHandler) initRoutes() {
	buildingGroup := h.group.Group("/buildings")

	buildingGroup.POST("", h.Create)
	buildingGroup.PUT("/:id", h.Update)
	buildingGroup.DELETE("/:id", h.Delete)
	buildingGroup.GET("/:id", h.GetByID)
	buildingGroup.GET("", h.List)
}

func (h *BuildingHandler) Create(c echo.Context) error {
	var building models.Building
	if err := c.Bind(&building); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&building); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, building)
}

func (h *BuildingHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var building models.Building
	if err := c.Bind(&building); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	building.ID = database.PID(id)
	if err := h.usecase.Update(&building); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, building)
}

func (h *BuildingHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *BuildingHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	building, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, building)
}

func (h *BuildingHandler) List(c echo.Context) error {
	var filters models.FetchBuildingRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	buildings, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, buildings)
}
