// file src/hand.rs defines project module "hand".

use std::mem::transmute;
use std::slice::Iter;

// #[derive(Eq, PartialEq, Hash, Copy, Clone)]
#[derive(Copy, Clone, PartialEq, Debug)]
#[repr(usize)]
pub enum HandOutcome {
    STAND,
    BUST,
    SURRENDER,
    DEALER_BLACKJACK,
    IN_PLAY,
}

impl HandOutcome {
    pub fn iterator() -> Iter<'static, HandOutcome> {
        static SUITES: [HandOutcome; 5] = [
            HandOutcome::STAND,
            HandOutcome::BUST,
            HandOutcome::SURRENDER,
            HandOutcome::DEALER_BLACKJACK,
            HandOutcome::IN_PLAY,
        ];
        SUITES.iter()
    }
    pub fn discriminant(&self) -> usize {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as usize
    }
    pub fn transmute(discrim: usize) -> HandOutcome {
        // FAILS: rank = CardRank(2);
        // FAILS: rank = 2 as CardRank;
        // WORKS: rank = unsafe { transmute(2 as u8) };
        unsafe { transmute(discrim) }
    }
    pub fn to_string(&self) -> String {
        static STRINGS: [&str; 5] = [
            "stand",            // HandOutcome::STAND
            "bust",             // HandOutcome::BUST
            "surrender",        // HandOutcome::SURRENDER
            "dealer-blackjack", // HandOutcome::DEALER_BLACKJACK
            "in-play",          // HandOutcome::IN_PLAY
        ];
        STRINGS[self.discriminant()].to_string()
    }
}

#[cfg(test)]
mod tests;
