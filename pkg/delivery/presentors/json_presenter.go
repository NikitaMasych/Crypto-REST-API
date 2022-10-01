package presentors

import (
	"GenesisTask/pkg/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PresentRateJSON(c *gin.Context, rate models.CurrencyRate, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid status value"})
	} else {
		c.JSON(http.StatusOK, gin.H{"rate": rate.GetPrice(), "timestamp": rate.GetTimestamp()})
	}
}

func PresentSubscriptionConflictJSON(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{"description": "Already subscribed"})
}

func PresentSubscriptionJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"description": "Successfully subscribed"})
}

func PresentEmailsSentJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"description": "Emails has been sent"})
}

func PresentErrorJSON(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"description": "Error has occured"})
}
