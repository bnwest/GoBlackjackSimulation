package main

import (
	"testing"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/game"

	// "github.com/bnwest/GoBlackjackSimulation/go/blackjack/strategy"

	"github.com/stretchr/testify/assert"
)

//
// PlayerHand
//

func TestCreatePlayerHand(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)
	assert.NotEmpty(t, hand, "CreatePlayerHand() must return a non-nil object")
}

func TestPlayerHandAddCard(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	for i := cards.ACE; i <= cards.KING; i++ {
		card := cards.Card{
			Suite: cards.CardSuite(cards.HEARTS),
			Rank: cards.CardRank(i),
		}
		hand.AddCard(card)
	}

	assert.Equal(t, 13, len(hand.Cards), "Failed to add cards to hand")
	assert.Equal(t, 1, hand.AcesCount(), "Hand should have one and only one Ace")
}

func TestPlayerHandAcesCount(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	assert.Equal(t, 0, hand.AcesCount(), "Empty hand should not have an Ace")

	for i := cards.TWO; i <= cards.KING; i++ {
		card := cards.Card{
			Suite: cards.CardSuite(cards.HEARTS),
			Rank: cards.CardRank(i),
		}
		hand.AddCard(card)
		assert.Equal(t, 0, hand.AcesCount(), "Hand should not have an Ace")
	}

	card := cards.Card{
		Suite: cards.CardSuite(cards.HEARTS),
		Rank: cards.CardRank(cards.CardRank(cards.ACE)),
	}
	hand.AddCard(card)

	assert.Equal(t, 1, hand.AcesCount(), "Hand should have one and only one Ace")

	card = cards.Card{
		Suite: cards.CardSuite(cards.DIAMONDS),
		Rank: cards.CardRank(cards.CardRank(cards.ACE)),
	}
	hand.AddCard(card)

	assert.Equal(t, 2, hand.AcesCount(), "Hand should have two Aces")
}

func TestPlayerHandHardCount(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	for i := cards.ACE; i <= cards.KING; i++ {
		card1 := cards.Card{Rank: cards.CardRank(i), Suite: cards.HEARTS}
		for j := cards.ACE; j <= cards.KING; j++ {
			card2 := cards.Card{Rank: cards.CardRank(j), Suite: cards.HEARTS}
			hand.Cards = []cards.Card{card1, card2}
			hard_count := hand.HardCount()
			assert.Condition(
				t,
				func() (success bool) { return 2 <= hard_count && hard_count <= 20},
				"Hard Count can exceed 20",
			)
		}
	}
}

func TestPlayerHandSoftCount(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	for i := cards.ACE; i <= cards.KING; i++ {
		card1 := cards.Card{Rank: cards.CardRank(i), Suite: cards.HEARTS}
		for j := cards.ACE; j <= cards.KING; j++ {
			card2 := cards.Card{Rank: cards.CardRank(j), Suite: cards.HEARTS}
			hand.Cards = []cards.Card{card1, card2}
			soft_count := hand.SoftCount()
			assert.Condition(
				t,
				func() (success bool) { return 4 <= soft_count && soft_count <= 21},
				"Soft Count can exceed 21",
			)
		}
	}
}

func TestPlayerHandIsNatural(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	for i := cards.ACE; i <= cards.KING; i++ {
		card1 := cards.Card{Rank: cards.CardRank(i), Suite: cards.HEARTS}
		for j := cards.ACE; j <= cards.KING; j++ {
			card2 := cards.Card{Rank: cards.CardRank(j), Suite: cards.HEARTS}
			hand.Cards = []cards.Card{card1, card2}
			is_natural_expected := (
				(card1.Rank == cards.ACE && cards.CardRankValue[card2.Rank] == 10) ||
				(card2.Rank == cards.ACE && cards.CardRankValue[card1.Rank] == 10))
			assert.Equalf(t, 
				is_natural_expected,
				hand.IsNatural(),
				"IsNatural() for cards[%v %v] did not return expect %v",
				card1.Rank,
				card2.Rank,
				is_natural_expected,
			)
		}
	}
}

func TestPlayerHandIsBust(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	card1 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	card2 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	card3 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}

	hand.Cards = []cards.Card{card1, card2}
    assert.Equal(t, false, hand.IsBust(), "Unexpected bust")

	hand.Cards = []cards.Card{card1, card2, card3}
    assert.Equal(t, true, hand.IsBust(), "Expected bust")

	card1 = cards.Card{Rank: cards.CardRank(cards.ACE), Suite: cards.HEARTS}
	card2 = cards.Card{Rank: cards.CardRank(cards.TEN), Suite: cards.HEARTS}

	hand.Cards = []cards.Card{card1, card2}
    assert.Equal(t, false, hand.IsBust(), "Natural can not be a bust")
}

func TestPlayerHandCardCount(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)
	assert.Equal(t, 0, hand.CardCount(), "Empty hand should have zero card count")

	card1 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1}
	assert.Equal(t, 1, hand.CardCount(), "Hand should have a card count of 1")

	card2 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2}
	assert.Equal(t, 2, hand.CardCount(), "Hand should have a card count of 2")

	card3 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2, card3}
	assert.Equal(t, 3, hand.CardCount(), "Hand should have a card count of 3")
}

