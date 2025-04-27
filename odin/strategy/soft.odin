package strategy

import "../cards"

soft_total_decision := [][13]Decision{
	// A    2    3    4    5    6    7    8    9   10    J    Q    K
	// soft total: 0  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// soft total: 1  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// soft total: 2  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// soft total: 3  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// soft total: 4  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 5  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 6  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 7  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 8  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 9  x  dealer top card
	{ .H,  .H, .Dh, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 10  x  dealer top card
	{ .H, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H},
	// soft total: 11  x  dealer top card
	{.Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh},
	// soft total: 12 (A, A)  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 13 (A, 2)  x  dealer top card
	{ .H,  .H,  .H,  .H, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 14 (A, 3)  x  dealer top card
	{ .H,  .H,  .H,  .H, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 15 (A, 4)  x  dealer top card
	{ .H,  .H,  .H, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 16 (A, 5)  x  dealer top card
	{ .H,  .H,  .H, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 17 (A, 6)  x  dealer top card
	{ .H,  .H, .Dh, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// soft total: 18 (A, 7)  x  dealer top card
	{ .H, .Ds, .Ds, .Ds, .Ds, .Ds,  .S,  .S,  .H,  .H,  .H,  .H,  .H},
	// soft total: 19 (A, 8)  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S, .Ds,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// soft total: 20 (A, 9)  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// soft total: 21 (A, 10)  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// A    2    3    4    5    6    7    8    9   10    J    Q    K
}

create_soft_total_decisions :: proc() -> [22]map[cards.CardRank]Decision {
	// Convert 2D array of Decisions into a list of maps:
	// decisions = {
	//     0: {cards.ACE: NO, ..,}, ...
	// }
	decisions := [22]map[cards.CardRank]Decision{}


	// initialize inner map
	for i := 0; i < 22; i += 1 {
		decisions[i] = make(map[cards.CardRank]Decision)
	}

	for i := 0; i < 22; i += 1 {
		total := i
		for dealer_top_card in cards.CardRank {
			decisions[total][dealer_top_card] = soft_total_decision[total][dealer_top_card]
		}
    }

    return decisions
}

// not exported
_SOFT_TOTAL_DECISIONS:[22]map[cards.CardRank]Decision = create_soft_total_decisions()

get_soft_total_decision :: proc(
    player_total: uint,
    dealer_top_card: cards.CardRank,
) -> Decision {
    decision := _SOFT_TOTAL_DECISIONS[player_total][dealer_top_card]
    return decision
}
