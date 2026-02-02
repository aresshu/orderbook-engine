package engine

import "time"

type Side string

type OrderStatus string

type Order struct {
	Id        string
	Side      Side
	Price     int64
	Quantity  int64
	Timestamp time.Time
}
