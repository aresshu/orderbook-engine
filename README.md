# Orderbook Engine

A high performance order matching engine written in Go with zero external dependencies. This project implements the core functionality of a financial exchange's orderbook, matching buy and sell orders based on **price-time** priority.

> **Status:** Work in Progress - This is an educational project demonstrating core orderbook concepts.

## Features

- **Zero Dependencies** - Pure Go implementation
- **Price-Time Priority Matching** - Orders matched at best prices first, FIFO within price levels
- **Full & Partial Fills** - Supports orders that match completely or partially
- **Trade Execution Tracking** - Records all executed trades with timestamps
- **Clean Architecture** - Separation between engine, orderbook, and matching logic

## How It Works

The engine maintains two sides of the market:
- **Bids** (Buy orders) - Sorted by price (highest first)
- **Asks** (Sell orders) - Sorted by price (lowest first)

When a new order arrives:
1. The engine attempts to match it against resting orders on the opposite side
2. Orders are matched if prices cross (bid >= ask)
3. Matches are executed at the resting order's price
4. Fully filled orders are removed from the book
5. Remaining quantity is added to the book as a resting order

## Quick Start

### Prerequisites
- Go 1.25.5 or higher

### Installation

```bash
git clone https://github.com/aresshu/orderbook-engine.git
cd orderbook-engine
```

### Running the Demo

```bash
go run cmd/main.go
```

### Example Output

```
--- Orderbook Engine Demo ---
--- Setting up initial sell orders ---
SELL 100 @ $50.00
Result: 0 trade(s) executed
  No trades executed
SELL 50 @ $50.50
Result: 0 trade(s) executed
  No trades executed

--- Test 1: Buy order that fully matches ---
BUY 75 @ $51.00
Result: 1 trade(s) executed
  Trades executed:
    - 75 shares @ $50.00 (Buy Order: 0, Sell Order: 0) at 3:30PM

--- Test 2: Buy order with price too low (no match) ---
BUY 30 @ $49.00
Result: 0 trade(s) executed
  No trades executed

--- Test 3: Buy order that partially matches ---
BUY 40 @ $50.75
Result: 2 trade(s) executed
  Trades executed:
    - 25 shares @ $50.00 (Buy Order: 0, Sell Order: 0) at 3:30PM
    - 15 shares @ $50.50 (Buy Order: 0, Sell Order: 1) at 3:30PM

--- Final Orderbook State ---
Bids:
  $49.00: 30 shares (1 orders)

Asks:
  $50.50: 35 shares (1 orders)
```

## Project Structure

```
orderbook-engine/
├── cmd/
│   └── main.go              # Demo application
├── internal/
│   └── engine/
│       ├── order.go         # Order types and enums
│       ├── book.go          # Orderbook data structure
│       ├── engine.go        # Main engine interface
│       └── matching.go      # Order matching logic
├── go.mod
└── README.md
```

## Roadmap

Future enhancements planned:

- [ ] **Unit Tests** - Comprehensive test coverage for matching logic
- [ ] **Benchmarks** - Performance testing and optimization
- [ ] **Price-Level Data Structure** - Improve FIFO guarantees at same price
- [ ] **Decimal Prices** - Replace `float64` with fixed-point arithmetic for precision
- [ ] **Market Orders** - Support for market price orders
- [ ] **Order Modification** - Ability to modify existing orders
- [ ] **REST API** - HTTP interface for order submission

## Known Limitations

This is an educational project with several simplifications:

- **Not Thread-Safe** - No concurrent access protection (single-threaded use only)
- **Floating-Point Prices** - Uses `float64` instead of fixed-point decimal (not production-ready for financial applications)
- **Limited Validation** - Limited input validation and error handling
- **No Order Types** - Only limit orders supported (no market, stop, IOC, FOK, etc)

## Potential Improvements

Areas for enhancement:

1. **Thread Safety** - Add mutex locks or use channels for concurrent access (WIP)
2. **Better Data Structures** - Use min/max heaps for price levels
3. **Order Validation** - Validate prices, quantities, and order parameters
4. **Error Handling** - Return errors instead of panicking

## Learning Goals

This project demonstrates:
- Understanding of financial market microstructure
- Data structure selection for performance
- Go language fundamentals (structs, interfaces, maps, slices)
- Clean code organization and separation of concerns
- Building systems with zero dependencies

## License

MIT License - feel free to use this for learning purposes.

## Author

I built this as a learning project to understand how exchange matching engines work