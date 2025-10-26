// file src/strategy/pairs.rs defines project module "strategy::pairs".

use std::collections::HashMap;

use super::Decision;
use super::Decision::Dh;
use super::Decision::Ds;
use super::Decision::Uh;
use super::Decision::Us;
use super::Decision::Usp;
use super::Decision::H;
use super::Decision::NO;
use super::Decision::S;
use super::Decision::SP;

use crate::cards;

// turn off formating for the entire module
#[rustfmt::skip]

const _PAIR_DECISIONS: [[Decision; 14]; 14] = [
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
	[NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
	// player pair card: Ace  x  dealer top card
	[NO, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],
	// player pair card: 2  x  dealer top card
	[NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],
	// player pair card: 3  x  dealer top card
	[NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],
	// player pair card: 4  x  dealer top card
	[NO,  H,  H,  H,  H, SP, SP,  H,  H,  H,  H,  H,  H,  H],
	// player pair card: 5  x  dealer top card
	[NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H],
	// player pair card: 6  x  dealer top card
	[NO,  H, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H,  H],
	// player pair card: 7  x  dealer top card
	[NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],
	// player pair card: 8  x  dealer top card
	[NO, Usp, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],
	// player pair card: 9  x  dealer top card
	[NO,  S, SP, SP, SP, SP, SP,  S, SP, SP,  S,  S,  S,  S],
	// player pair card: 10  x  dealer top card
	[NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
	// player pair card: J  x  dealer top card
	[NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
	// player pair card: Q  x  dealer top card
	[NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
	// player pair card: K  x  dealer top card
	[NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
];

use lazy_static::lazy_static;

lazy_static! {
    static ref PAIR_DECISIONS: HashMap<cards::CardRank, HashMap<cards::CardRank, Decision>> = {
        let mut pair_decisions: HashMap<cards::CardRank, HashMap<cards::CardRank, Decision>> = HashMap::new();
        for split_rank in cards::CardRank::iterator() {
            // split_rank: &cards::CardRank
            let mut decisions_row: HashMap<cards::CardRank, Decision> = HashMap::new();
            for rank in cards::CardRank::iterator() {
                // rank: &cards::CardRank
                let decision: Decision = _PAIR_DECISIONS[split_rank.discriminant() as usize][rank.discriminant() as usize];
                decisions_row.insert(*rank, decision);
                // the trait `Eq` is not implemented for `CardRank`, which is required by `HashMap<_, _, _>: Index<&_>`
                // the trait `Hash` is not implemented for `CardRank`, which is required by `HashMap<_, _, _>: Index<&_>`
            }
            // hard_total_decisions[i] = decisions_row;
            pair_decisions.insert(*split_rank, decisions_row);
        }
        pair_decisions
    };
}

pub fn get_pair_decision(
    player_hand_pair_rank: cards::CardRank,
    dealer_top_card_rank: cards::CardRank,
) -> Decision {
    let decision: Decision = PAIR_DECISIONS[&player_hand_pair_rank][&dealer_top_card_rank];
    return decision;
}
