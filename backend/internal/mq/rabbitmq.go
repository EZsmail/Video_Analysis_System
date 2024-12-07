package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	broker *amqp.Connection
}

func ConnectRabbitMQ(mqUrl string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(mqUrl)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn}, nil
}

func (conn *RabbitMQ) Close() {
	conn.broker.Close()
}

func (conn *RabbitMQ) SendTask(queue string, message []byte) error {
	const op = "mq.sendtask"

	ch, err := conn.broker.Channel()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = ch.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        message,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
