package main

import (
	"testing"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"

	"github.com/stretchr/testify/assert"
)

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

    assert.Equal(t, cards.CardSuiteValue[hearts],   "♥️",   "card suite string should match")
    assert.Equal(t, cards.CardSuiteValue[diamonds], "♦️", "card suite string should match")
    assert.Equal(t, cards.CardSuiteValue[spades],   "♠️",   "card suite string should match")
    assert.Equal(t, cards.CardSuiteValue[clubs],    "♣️",    "card suite string should match")
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
