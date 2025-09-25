package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rubenalves-dev/taskuik/adapters/http/routes"
	"github.com/rubenalves-dev/taskuik/adapters/sqlite"
	"github.com/rubenalves-dev/taskuik/internal/config"
	"github.com/rubenalves-dev/taskuik/internal/middleware"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	dbClient := sqlite.NewClient()

	router := routes.NewRouter(gin.Default())

	router.Engine.Use(middleware.CORSMiddleware())

	router.AddDbClient(dbClient)
	router.RegisterRoutes()

	router.Run(config.HttpPort)
}
