package handlers

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/presentors"
	"GenesisTask/pkg/domain/models"

	"github.com/gin-gonic/gin"
)

type RateHandler struct {
	rateRepository *application.RateRepository
}

func NewRateHandler(r *application.RateRepository) *RateHandler {
	return &RateHandler{r}
}

type rateRequest struct {
	Base  string `json:"base"  binding:"required"`
	Quote string `json:"quote" binding:"required"`
}

func (h *RateHandler) GetRate(c *gin.Context) {
	var requestData rateRequest
	if err := c.BindJSON(&requestData); err != nil {
		presentors.PresentErrorJSON(c)
	}
	pair := *models.NewCurrencyPair(requestData.Base, requestData.Quote)
	rate, err := h.rateRepository.GetRate(pair)
	presentors.PresentRateJSON(c, rate, err)
}
