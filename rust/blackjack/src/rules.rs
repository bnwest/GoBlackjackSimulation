// file src/house.rs defines project module "house".

// instead of having a bag of constants in a struct,
// have a bag of constants via a module namespace.

pub const DECKS_IN_SHOE: usize = 6;

pub const FORCE_RESHUFFLE: usize = ((52 * DECKS_IN_SHOE) * 3) / 4;

// True => Must stand after the Ace split (stand on the Ace plus the one card dealt after split)
// True => no double down after the Ace split,
// True => no splitting Aces after the Ace split
pub const NO_MORE_CARDS_AFTER_SPLITTING_ACES: bool = true;

// [9, 10, 11] aka range(9, 12) => "Reno Rules"
//var double_down_on_total []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
pub const CAN_DOUBLE_DOWN_ON_TOTAL: [usize; 21] = [
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
];

pub fn can_double_down(top_card_value: usize) -> bool {
    //return top_card_value in CAN_DOUBLE_DOWN_ON_TOTAL;
    for elem in CAN_DOUBLE_DOWN_ON_TOTAL.iter() {
        if *elem == top_card_value {
            return true;
        }
    }
    return false;
}

// Does not apply to Aces if NO_MORE_CARDS_AFTER_SPLITTING_ACES is true
pub const DOUBLE_DOWN_AFTER_SPLIT: bool = true;

// 3 => turn one hand into no more than 4 hands
pub const SPLITS_PER_HAND: usize = 3;

// rank match like K-K always can split, values match allows K-10 split
pub const SPLIT_ON_VALUE_MATCH: bool = true;

// Hit on soft 17 (6/8 decks) is more common on low bet tables.
pub const DEALER_HITS_HARD_ON: usize = 16; // or less
pub const DEALER_HITS_SOFT_ON: usize = 17; // or less

// 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
// 6 to 5 is more common in two deck games
pub const NATURAL_BLACKJACK_PAYOUT: f32 = 1.5;

// Usually 8 deck game, no Ace re-splitting, 50-100 minimum bet ...
// Setting True here since I am a high roller ;) and want to shake out the code.
pub const SURRENDER_ALLOWED: bool = true;
