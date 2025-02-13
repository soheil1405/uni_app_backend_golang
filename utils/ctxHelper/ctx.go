package ctxHelper

import (
	"uni_app/database"

	"github.com/labstack/echo/v4"
)

func GetIDFromContxt(c echo.Context) (ID database.PID, err error) {
	ID = database.Parse(c.Param("id"))
	if !ID.IsValid() {
		return database.NilPID, database.ErrInvalidPID
	}

	return ID, nil
}
