package handlers

import (
	"customers/types"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const globalID = "gid"

type RegisterCustomerHandler struct {
	db *gorm.DB
}

type registerCustomerRequest struct {
	EmailAddress string `json:"emailAddress"`
}

func NewRegisterCustomerHandler(db *gorm.DB) *RegisterCustomerHandler {
	return &RegisterCustomerHandler{db}
}

func (h *RegisterCustomerHandler) RegisterCustomer(c *gin.Context) interface{} {
	var customer registerCustomerRequest
	err := c.BindJSON(&customer)
	if err != nil {
		return dtmcli.ErrFailure
	}

	transactionId := c.Query(globalID)
	if err = h.db.Create(&types.Order{
		IDTransaction: transactionId,
		EmailAddress:  customer.EmailAddress,
		Status:        "Created",
	}).Error; err != nil {
		return err
	}

	return h.db.Save(&types.Customer{EmailAddress: customer.EmailAddress}).Error
}
