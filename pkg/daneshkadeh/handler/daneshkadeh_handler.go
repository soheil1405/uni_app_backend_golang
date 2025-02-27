package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/daneshkadeh/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
)

type FacultyHandler struct {
	usecase usecases.FacultyUsecase
}

func NewDaneshKadehHandler(usecase usecases.FacultyUsecase, e echo.Group) {
	facultyHandler := &FacultyHandler{usecase}

	daneshKadehRouteGroups := e.Group("/daneshkadeh")
	daneshKadehRouteGroups.GET("", facultyHandler.GetAlldaneshkadeha)
	daneshKadehRouteGroups.GET("/:id", facultyHandler.GetDaneshKadehByID)
	daneshKadehRouteGroups.POST("", facultyHandler.CreateDaneshKadeh)
	daneshKadehRouteGroups.PUT("/:id", facultyHandler.UpdateDaneshKadeh)
	daneshKadehRouteGroups.DELETE("/:id", facultyHandler.DeleteDaneshKadeh)
}

func (h *FacultyHandler) CreateDaneshKadeh(c echo.Context) error {
	var faculty models.DaneshKadeh
	if err := c.Bind(&faculty); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateDaneshKadeh(&faculty); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, faculty)
}

func (h *FacultyHandler) GetDaneshKadehByID(c echo.Context) error {
	var (
		MyErr helpers.MyError
		resp  map[string]interface{}
	)

	id, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	resp, MyErr = h.usecase.GetDaneshKadehByID(c, id, false)
	if MyErr.Err != nil {
		return helpers.Reply(c, MyErr.Code, MyErr.Err, resp, nil)
	}
	
	return helpers.Reply(c, MyErr.Code, MyErr.Err, resp, nil)
}

func (h *FacultyHandler) UpdateDaneshKadeh(c echo.Context) error {
	var (
		err         error
		daneshkadeh models.DaneshKadeh
	)

	if err := c.Bind(&daneshkadeh); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if daneshkadeh.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.UpdateDaneshKadeh(&daneshkadeh); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, daneshkadeh)
}

func (h *FacultyHandler) DeleteDaneshKadeh(c echo.Context) error {
	var (
		err error
		id  database.PID
	)
	if id, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteDaneshKadeh(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Faculty deleted"})
}

func (h *FacultyHandler) GetAlldaneshkadeha(c echo.Context) error {
	var (
		err  helpers.MyError
		req  *models.FetchRequest
		resp map[string]interface{}
		meta *templates.PaginateTemplate
	)

	if req, err = helpers.BindFetchRequestFromCtx(c); err.Err != nil {
		return helpers.Reply(c, err.Code, err.Err, nil, nil)
	}

	resp, meta, err = h.usecase.GetAllDaneshKadeha(c, *req)
	return helpers.Reply(c, err.Code, err.Err, resp, meta)
}
