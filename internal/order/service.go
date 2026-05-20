package order

import (
	"fmt"
	"time"
)

func NewOrder(email string, amount, price int) *Order {
	return &Order{
		ID:            fmt.Sprintf("ID - %d", time.Now().Unix()),
		CustomerEmail: email,
		CreatedAt:     time.Now(),
		Amount:        amount,
		TotalPrice:    amount * price,
		Status:        "pending",
	}
}
