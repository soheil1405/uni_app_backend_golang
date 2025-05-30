package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/major/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type MajorHandler struct {
	usecase usecase.MajorUsecase
}

func NewMajorHandler(usecase usecase.MajorUsecase, e echo.Group) {
	majorHandler := &MajorHandler{usecase}

	majorsRouteGroup := e.Group("/majors")
	majorsRouteGroup.POST("", majorHandler.CreateMajor)
	majorsRouteGroup.GET("/:id", majorHandler.GetMajorByID)
	majorsRouteGroup.PUT("/:id", majorHandler.UpdateMajor)
	majorsRouteGroup.DELETE("/:id", majorHandler.DeleteMajor)
	majorsRouteGroup.GET("", majorHandler.GetAllMajors)
}

func (h *MajorHandler) CreateMajor(c echo.Context) error {
	var major models.Major
	if err := c.Bind(&major); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateMajor(&major); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"major": major}, nil)
}

func (h *MajorHandler) GetMajorByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	major, err := h.usecase.GetMajorByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"major": major}, nil)
}

func (h *MajorHandler) UpdateMajor(c echo.Context) (err error) {
	var major models.Major
	if major.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateMajor(&major); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"major": major}, nil)
}

func (h *MajorHandler) DeleteMajor(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteMajor(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Major deleted"}, nil)
}

func (h *MajorHandler) GetAllMajors(c echo.Context) error {
	var request models.FetchMajorRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	majors, paginate, err := h.usecase.GetAllMajors(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"majors": majors}, paginate)
}
