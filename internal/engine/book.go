package engine

type OrderBook struct {
	Bids map[float64][]Order
	Asks map[float64][]Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids: make(map[float64][]Order),
		Asks: make(map[float64][]Order),
	}
}

func (ob *OrderBook) AddOrder() {

}

func (ob *OrderBook) RemoveOrder() {

}

func (ob *OrderBook) GetBestBid() {

}

func (ob *OrderBook) GetBestAsk() {

}
