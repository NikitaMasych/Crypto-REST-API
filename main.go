package main

import (
	"GenesisTask/config"
	"GenesisTask/emails"
	"GenesisTask/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	emails.EnsureFileExists()
	router := gin.Default()

	router.GET("/api/rate", routes.GetRate)
	router.POST("/api/subscribe", routes.PostSubscribe)
	router.POST("/api/sendEmails", routes.PostSendMessage)

	router.Run(config.Get().ServerURL)
}
