package handlers

import (
	"GenesisTask/pkg/cache"
	"GenesisTask/pkg/crypto"
	"GenesisTask/pkg/emails"
	"GenesisTask/pkg/model"
	"GenesisTask/pkg/presentation"
	"GenesisTask/pkg/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func GetRate(c *gin.Context) {
	rate, err := cache.GetConfigCurrencyRateFromCache()
	if err != nil {
		rate, err = crypto.GetConfigCurrencyRate()
	}

	presentation.PresentRateJSON(c, rate, err)
}

func Subscribe(c *gin.Context) {
	email := c.PostForm("email")
	user := model.NewUser(email)
	userRepo := c.MustGet("userRepo").(repository.UserRepository)

	if userRepo.IsExist(user) {
		presentation.PresentUserConflictJSON(c)
	} else {
		err := userRepo.Add(user)
		if err != nil {
			log.Fatal(err)
		}
		presentation.PresentUserSubscriptionJSON(c)
	}
}

func SendMessage(c *gin.Context) {
	userRepo := c.MustGet("userRepo").(repository.UserRepository)
	users := userRepo.GetUsers()
	emails.SendEmails(users)
	presentation.PresentEmailsSentJSON(c)
}
