package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/address/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type AddressHandler struct {
	usecase usecase.AddressUsecase
}

func NewAddressHandler(usecase usecase.AddressUsecase, e echo.Group) {
	addressHandler := &AddressHandler{usecase}

	addressesRouteGroup := e.Group("/addresses")
	addressesRouteGroup.POST("", addressHandler.CreateAddress)
	addressesRouteGroup.GET("/:id", addressHandler.GetAddressByID)
	addressesRouteGroup.PUT("/:id", addressHandler.UpdateAddress)
	addressesRouteGroup.DELETE("/:id", addressHandler.DeleteAddress)
	addressesRouteGroup.GET("", addressHandler.GetAllAddresses)
}

func (h *AddressHandler) CreateAddress(c echo.Context) error {
	var address models.Address
	if err := c.Bind(&address); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateAddress(&address); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"address": address}, nil)
}

func (h *AddressHandler) GetAddressByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	address, err := h.usecase.GetAddressByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"address": address}, nil)
}

func (h *AddressHandler) UpdateAddress(c echo.Context) (err error) {
	var address models.Address
	if address.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateAddress(&address); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"address": address}, nil)
}

func (h *AddressHandler) DeleteAddress(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteAddress(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Address deleted"}, nil)
}

func (h *AddressHandler) GetAllAddresses(c echo.Context) error {
	var request models.FetchAddressRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	addresses, paginate, err := h.usecase.GetAllAddresses(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"addresses": addresses}, paginate)
}
