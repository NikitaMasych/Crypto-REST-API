package routes

import (
	"customers/delivery/handlers"

	"github.com/dtm-labs/dtm/dtmutil"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, h *handlers.Handlers) {
	router.POST("/register-customer", dtmutil.WrapHandler2(h.RegisterCustomerHandler.RegisterCustomer))
	router.POST("/register-customer-compensate",
		dtmutil.WrapHandler2(h.RegisterCustomerCompensateHandler.RegisterCustomerCompensate))
}
