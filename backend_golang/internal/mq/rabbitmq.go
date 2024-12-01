package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ(mqUrl string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(mqUrl)
	if err != nil {
		fmt.Println(mqUrl)
		return nil, err
	}

	return conn, nil
}

func SendTask(conn *amqp.Connection, queue string, message []byte) error {
	const op = "mq.sendtask"

	ch, err := conn.Channel()
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
