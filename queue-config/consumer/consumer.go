package consumer

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewConsumer() (*Consumer, error) {
	consumer, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := consumer.Channel()
	if err != nil {
		return nil, err
	}
	return &Consumer{conn: consumer, channel: ch}, nil
}

func (c *Consumer) ProcessMessage(queueName string, T any) error {
	q, err := c.channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}
	msgs, err := c.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	log.Printf("Consumer got %d messages", len(msgs))
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			err = json.Unmarshal(d.Body, T)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("Consumer got message: %s", d.Body)
			processAny(T)
		}
	}()
	<-forever
	return nil
}

func processAny(T any) {
	log.Printf("Processing email: %d", T) // simulate
	time.Sleep(10 * time.Second)
}

func (c *Consumer) Close() {
	c.conn.Close()
	c.channel.Close()
}
