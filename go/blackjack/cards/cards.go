package cards

// import {}

// enums are not first calss citizens in Go.

type CardSuite int

const (
	HEARTS CardSuite = iota
	DIAMONDS
	SPADES
	CLUBS
)

var CardSuiteValue = map[CardSuite]string{
	HEARTS:   "♥️", // aka U+2665 + U+fe0f
	DIAMONDS: "♦️", // aka U+2666 + U+fe0f
	SPADES:   "♠️", // aka U+2660 + U+fe0f
	CLUBS:    "♣️", // aka U+2663 + U+fe0f
}

type CardRank int

const (
	ACE CardRank = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)
