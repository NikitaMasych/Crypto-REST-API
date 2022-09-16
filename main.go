package main

import (
	"GenesisTask/cache"
	"GenesisTask/config"
	"GenesisTask/platform"
	"GenesisTask/repository"
	"GenesisTask/routes"

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

	router.GET("/api/rate", routes.GetRate)
	router.POST("/api/subscribe", routes.PostSubscribe)
	router.POST("/api/sendEmails", routes.PostSendMessage)

	router.Run(config.Get().ServerURL)
}
