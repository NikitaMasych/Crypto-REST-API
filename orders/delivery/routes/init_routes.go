package routes

import (
	"orders/delivery/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/create-customer", handlers.CreateCustomer)
}
