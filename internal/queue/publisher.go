package queue

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) Publish(queueName string, data interface{}) error {

	q, err := r.Channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto delete
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.Channel.Publish(
		"",     // default exchange
		q.Name, // queue name
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
