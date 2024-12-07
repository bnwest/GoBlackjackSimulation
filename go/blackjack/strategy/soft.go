package strategy

import (
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
)

var softTotalDecision = [22][14]Decision{
	//0   A   2   3   4   5   6   7   8   9  10   J   Q   K
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 1  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 2  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 3  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 4  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 5  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 6  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 7  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 8  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 9  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 10  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 11  x  dealer top card
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// soft total: 12 (A, A)  x  dealer top card
	{NO,  H,  H,  H,  H,  H, Dh,  H,  H,  H,  H,  H,  H,  H},
	// soft total: 13 (A, 2)  x  dealer top card
	{NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H},
	// soft total: 14 (A, 3)  x  dealer top card
	{NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H},
	// soft total: 15 (A, 4)  x  dealer top card
	{NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H},
	// soft total: 16 (A, 5)  x  dealer top card
	{NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H},
	// soft total: 17 (A, 6)  x  dealer top card
	{NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H},
	// soft total: 18 (A, 7)  x  dealer top card
	{NO,  H, Ds, Ds, Ds, Ds, Ds,  S,  S,  H,  H,  H,  H,  H},
	// soft total: 19 (A, 8)  x  dealer top card
	{NO,  S,  S,  S,  S,  S, Ds,  S,  S,  S,  S,  S,  S,  S},
	// soft total: 20 (A, 9)  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// soft total: 21 (A, 10)  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	//0   A   2   3   4   5   6   7   8   9  10   J   Q   K
}

func createSoftTotalDecisions() [22]map[cards.CardRank]Decision {
	// Convert 2D array of Decisions into a list of maps:
	// decisions = {
	//     0: {cards.ACE: NO, ..,}, ...
	// }
	decisions := [22]map[cards.CardRank]Decision{}

	// initialize iner map
	for i := 0; i < 22; i++ {
		decisions[i] = make(map[cards.CardRank]Decision)
	}

	for i := 0; i < 22; i++ {
		total := i
		for j := cards.ACE; j <= cards.KING; j++ {
			dealerTopCard := cards.CardRank(j)
			decisions[total][dealerTopCard] = softTotalDecision[total][dealerTopCard]
		}
	}

	return decisions
}

// not exported
var _SOFT_TOTAL_DECISIONS [22]map[cards.CardRank]Decision = createSoftTotalDecisions()

func GetSoftTotalDecision(
	playerTotal int, 
	dealerTopCard cards.CardRank,
) Decision {
	var decision Decision = _SOFT_TOTAL_DECISIONS[playerTotal][dealerTopCard]
	return decision
}
