package matchengine

import "time"

type MatchRecord struct {
	price      int64
	amount     int64
	takerOrder *Order
	makerOrder *Order
	createAt   time.Time
}

type MatchResult struct {
	takerOrder   *Order
	matchRecords []MatchRecord
}

func NewMatchResult(takerOrder *Order) *MatchResult {
	return &MatchResult{
		takerOrder: takerOrder,
	}
}

func (mr *MatchResult) Add(price int64, matchedAmount int64, takerOrder *Order, makerOrder *Order) {
	mr.matchRecords = append(mr.matchRecords, MatchRecord{
		price:      price,
		amount:     matchedAmount,
		takerOrder: takerOrder,
		makerOrder: makerOrder,
		createAt:   time.Now(),
	})
}
