package matchengine

import (
	"fmt"
	"strings"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type OrderBook struct {
	direction Direction
	treeMap   *rbt.Tree
}

func NewOrderBook(direction Direction) *OrderBook {
	return &OrderBook{
		direction: direction,
		treeMap:   rbt.NewWith(NewOrderKeyComparator(direction)),
	}
}

func (ob *OrderBook) GetFirst() *Order {
	node := ob.treeMap.Left()

	if node == nil {
		return nil
	}

	return node.Value.(*Order)
}

func (ob *OrderBook) Add(order *Order) {
	ob.treeMap.Put(order.GetOrderKey(), order)
}

func (ob *OrderBook) Remove(order *Order) {
	ob.treeMap.Remove(order.GetOrderKey())
}

func (ob *OrderBook) String() string {
	s := make([]string, 0, ob.treeMap.Size())

	itr := ob.treeMap.Iterator()

	for itr.Next() {
		order := itr.Value().(*Order)
		s = append(s, fmt.Sprintf("%v\t%v\t%v\t%v", order.price, order.amount, order.trader, order.ticket))
	}

	if ob.direction == Sell {
		ReverseSlice(s)
	}

	return strings.Join(s, "\n")
}
