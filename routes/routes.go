package routes

import (
	"GenesisTask/crypto"
	"GenesisTask/emails"
	"GenesisTask/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRate(c *gin.Context) {
	price, err := crypto.GetConfigCurrencyRate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid status value"})
	} else {
		c.JSON(http.StatusOK, gin.H{"description": price})
	}
}

func PostSubscribe(c *gin.Context) {
	email := c.PostForm("email")
	err := emails.AddEmail(email)
	if err == errors.ErrAlreadyExists {
		c.JSON(http.StatusConflict, gin.H{"description": "Email is already subscribed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"description": "Email successfully subscribed"})
	}
}

func PostSendMessage(c *gin.Context) {
	emails.SendEmails()
	c.JSON(http.StatusOK, gin.H{"description": "Emails has been sent"})
}
