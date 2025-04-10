package adapters

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitProducer struct {
	conn *amqp.Connection
}

func NewRabbitProducer(url string) (*RabbitProducer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitProducer{conn: conn}, nil
}

func (p *RabbitProducer) Publish(idProduct int32, quantity int32, totalPrice float64, status string) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"order",  // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := map[string]interface{}{
		"idProduct":  idProduct,
		"quantity":   quantity,
		"totalPrice": totalPrice,
		"status":     status,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		"order", // exchange
		"123",   // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}
