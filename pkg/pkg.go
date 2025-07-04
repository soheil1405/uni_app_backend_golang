package pkg

import (
	"fmt"
	"uni_app/models"
	"uni_app/pkg/address"
	"uni_app/pkg/city"
	"uni_app/pkg/comment/comment"
	"uni_app/pkg/major"
	"uni_app/pkg/major_chart"
	"uni_app/pkg/notification"
	"uni_app/pkg/place"
	"uni_app/pkg/place_type"
	"uni_app/pkg/rating"
	"uni_app/pkg/role"
	"uni_app/pkg/route"
	"uni_app/pkg/student"
	"uni_app/pkg/student_passed_lesson"
	"uni_app/pkg/uni"
	"uni_app/pkg/user"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPkgs(db *gorm.DB, e echo.Group, cfg *env.Config) {
	migrateModels(db, cfg)
	address.Init(db, e, cfg)
	city.Init(db, e, cfg)
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
	user.Init(db, e, cfg)
	notification.InitNotification(e, db)
	comment.Init(db, e, cfg)
}

func migrateModels(db *gorm.DB, config *env.Config) {
	var (
		err error
	)
	if migration := config.GetBool("migration"); migration {
		fmt.Println("Migrating database...")
		if err = db.Debug().AutoMigrate(
			&models.Role{},
			&models.Place{},
			&models.PlaceType{},
			&models.Student{},
			&models.Token{},
			&models.Uni{},
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
			&models.Major{},
			&models.MajorsChart{},
			&models.Rating{},
			&models.Notification{},
			&models.NotificationTemplate{},
			&models.NotificationPreference{},
			&models.Comment{},
		); err != nil {
			panic(err)
		}

		fmt.Println("Migrating done ...")
	}

}
