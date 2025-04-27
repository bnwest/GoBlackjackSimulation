package strategy

import "../cards"

/*
Odin wants
    Error: Enumerated array literals must only have 'field = value' elements, 
    bare elements are not allowed 

foo := [cards.CardRank]Decision{
    .ACE=.NO, 
    .TWO=.NO, 
    .THREE=.NO, 
    .FOUR=.NO, 
    .FIVE=.NO, 
    .SIX=.NO, 
    .SEVEN=.NO, 
    .EIGHT=.NO, 
    .NINE=.NO, 
    .TEN=.NO, 
    .JACK=.NO, 
    .QUEEN=.NO, 
    .KING=.NO
}

foo2 := [13]Decision{
    .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO
}
*/

hard_total_decision := [][13]Decision{
	// A    2    3    4    5    6    7    8    9   10    J    Q    K
	// 0  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// 1  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// 2  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// 3  x  dealer top card
	{.NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO, .NO},
	// 4  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 5  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 6  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 7  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 8  x  dealer top card
	{ .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 9  x  dealer top card
	{ .H,  .H, .Dh, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 10  x  dealer top card
	{ .H, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh,  .H,  .H,  .H,  .H},
	// 11  x  dealer top card
	{.Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh, .Dh},
	// 12  x  dealer top card
	{ .H,  .H,  .H,  .S,  .S,  .S,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 13  x  dealer top card
	{ .H,  .S,  .S,  .S,  .S,  .S,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 14  x  dealer top card
	{ .H,  .S,  .S,  .S,  .S,  .S,  .H,  .H,  .H,  .H,  .H,  .H,  .H},
	// 15  x  dealer top card
	{.Uh,  .S,  .S,  .S,  .S,  .S,  .H,  .H,  .H, .Uh, .Uh, .Uh, .Uh},
	// 16  x  dealer top card
	{.Uh,  .S,  .S,  .S,  .S,  .S,  .H,  .H, .Uh, .Uh, .Uh, .Uh, .Uh},
	// 17  x  dealer top card
	{.Us,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// 18  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// 19  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// 20  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// 21  x  dealer top card
	{ .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S,  .S},
	// A    2    3    4    5    6    7    8    9   10    J    Q    K
}

create_hard_total_decisions :: proc() -> [22]map[cards.CardRank]Decision {
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
			decisions[total][dealer_top_card] = hard_total_decision[total][dealer_top_card]
		}
    }

    return decisions
}

// not exported
_HARD_TOTAL_DECISIONS:[22]map[cards.CardRank]Decision = create_hard_total_decisions()

get_hard_total_decision :: proc(
    player_total: uint,
    dealer_top_card: cards.CardRank,
) -> Decision {
    decision := _HARD_TOTAL_DECISIONS[player_total][dealer_top_card]
    return decision
}
