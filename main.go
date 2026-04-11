package main

import (
	"log"

	"clientcompany/config"
	"clientcompany/db"
	"clientcompany/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	db.Connect()

	if config.App.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routes.Register(r)

	log.Printf("starting server on :%s", config.App.AppPort)
	if err := r.Run(":" + config.App.AppPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
