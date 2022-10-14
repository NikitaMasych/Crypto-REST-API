package main

import (
	"orders/config"
	"orders/delivery/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(config.ServerUrl)
}
