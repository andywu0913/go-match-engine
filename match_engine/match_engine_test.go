package matchengine

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	me := NewMatchEngine()

	fmt.Printf("%+v\n", me.PlaceOrder(NewBuyOrder(1, 2, "Andy")))
	fmt.Printf("%+v\n", me.PlaceOrder(NewBuyOrder(3, 5, "Bryan")))
	fmt.Printf("%+v\n", me.PlaceOrder(NewBuyOrder(1, 3, "Davis")))

	fmt.Printf("%+v\n", me.PlaceOrder(NewSellOrder(4, 2, "Casper")))
	fmt.Printf("%+v\n", me.PlaceOrder(NewSellOrder(6, 7, "Elaine")))
	fmt.Printf("%+v\n", me.PlaceOrder(NewSellOrder(5, 9, "Frank")))
	fmt.Printf("%+v\n", me.PlaceOrder(NewSellOrder(6, 1, "Gary")))

	fmt.Printf("%+v\n", me.String())

	fmt.Printf("%+v\n", me.PlaceOrder(NewSellOrder(2, 9, "Andy")))
	fmt.Printf("%+v\n", me.String())

	fmt.Printf("%+v\n", me.PlaceOrder(NewBuyOrder(5, 5, "Andy")))
	fmt.Printf("%+v\n", me.String())
}
