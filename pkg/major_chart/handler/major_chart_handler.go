package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/major_chart/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type ChartHandler struct {
	usecase usecase.ChartUsecase
}

func NewChartHandler(usecase usecase.ChartUsecase, e echo.Group) {
	chartHandler := &ChartHandler{usecase}

	chartsRouteGroup := e.Group("/charts")
	chartsRouteGroup.POST("", chartHandler.CreateChart)
	chartsRouteGroup.GET("/:id", chartHandler.GetChartByID)
	chartsRouteGroup.PUT("/:id", chartHandler.UpdateChart)
	chartsRouteGroup.DELETE("/:id", chartHandler.DeleteChart)
	chartsRouteGroup.GET("", chartHandler.GetAllCharts)
}

func (h *ChartHandler) CreateChart(c echo.Context) error {
	var chart models.MajorsChart
	if err := c.Bind(&chart); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateChart(&chart); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"chart": chart}, nil)
}

func (h *ChartHandler) GetChartByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	chart, err := h.usecase.GetChartByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"chart": chart}, nil)
}

func (h *ChartHandler) UpdateChart(c echo.Context) (err error) {
	var chart models.MajorsChart
	if chart.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateChart(&chart); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"chart": chart}, nil)
}

func (h *ChartHandler) DeleteChart(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteChart(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Chart deleted"}, nil)
}

func (h *ChartHandler) GetAllCharts(c echo.Context) error {
	charts, err := h.usecase.GetAllCharts()
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"charts": charts}, nil)
}
