package producer

import amqp "github.com/rabbitmq/amqp091-go"

func connectToBroker() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@localhost:5672/")
}

func getChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	return conn.Channel()
}

func declareTopicExchange(channel *amqp.Channel, exchangeName string) error {
	return channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
