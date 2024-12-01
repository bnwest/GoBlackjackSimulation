package main

import (
	"testing"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/strategy"

	"github.com/stretchr/testify/assert"
)

func TestDecisions(t *testing.T) {
	assert.NotEmpty(t, strategy.S,   "Decision %s not found", strategy.S)
	assert.NotEmpty(t, strategy.H,   "Decision %s not found", strategy.H)
	assert.NotEmpty(t, strategy.Dh,  "Decision %s not found", strategy.Dh)
	assert.NotEmpty(t, strategy.Ds,  "Decision %s not found", strategy.Ds)
	assert.NotEmpty(t, strategy.SP,  "Decision %s not found", strategy.SP)
	assert.NotEmpty(t, strategy.Uh,  "Decision %s not found", strategy.Uh)
	assert.NotEmpty(t, strategy.Us,  "Decision %s not found", strategy.Us)
	assert.NotEmpty(t, strategy.Usp, "Decision %s not found", strategy.Usp)
	assert.NotEmpty(t, strategy.NO,  "Decision %s not found", strategy.NO)
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
				func(){
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
				func(){
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
				func(){
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
