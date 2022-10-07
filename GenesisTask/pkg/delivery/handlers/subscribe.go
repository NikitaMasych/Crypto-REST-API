package handlers

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/presentors"
	"GenesisTask/pkg/domain/models"
	custom "GenesisTask/pkg/errors"
	"errors"

	"github.com/gin-gonic/gin"
)

type SubscribeHandler struct {
	subscriptionRepository application.SubscriptionRepository
}

func NewSubscribeHandler(r application.SubscriptionRepository) *SubscribeHandler {
	return &SubscribeHandler{r}
}

type subscriptionRequest struct {
	EmailAddress string `json:"email" binding:"required"`
	Base         string `json:"base"  binding:"required"`
	Quote        string `json:"quote" binding:"required"`
}

func (h *SubscribeHandler) Subscribe(c *gin.Context) {
	var requestData subscriptionRequest
	if err := c.BindJSON(&requestData); err != nil {
		presentors.PresentErrorJSON(c)
	}
	address := models.NewEmailAddress(requestData.EmailAddress)
	pair := models.NewCurrencyPair(requestData.Base, requestData.Quote)
	user := models.NewUser(*address, []models.CurrencyPair{*pair})
	err := h.subscriptionRepository.Subscribe(*user)
	if err == nil {
		presentors.PresentSubscriptionJSON(c)
	} else {
		if errors.Is(err, custom.ErrAlreadySubscribed) {
			presentors.PresentSubscriptionConflictJSON(c)
		} else {
			presentors.PresentErrorJSON(c)
		}
	}
}
