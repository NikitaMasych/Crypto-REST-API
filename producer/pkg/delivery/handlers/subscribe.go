package handlers

import (
	"errors"
	"producer/pkg/application"
	"producer/pkg/delivery/presentors"
	"producer/pkg/domain/models"
	custom "producer/pkg/errors"
	"producer/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SubscribeHandler struct {
	subscriptionRepository application.SubscriptionRepository
	logger                 application.Logger
}

func NewSubscribeHandler(r application.SubscriptionRepository,
	l application.Logger) *SubscribeHandler {
	return &SubscribeHandler{r, l}
}

type subscriptionRequest struct {
	EmailAddress string `json:"email" binding:"required"`
	Base         string `json:"base"  binding:"required"`
	Quote        string `json:"quote" binding:"required"`
}

func (h *SubscribeHandler) Subscribe(c *gin.Context) {
	var requestData subscriptionRequest
	if err := c.BindJSON(&requestData); err != nil {
		h.logger.LogError(err)
		presentors.PresentErrorJSON(c)
		h.logger.LogDebug("Presented JSON error")
	}
	address := models.NewEmailAddress(requestData.EmailAddress)
	pair := models.NewCurrencyPair(requestData.Base, requestData.Quote)
	user := models.NewUser(*address, []models.CurrencyPair{*pair})
	h.subscribe(c, user)
	utils.CreateCustomer(*address, h.logger)
}

func (h *SubscribeHandler) subscribe(c *gin.Context, user *models.User) {
	err := h.subscriptionRepository.Subscribe(*user)
	if err == nil {
		presentors.PresentSubscriptionJSON(c)
		h.logger.LogDebug("Presented subscription JSON")
	} else {
		h.logger.LogError(err)
		if errors.Is(err, custom.ErrAlreadySubscribed) {
			presentors.PresentSubscriptionConflictJSON(c)
			h.logger.LogDebug("Presented subscription conflict JSON")
		} else {
			presentors.PresentErrorJSON(c)
			h.logger.LogDebug("Presented JSON error")
		}
	}
}
