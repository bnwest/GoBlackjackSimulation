package cards

import (
	"fmt"
	"math/rand"
	"slices"

	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"
)

// enums are not first calss citizens in Go.

type CardSuite int

const (
	HEARTS CardSuite = iota
	DIAMONDS
	SPADES
	CLUBS
)

var CardSuiteValue = map[CardSuite]string {
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

var CardRankValue = map[CardRank]int {
    ACE:    1,
    TWO:    2,
    THREE:  3,
    FOUR:   4,
    FIVE:   5,
    SIX:    6,
    SEVEN:  7,
    EIGHT:  8,
    NINE:   9,
    TEN:   10,
    JACK:  10,
    QUEEN: 10,
    KING:  10,
}

var CardRankString = map[CardRank]string {
    ACE:    "A",
    TWO:    "2",
    THREE:  "3",
    FOUR:   "4",
    FIVE:   "5",
    SIX:    "6",
    SEVEN:  "7",
    EIGHT:  "8",
    NINE:   "9",
    TEN:   "10",
    JACK:  "J",
    QUEEN: "Q",
    KING:  "K",
}

type Card struct {
	Suite CardSuite
	Rank CardRank
}

var UNSHUFFLED_DECK = []Card {
	// HEARTS
	Card{Suite: HEARTS, Rank: ACE},
	Card{Suite: HEARTS, Rank: TWO},
	Card{Suite: HEARTS, Rank: THREE},
	Card{Suite: HEARTS, Rank: FOUR},
	Card{Suite: HEARTS, Rank: FIVE},
	Card{Suite: HEARTS, Rank: SIX},
	Card{Suite: HEARTS, Rank: SEVEN},
	Card{Suite: HEARTS, Rank: EIGHT},
	Card{Suite: HEARTS, Rank: NINE},
	Card{Suite: HEARTS, Rank: TEN},
	Card{Suite: HEARTS, Rank: JACK},
	Card{Suite: HEARTS, Rank: QUEEN},
	Card{Suite: HEARTS, Rank: KING},
	// DIAMONDS
	Card{Suite: DIAMONDS, Rank: ACE},
	Card{Suite: DIAMONDS, Rank: TWO},
	Card{Suite: DIAMONDS, Rank: THREE},
	Card{Suite: DIAMONDS, Rank: FOUR},
	Card{Suite: DIAMONDS, Rank: FIVE},
	Card{Suite: DIAMONDS, Rank: SIX},
	Card{Suite: DIAMONDS, Rank: SEVEN},
	Card{Suite: DIAMONDS, Rank: EIGHT},
	Card{Suite: DIAMONDS, Rank: NINE},
	Card{Suite: DIAMONDS, Rank: TEN},
	Card{Suite: DIAMONDS, Rank: JACK},
	Card{Suite: DIAMONDS, Rank: QUEEN},
	Card{Suite: DIAMONDS, Rank: KING},
	// SPADES
	Card{Suite: SPADES, Rank: ACE},
	Card{Suite: SPADES, Rank: TWO},
	Card{Suite: SPADES, Rank: THREE},
	Card{Suite: SPADES, Rank: FOUR},
	Card{Suite: SPADES, Rank: FIVE},
	Card{Suite: SPADES, Rank: SIX},
	Card{Suite: SPADES, Rank: SEVEN},
	Card{Suite: SPADES, Rank: EIGHT},
	Card{Suite: SPADES, Rank: NINE},
	Card{Suite: SPADES, Rank: TEN},
	Card{Suite: SPADES, Rank: JACK},
	Card{Suite: SPADES, Rank: QUEEN},
	Card{Suite: SPADES, Rank: KING},
	// CLUBS
	Card{Suite: CLUBS, Rank: ACE},
	Card{Suite: CLUBS, Rank: TWO},
	Card{Suite: CLUBS, Rank: THREE},
	Card{Suite: CLUBS, Rank: FOUR},
	Card{Suite: CLUBS, Rank: FIVE},
	Card{Suite: CLUBS, Rank: SIX},
	Card{Suite: CLUBS, Rank: SEVEN},
	Card{Suite: CLUBS, Rank: EIGHT},
	Card{Suite: CLUBS, Rank: NINE},
	Card{Suite: CLUBS, Rank: TEN},
	Card{Suite: CLUBS, Rank: JACK},
	Card{Suite: CLUBS, Rank: QUEEN},
	Card{Suite: CLUBS, Rank: KING},
}

func CreateShoe() []Card {
    var shoe []Card = []Card{}
	switch house_rules.DECKS_IN_SHOE {
	case 1:
		shoe = UNSHUFFLED_DECK
	case 2:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	case 3:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	case 4:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	case 5:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	case 6:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	case 7:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	case 8:
		shoe = slices.Concat(UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK, UNSHUFFLED_DECK)
	default:
		shoe = UNSHUFFLED_DECK
	}
    return shoe
}

func ShuffleShoe(shoe []Card) {
	for i := 0; i < len(shoe); i++ {
		rand.Shuffle(
			len(shoe), 
			func(i, j int) {
				shoe[i], shoe[j] = shoe[j], shoe[i]
			},
		)
	}
}

func DisplayShoe(shoe []Card) {
	for i := 0; i < len(shoe); i++ {
		card := shoe[i]
		fmt.Printf("%v%v\n", CardRankString[card.Rank], CardSuiteValue[card.Suite])
	}
}
