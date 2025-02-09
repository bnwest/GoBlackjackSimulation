// file src/strategy/soft.rs defines project module "strategy::soft".

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

const _SOFT_TOTAL_DECISIONS: [[Decision; 14]; 22] = [
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 1 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 2 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 3 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 4 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 5 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 6 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 7 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 8 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 9 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 10 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 11 x dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft hand total 12 x dealer top card
	[NO,  H,  H,  H,  H,  H, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft hand total 13 x dealer top card
	[NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft hand total 14 x dealer top card
	[NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft hand total 15 x dealer top card
	[NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft hand total 16 x dealer top card
	[NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft hand total 17 x dealer top card
	[NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft hand total 18 x dealer top card
	[NO,  H, Ds, Ds, Ds, Ds, Ds,  S,  S,  H,  H,  H,  H,  H],
    // soft hand total 19 x dealer top card
	[NO,  S,  S,  S,  S,  S, Ds,  S,  S,  S,  S,  S,  S,  S],
    // soft hand total 20 x dealer top card
	[NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // soft hand total 21 x dealer top card
	[NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
];

use lazy_static::lazy_static;

lazy_static! {
    static ref SOFT_TOTAL_DECISIONS: Vec<HashMap<cards::CardRank, Decision>> = {
        let mut soft_total_decisions: Vec<HashMap<cards::CardRank, Decision>> = vec![];
        for i in 0..22 {
            let mut decisions_row: HashMap<cards::CardRank, Decision> = HashMap::new();
            for rank in cards::CardRank::iterator() {
                // rank: &cards::CardRank
                let decision: Decision = _SOFT_TOTAL_DECISIONS[i][rank.discriminant()];
                decisions_row.insert(*rank, decision);
                // the trait `Eq` is not implemented for `CardRank`, which is required by `HashMap<_, _, _>: Index<&_>`
                // the trait `Hash` is not implemented for `CardRank`, which is required by `HashMap<_, _, _>: Index<&_>`
            }
            // hard_total_decisions[i] = decisions_row;
            soft_total_decisions.push(decisions_row);
        }
        soft_total_decisions
    };
}

pub fn get_soft_total_decision(
    player_hand_total: usize,
    dealer_top_card: &cards::Card,
) -> Decision {
    let dealer_top_card_rank: cards::CardRank = dealer_top_card.rank;
    let decision: Decision = SOFT_TOTAL_DECISIONS[player_hand_total][&dealer_top_card_rank];
    return decision;
}
