package logger

import (
	"GenesisTask/pkg/application"
	log_types "GenesisTask/pkg/infrastructure/logger/common"
	"os"

	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQLogger struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQLogger() application.Logger {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connection, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatal(err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	declareQueue(channel, log_types.Debug)
	declareQueue(channel, log_types.Error)
	declareQueue(channel, log_types.Info)

	return &RabbitMQLogger{connection, channel}
}

func (l *RabbitMQLogger) CloseConnectionAndChannel() {
	l.channel.Close()
	l.connection.Close()
}

func declareQueue(channel *amqp.Channel, name string) {
	_, err := channel.QueueDeclare(
		name,  // queue name
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (l *RabbitMQLogger) LogDebug(v ...any) {
	l.publishMessage(log_types.Debug, fmt.Sprint(v))
}

func (l *RabbitMQLogger) LogError(v ...any) {
	l.publishMessage(log_types.Error, fmt.Sprint(v))
}

func (l *RabbitMQLogger) LogInfo(v ...any) {
	l.publishMessage(log_types.Info, fmt.Sprint(v))
}

func (l *RabbitMQLogger) publishMessage(logType string, logMsg string) {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(logMsg),
	}
	if err := l.channel.Publish(
		"",      // exchange
		logType, // routing key
		false,   // mandatory
		false,   // immediate
		message); err != nil {
		log.Fatal(err)
	}
}
