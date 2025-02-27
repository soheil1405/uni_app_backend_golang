package pkg

import (
	"fmt"
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
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPkgs(db *gorm.DB, e echo.Group, cfg *env.Config) {
	migrateModels(db, cfg)
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
func migrateModels(db *gorm.DB, config *env.Config) {
	var (
	// interfaces []interface{}
	// err error
	)
	if migration := config.Viper.GetBool("migration"); migration {
		// interfaces = []interface{}{
		// 	&models.Uni{},
		// 	&models.User{},
		// 	&models.Route{},
		// 	&models.RouteGroup{},
		// 	&models.AuthRules{},
		// 	&models.User{},
		// 	&models.Config{},
		// 	&models.ContactWay{},
		// 	&models.DaneshKadeh{},
		// 	&models.DaneshKadehType{},
		// 	&models.UniMajor{},
		// 	&models.FailedJob{},
		// 	&models.Major{},
		// 	&models.MajorsChart{},
		// 	&models.Phone{},
		// 	&models.Place{},
		// 	&models.PlaceType{},
		// 	// &models.StudentRole{},
		// 	&models.Student{},
		// 	&models.UserRole{},
		// 	&models.Token{},
		// }
		fmt.Println("Migrating database...")
		// db.Debug().AutoMigrate(&models.UserRole{})
		// db.Debug().AutoMigrate(&models.Role{})
		// // db.Debug().AutoMigrate(&models.DaneshKadeh{})
		// // db.Debug().AutoMigrate(&models.Uni{})
		db.Debug().AutoMigrate(&models.User{})
		// db.Debug().AutoMigrate(&models.Address{})
		// db.Debug().AutoMigrate(&models.Route{})
		// db.Debug().AutoMigrate(&models.RouteGroup{})
		// db.Debug().AutoMigrate(&models.AuthRules{})
		// db.Debug().AutoMigrate(&models.ContactWay{})
		// db.Debug().AutoMigrate(&models.Major{})
		// db.Debug().AutoMigrate(&models.MajorsChart{})
		// db.Debug().AutoMigrate(&models.Phone{})
		// db.Debug().AutoMigrate(&models.Place{})
		// db.Debug().AutoMigrate(&models.PlaceType{})
		// db.Debug().AutoMigrate(&models.Token{})
		// db.Debug().AutoMigrate(&models.UniMajor{})

		fmt.Println("Migrating done ...")

	}

	// database.DoMigrate(db, interfaces)
}
