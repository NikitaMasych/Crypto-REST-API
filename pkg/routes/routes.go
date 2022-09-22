package routes

import (
	"GenesisTask/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/api/rate", handlers.GetRate)
	router.POST("/api/subscribe", handlers.Subscribe)
	router.POST("/api/sendEmails", handlers.SendMessage)
}
