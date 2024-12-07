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
	var playerSplitCard cards.CardRank = cards.ACE
	for i := cards.ACE; i <= cards.KING; i++ {
		var dealerTopCard cards.CardRank = cards.CardRank(i)
		var decision strategy.Decision = strategy.GetPairSplitDecision(playerSplitCard, dealerTopCard)
		assert.Equal(t, strategy.SP, decision, "ACEs should always be split")
	}

	for i := cards.ACE; i <= cards.KING; i++ {
		var playerSplitCard cards.CardRank = cards.CardRank(i)
		for j := cards.ACE; j <= cards.KING; j++ {
			var dealerTopCard cards.CardRank = cards.CardRank(j)
			assert.NotPanics(
				t,
				func() {
					var decision strategy.Decision = strategy.Decision(strategy.GetPairSplitDecision(playerSplitCard, dealerTopCard))
					if !strategy.IsValidDecision(decision) {
						panic("GetPairSplitDecision() returned an invalid decision")
					}
				},
				"GetPairSplitDecision(%v, %v) should not panic given valid inputs",
				playerSplitCard,
				dealerTopCard,
			)
		}
	}
}

func TestGetHardTotalDecision(t *testing.T) {
	playerTotal := 21
	for i := cards.ACE; i <= cards.KING; i++ {
		var dealerTopCard cards.CardRank = cards.CardRank(i)
		var decision strategy.Decision = strategy.GetHardTotalDecision(playerTotal, dealerTopCard)
		assert.Equal(t, strategy.S, decision, "Player should always stand on a 21")
	}

	for i := 4; i <= 21; i++ {
		playerTotal := i
		for j := cards.ACE; j <= cards.KING; j++ {
			var dealerTopCard cards.CardRank = cards.CardRank(j)
			assert.NotPanicsf(
				t,
				func() {
					var decision strategy.Decision = strategy.Decision(strategy.GetHardTotalDecision(playerTotal, dealerTopCard))
					if !strategy.IsValidDecision(decision) {
						panic("GetHardTotalDecision() returned an invalid decision")
					}
				},
				"GetHardTotalDecision(%v, %v) should not panic given valid inputs",
				playerTotal,
				dealerTopCard,
			)
		}
	}
}

func TestGetSoftTotalDecision(t *testing.T) {
	for i := 20; i <= 21; i++ {
		playerTotal := i
		for j := cards.ACE; j <= cards.KING; j++ {
			var dealerTopCard cards.CardRank = cards.CardRank(j)
			var decision strategy.Decision = strategy.GetHardTotalDecision(playerTotal, dealerTopCard)
			assert.Equal(t, strategy.S, decision, "Player should always stand on a S20 or S21")
		}
	}

	for i := 12; i <= 21; i++ {
		playerTotal := i
		for j := cards.ACE; j <= cards.KING; j++ {
			var dealerTopCard cards.CardRank = cards.CardRank(j)
			assert.NotPanicsf(
				t,
				func() {
					var decision strategy.Decision = strategy.Decision(strategy.GetSoftTotalDecision(playerTotal, dealerTopCard))
					if !strategy.IsValidDecision(decision) {
						panic("GetSoftTotalDecision() returned an invalid decision")
					}
				},
				"GetSoftTotalDecision(%v, %v) should not panic given valid inputs",
				playerTotal,
				dealerTopCard,
			)
		}
	}
}

func determineOneBasicDecisionPlay(
	dealerTopCard cards.Card,
	playerCard1 cards.Card,
	playerCard2 cards.Card,
	bet int,
	fromSplit bool,
	handAllowMoreSplits bool,
) bool {
	var playerHand *game.PlayerHand = game.CreatePlayerHand(fromSplit, bet)
	playerHand.AddCard(playerCard1)
	playerHand.AddCard(playerCard2)

	var playerDecision strategy.PlayerDecision = strategy.DetermineBasicStrategyPlay(
		dealerTopCard, playerHand, handAllowMoreSplits,
	)

	// end running lazy golang decision to prevent as many newlines as possible
	var decisionOk bool
	decisionOk = strategy.IsValidPlayerDecision(playerDecision)
	return decisionOk
}

