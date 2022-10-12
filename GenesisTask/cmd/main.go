package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/infrastructure/logger"
	"GenesisTask/pkg/utils"
)

func main() {
	logger := logger.CreateLogger(logger.RabbitMQLoggerType)

	utils.EnsureFileExists(config.StorageFile)
	logger.LogInfo("Email storage file existance ensured")

	LaunchEngine(logger)
}
