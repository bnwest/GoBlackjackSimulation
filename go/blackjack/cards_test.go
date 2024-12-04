package main

import (
	"testing"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"

	"github.com/stretchr/testify/assert"
)

// For a unit test to be recognized as such, the name 
// of the test function should start with a prefix “Test”
// and should take *testing. T as its only parameter.

func TestCardSuite(t *testing.T) {

	var hearts cards.CardSuite = cards.HEARTS
	var diamonds cards.CardSuite = cards.DIAMONDS
	var spades cards.CardSuite = cards.SPADES
	var clubs cards.CardSuite = cards.CLUBS

	assert.Equal(t, hearts,   cards.CardSuite(0), "card suite enum should be int")
	assert.Equal(t, diamonds, cards.CardSuite(1), "card suite enum should be int")
	assert.Equal(t, spades,   cards.CardSuite(2), "card suite enum should be int")
	assert.Equal(t, clubs,    cards.CardSuite(3), "card suite enum should be int")

	assert.NotEqual(t, cards.CardSuiteValue[hearts],   "hearts",   "card suite string should match")
	assert.NotEqual(t, cards.CardSuiteValue[diamonds], "diamonds", "card suite string should match")
	assert.NotEqual(t, cards.CardSuiteValue[spades],   "spades",   "card suite string should match")
	assert.NotEqual(t, cards.CardSuiteValue[clubs],    "clubs",    "card suite string should match")

	assert.Equal(t, cards.CardSuiteValue[hearts],   "♥️", "card suite string should match")
	assert.Equal(t, cards.CardSuiteValue[diamonds], "♦️", "card suite string should match")
	assert.Equal(t, cards.CardSuiteValue[spades],   "♠️", "card suite string should match")
	assert.Equal(t, cards.CardSuiteValue[clubs],    "♣️",  "card suite string should match")
}

func TestCardRank(t *testing.T) {
	var ace cards.CardRank = cards.ACE
	var two cards.CardRank = cards.TWO
	var three cards.CardRank = cards.THREE
	var four cards.CardRank = cards.FOUR
	var five cards.CardRank = cards.FIVE
	var six cards.CardRank = cards.SIX
	var seven cards.CardRank = cards.SEVEN
	var eight cards.CardRank = cards.EIGHT
	var nine cards.CardRank = cards.NINE
	var ten cards.CardRank = cards.TEN
	var jack cards.CardRank = cards.JACK
	var queen cards.CardRank = cards.QUEEN
	var king cards.CardRank = cards.KING

	assert.Equal(t, ace, cards.CardRank(1), "card rank enum is expected int")
	assert.Equal(t, two, cards.CardRank(2), "card rank enum is expected int")
	assert.Equal(t, three, cards.CardRank(3), "card rank enum is expected int")
	assert.Equal(t, four, cards.CardRank(4), "card rank enum is expected int")
	assert.Equal(t, five, cards.CardRank(5), "card rank enum is expected int")
	assert.Equal(t, six, cards.CardRank(6), "card rank enum is expected int")
	assert.Equal(t, seven, cards.CardRank(7), "card rank enum is expected int")
	assert.Equal(t, eight, cards.CardRank(8), "card rank enum is expected int")
	assert.Equal(t, nine, cards.CardRank(9), "card rank enum is expected int")
	assert.Equal(t, ten, cards.CardRank(10), "card rank enum is expected int")
	assert.Equal(t, jack, cards.CardRank(11), "card rank enum is expected int")
	assert.Equal(t, queen, cards.CardRank(12), "card rank enum is expected int")
	assert.Equal(t, king, cards.CardRank(13), "card rank enum is expected int")

	assert.Equal(t, int(ace), 1, "card rank enum is expected int")
	assert.Equal(t, int(two), 2, "card rank enum is expected int")
	assert.Equal(t, int(three), 3, "card rank enum is expected int")
	assert.Equal(t, int(four), 4, "card rank enum is expected int")
	assert.Equal(t, int(five), 5, "card rank enum is expected int")
	assert.Equal(t, int(six), 6, "card rank enum is expected int")
	assert.Equal(t, int(seven), 7, "card rank enum is expected int")
	assert.Equal(t, int(eight), 8, "card rank enum is expected int")
	assert.Equal(t, int(nine), 9, "card rank enum is expected int")
	assert.Equal(t, int(ten), 10, "card rank enum is expected int")
	assert.Equal(t, int(jack), 11, "card rank enum is expected int")
	assert.Equal(t, int(queen), 12, "card rank enum is expected int")
	assert.Equal(t, int(king), 13, "card rank enum is expected int")

	ranks := []interface{} {
		nil,
		cards.ACE,
		cards.TWO,
		cards.THREE,
		cards.FOUR,
		cards.FIVE,
		cards.SIX,
		cards.SEVEN,
		cards.EIGHT,
		cards.NINE,
		cards.TEN,
		cards.JACK,
		cards.QUEEN,
		cards.KING,
	}

	for i:= 0; i < len(ranks); i++ {
		maybe_rank := ranks[i]
		// To test whether an interface value holds a specific type, 
		// a type assertion can return two values: the underlying value 
		// and a boolean value that reports whether the assertion succeeded.
		rank, ok := maybe_rank.(cards.CardRank)
		if ok {
			assert.Equal(t, int(rank), i, "card rank enum is expected int")
			assert.Equal(t, rank, cards.CardRank(i), "card rank enum is expected int")
		}
	}
}

