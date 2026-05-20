package main

import (
	"rabbit/internal/order"
	"rabbit/queue-config/consumer"
	"rabbit/queue-config/producer"
)

var (
	email     = "test@2.com"
	price     = 100
	amount    = 3
	queueName = "test"
)

func main() {
	NewOrder := order.NewOrder(email, price, amount)
	newProducer, err := producer.NewProducer(queueName)
	if err != nil {
		panic(err)
	}
	defer newProducer.Close()

	newConsumer, err := consumer.NewConsumer()
	if err != nil {
		panic(err)
	}
	defer newConsumer.Close()
	go func() {
		err = newConsumer.ProcessMessage(queueName, NewOrder)
		if err != nil {
			panic(err)
		}
	}()

	err = newProducer.SendMessage(NewOrder)
	if err != nil {
		panic(err)
	}
}
