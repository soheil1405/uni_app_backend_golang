package main

import (
	"flag"
	"fmt"
	"log"
	"uni_app/database"
	"uni_app/pkg"
	"uni_app/services/env"
	"uni_app/utils/helpers"
	"uni_app/utils/middleware"

	mw "github.com/labstack/echo/v4/middleware"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func main() {
	configPath := flag.String("c", "config.json", "Path to the configuration file")
	config := env.NewViperConfig(*configPath)
	flag.Parse()
	dbCfg := &database.Database{
		Host:     helpers.GetEnvDefault("DB_HOST", env.GetString("database.host")),
		Port:     helpers.GetEnvDefault("DB_PORT", env.GetString("database.port")),
		User:     helpers.GetEnvDefault("DB_USER", env.GetString("database.user")),
		Password: helpers.GetEnvDefault("DB_PASS", env.GetString("database.pass")),
		DBName:   helpers.GetEnvDefault("DB_NAME", env.GetString("database.name")),
		SSLMode:  helpers.GetEnvDefault("DB_NAME", env.GetString("database.sslmode")),
	}

	db, err := database.Connection(dbCfg)
	if err != nil {
		log.Fatal(err)
	}

	echo := echo.New()
	apiVersion := env.GetString("api_version")
	e := echo.Group(apiVersion)
	// فعال کردن CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"}, // دامنه فرانت‌اند
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	
	if auth := env.GetBool("service.auth.active"); auth {
		jwtSecret := env.GetString("service.auth.secret")
		middle := middleware.InitMiddleware(e, db, config)
		e.Use(mw.JWTWithConfig(mw.JWTConfig{
			SigningKey: []byte(jwtSecret),
			Claims:     &jwt.StandardClaims{},
			Skipper:    middleware.RegisterSkipper,
		}),
			middle.SkipSetContext(middleware.RegisterSkipper), // Skip context middleware too
		)
	}

	pkg.InitPkgs(db, *e, config)

	for _, route := range echo.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}

	echo.Start(":" + env.GetString("port"))
}
