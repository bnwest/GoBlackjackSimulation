package main

import (
	"fmt"
	"testing"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/game"

	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"

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
	assert.Equal(t, 0, hand.Num_Cards(), "Empty hand should have zero card count")

	card1 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1}
	assert.Equal(t, 1, hand.Num_Cards(), "Hand should have a card count of 1")

	card2 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2}
	assert.Equal(t, 2, hand.Num_Cards(), "Hand should have a card count of 2")

	card3 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2, card3}
	assert.Equal(t, 3, hand.Num_Cards(), "Hand should have a card count of 3")
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
	assert.Equal(t, 0, hand.Num_Cards(), "Empty hand should have zero card count")

	card1 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1}
	assert.Equal(t, 1, hand.Num_Cards(), "Hand should have a card count of 1")

	card2 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2}
	assert.Equal(t, 2, hand.Num_Cards(), "Hand should have a card count of 2")

	card3 := cards.Card{Rank: cards.CardRank(8), Suite: cards.HEARTS}
	hand.Cards = []cards.Card{card1, card2, card3}
	assert.Equal(t, 3, hand.Num_Cards(), "Hand should have a card count of 3")
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

//
// PlayerMasterHand
//

func TestCreatePlayerMasterHand(t *testing.T) {
	master_hand := game.CreatePlayerMasterHand()
	assert.NotEmpty(t, master_hand, "CreatePlayerMasterHand() must return a non-nil object")
	assert.Equal(t, 0, master_hand.Num_Hands(), "Need to start without a single hand in the master hand")
	assert.Equal(
		t,
		house_rules.SPLITS_PER_HAND+1, 
		master_hand.HANDS_LIMIT, 
		"Master hand can only be split %v times", 
		house_rules.SPLITS_PER_HAND,
	)

	bet := 2
	master_hand.AddStartHand(bet)
	assert.Equal(t, 1, master_hand.Num_Hands(), "AddStartHand() did not add a hand to the master hand")
}

func dump_hands(intro string, expected []cards.Card, actual []cards.Card) {
    fmt.Printf(
		"%v: expected [ %v%v , %v%v ] actual [ %v%v , %v%v ]\n",
		intro,
		expected[0].Rank, cards.CardSuiteValue[expected[0].Suite],
		expected[1].Rank, cards.CardSuiteValue[expected[1].Suite],
		actual[0].Rank,   cards.CardSuiteValue[actual[0].Suite],
		actual[1].Rank,   cards.CardSuiteValue[actual[1].Suite],
	)
}

