package main

import (
	"customers/config"
	"customers/delivery/handlers"
	"customers/delivery/routes"
	"customers/storage"

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
	db := storage.NewMySqlDB(config.DatabaseUrl)

	h1 := handlers.NewRegisterCustomerHandler(db)
	h2 := handlers.NewRegisterCustomerCompensateHandler(db)

	return handlers.NewHandlers(h1, h2)
}
