package pkg

import (
	"fmt"
	"uni_app/models"
	"uni_app/pkg/address"
	"uni_app/pkg/auth"
	"uni_app/pkg/city"
	daneshkadeh "uni_app/pkg/daneshkadeh"
	"uni_app/pkg/lesson"
	"uni_app/pkg/major"
	"uni_app/pkg/major_chart"
	"uni_app/pkg/place"
	"uni_app/pkg/place_type"
	"uni_app/pkg/rating"
	"uni_app/pkg/role"
	"uni_app/pkg/route"
	"uni_app/pkg/student"
	"uni_app/pkg/student_passed_lesson"
	"uni_app/pkg/uni"
	"uni_app/pkg/uni_major"
	"uni_app/pkg/user"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPkgs(db *gorm.DB, e echo.Group, cfg *env.Config) {
	migrateModels(db, cfg)
	address.Init(db, e, cfg)
	auth.Init(db, e, cfg)
	city.Init(db, e, cfg)
	daneshkadeh.Init(db, e, cfg)
	lesson.Init(db, e, cfg)
	major.Init(db, e, cfg)
	major_chart.Init(db, e, cfg)
	place.Init(db, e, cfg)
	place_type.Init(db, e, cfg)
	rating.Init(db, e, cfg)
	role.Init(db, e, cfg)
	route.Init(db, e, cfg)
	student.Init(db, e, cfg)
	student_passed_lesson.Init(db, e, cfg)
	uni.Init(db, e, cfg)
	uni_major.Init(db, e, cfg)
	user.Init(db, e, cfg)
}

func migrateModels(db *gorm.DB, config *env.Config) {
	var (
		err error
	)
	if migration := config.GetBool("migration"); migration {
		fmt.Println("Migrating database...")
		if err = db.Debug().AutoMigrate(
			&models.UserRole{},
			&models.Role{},
			&models.DaneshKadeh{},
			&models.DaneshKadehType{},
			&models.Place{},
			&models.PlaceType{},
			&models.Student{},
			&models.Token{},
			&models.Uni{},
			&models.UniMajor{},
			&models.User{},
			&models.Major{},
			&models.MajorsChart{},
			&models.Phone{},
			&models.Route{},
			&models.RouteGroup{},
			&models.AuthRules{},
			&models.ContactWay{},
			&models.City{},
			&models.Address{},
			&models.Lesson{},
			&models.Major{},
			&models.MajorsChart{},
			&models.StudentPassedLesson{},
			&models.Rating{},
		); err != nil {
			panic(err)
		}

		fmt.Println("Migrating done ...")
	}

}
