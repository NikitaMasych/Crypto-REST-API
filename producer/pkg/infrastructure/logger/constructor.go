package logger

import (
	"log"
	"producer/pkg/application"
	rabbitmqlogger "producer/pkg/infrastructure/logger/rabbitmq"
	txtlogger "producer/pkg/infrastructure/logger/txt"
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
