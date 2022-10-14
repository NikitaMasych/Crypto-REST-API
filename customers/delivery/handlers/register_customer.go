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

type RegisterCustomerRequest struct {
	CustomerId   string `json:"customerId"`
	EmailAddress string `json:"emailAddress"`
}

func NewRegisterCustomerHandler(db *gorm.DB) *RegisterCustomerHandler {
	return &RegisterCustomerHandler{db}
}

func (h *RegisterCustomerHandler) RegisterCustomer(c *gin.Context) interface{} {
	var customer RegisterCustomerRequest
	err := c.BindJSON(&customer)
	if err != nil {
		return dtmcli.ErrFailure
	}

	transactionId := c.Query(globalID)
	if err = h.db.Create(&types.Order{
		TransactionId: transactionId,
		CustomerId:    customer.CustomerId,
		EmailAddress:  customer.EmailAddress,
		Status:        "Created",
	}).Error; err != nil {
		return err
	}

	err = h.db.Save(&types.Customer{CustomerId: customer.CustomerId,
		EmailAddress: customer.EmailAddress}).Error

	return err
}
