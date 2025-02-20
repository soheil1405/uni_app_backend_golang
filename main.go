package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg"

	"github.com/labstack/echo/v4"
)

// loadConfig reads the configuration file
func loadConfig(path string) (*models.Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	configPath := flag.String("c", "config.json", "Path to the configuration file")
	flag.Parse()
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.Connection(&config.Database)
	if err != nil {
		log.Fatal(err)
	}

	echo := echo.New()
	e := echo.Group(config.ApiVersion)

	pkg.InitPkgs(db, *e, config)

	for _, route := range echo.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}

	echo.Start(":" + config.Port)
}
