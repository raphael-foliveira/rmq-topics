package producer

import (
	"context"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var exchangeName = "topics"

type Publisher struct {
	amqpCh *amqp.Channel
}

func NewProducer() (*Publisher, error) {
	ch, err := startAmqp()
	if err != nil {
		return nil, err
	}
	return &Publisher{ch}, nil
}

func (p *Publisher) Publish(message *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.amqpCh.PublishWithContext(
		ctx,
		exchangeName,
		message.TopicName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message.Content,
		})
}

func getQueueConnection() (*amqp.Connection, error) {
	amqpUrl := os.Getenv("AMQP_URL")
	return amqp.Dial(amqpUrl)
}

func startAmqp() (*amqp.Channel, error) {
	conn, err := getQueueConnection()
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
