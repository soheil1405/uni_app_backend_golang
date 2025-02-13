package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/major_chart/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type ChartHandler struct {
	usecase usecases.ChartUsecase
}

func NewChartHandler(usecase usecases.ChartUsecase, e echo.Group) {
	chartHandler := &ChartHandler{usecase}

	e.POST("/charts", chartHandler.CreateChart)
	e.GET("/charts/:id", chartHandler.GetChartByID)
	e.PUT("/charts/:id", chartHandler.UpdateChart)
	e.DELETE("/charts/:id", chartHandler.DeleteChart)
	e.GET("/charts", chartHandler.GetAllCharts)

}

func (h *ChartHandler) CreateChart(c echo.Context) error {
	var chart models.MajorsChart
	if err := c.Bind(&chart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateChart(&chart); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, chart)
}

func (h *ChartHandler) GetChartByID(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	chart, err := h.usecase.GetChartByID(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chart)
}

func (h *ChartHandler) UpdateChart(c echo.Context) error {
	var (
		err   error
		ID    database.PID
		chart models.MajorsChart
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Bind(&chart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	chart.ID = ID
	if err := h.usecase.UpdateChart(&chart); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chart)
}

func (h *ChartHandler) DeleteChart(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	if err := h.usecase.DeleteChart(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Chart deleted"})
}

func (h *ChartHandler) GetAllCharts(c echo.Context) error {
	charts, err := h.usecase.GetAllCharts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, charts)
}
