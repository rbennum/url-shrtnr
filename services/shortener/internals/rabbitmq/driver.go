package message_broker

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rbennum/url-shrtnr/utils"
	"github.com/rs/zerolog/log"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
)

func NewRabbitMQConnection(config *utils.CommonConfig) {
	// declare main connection
	conn, err := amqp.Dial(config.RabbitMQURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect with RabbitMQ")
	}
	// declare new channel
	channel, err = conn.Channel()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open a new channel.")
	}
	// declare new queue
	queue, err = channel.QueueDeclare(
		"url-shrtnr", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open a new queue.")
	}
}

func CloseConnection() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func SendMessage(message []byte, ctx context.Context) {
	log.Debug().Msg(queue.Name)
	err := channel.PublishWithContext(
		ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Error().Err(err).Msg("Unable to send a message.")
	} else {
		log.Info().Msgf("Message sent")
	}
}
