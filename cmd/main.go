package main

import (
	"os"
	"rabbit/internal/order"
)

var (
	email  = "test@2.com"
	price  = 100
	amount = 3
)

func main() {
	NewOrder := order.NewOrder(email, price, amount)

}
