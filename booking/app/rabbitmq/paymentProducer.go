package main

import (
	"log"

	"github.com/coroo/go-starter/app/entity"
	"github.com/streadway/amqp"
)

type PaymentProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewPaymentProducer(amqpURL, queueName string) (*PaymentProducer, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &PaymentProducer{
		connection: conn,
		channel:    ch,
		queue:      q,
	}, nil
}

func (p *PaymentProducer) PublishPayment(payment string) error {
	err := p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payment),
		},
	)
	if err != nil {
		return err
	}
	log.Printf("Payment message sent: %s", payment)
	return nil
}

func (p *PaymentProducer) Close() {
	p.channel.Close()
	p.connection.Close()
}

func publishPayment(ticket entity.Ticket) {
	producer, err := NewPaymentProducer("amqp://guest:guest@localhost:5672/", "payments")
	if err != nil {
		log.Fatalf("Failed to create payment producer: %v", err)
	}
	defer producer.Close()

	err = producer.PublishPayment("Payment details here")
	if err != nil {
		log.Fatalf("Failed to publish payment: %v", err)
	}
}
