package main

import (
	"testing"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/game"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/strategy"

	"github.com/stretchr/testify/assert"
)

func TestDecisions(t *testing.T) {
	assert.NotEmpty(t, strategy.S,  "Decision %s not found", strategy.S)
	assert.NotEmpty(t, strategy.H,  "Decision %s not found", strategy.H)
	assert.NotEmpty(t, strategy.Dh, "Decision %s not found", strategy.Dh)
	assert.NotEmpty(t, strategy.Ds, "Decision %s not found", strategy.Ds)
	assert.NotEmpty(t, strategy.SP, "Decision %s not found", strategy.SP)
	assert.NotEmpty(t, strategy.Uh, "Decision %s not found", strategy.Uh)
	assert.NotEmpty(t, strategy.Us, "Decision %s not found", strategy.Us)
	assert.NotEmpty(t, strategy.Usp, "Decision %s not found", strategy.Usp)
	assert.NotEmpty(t, strategy.NO, "Decision %s not found", strategy.NO)
}

func TestPlayerDecisions(t *testing.T) {
	assert.NotEmpty(t, strategy.STAND,     "Player Decision %s not found", strategy.STAND)
	assert.NotEmpty(t, strategy.HIT,       "Player Decision %s not found", strategy.HIT)
	assert.NotEmpty(t, strategy.DOUBLE,    "Player Decision %s not found", strategy.DOUBLE)
	assert.NotEmpty(t, strategy.SPLIT,     "Player Decision %s not found", strategy.Ds)
	assert.NotEmpty(t, strategy.SP,        "Player Decision %s not found", strategy.SPLIT)
	assert.NotEmpty(t, strategy.SURRENDER, "Player Decision %s not found", strategy.SURRENDER)
}

func TestGetPairSplitDecision(t *testing.T) {
	player_split_card := cards.ACE
	for i := cards.ACE; i <= cards.KING; i++ {
		dealer_top_card := cards.CardRank(i)
		decision := strategy.GetPairSplitDecision(player_split_card, dealer_top_card)
		assert.Equal(t, strategy.SP, decision, "ACEs should always be split")
	}

	for i := cards.ACE; i <= cards.KING; i++ {
		player_split_card := cards.CardRank(i)
		for j := cards.ACE; j <= cards.KING; j++ {
			dealer_top_card := cards.CardRank(j)
			assert.NotPanics(
				t,
				func() {
					decision := strategy.Decision(strategy.GetPairSplitDecision(player_split_card, dealer_top_card))
					if !strategy.IsValidDecision(decision) {
						panic("GetPairSplitDecision() returned an invalid decision")
					}
				},
				"GetPairSplitDecision(%v, %v) should not panic given valid inputs",
				player_split_card,
				dealer_top_card,
			)
		}
	}
}

func TestGetHardTotalDecision(t *testing.T) {
	player_total := 21
	for i := cards.ACE; i <= cards.KING; i++ {
		dealer_top_card := cards.CardRank(i)
		decision := strategy.GetHardTotalDecision(player_total, dealer_top_card)
		assert.Equal(t, strategy.S, decision, "Player should always stand on a 21")
	}

	for i := 4; i <= 21; i++ {
		player_total := i
		for j := cards.ACE; j <= cards.KING; j++ {
			dealer_top_card := cards.CardRank(j)
			assert.NotPanicsf(
				t,
				func() {
					decision := strategy.Decision(strategy.GetHardTotalDecision(player_total, dealer_top_card))
					if !strategy.IsValidDecision(decision) {
						panic("GetHardTotalDecision() returned an invalid decision")
					}
				},
				"GetHardTotalDecision(%v, %v) should not panic given valid inputs",
				player_total,
				dealer_top_card,
			)
		}
	}
}

func TestGetSoftTotalDecision(t *testing.T) {
	for i := 20; i <= 21; i++ {
		player_total := i
		for j := cards.ACE; j <= cards.KING; j++ {
			dealer_top_card := cards.CardRank(j)
			decision := strategy.GetHardTotalDecision(player_total, dealer_top_card)
			assert.Equal(t, strategy.S, decision, "Player should always stand on a S20 or S21")
		}
	}

	for i := 12; i <= 21; i++ {
		player_total := i
		for j := cards.ACE; j <= cards.KING; j++ {
			dealer_top_card := cards.CardRank(j)
			assert.NotPanicsf(
				t,
				func() {
					decision := strategy.Decision(strategy.GetSoftTotalDecision(player_total, dealer_top_card))
					if !strategy.IsValidDecision(decision) {
						panic("GetSoftTotalDecision() returned an invalid decision")
					}
				},
				"GetSoftTotalDecision(%v, %v) should not panic given valid inputs",
				player_total,
				dealer_top_card,
			)
		}
	}
}

