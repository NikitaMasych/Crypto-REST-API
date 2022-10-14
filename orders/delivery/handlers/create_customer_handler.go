package handlers

import (
	"net/http"
	"orders/config"
	"orders/types"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer types.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	globalTransactionId := dtmcli.MustGenGid(config.DtmCoordinatorAddress)
	err := dtmcli.
		NewSaga(config.DtmCoordinatorAddress, globalTransactionId).
		Add(config.CustomersServerURL+"/register-customer",
			config.CustomersServerURL+"/register-customer-compensate", customer).
		Submit()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"gid": globalTransactionId})
	} else {
		c.JSON(http.StatusOK, gin.H{"gid": globalTransactionId})
	}
}
