package matchengine

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/emirpasic/gods/utils"
)

var (
	_buySeq  uint64
	_sellSeq uint64
)

type Order struct {
	sequenceID uint64
	ticket     string
	direction  Direction
	price      int64
	amount     int64
	trader     string
	createAt   time.Time
}

func NewBuyOrder(price int64, amount int64, trader string) *Order {
	seq := atomic.AddUint64(&_buySeq, uint64(1))
	createAt := time.Now()
	ticket := generateBuyOrderTicket(seq, createAt)

	return newOrder(seq, ticket, Buy, price, amount, trader, createAt)
}

func NewSellOrder(price int64, amount int64, trader string) *Order {
	seq := atomic.AddUint64(&_sellSeq, uint64(1))
	createAt := time.Now()
	ticket := generateSellOrderTicket(seq, createAt)

	return newOrder(seq, ticket, Sell, price, amount, trader, createAt)
}

func newOrder(sequenceID uint64, ticket string, direction Direction, price int64, amount int64, trader string, createAt time.Time) *Order {
	return &Order{
		sequenceID: sequenceID,
		ticket:     ticket,
		direction:  direction,
		price:      price,
		amount:     amount,
		trader:     trader,
		createAt:   createAt,
	}
}

func (order *Order) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v)",
		order.sequenceID,
		order.direction,
		order.price,
		order.amount,
	)
}

type OrderKey struct {
	price      int64
	sequenceID uint64
}

func (order *Order) GetOrderKey() OrderKey {
	return OrderKey{
		price:      order.price,
		sequenceID: order.sequenceID,
	}
}

func NewOrderKeyComparator(direction Direction) utils.Comparator {
	return func(key1, key2 any) int {
		_key1 := key1.(OrderKey)
		_key2 := key2.(OrderKey)

		if _key1.price == _key2.price {
			return int(_key1.sequenceID - _key2.sequenceID)
		}

		return int(_key1.price-_key2.price) * int(direction)
	}
}

func generateBuyOrderTicket(seq uint64, createAt time.Time) string {
	return fmt.Sprintf("B%v%06d", createAt.UnixNano(), seq)
}

func generateSellOrderTicket(seq uint64, createAt time.Time) string {
	return fmt.Sprintf("S%v%06d", createAt.UnixNano(), seq)
}
