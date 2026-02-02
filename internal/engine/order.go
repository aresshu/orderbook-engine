package engine

import "time"

type Side int

const (
	Bid Side = iota
	Ask
)

type Order struct {
	Id        int
	Side      Side
	Price     float64
	Quantity  int64
	Timestamp time.Time
}
