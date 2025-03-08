package core

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetRabbitMQConnection() *amqp.Connection {
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")

	connStr := "amqp://" + rabbitMQUser + ":" + rabbitMQPassword + "@" + rabbitMQHost + ":" + rabbitMQPort + "/"
	conn, err := amqp.Dial(connStr)
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
