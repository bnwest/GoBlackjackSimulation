package strategy

import (
	// "fmt"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"
)

// created interface to fix circular import between packages strategy and game
type PlayerHandInterface interface {
	HardCount() int
	SoftCount() int
	NumCards() int
	IsFromSplit() bool
	GetCard(cardIndex int) cards.Card
}

func convertToPlayerDecision(
	decision Decision,
	playerHand PlayerHandInterface,
) PlayerDecision {
	// Decision sometimes return Xy, which translates to do X if allowed else do y.
	// Determine the X or the y here.

	isFirstDecision := playerHand.NumCards() == 2
	isFirstPostSplitDecision := isFirstDecision && playerHand.IsFromSplit()

	var playerDecision PlayerDecision

	hardCount := playerHand.HardCount()
	softCount := playerHand.SoftCount()

	if decision == Decision(S) {
		return PlayerDecision(STAND)

	} else if decision == Decision(H) {
		return PlayerDecision(HIT)

	} else if decision == Decision(Dh) || decision == Decision(Ds) {
		// may be only allow to down on hand totals [9, 10, 11] or some such
		// basic stratgey wants to double down on
		//     hand hard totals [9, 10, 11]
		//     hand soft totals [12, 13,14, 15, 16, 17, 18, 19]
		var nondoubleDownDecision PlayerDecision
		if decision == Decision(Dh) {
			playerDecision = PlayerDecision(HIT)
		} else {
			playerDecision = PlayerDecision(STAND)
		}

		var canDoubleDown bool
		if isFirstDecision {
			if isFirstPostSplitDecision {
				if house_rules.DOUBLE_DOWN_AFTER_SPLIT {
					canDoubleDown = true
				} else {
					canDoubleDown = false
				}
			} else {
				canDoubleDown = true
			}
		} else {
			canDoubleDown = false
		}

		if canDoubleDown {
			var doubleDown bool
			if house_rules.CanDoubleDown(hardCount) {
				doubleDown = true
			} else if house_rules.CanDoubleDown(softCount) {
				doubleDown = true
			} else {
				doubleDown = false
			}
			if doubleDown {
				playerDecision = PlayerDecision(DOUBLE)
			} else {
				playerDecision = nondoubleDownDecision
			}
		} else {
			playerDecision = nondoubleDownDecision
		}

	} else if decision == Decision(SP) {
		playerDecision = PlayerDecision(SPLIT)

	} else if decision == Decision(Uh) || decision == Decision(Us) || decision == Decision(Usp) {
		// surrent decision must be allowed in the House Rules and
		// must be a first decision (before splitting)
		var nonsurrenderDecision PlayerDecision
		switch decision {
		case Decision(Uh):
			nonsurrenderDecision = PlayerDecision(HIT)
		case Decision(Us):
			nonsurrenderDecision = PlayerDecision(STAND)
		case Decision(Usp):
			nonsurrenderDecision = PlayerDecision(SPLIT)
		default:
			// should never get here
			panic("convertToPlayerDecision() ran into a little trouble in town.")
		}

		surrenderCanBePlayed := isFirstDecision && !isFirstPostSplitDecision && house_rules.SURRENDER_ALLOWED
		if surrenderCanBePlayed {
			playerDecision = PlayerDecision(SURRENDER)
		} else {
			playerDecision = nonsurrenderDecision
		}
	}

	return playerDecision
}

func DetermineBasicStrategyPlay(
	dealerTopCard cards.Card,
	playerHand PlayerHandInterface,
	handAllowsMoreSplits bool,
) PlayerDecision {
	// isFirstDecision := playerHand.NumCards() == 2
	// isFirstPostSplitDecision := isFirstDecision && playerHand.FromSplit

	playerCard1 := playerHand.GetCard(0)
	playerCard2 := playerHand.GetCard(1)

	var decision Decision
	var playerDecision PlayerDecision

	var gotPairs bool
	if house_rules.SPLIT_ON_VALUE_MATCH {
		gotPairs = cards.CardRankValue[playerCard1.Rank] == cards.CardRankValue[playerCard2.Rank]
	} else {
		gotPairs = playerCard1.Rank == playerCard2.Rank
	}

	if gotPairs && handAllowsMoreSplits {
		// Determine if the pairs can be split.
		// Note all of the non-split decisions that are ignored below
		// will not contradict the hard/soft total decision.
		var pairRank cards.CardRank
		if cards.CardRankValue[playerCard1.Rank] == 10 {
			pairRank = cards.CardRank(10)
		} else {
			pairRank = playerCard1.Rank
		}

		decision = GetPairSplitDecision(pairRank, dealerTopCard.Rank)
		playerDecision = convertToPlayerDecision(decision, playerHand)
		if playerDecision == PlayerDecision(SPLIT) {
			return PlayerDecision(SPLIT)
		}
	}

	hardCount := playerHand.HardCount()
	softCount := playerHand.SoftCount()
	useSoftTotal := hardCount < softCount && softCount <= 21

	if useSoftTotal {
		decision = GetSoftTotalDecision(softCount, dealerTopCard.Rank)
		playerDecision = convertToPlayerDecision(decision, playerHand)
		return playerDecision

	} else {
		decision = GetHardTotalDecision(hardCount, dealerTopCard.Rank)
		playerDecision = convertToPlayerDecision(decision, playerHand)
		return playerDecision
	}

	// panic("DetermineBasicStrategyPlay() ran into a little trouble in town.")
}