func TestDetermineBasicStrategyPlay(t *testing.T) {
	// strategy.DetermineBasicStrategyPlay
	//     dealerTopCard cards.Card,
	//     playerHand game.PlayerHand,
	//     handAllowMoreSplits bool,
	// )
	const bet int = 100
	var dealerTopCard cards.Card
	var playerCard1 cards.Card
	var playerCard2 cards.Card
	for i := cards.ACE; i <= cards.KING; i++ {
		dealerTopCard = cards.Card{Rank: cards.CardRank(i), Suite: cards.DIAMONDS}

		for j := cards.ACE; j <= cards.KING; j++ {
			playerCard1 = cards.Card{Rank: cards.CardRank(i), Suite: cards.DIAMONDS}

			for k := cards.ACE; k <= cards.KING; k++ {
				playerCard2 = cards.Card{Rank: cards.CardRank(i), Suite: cards.DIAMONDS}

				var fromSplit bool
				var handAllowMoreSplits bool
				var decisionOk bool

				fromSplit = false
				handAllowMoreSplits = false
				decisionOk = determineOneBasicDecisionPlay(
					dealerTopCard, playerCard1, playerCard2,
					bet, fromSplit, handAllowMoreSplits,
				)
				assert.Equal(t, true, decisionOk, "Verifying basic strategy API returns valid decision")

				fromSplit = false
				handAllowMoreSplits = true
				decisionOk = determineOneBasicDecisionPlay(
					dealerTopCard, playerCard1, playerCard2,
					bet, fromSplit, handAllowMoreSplits,
				)
				assert.Equal(t, true, decisionOk, "Verifying basic strategy API returns valid decision")

				fromSplit = true
				handAllowMoreSplits = false
				decisionOk = determineOneBasicDecisionPlay(
					dealerTopCard, playerCard1, playerCard2,
					bet, fromSplit, handAllowMoreSplits,
				)
				assert.Equal(t, true, decisionOk, "Verifying basic strategy API returns valid decision")

				fromSplit = true
				handAllowMoreSplits = true
				decisionOk = determineOneBasicDecisionPlay(
					dealerTopCard, playerCard1, playerCard2,
					bet, fromSplit, handAllowMoreSplits,
				)
				assert.Equal(t, true, decisionOk, "Verifying basic strategy API returns valid decision")
			}
		}
	}
}

func TestPlayerHandInterface(t *testing.T) {
	var playerCard1 cards.Card = cards.Card{Rank: cards.ACE, Suite: cards.HEARTS}
	var playerCard2 cards.Card = cards.Card{Rank: cards.ACE, Suite: cards.DIAMONDS}
	fromSplit := false
	bet := 100
	var playerHand *game.PlayerHand = game.CreatePlayerHand(fromSplit, bet)
	playerHand.AddCard(playerCard1)
	playerHand.AddCard(playerCard2)

	// assign (*game.PlayerHand) to (strategy.PlayerHandInterface)
	var playerHandInterface strategy.PlayerHandInterface = playerHand

	assert.Equal(t, false, playerHandInterface.IsFromSplit(), "PlayerHandInterface IsFromSplit() failed")
	assert.Equal(t, 2, playerHandInterface.NumCards(), "PlayerHandInterface NumCards() failed")
	assert.Equal(t, 2, playerHandInterface.HardCount(), "PlayerHandInterface HardCount() failed")
	assert.Equal(t, 12, playerHandInterface.SoftCount(), "PlayerHandInterface SoftCount() failed")
	assert.Equal(t, playerCard1, playerHandInterface.GetCard(0), "PlayerHandInterface GetCard() failed")
}
