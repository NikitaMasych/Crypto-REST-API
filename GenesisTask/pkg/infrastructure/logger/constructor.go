package logger

import (
	"GenesisTask/pkg/application"
	rabbitmqlogger "GenesisTask/pkg/infrastructure/logger/rabbitmq"
	txtlogger "GenesisTask/pkg/infrastructure/logger/txt"
	"log"
)

const (
	RabbitMQLoggerType = "RabbitMQLoggerType"
	TxtLoggerType      = "TxtLoggerType"
)

type loggerType string

func CreateLogger(class loggerType) application.Logger {
	switch class {
	case RabbitMQLoggerType:
		return rabbitmqlogger.NewRabbitMQLogger()
	case TxtLoggerType:
		return txtlogger.CreateTxtLoggerWithConfigSpecs()
	default:
		log.Fatal("Unknown logger type\n")
		panic("Compiler requirement, not executed")
	}
}
