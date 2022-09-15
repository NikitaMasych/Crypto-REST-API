package main

import (
	"GenesisTask/cache"
	"GenesisTask/config"
	"GenesisTask/platform"
	"GenesisTask/repository"
	"GenesisTask/routes"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cache.InitCache()

	platform.EnsureFileExists(config.Get().LoggerFile)
	platform.EnsureFileExists(config.Get().StorageFile)

	userRepo := repository.New()

	router := gin.Default()

	router.Use(repository.AttachRepository(userRepo))

	router.GET("/api/rate", routes.GetRate)
	router.POST("/api/subscribe", routes.PostSubscribe)
	router.POST("/api/sendEmails", routes.PostSendMessage)

	router.Run(config.Get().ServerURL)
}
