// file src/strategy.rs defines project module "strategy".

use std::mem::transmute;
use std::slice::Iter;

#[derive(Copy, Clone, PartialEq, Debug)]
#[repr(u8)]
pub enum Decision {
    S = 0,
    H = 1,
    Dh = 2,
    Ds = 3,
    SP = 4,
    // U => Surrender, in a world of too many S words
    Uh = 5,
    Us = 6,
    Usp = 7,
    NO = 8,
}

impl Decision {
    pub fn iterator() -> Iter<'static, Decision> {
        static DECISIONS: [Decision; 9] = [
            Decision::S,
            Decision::H,
            Decision::Dh,
            Decision::Ds,
            Decision::SP,
            Decision::Uh,
            Decision::Us,
            Decision::Usp,
            Decision::NO,
        ];
        DECISIONS.iter()
    }
    pub fn discriminant(&self) -> u8 {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as u8
        // ^^^^^ move occurs because `*self` has type `Decision`,
        // which does not implement the `Copy` trait
    }
    pub fn transmute(discrim: u8) -> Decision {
        // FAILS: rank = Decision(2);
        // FAILS: rank = 2 as Decision;
        // WORKS POORLY: rank = unsafe { transmute(2 as usize) };
        match discrim {
            0 => Decision::S,
            1 => Decision::H,
            2 => Decision::Dh,
            3 => Decision::Ds,
            4 => Decision::SP,
            5 => Decision::Uh,
            6 => Decision::Us,
            7 => Decision::Usp,
            8 => Decision::NO,
            _ => Decision::S, // Default fallback
        }
    }
    pub fn to_string(&self) -> String {
        static STRINGS: [&str; 9] = [
            "stand",                           // Decision::S
            "hit",                             // Decision::H
            "double-down-if-allowed-or-hit",   // Decision:Dh
            "double-down-if-allowed-or-stand", // Decision::Ds
            "split",                           // Decision::SP
            "surrender-if-allowed-or-hit",     // Decision::Uh
            "surrender-if-allowed-or-stand",   // Decision::Us
            "surrender-if-allowed-or-split",   // Decision::Usp
            "no-decision",                     // Decision::NO
        ];
        STRINGS[self.discriminant() as usize].to_string()
    }
}

#[derive(Copy, Clone, PartialEq, Debug)]
#[repr(u8)]
pub enum PlayerDecision {
    STAND = 0,
    HIT = 1,
    DOUBLE = 2,
    SPLIT = 3,
    SURRENDER = 4,
}

impl PlayerDecision {
    pub fn iterator() -> Iter<'static, PlayerDecision> {
        static PLAYER_DECISIONS: [PlayerDecision; 5] = [
            PlayerDecision::STAND,
            PlayerDecision::HIT,
            PlayerDecision::DOUBLE,
            PlayerDecision::SPLIT,
            PlayerDecision::SURRENDER,
        ];
        PLAYER_DECISIONS.iter()
    }
    pub fn discriminant(&self) -> u8 {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as u8
    }
    pub fn transmute(discrim: u8) -> PlayerDecision {
        // FAILS: rank = Decision(2);
        // FAILS: rank = 2 as Decision;
        // WORKS POORLY: rank = unsafe { transmute(2 as usize) };
        match discrim {
            0 => PlayerDecision::STAND,
            1 => PlayerDecision::HIT,
            2 => PlayerDecision::DOUBLE,
            3 => PlayerDecision::SPLIT,
            4 => PlayerDecision::SURRENDER,
            _ => PlayerDecision::STAND, // Default fallback
        }
    }
    pub fn to_string(&self) -> String {
        static STRINGS: [&str; 5] = [
            "stand",       // PlayerDecision::STAND
            "hit",         // PlayerDecision::HIT
            "double-down", // PlayerDecision:DOUBLE
            "split",       // PlayerDecision::SPLIT
            "surrender",   // PlayerDecision::SURRENDER
        ];
        STRINGS[self.discriminant() as usize].to_string()
    }
}

mod hard;
mod pairs;
mod soft;

use crate::cards::{self, Card};
use crate::hand::PlayerHand;
use crate::rules;