func TestPlayerMasterHandSplitHand(t *testing.T) {
	var master_hand *game.PlayerMasterHand
	master_hand = game.CreatePlayerMasterHand()

	bet := 2
	master_hand.AddStartHand(bet)

	// add pair to start hand in the master hand
	for i := cards.ACE; i <= cards.KING; i++ {
		// reset master hand back to have just a single hand
		master_hand.Hands = []*game.PlayerHand{master_hand.Hands[0]}

		card1 := cards.Card{Suite: cards.HEARTS, Rank: cards.CardRank(i)}
		card2 := cards.Card{Suite: cards.SPADES, Rank: cards.CardRank(i)}

		// have a card pair to split to the first hand
		master_hand.Hands[0].Cards = []cards.Card{card1, card2}

		//
		// SPLIT #1:
		// hand[0] [A♥️, A♠️]
		// splits into 
		// hand[0] [A♥️, A♦️] and hand[1] [A♠️, A♣️]
		//

		// create two new cards to add second to each split hand
		new_card1 := cards.Card{Suite: cards.DIAMONDS, Rank: cards.CardRank(i)}
		new_card2 := cards.Card{Suite: cards.CLUBS,    Rank: cards.CardRank(i)}
		cards_to_add := [2]cards.Card{new_card1, new_card2}

		hand_index := 0
		assert.Equal(t, true, master_hand.CanSplit(hand_index), "Hand should be split-able")

		new_hand_index := master_hand.SplitHand(hand_index, cards_to_add)
		assert.Equal(t, 2, master_hand.Num_Hands(), "Master Hand should now have 2 hands")
		assert.Equal(t, 1, new_hand_index, "New split hand got added as expected")

		//
		// SPLIT #2:
		// hand[0] [A♥️, A♦️]
		// splits into
		// hand[0] [A♥️, A♣️] and hand[2] [A♦️, A♠️]
		//

		// create two new cards to add second to each split hand
		new_card1 = cards.Card{Suite: cards.CLUBS, Rank: cards.CardRank(i)}
		new_card2 = cards.Card{Suite: cards.SPADES, Rank: cards.CardRank(i)}
		cards_to_add = [2]cards.Card{new_card1, new_card2}

		hand_index = 0
		assert.Equal(t, true, master_hand.CanSplit(hand_index), "Hand should be split-able")

		new_hand_index = master_hand.SplitHand(hand_index, cards_to_add)
		assert.Equal(t, 3, master_hand.Num_Hands(), "Master Hand should now have 2 hands")
		assert.Equal(t, 2, new_hand_index, "New split hand got added as expected")

		//
		// SPLIT #3:
		// hand[1] [A♠️, A♣️]
		// splits into
		// hand[1] [A♠️, A♦️] and hand[3] [A♣️, A♥️]
		//

		// create two new cards to add second to each split hand
		new_card1 = cards.Card{Suite: cards.DIAMONDS, Rank: cards.CardRank(i)}
		new_card2 = cards.Card{Suite: cards.HEARTS, Rank: cards.CardRank(i)}
		cards_to_add = [2]cards.Card{new_card1, new_card2}

		hand_index = 1
		assert.Equal(t, true, master_hand.CanSplit(hand_index), "Hand should be split-able")

		new_hand_index = master_hand.SplitHand(hand_index, cards_to_add)
		assert.Equal(t, 4, master_hand.Num_Hands(), "Master Hand should now have 2 hands")
		assert.Equal(t, 3, new_hand_index, "New split hand got added as expected")

		//
		// A master should only be able to be split 3 times
		// => a single master hand can turn into no more than four hands
		//

		assert.Equal(t, false, master_hand.CanSplit(0), "Hand should not be split-able")
		assert.Equal(t, false, master_hand.CanSplit(1), "Hand should not be split-able")
		assert.Equal(t, false, master_hand.CanSplit(2), "Hand should not be split-able")
		assert.Equal(t, false, master_hand.CanSplit(3), "Hand should not be split-able")

		//
		// Verify that thee four hand in the master hands have the expected cards.
		//

		// master_hand.Hands[0]  [A♥️, A♣️]
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[0].Cards[0].Rank,
		    "Hand[0] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[1].Suite],
		)
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[0].Cards[1].Rank,
		    "Hand[0] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[1].Suite],
		)
		assert.Equal(
			t,
			cards.CardSuite(cards.HEARTS),
			master_hand.Hands[0].Cards[0].Suite,
		    "Hand[0] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[1].Suite],
		)
		assert.Equal(
			t, 
			cards.CardSuite(cards.CLUBS),  
			master_hand.Hands[0].Cards[1].Suite, 
		    "Hand[0] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[0].Cards[1].Suite],
		)

		// master_hand.Hands[1]  [A♠️, A♦️]
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[1].Cards[0].Rank,
		    "Hand[1] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[1].Suite],
		)
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[1].Cards[1].Rank,
		    "Hand[1] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[1].Suite],
		)
		assert.Equal(
			t,
			cards.CardSuite(cards.SPADES),
			master_hand.Hands[1].Cards[0].Suite,
		    "Hand[1] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[1].Suite],
		)
		assert.Equal(
			t, 
			cards.CardSuite(cards.DIAMONDS),  
			master_hand.Hands[1].Cards[1].Suite, 
		    "Hand[1] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[1].Cards[1].Suite],
		)

		// master_hand.Hands[2]  [A♦️, A♠️]
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[2].Cards[0].Rank,
		    "Hand[2] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[1].Suite],
		)
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[2].Cards[1].Rank,
		    "Hand[2] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[1].Suite],
		)
		assert.Equal(
			t,
			cards.CardSuite(cards.DIAMONDS),
			master_hand.Hands[2].Cards[0].Suite,
		    "Hand[2] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[1].Suite],
		)
		assert.Equal(
			t, 
			cards.CardSuite(cards.SPADES),  
			master_hand.Hands[2].Cards[1].Suite, 
		    "Hand[2] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.DIAMONDS],
			cards.CardRank(i), cards.CardSuiteValue[cards.SPADES],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[2].Cards[1].Suite],
		)

		// master_hand.Hands[3]  [A♣️, A♥️]
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[3].Cards[0].Rank,
		    "Hand[3] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[1].Suite],
		)
		assert.Equalf(
			t,
			cards.CardRank(i),
			master_hand.Hands[3].Cards[1].Rank,
		    "Hand[3] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[1].Suite],
		)
		assert.Equal(
			t,
			cards.CardSuite(cards.CLUBS),
			master_hand.Hands[3].Cards[0].Suite,
		    "Hand[3] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[1].Suite],
		)
		assert.Equal(
			t, 
			cards.CardSuite(cards.HEARTS),  
			master_hand.Hands[3].Cards[1].Suite, 
		    "Hand[3] expected [ %v%v , %v%v ] got [ %v%v , %v%v ]",
			cards.CardRank(i), cards.CardSuiteValue[cards.CLUBS],
			cards.CardRank(i), cards.CardSuiteValue[cards.HEARTS],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[0].Suite],
			cards.CardRank(i), cards.CardSuiteValue[master_hand.Hands[3].Cards[1].Suite],
		)
	}
}

