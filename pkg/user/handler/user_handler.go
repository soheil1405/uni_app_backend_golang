package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/user/usecase"
	"uni_app/utils/ctxHelper"
	helper "uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase, e echo.Group) {
	userHandler := &UserHandler{usecase}
	userRoutes := e.Group("/users")

	userRoutes.POST("/login", userHandler.Login)

	userRoutes.POST("", userHandler.CreateUser)
	userRoutes.GET("/:id", userHandler.GetUserByID)
	userRoutes.PUT("/:id", userHandler.UpdateUser)
	userRoutes.DELETE("/:id", userHandler.DeleteUser)
	userRoutes.GET("", userHandler.GetAllUsers)

}

func (h *UserHandler) Login(ctx echo.Context) error {
	var (
		request  models.UserLoginRequst
		response *models.Token
		err      error
	)

	if err = ctx.Bind(&request); err != nil {
		return helper.Reply(ctx, http.StatusBadRequest, err, nil, nil)
	}

	if response, err = h.usecase.Login(ctx, &request); err != nil {
		return helper.Reply(ctx, http.StatusBadRequest, err, nil, nil)
	}

	ctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   response.TokenKey,
		Path:    "/",
		Expires: response.ExpireTime,
		Secure:  false,
	})

	return helper.Reply(
		ctx,
		http.StatusOK,
		nil,
		map[string]interface{}{
			"token":     response.TokenKey,
			"token_new": response,
		},
		nil,
	)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return helper.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		return helper.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helper.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	user, err := h.usecase.GetUserByID(ID)
	if err != nil {
		return helper.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	return helper.Reply(c, http.StatusOK, nil, map[string]interface{}{"user": user}, nil)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var (
		err  error
		ID   database.PID
		user models.User
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helper.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := c.Bind(&user); err != nil {
		return helper.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	user.ID = ID
	if err := h.usecase.UpdateUser(&user); err != nil {
		return helper.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helper.Reply(c, http.StatusOK, nil, map[string]interface{}{"user": user}, nil)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var (
		err error
		ID  database.PID
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helper.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteUser(ID); err != nil {
		return helper.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helper.Reply(c, http.StatusOK, nil, map[string]interface{}{"id": ID}, nil)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		return helper.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helper.Reply(c, http.StatusOK, nil, map[string]interface{}{"users": users}, nil)

}