func TestCardRankValue(t *testing.T) {
	for i := cards.ACE; i <= cards.NINE; i++ {
		assert.Equal(t, int(i), cards.CardRankValue[i], "card rank value be an int 1..10.")
	}
	for i := cards.TEN; i <= cards.KING; i++ {
		assert.Equal(t, 10, cards.CardRankValue[i], "card rank value be an int 1..10.")
	}
}

func TestUnshuffledDeck(t *testing.T) {
	hearts_count := 0
	diamonds_count := 0
	spades_count := 0
	clubs_count := 0

	rank_counts := map[cards.CardRank]int {
		cards.ACE: 0,
		cards.TWO: 0,
		cards.THREE: 0,
		cards.FOUR: 0,
		cards.FIVE: 0,
		cards.SIX: 0,
		cards.SEVEN: 0,
		cards.EIGHT: 0,
		cards.NINE: 0,
		cards.TEN: 0,
		cards.JACK: 0,
		cards.QUEEN: 0,
		cards.KING: 0,
	}

	for i := 0; i < len(cards.UNSHUFFLED_DECK); i++ {
		var card cards.Card = cards.UNSHUFFLED_DECK[i]
		if card.Suite == cards.HEARTS {
			hearts_count++
		} else if card.Suite == cards.DIAMONDS {
			diamonds_count++
		} else if card.Suite == cards.SPADES {
			spades_count++
		} else if card.Suite == cards.CLUBS {
			clubs_count++
		}
		rank_counts[card.Rank]++
	}

	assert.Equal(t, hearts_count,   13, "deck must 13 cards from each suite")
	assert.Equal(t, diamonds_count, 13, "deck must 13 cards from each suite")
	assert.Equal(t, spades_count,   13, "deck must 13 cards from each suite")
	assert.Equal(t, clubs_count,	13, "deck must 13 cards from each suite")

	for rank := cards.ACE; rank <= cards.KING; rank++ {
		assert.Equal(t, rank_counts[rank], 4, "deck must have four of each rank")
	}
}

func TestCreateShoe(t *testing.T) {
	shoe := cards.CreateShoe()
	assert.Equal(
		t, 
		len(shoe), 
		house_rules.DECKS_IN_SHOE * len(cards.UNSHUFFLED_DECK),
		"shoe must have the correctnumber of cards",
	)
}


func TestDisplayShoe(t *testing.T) {
	shoe := cards.CreateShoe()
	assert.Equal(
		t, 
		len(shoe), 
		house_rules.DECKS_IN_SHOE * len(cards.UNSHUFFLED_DECK),
		"shoe must have the correct number of cards",
	)
	cards.ShuffleShoe(shoe)
	assert.Equal(
		t, 
		len(shoe), 
		house_rules.DECKS_IN_SHOE * len(cards.UNSHUFFLED_DECK),
		"shoe must have the correct number of cards",
	)
	// cards.DisplayShoe(shoe)
}
