package handlers

import (
	"customers/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterCustomerCompensateHandler struct {
	db *gorm.DB
}

func NewRegisterCustomerCompensateHandler(db *gorm.DB) *RegisterCustomerCompensateHandler {
	return &RegisterCustomerCompensateHandler{db}
}

func (h *RegisterCustomerCompensateHandler) RegisterCustomerCompensate(c *gin.Context) interface{} {
	transactionId := c.Query(globalID)

	return h.db.
		Model(&types.Order{}).
		Where("id_transaction = ?", transactionId).
		Update("status", "Failed").
		Limit(1).
		Error
}
