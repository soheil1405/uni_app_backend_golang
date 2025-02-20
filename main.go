package main

import (
	"encoding/json"
	"flag"
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
	// Define and parse the flag
	configPath := flag.String("-c", "config.json", "Path to the configuration file")
	flag.Parse()
	// Load configuration
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	echo := echo.New()

	db, err := database.Connection(&config.Database)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.Group("/api/v1")

	pkg.InitPkgs(db, *e)

	echo.Start(":8080")
}
