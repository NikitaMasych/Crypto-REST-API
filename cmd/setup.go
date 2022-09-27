package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/handlers"
	"GenesisTask/pkg/infrastructure/crypto"
	"GenesisTask/pkg/infrastructure/email"
	"GenesisTask/pkg/infrastructure/storage/cache"
	storage "GenesisTask/pkg/infrastructure/storage/emails_repository"

	"github.com/gin-gonic/gin"
)

func LaunchEngine() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	handlers := createHandlers()
	initRoutes(router, handlers)
	router.Run(config.Get().ServerURL)
}

func initRoutes(router *gin.Engine, h *handlers.Handlers) {
	router.GET("/api/rate", h.GetRate)
	router.POST("/api/subscribe", h.Subscribe)
	router.POST("/api/sendEmails", h.SendRateEmails)
}

func createHandlers() (h *handlers.Handlers) {
	providersChain := crypto.NewProvidersChain()
	emailSender := email.NewGomailSender()
	cache := cache.NewGoCache()
	emailAddressesStorage := storage.NewFileRepository()
	pairSource := config.NewConfigPairSource()

	r1 := application.NewRateRepository(*providersChain, cache)
	r2 := application.NewSubscriptionRepository(emailAddressesStorage)
	r3 := application.NewEmailSenderRepository(emailAddressesStorage, emailSender, *r1, pairSource)

	h1 := handlers.NewRateHandler(r1, &pairSource)
	h2 := handlers.NewSubscribeHandler(*r2)
	h3 := handlers.NewSendRateEmailsHandler(*r3)

	return handlers.NewHandlers(h1, h2, h3)
}
