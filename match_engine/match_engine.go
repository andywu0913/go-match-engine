package matchengine

import (
	"fmt"
	"strings"
	"sync"
)

type MatchEngine struct {
	buyBook     *OrderBook
	sellBook    *OrderBook
	latestPrice int64
	mu          sync.RWMutex
}

func NewMatchEngine() *MatchEngine {
	return &MatchEngine{
		buyBook:  NewOrderBook(Buy),
		sellBook: NewOrderBook(Sell),
	}
}

func (me *MatchEngine) PlaceOrder(order *Order) *MatchResult {
	switch order.direction {
	case Buy:
		// taker: buy, maker: sell
		return me.doPlaceOrder(order, me.buyBook, me.sellBook)
	case Sell:
		// taker: sell, maker: buy
		return me.doPlaceOrder(order, me.sellBook, me.buyBook)
	}
	return nil
}

func (me *MatchEngine) doPlaceOrder(takerOrder *Order, takerBook *OrderBook, makerBook *OrderBook) *MatchResult {
	matchResult := NewMatchResult(takerOrder)

	me.mu.Lock()
	defer me.mu.Unlock()

	for {
		makerOrder := makerBook.GetFirst()

		if makerOrder == nil {
			break
		}

		// if price lower than bid price then it's not possible to close the order
		if takerOrder.direction == Buy && takerOrder.price < makerOrder.price {
			break
		}
		// if price higher than ask price then it's not possible to close the order
		if takerOrder.direction == Sell && takerOrder.price > makerOrder.price {
			break
		}

		// close the order with the price from maker
		me.latestPrice = makerOrder.price

		// the intersection amount from taker and maker
		matchedAmount := Min(takerOrder.amount, makerOrder.amount)

		// write to match record
		matchResult.Add(makerOrder.price, matchedAmount, takerOrder, makerOrder)

		// update unclosed order
		takerOrder.amount = takerOrder.amount - matchedAmount
		makerOrder.amount = makerOrder.amount - matchedAmount

		// remove order from maker book if all the order amount has been closed
		if makerOrder.amount == 0 {
			makerBook.Remove(makerOrder)
		}

		// stop matching if all the order amount from taker has been closed
		if takerOrder.amount == 0 {
			break
		}
	}

	// add to taker book if there are remaining order amount
	if takerOrder.amount > 0 {
		takerBook.Add(takerOrder)
	}

	return matchResult
}

func (me *MatchEngine) CancelOrder(order *Order) {
	me.mu.Lock()
	defer me.mu.Unlock()

	switch order.direction {
	case Buy:
		me.buyBook.Remove(order)
	case Sell:
		me.sellBook.Remove(order)
	}
}

func (me *MatchEngine) String() string {
	var sb strings.Builder

	me.mu.RLock()
	defer me.mu.RUnlock()

	sb.WriteString("Price===Quantity=\n")
	sb.WriteString(me.sellBook.String())
	sb.WriteString("\n-----------------\n")
	sb.WriteString(fmt.Sprintf("%+v", me.latestPrice))
	sb.WriteString("\n-----------------\n")
	sb.WriteString(me.buyBook.String())
	sb.WriteString("\n=================\n")

	return sb.String()
}
