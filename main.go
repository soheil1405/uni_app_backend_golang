package main

import (
	"flag"
	"fmt"
	"log"
	"uni_app/database"
	"uni_app/pkg"
	"uni_app/services/env"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

func main() {
	configPath := flag.String("c", "config.json", "Path to the configuration file")
	config := env.NewViperConfig(*configPath)
	dbConn := &database.Database{
		Host:     helpers.GetEnvDefault("DB_HOST", env.GetString("database.host")),
		Port:     helpers.GetEnvDefault("DB_PORT", env.GetString("database.port")),
		User:     helpers.GetEnvDefault("DB_USER", env.GetString("database.user")),
		Password: helpers.GetEnvDefault("DB_PASS", env.GetString("database.pass")),
		DBName:   helpers.GetEnvDefault("DB_NAME", env.GetString("database.name")),
		SSLMode:  helpers.GetEnvDefault("DB_NAME", env.GetString("database.sslmode")),
	}
	flag.Parse()

	db, err := database.Connection(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	echo := echo.New()
	e := echo.Group(env.GetString("api_version"))

	pkg.InitPkgs(db, *e, config)

	for _, route := range echo.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}

	echo.Start(":" + env.GetString("port"))
}
