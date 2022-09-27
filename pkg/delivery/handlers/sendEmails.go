package handlers

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/presentors"

	"github.com/gin-gonic/gin"
)

type SendRateEmailsHandler struct {
	emailSenderRepository application.EmailSenderRepository
}

func NewSendRateEmailsHandler(r application.EmailSenderRepository) *SendRateEmailsHandler {
	return &SendRateEmailsHandler{r}
}

func (h *SendRateEmailsHandler) SendRateEmails(c *gin.Context) {
	err := h.emailSenderRepository.SendRateEmails()
	if err != nil {
		presentors.PresentErrorJSON(c)
	} else {
		presentors.PresentEmailsSentJSON(c)
	}
}
