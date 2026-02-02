package engine

import "time"

type Trade struct {
	BuyOrderId  int
	SellOrderId int
	Price       float64
	Quantity    int64
	Timestamp   time.Time
}

func (ob *OrderBook) Match(order Order) []Trade {
	var trades []Trade
	remainingQty := order.Quantity

	for remainingQty > 0 {
		var bestPrice float64
		var found bool

		// Get the best price from the opposite side.
		if order.Side == Bid {
			bestPrice, found = ob.GetBestAsk()
		} else {
			bestPrice, found = ob.GetBestBid()
		}

		// If there are no more orders to match against
		if !found {
			break
		}

		// Check if price allows matching
		canMatch := false
		if order.Side == Bid && order.Price >= bestPrice {
			canMatch = true
		} else if order.Side == Ask && order.Price <= bestPrice {
			canMatch = true
		}

		if !canMatch {
			break
		}

		// Get the orders at the best price
		var ordersAtPrice []Order
		if order.Side == Bid {
			ordersAtPrice = ob.Asks[bestPrice]
		} else {
			ordersAtPrice = ob.Bids[bestPrice]
		}

		// Match against orders at this price level (First In First Out)
		for len(ordersAtPrice) > 0 && remainingQty > 0 {
			restingOrder := &ordersAtPrice[0]

			// Determine the fill quantity
			fillQty := remainingQty
			if restingOrder.Quantity < fillQty {
				fillQty = restingOrder.Quantity
			}

			// Create the trade
			trade := Trade{
				Price:     bestPrice,
				Quantity:  fillQty,
				Timestamp: time.Now(),
			}

			// Set the buyer and the seller IDS
			if order.Side == Bid {
				trade.BuyOrderId = order.Id
				trade.SellOrderId = restingOrder.Id
			} else {
				trade.BuyOrderId = restingOrder.Id
				trade.SellOrderId = order.Id
			}

			trades = append(trades, trade)

			// Update the quantities
			remainingQty -= fillQty
			restingOrder.Quantity -= fillQty

			// Remove the fully filled orders
			if restingOrder.Quantity == 0 {
				ob.RemoveOrder(restingOrder.Id)
			}

			// Refresh the orders at this price level after each match
			if order.Side == Bid {
				ordersAtPrice = ob.Asks[bestPrice]
			} else {
				ordersAtPrice = ob.Bids[bestPrice]
			}
		}
	}
	// Add the remaining quantity to the book
	if remainingQty > 0 {
		order.Quantity = remainingQty
		ob.AddOrder(order)
	}

	return trades
}
