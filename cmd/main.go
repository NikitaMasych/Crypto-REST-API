package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/cache"
	"GenesisTask/pkg/platform"
	"GenesisTask/pkg/repository"
	"GenesisTask/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cache.InitCache()

	platform.EnsureFileExists(config.Get().LoggerFile)
	platform.EnsureFileExists(config.Get().StorageFile)

	userRepo := repository.New()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(repository.AttachRepository(userRepo))

	routes.InitRoutes(router)

	router.Run(config.Get().ServerURL)
}