func TestPlayerHandIsHandOver(t *testing.T) {
	hand := game.CreatePlayerHand(false, 100)

	hand.OutCome = game.IN_PLAY
	assert.Equalf(
		t,
		false,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.IN_PLAY,
	)

	hand.OutCome = game.STAND
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.STAND,
	)

	hand.OutCome = game.BUST
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.BUST,
	)

	hand.OutCome = game.SURRENDER
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.SURRENDER,
	)

	hand.OutCome = game.DEALER_BLACKJACK
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.DEALER_BLACKJACK,
	)
}

//
// DealerHand
//

func TestCreateDealerHand(t *testing.T) {
	hand := game.CreateDealerHand()
	assert.NotEmpty(t, hand, "CreateDeaalerHand() must return a non-nil object")
}

func TestDealerHandAddCard(t *testing.T) {
	hand := game.CreateDealerHand()

	for i := cards.ACE; i <= cards.KING; i++ {
		card := cards.Card{
			Suite: cards.CardSuite(cards.HEARTS),
			Rank: cards.CardRank(i),
		}
		hand.AddCard(card)
	}

	assert.Equal(t, 13, len(hand.Cards), "Failed to add cards to hand")
}

func TestDealerrHandHardCount(t *testing.T) {
	hand := game.CreateDealerHand()

	for i := cards.ACE; i <= cards.KING; i++ {
		card1 := cards.Card{Rank: cards.CardRank(i), Suite: cards.HEARTS}
		for j := cards.ACE; j <= cards.KING; j++ {
			card2 := cards.Card{Rank: cards.CardRank(j), Suite: cards.HEARTS}
			hand.Cards = []cards.Card{card1, card2}
			hard_count := hand.HardCount()
			assert.Condition(
				t,
				func() (success bool) { return 2 <= hard_count && hard_count <= 20},
				"Hard Count can exceed 20",
			)
		}
	}
}

func TestDealerHandSoftCount(t *testing.T) {
	hand := game.CreateDealerHand()

	for i := cards.ACE; i <= cards.KING; i++ {
		card1 := cards.Card{Rank: cards.CardRank(i), Suite: cards.HEARTS}
		for j := cards.ACE; j <= cards.KING; j++ {
			card2 := cards.Card{Rank: cards.CardRank(j), Suite: cards.HEARTS}
			hand.Cards = []cards.Card{card1, card2}
			soft_count := hand.SoftCount()
			assert.Condition(
				t,
				func() (success bool) { return 4 <= soft_count && soft_count <= 21},
				"Soft Count can exceed 21",
			)
		}
	}
}

func TestDealerHandIsNatural(t *testing.T) {
	hand := game.CreateDealerHand()

	for i := cards.ACE; i <= cards.KING; i++ {
		card1 := cards.Card{Rank: cards.CardRank(i), Suite: cards.HEARTS}
		for j := cards.ACE; j <= cards.KING; j++ {
			card2 := cards.Card{Rank: cards.CardRank(j), Suite: cards.HEARTS}
			hand.Cards = []cards.Card{card1, card2}
			is_natural_expected := (
				(card1.Rank == cards.ACE && cards.CardRankValue[card2.Rank] == 10) ||
				(card2.Rank == cards.ACE && cards.CardRankValue[card1.Rank] == 10))
			assert.Equalf(t, 
				is_natural_expected,
				hand.IsNatural(),
				"IsNatural() for cards[%v %v] did not return expect %v",
				card1.Rank,
				card2.Rank,
				is_natural_expected,
			)
		}
	}
}

func TestDealerHandIsBust(t *testing.T) {
	hand := game.CreateDealerHand()

	card1 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	card2 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	card3 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}

	hand.Cards = []cards.Card{card1, card2}
    assert.Equal(t, false, hand.IsBust(), "Unexpected bust")

	hand.Cards = []cards.Card{card1, card2, card3}
    assert.Equal(t, true, hand.IsBust(), "Expected bust")

	card1 = cards.Card{Rank: cards.CardRank(cards.ACE), Suite: cards.HEARTS}
	card2 = cards.Card{Rank: cards.CardRank(cards.TEN), Suite: cards.HEARTS}

	hand.Cards = []cards.Card{card1, card2}
    assert.Equal(t, false, hand.IsBust(), "Natural can not be a bust")
}

func TestDealerHandCardCount(t *testing.T) {
	hand := game.CreateDealerHand()
	assert.Equal(t, 0, hand.CardCount(), "Empty hand should have zero card count")

	card1 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1}
	assert.Equal(t, 1, hand.CardCount(), "Hand should have a card count of 1")

	card2 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2}
	assert.Equal(t, 2, hand.CardCount(), "Hand should have a card count of 2")

	card3 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2, card3}
	assert.Equal(t, 3, hand.CardCount(), "Hand should have a card count of 3")
}

func TestDealerHandIsHandOver(t *testing.T) {
	hand := game.CreateDealerHand()

	hand.OutCome = game.IN_PLAY
	assert.Equalf(
		t,
		false,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.IN_PLAY,
	)

	hand.OutCome = game.STAND
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.STAND,
	)

	hand.OutCome = game.BUST
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.BUST,
	)

	hand.OutCome = game.DEALER_BLACKJACK
	assert.Equalf(
		t,
		true,
		hand.IsHandOver(),
		"Is hand over? expected %v for outcome %v",
		false,
		game.DEALER_BLACKJACK,
	)
}
