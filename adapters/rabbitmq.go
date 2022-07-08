package adapters

import (
	"github.com/rs/zerolog/log"

	"github.com/streadway/amqp"
)

type QueueAdapter interface {
	PublishMessage(message string) error
}

type RabbitMQAdapter struct {
	Queue            string
	ConnectionString string
}

func NewRabbbitMqAdapter(queue, connectionString string) QueueAdapter {
	return &RabbitMQAdapter{
		queue,
		connectionString,
	}
}

func (rq *RabbitMQAdapter) OnError(err error, msg string) {
	if err != nil {
		log.Err(err).Msgf("Error occurred while publishing message on '%s' queue. Error message: %s", rq.Queue, msg)
	}
}

func (rq *RabbitMQAdapter) PublishMessage(message string) error {
	conn, err := amqp.Dial(rq.ConnectionString)
	rq.OnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	rq.OnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rq.Queue, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	rq.OnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	rq.OnError(err, "Failed to publish a message")

	if err != nil {
		return err
	}

	return nil
}
