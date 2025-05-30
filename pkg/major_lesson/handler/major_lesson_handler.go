package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/major_lesson/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type MajorLessonHandler struct {
	usecase usecase.MajorLessonUsecase
}

func NewMajorLessonHandler(usecase usecase.MajorLessonUsecase, e echo.Group) {
	majorLessonHandler := &MajorLessonHandler{usecase}

	majorLessonsRouteGroup := e.Group("/major-lessons")
	majorLessonsRouteGroup.POST("", majorLessonHandler.CreateMajorLesson)
	majorLessonsRouteGroup.GET("/:id", majorLessonHandler.GetMajorLessonByID)
	majorLessonsRouteGroup.PUT("/:id", majorLessonHandler.UpdateMajorLesson)
	majorLessonsRouteGroup.DELETE("/:id", majorLessonHandler.DeleteMajorLesson)
	majorLessonsRouteGroup.GET("", majorLessonHandler.GetAllMajorLessons)
}

func (h *MajorLessonHandler) CreateMajorLesson(c echo.Context) error {
	var majorLesson models.MajorLesson
	if err := c.Bind(&majorLesson); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateMajorLesson(&majorLesson); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"major_lesson": majorLesson}, nil)
}

func (h *MajorLessonHandler) GetMajorLessonByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	majorLesson, err := h.usecase.GetMajorLessonByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"major_lesson": majorLesson}, nil)
}

func (h *MajorLessonHandler) UpdateMajorLesson(c echo.Context) (err error) {
	var majorLesson models.MajorLesson
	if majorLesson.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateMajorLesson(&majorLesson); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"major_lesson": majorLesson}, nil)
}

func (h *MajorLessonHandler) DeleteMajorLesson(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteMajorLesson(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Major lesson deleted"}, nil)
}

func (h *MajorLessonHandler) GetAllMajorLessons(c echo.Context) error {
	var request models.FetchMajorLessonRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	majorLessons, paginate, err := h.usecase.GetAllMajorLessons(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"major_lessons": majorLessons}, paginate)
}