pub fn convert_to_player_decision(decision: Decision, player_hand: &PlayerHand) -> PlayerDecision {
    // Decision sometimes return Xy, which translates to do X if allowed else do y.
    // Determine the X or the y here
    let is_first_decision: bool = player_hand.num_cards() == 2;
    let is_first_post_split_decision: bool = is_first_decision && player_hand.is_from_split();

    let hard_count: usize = player_hand.hard_count();
    let soft_count: usize = player_hand.soft_count();

    let player_decision: PlayerDecision;
    if decision == Decision::S {
        player_decision = PlayerDecision::STAND;
    } else if decision == Decision::H {
        player_decision = PlayerDecision::HIT;
    } else if decision == Decision::Dh || decision == Decision::Ds {
        let non_double_down_decision: PlayerDecision;
        if decision == Decision::Dh {
            non_double_down_decision = PlayerDecision::HIT;
        } else {
            non_double_down_decision = PlayerDecision::STAND;
        }

        let can_double_down: bool;
        if is_first_decision {
            if is_first_post_split_decision {
                if rules::DOUBLE_DOWN_AFTER_SPLIT {
                    can_double_down = true;
                } else {
                    can_double_down = false;
                }
            } else {
                can_double_down = true;
            }
        } else {
            can_double_down = false;
        }

        if can_double_down {
            let double_down: bool;
            if rules::can_double_down(hard_count) {
                double_down = true;
            } else if rules::can_double_down(soft_count) {
                double_down = true;
            } else {
                double_down = false;
            }
            if double_down {
                player_decision = PlayerDecision::DOUBLE;
            } else {
                player_decision = non_double_down_decision;
            }
        } else {
            player_decision = non_double_down_decision;
        }
    } else if decision == Decision::SP {
        player_decision = PlayerDecision::SPLIT;
    } else if decision == Decision::Uh || decision == Decision::Us || decision == Decision::Usp {
        let non_surrender_decision: PlayerDecision;
        match decision {
            Decision::Uh => non_surrender_decision = PlayerDecision::HIT,
            Decision::Us => non_surrender_decision = PlayerDecision::STAND,
            Decision::Usp => non_surrender_decision = PlayerDecision::SPLIT,
            _ => panic!("convertToPlayerDecision() ran into a little trouble in town."),
        }

        let can_surrender_be_played: bool;
        can_surrender_be_played =
            is_first_decision && !is_first_post_split_decision && rules::SURRENDER_ALLOWED;
        if can_surrender_be_played {
            player_decision = PlayerDecision::SURRENDER;
        } else {
            player_decision = non_surrender_decision;
        }
    } else {
        // rust compiler errors out with this else clause:
        //     `player_decision` used here but it is possibly-uninitialized
        // rust compiler is wrong but has a very strong opinion.
        panic!("convertToPlayerDecision() ran into a little trouble in town.");
    }

    return player_decision;
}

pub fn determine_basic_strategy(
    dealer_top_card: &cards::Card,
    player_hand: &PlayerHand,
    hand_allow_more_splits: bool,
) -> PlayerDecision {
    let is_first_decision: bool = player_hand.num_cards() == 2;

    let player_card1: cards::Card = player_hand.get_card(0);
    let player_card2: cards::Card = player_hand.get_card(1);

    let decision: Decision;
    let player_decision: PlayerDecision;

    let got_pairs: bool;
    if is_first_decision {
        if rules::SPLIT_ON_VALUE_MATCH {
            got_pairs = player_card1.rank.value() == player_card2.rank.value();
        } else {
            got_pairs = player_card1.rank == player_card2.rank;
        }
    } else {
        got_pairs = false;
    }

    if got_pairs && hand_allow_more_splits {
        // Determine if the pairs can be split.
        // Note all of the non-split decisions that are ignored below
        // will not contradict the hard/soft total decision.
        let pair_rank: cards::CardRank;
        if player_card1.rank.value() == 10 {
            pair_rank = cards::CardRank::TEN;
        } else {
            pair_rank = player_card1.rank;
        }

        decision = pairs::get_pair_decision(pair_rank, dealer_top_card.rank);
        player_decision = convert_to_player_decision(decision, player_hand);
        if player_decision == PlayerDecision::SPLIT {
            return PlayerDecision::SPLIT;
        }
    }

    let hard_count: usize = player_hand.hard_count();
    let soft_count: usize = player_hand.soft_count();
    let use_soft_count: bool = hard_count < soft_count && soft_count <= 21;

    // define these two, since rust only allow immutable variable to declare once
    // and we could have fallen through from the above pairs decisions.
    let decision: Decision;
    let player_decision: PlayerDecision;

    if use_soft_count {
        decision = soft::get_soft_total_decision(soft_count, dealer_top_card);
    } else {
        decision = hard::get_hard_total_decision(hard_count, dealer_top_card);
    }

    player_decision = convert_to_player_decision(decision, player_hand);
    return player_decision;
}

#[cfg(test)]
mod tests;
