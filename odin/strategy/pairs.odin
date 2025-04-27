package strategy

import "../cards"

pairs_decision := [][13]Decision{
	// A    2    3    4    5    6    7    8    9   10    J    Q    K
	// player pair card: Ace  x  dealer top card
	{.SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP},
	// player pair card: 2  x  dealer top card
	{ .H, .SP, .SP, .SP, .SP, .SP, .SP,  .H,  .H,  .H,  .H,  .H,  .H},
	// player pair card: 3  x  dealer top card
	{ .H, .SP, .SP, .SP, .SP, .SP, .SP,  .H,  .H,  .H,  .H,  .H,  .H},
	// player pair card: 4  x  dealer top card
	{ .H,  .H,  .H,  .H, .SP, .SP,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// player pair card: 5  x  dealer top card
	{ .H, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H},
	// player pair card: 6 x  dealer top card
	{ .H, .SP, .SP, .SP, .SP, .SP,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// player pair card: 7  x  dealer top card
	{ .H, .SP, .SP, .SP, .SP, .SP, .SP,  .H,  .H,  .H,  .H,  .H,  .H},
	// player pair card: 8  x  dealer top card
	{.Usp, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP, .SP},
	// player pair card: 9  x  dealer top card
	{ .S, .SP, .SP, .SP, .SP, .SP,  .S, .SP, .SP,  .S,  .S,  .S,  .S},
	// player pair card: 10  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// player pair card: J  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// player pair card: Q  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// player pair card: K  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// A    2    3    4    5    6    7    8    9   10    J    Q    K
}

create_pair_decisions :: proc() -> map[cards.CardRank]map[cards.CardRank]Decision {
    // Convert 2D array into a map of a map of Decisions.
    // decisions = {
    //     cards.ACE: {cards.ACE: S, ...}, ...
    // }
    decisions: map[cards.CardRank]map[cards.CardRank]Decision
    decisions = map[cards.CardRank]map[cards.CardRank]Decision{}

    for player_pair_rank in cards.CardRank {
		// checks to see if key exists 
		// (for demonstrative purposes only, since key will not exist)
		// _, ok := decisions[player_pair_rank]
        ok := player_pair_rank in decisions

        this_pair_rank := map[cards.CardRank]Decision{}

        for dealer_top_card in cards.CardRank {
            // Error: Cannot assign to the value of a map 'decisions[player_pair_rank][dealer_top_card]' 
			// decisions[player_pair_rank][dealer_top_card] = pairs_decision[player_pair_rank][dealer_top_card]
            this_pair_rank[dealer_top_card] = pairs_decision[player_pair_rank][dealer_top_card]
        }

        decisions[player_pair_rank] = this_pair_rank
    }

    return decisions
}

_PAIR_DECISIONS: map[cards.CardRank]map[cards.CardRank]Decision = create_pair_decisions()

get_pairs_split_decision :: proc(
    player_rank_for_pair: cards.CardRank,
    dealer_top_card: cards.CardRank,
) -> Decision {
    decision := _PAIR_DECISIONS[player_rank_for_pair][dealer_top_card]
    return decision
}
