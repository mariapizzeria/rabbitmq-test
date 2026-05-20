package producer

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewProducer(QueueName string) (*Producer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(QueueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &Producer{
		conn:    conn,
		channel: ch,
		queue:   q}, nil
}

func (p *Producer) SendMessage(T any) error {
	body, err := json.Marshal(T)
	if err != nil {
		return err
	}
	err = p.channel.Publish("", p.queue.Name, false, false, amqp.Publishing{
		Body:        body,
		ContentType: "application/json",
	})
	if err != nil {
		return err
	}
	log.Printf("Sent Message: %s", string(body))
	return nil
}

func (p *Producer) Close() error {
	return p.channel.Close()
}
