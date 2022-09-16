package routes

import (
	"GenesisTask/cache"
	"GenesisTask/crypto"
	"GenesisTask/emails"
	"GenesisTask/model"
	"GenesisTask/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRate(c *gin.Context) {
	price, err := cache.GetConfigCurrencyRateFromCache()
	if err != nil {
		price, err = crypto.GetConfigCurrencyRate()
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid status value"})
	} else {
		c.JSON(http.StatusOK, gin.H{"description": price})
	}
}

func PostSubscribe(c *gin.Context) {
	email := c.PostForm("email")
	user := model.NewUser(email)
	userRepo := c.MustGet("userRepo").(repository.UserRepository)

	if userRepo.IsExist(user) {
		c.JSON(http.StatusConflict, gin.H{"description": "User is already subscribed"})
	} else {
		if err := userRepo.Add(user); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"description": "User successfully subscribed"})
	}
}

func PostSendMessage(c *gin.Context) {
	userRepo := c.MustGet("userRepo").(repository.UserRepository)
	users := userRepo.GetUsers()
	emails.SendEmails(users)
	c.JSON(http.StatusOK, gin.H{"description": "Emails has been sent"})
}
