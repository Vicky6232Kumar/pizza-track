package queue

import (
	"log"
)

func (r *RabbitMQ) Consume(queueName string, handler func([]byte)) {

	q, err := r.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := r.Channel.Consume(
		q.Name,
		"",
		true, // auto ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		handler(msg.Body)
	}
}
