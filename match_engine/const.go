package matchengine

type Direction int

const (
	Buy  Direction = -1 // sort buy book from highest to lowest
	Sell Direction = 1  // sort sell book from lowest to highest
)
