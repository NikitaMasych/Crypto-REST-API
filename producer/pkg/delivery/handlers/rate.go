package handlers

import (
	"producer/pkg/application"
	"producer/pkg/delivery/presentors"
	"producer/pkg/domain/models"

	"github.com/gin-gonic/gin"
)

type RateHandler struct {
	rateRepository *application.RateRepository
	logger         application.Logger
}

func NewRateHandler(r *application.RateRepository, l application.Logger) *RateHandler {
	return &RateHandler{r, l}
}

type rateRequest struct {
	Base  string `json:"base"  binding:"required"`
	Quote string `json:"quote" binding:"required"`
}

func (h *RateHandler) GetRate(c *gin.Context) {
	var requestData rateRequest
	if err := c.BindJSON(&requestData); err != nil {
		h.logger.LogError(err)
		presentors.PresentErrorJSON(c)
		h.logger.LogDebug("Presented JSON error")
	}
	pair := *models.NewCurrencyPair(requestData.Base, requestData.Quote)
	rate, err := h.rateRepository.GetRate(pair)
	h.logger.LogInfo()
	if err != nil {
		h.logger.LogError(err)
		presentors.PresentErrorJSON(c)
		h.logger.LogDebug("Presented JSON error")
	} else {
		presentors.PresentRateJSON(c, rate)
		h.logger.LogDebug("Presented JSON rate")
	}
}
