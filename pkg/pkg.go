package pkg

import (
	"uni_app/models"
	"uni_app/pkg/city"
	daneshkadeh "uni_app/pkg/daneshkadeh"
	"uni_app/pkg/major"
	"uni_app/pkg/major_chart"
	"uni_app/pkg/place"
	"uni_app/pkg/place_type"
	"uni_app/pkg/role"
	"uni_app/pkg/student"
	"uni_app/pkg/uni"
	"uni_app/pkg/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPkgs(db *gorm.DB, e echo.Group, cfg *models.Config) {
	uni.Init(db, e, cfg)
	city.Init(db, e, cfg)
	daneshkadeh.Init(db, e, cfg)
	major.Init(db, e, cfg)
	major_chart.Init(db, e, cfg)
	place.Init(db, e, cfg)
	place_type.Init(db, e, cfg)
	role.Init(db, e, cfg)
	user.Init(db, e, cfg)
	student.Init(db, e, cfg)
}
