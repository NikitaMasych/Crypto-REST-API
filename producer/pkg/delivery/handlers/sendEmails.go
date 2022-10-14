package handlers

import (
	"producer/pkg/application"
	"producer/pkg/delivery/presentors"

	"github.com/gin-gonic/gin"
)

type SendRateEmailsHandler struct {
	emailSenderRepository application.EmailSenderRepository
	logger                application.Logger
}

func NewSendRateEmailsHandler(r application.EmailSenderRepository, l application.Logger) *SendRateEmailsHandler {
	return &SendRateEmailsHandler{r, l}
}

func (h *SendRateEmailsHandler) SendEmailsToUsers(c *gin.Context) {
	err := h.emailSenderRepository.SendEmailsToUsers()
	if err != nil {
		presentors.PresentErrorJSON(c)
		h.logger.LogDebug("Presented JSON error")
	} else {
		presentors.PresentEmailsSentJSON(c)
		h.logger.LogDebug("Presented \"emails sent\" JSON")
	}
}
