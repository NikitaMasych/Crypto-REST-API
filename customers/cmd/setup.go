package main

import (
	"customers/config"
	"customers/delivery/handlers"
	"customers/delivery/routes"
	"customers/storage"
	"os"

	"github.com/gin-gonic/gin"
)

func LaunchEngine() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	handlers := createHandlers()
	routes.InitRoutes(router, handlers)
	router.Run(config.ServerUrl)
}

func createHandlers() *handlers.Handlers {
	dbURL := os.Getenv("MYSQL_DSN")
	db := storage.NewMySqlDB(dbURL)

	h1 := handlers.NewRegisterCustomerHandler(db)
	h2 := handlers.NewRegisterCustomerCompensateHandler(db)

	return handlers.NewHandlers(h1, h2)
}
