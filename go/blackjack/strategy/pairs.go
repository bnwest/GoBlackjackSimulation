package strategy

import (
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
)

// Use NO since we live in a zero index world.
// Only use the SP decision from this table.  The other decisions mirror exactly
// what the hard/soft total decision tables yield.

var pairsDecisions = [14][14]Decision{
	//0   A   2   3   4   5   6   7   8   9  10   J   Q   K
	{NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO},
	// player pair card: Ace  x  dealer top card
	{NO, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP},
	// player pair card: 2  x  dealer top card
	{NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H},
	// player pair card: 3  x  dealer top card
	{NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H},
	// player pair card: 4  x  dealer top card
	{NO,  H,  H,  H,  H, SP, SP,  H,  H,  H,  H,  H,  H,  H},
	// player pair card: 5  x  dealer top card
	{NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H},
	// player pair card: 6 x  dealer top card
	{NO,  H, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H,  H},
	// player pair card: 7  x  dealer top card
	{NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H},
	// player pair card: 8  x  dealer top card
	{NO, Usp, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP},
	// player pair card: 9  x  dealer top card
	{NO,  S, SP, SP, SP, SP, SP,  S, SP, SP,  S,  S,  S,  S},
	// player pair card: 10  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// player pair card: J  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// player pair card: Q  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	// player pair card: K  x  dealer top card
	{NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S},
	//0   A   2   3   4   5   6   7   8   9  10   J   Q   K
}

func createPairDecisions() map[cards.CardRank]map[cards.CardRank]Decision {
	// Convert 2D array into a map of a map of Decisions.
	// decisions = {
	//     cards.ACE: {cards.ACE: S, ...}, ...
	// }
	decisions := make(map[cards.CardRank]map[cards.CardRank]Decision)

	// initialize the inner map
	for i := cards.ACE; i <= cards.KING; i++ {
		var playerPairRank cards.CardRank =  cards.CardRank(i)

		// checks to see if key exists 
		// (for demonstrative purposes only, since key will not exist)
		_, ok := decisions[playerPairRank]
		if !ok {
			// create the inner map
			decisions[playerPairRank] = map[cards.CardRank]Decision{}
		}
	}

	for i := cards.ACE; i <= cards.KING; i++ {
		var playerPairRank cards.CardRank =  cards.CardRank(i)
		for j := cards.ACE; j <= cards.KING; j++ {
			var dealerTopCard_rank cards.CardRank = cards.CardRank(j)
			decisions[playerPairRank][dealerTopCard_rank] = pairsDecisions[i][j]
		}
	}

	return decisions
}

// not exported
var _PAIR_DECISIONS map[cards.CardRank]map[cards.CardRank]Decision = createPairDecisions()

func GetPairSplitDecision(
	playerPairRank cards.CardRank,
	dealerTopCard cards.CardRank,
) Decision {
	var decision Decision = _PAIR_DECISIONS[playerPairRank][dealerTopCard]
	return decision
}
