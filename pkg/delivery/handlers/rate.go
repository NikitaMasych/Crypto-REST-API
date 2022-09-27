package handlers

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/presentors"

	"github.com/gin-gonic/gin"
)

type RateHandler struct {
	rateRepository *application.RateRepository
	pairSource     *application.PairSource
}

func NewRateHandler(r *application.RateRepository, s *application.PairSource) *RateHandler {
	return &RateHandler{r, s}
}

func (h *RateHandler) GetRate(c *gin.Context) {
	pair := (*h.pairSource).GetPair()
	rate, err := h.rateRepository.GetRate(pair)
	presentors.PresentRateJSON(c, rate, err)
}
