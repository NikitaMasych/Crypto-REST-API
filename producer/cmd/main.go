package main

import (
	"producer/config"
	rabbitmqlogger "producer/pkg/infrastructure/logger/rabbitmq"
	"producer/pkg/utils"
)

func main() {
	logger := rabbitmqlogger.NewRabbitMQLogger()
	defer logger.CloseConnectionAndChannel()

	utils.EnsureFileExists(config.StorageFile)
	logger.LogInfo("Email storage file existance ensured")

	LaunchEngine(logger)
}
