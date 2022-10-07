package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connection, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	const errorLogType = "ERROR"
	messages, err := channel.Consume(
		errorLogType, // queue name
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)
	go func() {
		for message := range messages {
			log.Printf("[error]: %s\n", message.Body)
		}
	}()

	<-forever
}
