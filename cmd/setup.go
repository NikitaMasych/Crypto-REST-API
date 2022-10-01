package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/handlers"
	"GenesisTask/pkg/infrastructure/crypto"
	"GenesisTask/pkg/infrastructure/email"

	//cache "GenesisTask/pkg/infrastructure/storage/cache/redis"
	cache "GenesisTask/pkg/infrastructure/storage/cache/go-cache"
	storage "GenesisTask/pkg/infrastructure/storage/subscription_repository"
	"time"

	"github.com/gin-gonic/gin"
)

func LaunchEngine() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	handlers := createHandlers()
	initRoutes(router, handlers)
	router.Run(config.ServerUrl)
}

func initRoutes(router *gin.Engine, h *handlers.Handlers) {
	router.POST("/api/rate", h.GetRate)
	router.POST("/api/subscribe", h.Subscribe)
	router.POST("/api/sendEmails", h.SendEmailsToUsers)
}

func createHandlers() (h *handlers.Handlers) {
	providersChain := crypto.NewProvidersChain()
	emailSender := email.NewGomailSender()
	cacheDuration := time.Duration(config.CacheDurationMins) * time.Minute
	// cache := cache.NewRedisCache(config.CacheHost, config.CacheDb, cacheDuration)
	// go-cache is useful for local launching on Windows
	cache := cache.NewGoCache(cacheDuration)
	emailAddressesStorage := storage.NewSubscriptionFileRepository()

	r1 := application.NewRateRepository(*providersChain, cache)
	r2 := application.NewSubscriptionRepository(emailAddressesStorage)
	r3 := application.NewEmailSenderRepository(emailAddressesStorage, emailSender, *r1)

	h1 := handlers.NewRateHandler(r1)
	h2 := handlers.NewSubscribeHandler(*r2)
	h3 := handlers.NewSendRateEmailsHandler(*r3)

	return handlers.NewHandlers(h1, h2, h3)
}
