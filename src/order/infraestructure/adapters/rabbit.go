package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Execute(idProduct int32, quantity int32, totalPrice float64, status string) {
	conn, err := amqp.Dial("amqp://leo:1234@34.235.202.211:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"rabbit", // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Crear un mapa con los datos
	data := map[string]interface{}{
		"idProduct":  idProduct,
		"quantity":   quantity,
		"totalPrice": totalPrice,
		"status":     status,
	}

	// Convertir los datos a JSON
	body, err := json.Marshal(data)
	failOnError(err, "Failed to marshal JSON")

	err = ch.PublishWithContext(ctx,
		"rabbit", // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
