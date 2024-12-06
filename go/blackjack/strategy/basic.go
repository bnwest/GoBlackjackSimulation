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
	player_hand PlayerHandInterface,
) PlayerDecision {
	// Decision sometimes return Xy, which translates to do X if allowed else do y.
	// Determine the X or the y here.

	is_first_decision := player_hand.NumCards() == 2
	is_first_postsplit_decision := is_first_decision && player_hand.IsFromSplit()

	var player_decision PlayerDecision

	hard_count := player_hand.HardCount()
	soft_count := player_hand.SoftCount()

	if decision == Decision(S) {
		return PlayerDecision(STAND)

	} else if decision == Decision(H) {
		return PlayerDecision(HIT)

	} else if decision == Decision(Dh) || decision == Decision(Ds) {
		// may be only allow to down on hand totals [9, 10, 11] or some such
		// basic stratgey wants to double down on
		//     hand hard totals [9, 10, 11]
		//     hand soft totals [12, 13,14, 15, 16, 17, 18, 19]
		var nondouble_down_decision PlayerDecision
		if decision == Decision(Dh) {
			player_decision = PlayerDecision(HIT)
		} else {
			player_decision = PlayerDecision(STAND)
		}

		var can_double_down bool
		if is_first_decision {
			if is_first_postsplit_decision {
				if house_rules.DOUBLE_DOWN_AFTER_SPLIT {
					can_double_down = true
				} else {
					can_double_down = false
				}
			} else {
				can_double_down = true
			}
		} else {
			can_double_down = false
		}

		if can_double_down {
			var double_down bool
			if house_rules.CanDoubleDown(hard_count) {
				double_down = true
			} else if house_rules.CanDoubleDown(soft_count) {
				double_down = true
			} else {
				double_down = false
			}
			if double_down {
				player_decision = PlayerDecision(DOUBLE)
			} else {
				player_decision = nondouble_down_decision
			}
		} else {
			player_decision = nondouble_down_decision
		}

	} else if decision == Decision(SP) {
		player_decision = PlayerDecision(SPLIT)

	} else if decision == Decision(Uh) || decision == Decision(Us) || decision == Decision(Usp) {
		// surrent decision must be allowed in the House Rules and
		// must be a first decision (before splitting)
		var nonsurrender_decision PlayerDecision
		switch decision {
		case Decision(Uh):
			nonsurrender_decision = PlayerDecision(HIT)
		case Decision(Us):
			nonsurrender_decision = PlayerDecision(STAND)
		case Decision(Usp):
			nonsurrender_decision = PlayerDecision(SPLIT)
		default:
			// should never get here
			panic("convertToPlayerDecision() ran into a little trouble in town.")
		}

		surrender_can_be_played := is_first_decision && !is_first_postsplit_decision && house_rules.SURRENDER_ALLOWED
		if surrender_can_be_played {
			player_decision = PlayerDecision(SURRENDER)
		} else {
			player_decision = nonsurrender_decision
		}
	}

	return player_decision
}

func DetermineBasicStrategyPlay(
	dealer_top_card cards.Card,
	player_hand PlayerHandInterface,
	hand_allows_more_splits bool,
) PlayerDecision {
	// is_first_decision := player_hand.NumCards() == 2
	// is_first_postsplit_decision := is_first_decision && player_hand.FromSplit

	player_card1 := player_hand.GetCard(0)
	player_card2 := player_hand.GetCard(1)

	var decision Decision
	var player_decision PlayerDecision

	var got_pairs bool
	if house_rules.SPLIT_ON_VALUE_MATCH {
		got_pairs = cards.CardRankValue[player_card1.Rank] == cards.CardRankValue[player_card2.Rank]
	} else {
		got_pairs = player_card1.Rank == player_card2.Rank
	}

	if got_pairs && hand_allows_more_splits {
		// Determine if the pairs can be split.
		// Note all of the non-split decisions that are ignored below
		// will not contradict the hard/soft total decision.
		var pair_rank cards.CardRank
		if cards.CardRankValue[player_card1.Rank] == 10 {
			pair_rank = cards.CardRank(10)
		} else {
			pair_rank = player_card1.Rank
		}

		decision = GetPairSplitDecision(pair_rank, dealer_top_card.Rank)
		player_decision = convertToPlayerDecision(decision, player_hand)
		if player_decision == PlayerDecision(SPLIT) {
			return PlayerDecision(SPLIT)
		}
	}

	hard_count := player_hand.HardCount()
	soft_count := player_hand.SoftCount()
	use_soft_total := hard_count < soft_count && soft_count <= 21

	if use_soft_total {
		decision = GetSoftTotalDecision(soft_count, dealer_top_card.Rank)
		player_decision = convertToPlayerDecision(decision, player_hand)
		return player_decision

	} else {
		decision = GetHardTotalDecision(hard_count, dealer_top_card.Rank)
		player_decision = convertToPlayerDecision(decision, player_hand)
		return player_decision
	}

	// panic("DetermineBasicStrategyPlay() ran into a little trouble in town.")
}
