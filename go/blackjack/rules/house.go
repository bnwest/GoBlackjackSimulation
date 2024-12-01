package rules

// instead of having a bag of constants in a struct,
// have a bag of constants via a package namespace.

const DECKS_IN_SHOE int = 6

const FORCE_RESHUFFLE int = ((52 * DECKS_IN_SHOE) * 3) / 4

// True => Must stand after the Ace split (stand on the Ace plus the one card dealt after split)
// True => no double down after the Ace split, no splitting Aces after the Ace split
const NO_MORE_CARDS_AFTER_SPLITTING_ACES bool = true

// [9, 10, 11] aka range(9, 12) => "Reno Rules"
var double_down_on_total []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
func CanDoubleDown(total int) bool {
	// Go does not support constant arrays, maps or slices.
	// GO also does not the "in" operator, eg total in double_down_on_total
	for i := 0; i < len(double_down_on_total); i++ {
		if double_down_on_total[i] == total {
			return true
		}
	}
	return false
}

// Does not apply tp Aces if NO_MORE_CARDS_AFTER_SPLITTING_ACES is true
const DOUBLE_DOWN_AFTER_SPLIT bool = true

// 3 => turn one hand into no more than 4 hands
const SPLITS_PER_HAND int = 3

// rank match like K-K always can split, values match allows K-10 split
const SPLIT_ON_VALUE_MATCH bool = true

// Hit on soft 17 (6/8 decks) is more common on low bet tables.
const DEALER_HITS_HARD_ON int = 16  // or less
const DEALER_HITS_SOFT_ON int = 17  // or less

// 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
// 6 to 5 is more common in two deck games
const NATURAL_BLACKJACK_PAYOUT float32 = 1.5

// Usually 8 deck game, no Ace re-splitting, 50-100 minimum bet ...
// Setting True here since I am a high roller ;) and want to shake out the code.
const SURRENDER_ALLOWED bool = true
