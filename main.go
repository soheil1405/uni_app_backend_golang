package main

import (
	"uni_app/database"
	"uni_app/pkg"

	"github.com/labstack/echo/v4"
)

func main() {
	echo := echo.New()

	db := database.Connection()

	e := echo.Group("/api/v1")

	pkg.InitPkgs(db, *e)

	echo.Start(":8080")
}
