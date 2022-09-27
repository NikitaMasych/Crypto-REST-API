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

func (h *SubscribeHandler) Subscribe(c *gin.Context) {
	email := c.PostForm("email")
	address := models.NewEmailAddress(email)
	user := models.NewUser(*address)
	err := h.subscriptionRepository.Subscribe(*user)
	if err == nil {
		presentors.PresentUserSubscriptionJSON(c)
	} else {
		if errors.Is(err, custom.ErrAlreadySubscribed) {
			presentors.PresentUserConflictJSON(c)
		} else {
			presentors.PresentErrorJSON(c)
		}
	}
}
