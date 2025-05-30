package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/uni_major/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

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
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateUniMajor(&uniMajor); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"uni_major": uniMajor}, nil)
}

func (h *UniMajorHandler) GetUniMajorByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	uniMajor, err := h.usecase.GetUniMajorByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"uni_major": uniMajor}, nil)
}

func (h *UniMajorHandler) UpdateUniMajor(c echo.Context) (err error) {
	var uniMajor models.UniMajor
	if uniMajor.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateUniMajor(&uniMajor); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"uni_major": uniMajor}, nil)
}

func (h *UniMajorHandler) DeleteUniMajor(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteUniMajor(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "UniMajor deleted"}, nil)
}

func (h *UniMajorHandler) GetAllUniMajors(c echo.Context) error {
	var request models.FetchUniMajorRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	uniMajors, paginate, err := h.usecase.GetAllUniMajors(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"uni_majors": uniMajors}, paginate)
}
