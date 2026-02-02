package engine

import "time"

type OrderBook struct {
	Bids        map[float64][]Order
	Asks        map[float64][]Order
	nextOrderID int
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids: make(map[float64][]Order),
		Asks: make(map[float64][]Order),
	}
}

func (ob *OrderBook) AddOrder(order Order) {
	order.Id = ob.nextOrderID
	order.Timestamp = time.Now()

	ob.nextOrderID++

	switch order.Side {
	case Bid:
		ob.Bids[order.Price] = append(ob.Bids[order.Price], order)
	case Ask:
		ob.Asks[order.Price] = append(ob.Asks[order.Price], order)
	default:
		panic("Invalid order side")
	}
}

func (ob *OrderBook) RemoveOrder(orderId int) bool {
	if ob.removeFromSide(orderId, ob.Bids) {
		return true
	}
	return ob.removeFromSide(orderId, ob.Asks)
}

func (ob *OrderBook) removeFromSide(orderId int, side map[float64][]Order) bool {
	for price, orders := range side {
		for i, order := range orders {
			if order.Id == orderId {
				// Remove order from slice
				side[price] = append(orders[:i], orders[i+1:]...)

				// Remove price key from map when no orders remain at that price
				if len(side[price]) == 0 {
					delete(side, price)
				}

				return true
			}
		}
	}
	return false
}

func (ob *OrderBook) GetBestBid() (float64, bool) {
	return ob.findBestPrice(ob.Bids, true)
}

func (ob *OrderBook) GetBestAsk() (float64, bool) {
	return ob.findBestPrice(ob.Asks, false)
}

func (ob *OrderBook) findBestPrice(side map[float64][]Order, findMax bool) (float64, bool) {
	var bestPrice float64
	found := false

	for price := range side {
		if !found {
			bestPrice = price
			found = true
			continue
		}

		if findMax && price > bestPrice {
			bestPrice = price
		} else if !findMax && price < bestPrice {
			bestPrice = price
		}
	}

	return bestPrice, found
}
