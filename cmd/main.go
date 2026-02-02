package main

import (
	"fmt"
	"time"

	"github.com/aresshu/orderbook-engine/internal/engine"
)

func main() {
	fmt.Println("--- Orderbook Engine Demo ---")

	e := engine.NewEngine()

	// Setup initial sell orders
	fmt.Println("--- Setting up initial sell orders ---")
	processAndPrint(e, engine.Order{Side: engine.Ask, Price: 50.00, Quantity: 100})
	processAndPrint(e, engine.Order{Side: engine.Ask, Price: 50.50, Quantity: 50})

	// Test Case 1: Full match
	fmt.Println("\n--- Test 1: Buy order that fully matches ---")
	processAndPrint(e, engine.Order{Side: engine.Bid, Price: 51.00, Quantity: 75})

	// Test Case 2: No match
	fmt.Println("\n--- Test 2: Buy order with price too low (no match) ---")
	processAndPrint(e, engine.Order{Side: engine.Bid, Price: 49.00, Quantity: 30})

	// Test Case 3: Partial match
	fmt.Println("\n--- Test 3: Buy order that partially matches ---")
	processAndPrint(e, engine.Order{Side: engine.Bid, Price: 50.75, Quantity: 40})

	// Final state
	fmt.Println("\n--- Final Orderbook State ---")
	printOrderbook(e)
}

func processAndPrint(e *engine.Engine, order engine.Order) {
	side := "BUY"
	if order.Side == engine.Ask {
		side = "SELL"
	}
	fmt.Printf("%s %d @ $%.2f\n", side, order.Quantity, order.Price)

	trades := e.ProcessOrder(order)
	fmt.Printf("Result: %d trade(s) executed\n", len(trades))
	printTrades(trades)
}

func printTrades(trades []engine.Trade) {
	if len(trades) == 0 {
		fmt.Println("  No trades executed")
		return
	}

	fmt.Println("  Trades executed:")
	for _, trade := range trades {
		fmt.Printf("    - %d shares @ $%.2f (Buy Order: %d, Sell Order: %d) at %s\n",
			trade.Quantity,
			trade.Price,
			trade.BuyOrderId,
			trade.SellOrderId,
			trade.Timestamp.Format(time.Kitchen))
	}
}

func printOrderbook(e *engine.Engine) {
	fmt.Println("Bids:")
	bids := e.GetBids()
	if len(bids) == 0 {
		fmt.Println("  None")
	} else {
		for price, orders := range bids {
			totalQty := int64(0)
			for _, order := range orders {
				totalQty += order.Quantity
			}
			fmt.Printf("  $%.2f: %d shares (%d orders)\n", price, totalQty, len(orders))
		}
	}

	fmt.Println("\nAsks:")
	asks := e.GetAsks()
	if len(asks) == 0 {
		fmt.Println("  None")
	} else {
		for price, orders := range asks {
			totalQty := int64(0)
			for _, order := range orders {
				totalQty += order.Quantity
			}
			fmt.Printf("  $%.2f: %d shares (%d orders)\n", price, totalQty, len(orders))
		}
	}
}
