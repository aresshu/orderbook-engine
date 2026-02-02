package engine

type Engine struct {
	book *OrderBook
}

func NewEngine() *Engine {
	return &Engine{
		book: NewOrderBook(),
	}
}
func (e *Engine) ProcessOrder(order Order) []Trade {
	return e.book.Match(order)
}

func (e *Engine) CancelOrder(orderId int) bool {
	return e.book.RemoveOrder(orderId)
}

func (e *Engine) GetBids() map[float64][]Order {
	return e.book.Bids
}

func (e *Engine) GetAsks() map[float64][]Order {
	return e.book.Asks
}
