package strategy

import (
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
)

// Expect to use the soft total decision table for: (A,A) and (A,2),
// which is the only way to get to hard totals 2 and 3.

var hard_total_decision = [22][14]Decision{
	//0   A   2   3   4   5   6   7   8   9  10   J   Q   K
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// 1  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// 2  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// 3  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// 4  x  dealer top card
	{NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H},
	// 5  x  dealer top card
	{NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H},
	// 6  x  dealer top card
	{NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H},
	// 7  x  dealer top card
	{NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H},
	// 8  x  dealer top card
	{NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H},
	// 9  x  dealer top card
	{NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H},
	// 10  x  dealer top card
	{NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H},
	// 11  x  dealer top card
	{NO, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh},
	// 12  x  dealer top card
	{NO,  H,  H,  H,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H},
	// 13  x  dealer top card
	{NO,  H,  S,  S,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H},
	// 14  x  dealer top card
	{NO,  H,  S,  S,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H},
	// 15  x  dealer top card
	{NO, Uh,  S,  S,  S,  S,  S,  H,  H,  H, Uh, Uh, Uh, Uh},
	// 16  x  dealer top card
	{NO, Uh,  S,  S,  S,  S,  S,  H,  H, Uh, Uh, Uh, Uh, Uh},
	// 17  x  dealer top card
	{NO, Us,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// 18  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// 19  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// 20  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// 21  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	//0   A   2   3   4   5   6   7   8   9  10   J   Q   K
}

func createHardTotalDecisions() [22]map[cards.CardRank]Decision {
	// Convert 2D array of Decisions into a list of maps:
	// decisions = {
	//     0: {cards.ACE: NO, ..,}, ...
	// }
	decisions := [22]map[cards.CardRank]Decision{}

	// initialize inner map
	for i := 0; i < 22; i++ {
		decisions[i] = make(map[cards.CardRank]Decision)
	}

	for i := 0; i < 22; i++ {
		total := i
		for j := cards.ACE; j <= cards.KING; j++ {
			dealerTopCard := cards.CardRank(j)
			decisions[total][dealerTopCard] = hard_total_decision[total][dealerTopCard]
		}
	}

	return decisions
}

// not exported
var _HARD_TOTAL_DECISIONS [22]map[cards.CardRank]Decision = createHardTotalDecisions()

func GetHardTotalDecision(
	playerTotal int, 
	dealerTopCard cards.CardRank,
) Decision {
	var decision Decision = _HARD_TOTAL_DECISIONS[playerTotal][dealerTopCard]
	return decision
}