func determine_one_basic_decision_play(
	dealer_top_card cards.Card,
	player_card1 cards.Card,
	player_card2 cards.Card,
	bet int,
	from_split bool,
	hand_allow_more_splits bool,
) bool {
	var player_hand *game.PlayerHand
	player_hand = game.CreatePlayerHand(from_split, bet)
	player_hand.AddCard(player_card1)
	player_hand.AddCard(player_card2)

	var player_decision strategy.PlayerDecision
	player_decision = strategy.DetermineBasicStrategyPlay(
		dealer_top_card, player_hand, hand_allow_more_splits,
	)

	// end running lazy golang decision to prevent as many newlines as possible
	var decision_ok bool = false
	decision_ok = strategy.IsValidPlayerDecision(player_decision)
	return decision_ok
}

func TestDetermineBasicStrategyPlay(t *testing.T) {
	// strategy.DetermineBasicStrategyPlay
	//     dealer_top_card cards.Card,
	//     player_hand game.PlayerHand,
	//     hand_allow_more_splits bool,
	// )
	const bet int = 100
	var dealer_top_card cards.Card
	var player_card1 cards.Card
	var player_card2 cards.Card
	for i := cards.ACE; i <= cards.KING; i++ {
		dealer_top_card = cards.Card{Rank: cards.CardRank(i), Suite: cards.DIAMONDS}

		for j := cards.ACE; j <= cards.KING; j++ {
			player_card1 = cards.Card{Rank: cards.CardRank(i), Suite: cards.DIAMONDS}

			for k := cards.ACE; k <= cards.KING; k++ {
				player_card2 = cards.Card{Rank: cards.CardRank(i), Suite: cards.DIAMONDS}

				var from_split bool
				var hand_allow_more_splits bool
				var decision_ok bool

				from_split = false
				hand_allow_more_splits = false
				decision_ok = determine_one_basic_decision_play(
					dealer_top_card, player_card1, player_card2,
					bet, from_split, hand_allow_more_splits,
				)
				assert.Equal(t, true, decision_ok, "Verifying basic strategy API returns valid decision")

				from_split = false
				hand_allow_more_splits = true
				decision_ok = determine_one_basic_decision_play(
					dealer_top_card, player_card1, player_card2,
					bet, from_split, hand_allow_more_splits,
				)
				assert.Equal(t, true, decision_ok, "Verifying basic strategy API returns valid decision")

				from_split = true
				hand_allow_more_splits = false
				decision_ok = determine_one_basic_decision_play(
					dealer_top_card, player_card1, player_card2,
					bet, from_split, hand_allow_more_splits,
				)
				assert.Equal(t, true, decision_ok, "Verifying basic strategy API returns valid decision")

				from_split = true
				hand_allow_more_splits = true
				decision_ok = determine_one_basic_decision_play(
					dealer_top_card, player_card1, player_card2,
					bet, from_split, hand_allow_more_splits,
				)
				assert.Equal(t, true, decision_ok, "Verifying basic strategy API returns valid decision")
			}
		}
	}
}

func TestPlayerHandInterface(t *testing.T) {
	var player_card1 cards.Card = cards.Card{Rank: cards.ACE, Suite: cards.HEARTS}
	var player_card2 cards.Card = cards.Card{Rank: cards.ACE, Suite: cards.DIAMONDS}
	from_split := false
	bet := 100
	var player_hand *game.PlayerHand = game.CreatePlayerHand(from_split, bet)
	player_hand.AddCard(player_card1)
	player_hand.AddCard(player_card2)

	// assign (*game.PlayerHand) to (strategy.PlayerHandInterface)
	var player_hand_interface strategy.PlayerHandInterface = player_hand

	assert.Equal(t, false, player_hand_interface.IsFromSplit(), "PlayerHandInterface IsFromSplit() failed")
	assert.Equal(t, 2, player_hand_interface.NumCards(), "PlayerHandInterface NumCards() failed")
	assert.Equal(t, 2, player_hand_interface.HardCount(), "PlayerHandInterface HardCount() failed")
	assert.Equal(t, 12, player_hand_interface.SoftCount(), "PlayerHandInterface SoftCount() failed")
	assert.Equal(t, player_card1, player_hand_interface.GetCard(0), "PlayerHandInterface GetCard() failed")
}
