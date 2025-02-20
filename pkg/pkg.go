package pkg

import (
	"uni_app/pkg/city"
	daneshkadeh "uni_app/pkg/daneshkadeh"
	"uni_app/pkg/major"
	"uni_app/pkg/major_chart"
	"uni_app/pkg/place"
	"uni_app/pkg/place_type"
	"uni_app/pkg/role"
	"uni_app/pkg/uni"
	"uni_app/pkg/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPkgs(db *gorm.DB, e echo.Group) {
	uni.Init(db, e)
	city.Init(db, e)
	daneshkadeh.Init(db, e)
	major.Init(db, e)
	major_chart.Init(db, e)
	place.Init(db, e)
	place_type.Init(db, e)
	role.Init(db, e)
	user.Init(db, e)
}
