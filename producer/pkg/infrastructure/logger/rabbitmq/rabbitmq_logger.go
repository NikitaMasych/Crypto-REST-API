package rabbitmqlogger

import (
	"log"
	"os"
	"producer/pkg/infrastructure/logger/logtypes"

	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQLogger struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQLogger() *RabbitMQLogger {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connection, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatal(err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	declareQueue(channel, logtypes.Debug)
	declareQueue(channel, logtypes.Error)
	declareQueue(channel, logtypes.Info)

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
	l.publishMessage(logtypes.Debug, fmt.Sprint(v))
}

func (l *RabbitMQLogger) LogError(v ...any) {
	l.publishMessage(logtypes.Error, fmt.Sprint(v))
}

func (l *RabbitMQLogger) LogInfo(v ...any) {
	l.publishMessage(logtypes.Info, fmt.Sprint(v))
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
