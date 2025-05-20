package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/address/usecase"
	"uni_app/utils/ctxHelper"

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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateAddress(&address); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, address)
}

func (h *AddressHandler) GetAddressByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	address, err := h.usecase.GetAddressByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, address)
}

func (h *AddressHandler) UpdateAddress(c echo.Context) (err error) {
	var address models.Address
	if address.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateAddress(&address); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, address)
}

func (h *AddressHandler) DeleteAddress(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteAddress(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Address deleted"})
}

func (h *AddressHandler) GetAllAddresses(c echo.Context) error {
	addresses, err := h.usecase.GetAllAddresses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, addresses)
} 