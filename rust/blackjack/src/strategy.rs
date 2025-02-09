// file src/strategy.rs defines project module "strategy".

use std::mem::transmute;
use std::slice::Iter;

#[derive(Copy, Clone, PartialEq, Debug)]
#[repr(usize)]
pub enum Decision {
    S,
    H,
    Dh,
    Ds,
    SP,
    // U => Surrender, in a world of too many S words
    Uh,
    Us,
    Usp,
    NO,
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
    pub fn discriminant(&self) -> usize {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as usize
        // ^^^^^ move occurs because `*self` has type `Decision`,
        // which does not implement the `Copy` trait
    }
    pub fn transmute(discrim: usize) -> Decision {
        // FAILS: rank = Decision(2);
        // FAILS: rank = 2 as Decision;
        // WORKS: rank = unsafe { transmute(2 as usize) };
        unsafe { transmute(discrim) }
        //    = note: source type: `usize` (64 bits)     << usize
        //    = note: target type: `Decision` (8 bits)   << default enum size
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
#[repr(usize)]
pub enum PlayerDecision {
    STAND,
    HIT,
    DOUBLE,
    SPLIT,
    SURRENDER,
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
    pub fn discriminant(&self) -> usize {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as usize
    }
    pub fn transmute(discrim: usize) -> PlayerDecision {
        // FAILS: rank = Decision(2);
        // FAILS: rank = 2 as Decision;
        // WORKS: rank = unsafe { transmute(2 as usize) };
        unsafe { transmute(discrim) }
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
mod soft;
mod pairs;

#[cfg(test)]
mod tests;
