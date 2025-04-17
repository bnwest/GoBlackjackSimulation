package rules

// instead of having a bag of constants in a struct,
// have a bag of constants via a package namespace.

// const DECKS_IN_SHOE: int = 6
// Only declarations are allowed at file scope, got expression statement 
DECKS_IN_SHOE :: 6

FORCE_RESHUFFLE :: ((52 * DECKS_IN_SHOE) * 3) / 4

// True => Must stand after the Ace split (stand on the Ace plus the one card dealt after split)
// True => no double down after the Ace split,
// True => no splitting Aces after the Ace split
NO_MORE_CARDS_AFTER_SPLITTING_ACES :: true

// [9, 10, 11] aka range(9, 12) => "Reno Rules"
DOUBLE_DOWN_ON_TOTAL :: []uint {
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21
}

can_double_down :: proc(total: uint) -> bool {
    for key in DOUBLE_DOWN_ON_TOTAL {
        if key == total {
            return true
        }
    }
    return false
}

// Does not apply tp Aces if NO_MORE_CARDS_AFTER_SPLITTING_ACES is true
DOUBLE_DOWN_AFTER_SPLIT :: true

// 3 => turn one hand into no more than 4 hands
SPLITS_PER_HAND :: 3

// rank match like K-K always can split, values match allows K-10 split
SPLIT_ON_VALUE_MATCH :: true

// Hit on soft 17 (6/8 decks) is more common on low bet tables.
DEALER_HITS_HARD_ON :: 16 // or less
DEALER_HITS_SOFT_ON :: 17 // or less

// 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
// 6 to 5 is more common in two deck games
NATURAL_BLACKJACK_PAYOUT :: 1.5

// Usually 8 deck game, no Ace re-splitting, 50-100 minimum bet ...
// Setting True here since I am a high roller ;) and want to shake out the code.
SURRENDER_ALLOWED :: true
