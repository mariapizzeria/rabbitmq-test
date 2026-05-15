package order

import "time"

type Order struct {
	ID            string    `json:"id"`
	CustomerEmail string    `json:"customer_email"`
	CreatedAt     time.Time `json:"created_at"`
	Amount        int       `json:"amount"`
	TotalPrice    int       `json:"total_price"`
	Status        string    `json:"status"`
}
