package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/handlers"
	"GenesisTask/pkg/infrastructure/crypto"
	"GenesisTask/pkg/infrastructure/email"
	logger "GenesisTask/pkg/infrastructure/logger/rabbitmq"
	cache "GenesisTask/pkg/infrastructure/storage/cache/redis"
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
	logger := logger.NewRabbitMQLogger() //createTxtLogger()
	providersChain := crypto.NewProvidersChain(logger)
	emailSender := email.NewGomailSender()
	cacheDuration := time.Duration(config.CacheDurationMins) * time.Minute
	cache := cache.NewRedisCache(config.CacheDb, cacheDuration)
	emailAddressesStorage := storage.NewSubscriptionFileRepository()

	r1 := application.NewRateRepository(*providersChain, cache, logger)
	r2 := application.NewSubscriptionRepository(emailAddressesStorage, logger)
	r3 := application.NewEmailSenderRepository(emailAddressesStorage, emailSender, *r1, logger)

	h1 := handlers.NewRateHandler(r1)
	h2 := handlers.NewSubscribeHandler(*r2)
	h3 := handlers.NewSendRateEmailsHandler(*r3)

	return handlers.NewHandlers(h1, h2, h3)
}

/*
func createTxtLogger() application.Logger {
	loggerFiles := logger.NewLoggerFiles(config.DebugLogFile,
		config.ErrorsLogFile, config.InfoLogFile)
	logger.EnsureLogFilesExist(loggerFiles)
	logger := logger.NewTxtLogger(loggerFiles)
	return logger
}
*/
