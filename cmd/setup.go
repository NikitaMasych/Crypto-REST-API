package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/handlers"
	"GenesisTask/pkg/infrastructure/crypto"
	"GenesisTask/pkg/infrastructure/email"
	"GenesisTask/pkg/infrastructure/storage/cache"
	storage "GenesisTask/pkg/infrastructure/storage/subscription_repository"
	"time"

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
	router.POST("/api/rate", h.GetRate)
	router.POST("/api/subscribe", h.Subscribe)
	router.POST("/api/sendEmails", h.SendEmailsToUsers)
}

func createHandlers() (h *handlers.Handlers) {
	cfg := config.Get()

	providersChain := crypto.NewProvidersChain()
	emailSender := email.NewGomailSender()
	cache := cache.NewRedisCache(cfg.CacheHost, cfg.CacheDb, time.Duration(cfg.CacheDurationMins)*time.Minute)
	emailAddressesStorage := storage.NewSubscriptionFileRepository()

	r1 := application.NewRateRepository(*providersChain, cache)
	r2 := application.NewSubscriptionRepository(emailAddressesStorage)
	r3 := application.NewEmailSenderRepository(emailAddressesStorage, emailSender, *r1)

	h1 := handlers.NewRateHandler(r1)
	h2 := handlers.NewSubscribeHandler(*r2)
	h3 := handlers.NewSendRateEmailsHandler(*r3)

	return handlers.NewHandlers(h1, h2, h3)
}