func TestCreatePlayer(t *testing.T) {
	dealer := game.CreateDealer()
	player1 := game.CreatePlayer("John")
	player2 := game.CreatePlayer("Jane")

	player1.SetGameBets([]int{2})
	player2.SetGameBets([]int{2, 2})

	assert.Equal(t, 1, player1.Num_Master_Hands(), "player 1 should have 1 master hand")
	assert.Equal(t, 2, player2.Num_Master_Hands(), "player 2 should have 2 master hand")

	// deal opening hand

	card := cards.Card{Suite: cards.SPADES, Rank: cards.ACE}

	for i := 0; i < 2; i++ {
		for j := 0; j < player1.Num_Master_Hands(); j++ {
			player1.PlayerMasterHands[j].Hands[0].AddCard(card)
		}
		for j := 0; j < player2.Num_Master_Hands(); j++ {
			player2.PlayerMasterHands[j].Hands[0].AddCard(card)
		}
		dealer.DealerHand.AddCard(card)
	}

	assert.Equal(t, 2, player1.PlayerMasterHands[0].Hands[0].Num_Cards(), "Player 1 should have 2 cards in hand")
	assert.Equal(t, 2, player2.PlayerMasterHands[0].Hands[0].Num_Cards(), "Player 2.1 should have 2 cards in hand")
	assert.Equal(t, 2, player2.PlayerMasterHands[1].Hands[0].Num_Cards(), "Player 2.2 should have 2 cards in hand")

	assert.Equal(t, 2, dealer.DealerHand.Num_Cards(), "Dealer should have 2 cards in hand")
}

func TestBlackJackGameStart(t *testing.T) {
	blackjack := game.CreateBlackJack()

	var player1 *game.Player = game.CreatePlayer("John")
	var player2 *game.Player = game.CreatePlayer("Jane")

	blackjack.SetPlayersForGame([]*game.Player{player1, player2})

	player1.SetGameBets([]int{2})
	player2.SetGameBets([]int{2, 2})

	assert.Equal(t, 1, player1.Num_Master_Hands(), "player 1 should have 1 master hand")
	assert.Equal(t, 2, player2.Num_Master_Hands(), "player 2 should have 2 master hand")

	assert.Equal(t, 1, blackjack.Players[0].Num_Master_Hands(), "player 1 should have 1 master hand")
	assert.Equal(t, 2, blackjack.Players[1].Num_Master_Hands(), "player 2 should have 2 master hand")
}

func TestBlackJackPlayGame(t *testing.T) {
	blackjack := game.CreateBlackJack()
	blackjack.PlayGame()
}
