package main

import (
	"imagecolors/internal/app/server"
	"imagecolors/internal/models"
	"log"
	"os"

	"imagecolors/internal/app/runner"

	"gopkg.in/yaml.v3"
)

func LoadAppConfig() models.AppConfig {
	data, err := os.ReadFile("internal/config/appconfig.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config models.AppConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	config := LoadAppConfig()

	serverAddr := config.Server.Host + ":" + config.Server.Port

	runner.SetupRoutes()

	server := server.SetupServer(serverAddr)
	log.Printf("Server is listening on : %s", serverAddr)
	server.ListenAndServe()
}
