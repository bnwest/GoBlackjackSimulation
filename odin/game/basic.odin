package game

import "../cards"
import "../strategy"
import house_rules "../rules"

convert_to_player_decision :: proc(
    decision: strategy.Decision,
    player_hand: ^PlayerHand,
) -> strategy.PlayerDecision {
	// Decision sometimes return Xy, which translates to do X if allowed else do y.
	// Determine the X or the y here.

    is_first_decision: bool = num_cards(player_hand) == 2
    is_first_post_split_decision: bool = (
        is_first_decision && is_from_split(player_hand)
    )

    player_decision: strategy.PlayerDecision

    hard_count := hard_count(player_hand)
    soft_count := soft_count(player_hand)

    if decision == strategy.Decision.S {
        player_decision = strategy.PlayerDecision.STAND

    } else if decision == strategy.Decision.H {
        player_decision = strategy.PlayerDecision.HIT

    } else if decision == strategy.Decision.Dh || decision == strategy.Decision.Ds {
		// may be only allow to down on hand totals [9, 10, 11] or some such
		// basic stratgey wants to double down on
		//     hand hard totals [9, 10, 11]
		//     hand soft totals [12, 13,14, 15, 16, 17, 18, 19]
		nondouble_down_decision: strategy.PlayerDecision
        nondouble_down_decision = (
            decision == strategy.Decision.Ds 
                ? strategy.PlayerDecision.STAND 
                : strategy.PlayerDecision.HIT
        )

        can_double_down: bool
        if is_first_decision {
            if is_first_post_split_decision {
                if house_rules.DOUBLE_DOWN_AFTER_SPLIT {
                    can_double_down = true
                } else {
                    can_double_down = false
                }
            } else {
                can_double_down = true
            }
        }

        if can_double_down {
            double_down: bool
            if house_rules.can_double_down(hard_count) {
                double_down = true
            } else if house_rules.can_double_down(soft_count) {
                double_down = true
            } else {
                double_down = false
            }
            if double_down {
                player_decision = strategy.PlayerDecision.DOUBLE
            } else {
                player_decision = nondouble_down_decision
            }
        } else {
            player_decision = nondouble_down_decision
        }

    } else if decision == strategy.Decision.SP {
        player_decision = strategy.PlayerDecision.SPLIT

    } else if decision == strategy.Decision.Uh || decision == strategy.Decision.Us || decision == strategy.Decision.Usp {
		// surrent decision must be allowed in the House Rules and
		// must be a first decision (before splitting)
		nonsurrender_decision: strategy.PlayerDecision
        nonsurrender_decision = (
            decision == strategy.Decision.Uh
                ? strategy.PlayerDecision.HIT
                : decision == strategy.Decision.Us
                    ? strategy.PlayerDecision.STAND
                    : strategy.PlayerDecision.SPLIT
        )

        surrender_can_be_played: bool
        surrender_can_be_played = (
            is_first_decision 
            && !is_first_post_split_decision 
            && house_rules.SURRENDER_ALLOWED
        )

        player_decision = (
            surrender_can_be_played 
                ? strategy.PlayerDecision.SURRENDER 
                : nonsurrender_decision
        )
    }

    return player_decision
}

determine_basic_strategy_play :: proc(
    dealer_top_card: cards.Card,
    player_hand: ^PlayerHand,
    hand_allows_more_splits: bool,
) -> strategy.PlayerDecision {
    is_first_decision: bool = num_cards(player_hand) == 2

    player_card1: cards.Card = get_card(player_hand, card_index=0)
    player_card2: cards.Card = get_card(player_hand, card_index=1)

    decision: strategy.Decision
    player_decision: strategy.PlayerDecision

    got_pairs: bool
    if is_first_decision {
        if house_rules.SPLIT_ON_VALUE_MATCH {
            // all cards with value 10 (10, J, Q, K) match
            got_pairs = cards.to_int(player_card1.rank) == cards.to_int(player_card2.rank)
        } else {
            got_pairs = player_card1.rank == player_card2.rank
        }
    } else {
        got_pairs = false
    }

    if got_pairs && hand_allows_more_splits {
		// Determine if the pairs can be split.
		// Note all of the non-split decisions that are ignored below
		// will not contradict the hard/soft total decision.
		pair_rank: cards.CardRank
        if cards.to_int(player_card1.rank) == 10 {
            pair_rank = cards.CardRank.TEN
        }
        else {
            pair_rank = player_card1.rank
        }

        decision = strategy.get_pairs_split_decision(pair_rank, dealer_top_card.rank)
        player_decision = convert_to_player_decision(decision, player_hand)
        if player_decision == strategy.PlayerDecision.SPLIT {
            return strategy.PlayerDecision.SPLIT
        }
    }

    hard_count := hard_count(player_hand)
    soft_count := soft_count(player_hand)

    use_soft_count: bool = hard_count < soft_count && soft_count <= 21
    if use_soft_count {
        decision = strategy.get_soft_total_decision(soft_count, dealer_top_card.rank)
        player_decision = convert_to_player_decision(decision, player_hand)
    } else {
        decision = strategy.get_hard_total_decision(hard_count, dealer_top_card.rank)
        player_decision = convert_to_player_decision(decision, player_hand)
    }

    return player_decision
}
