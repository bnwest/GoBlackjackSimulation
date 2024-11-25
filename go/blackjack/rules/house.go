package rules

type HouseRules struct {
    Name string
    DECKS_IN_SHOE int
	FORCE_RESHUFFLE int
    NO_MORE_CARDS_AFTER_SPLITTING_ACES bool
    DOUBLE_DOWN_AFTER_SPLIT bool
	DOUBLE_DOWN_ON_TOTAL []int
    SPLITS_PER_HAND int
    SPLIT_ON_VALUE_MATCH bool
    DEALER_HITS_HARD_ON int
    DEALER_HITS_SOFT_ON int
    NATURAL_BLACKJACK_PAYOUT float32
    SURRENDER_ALLOWED bool
}

const Decks_in_Shoe int = 6

var HOUSE_RULES HouseRules = HouseRules{
    Name: "Alice",
    // Hit/Stand on a soft 17 and 3:2 black jack payouts
    // are what casinos advetize wrt their BJ tables:
    // 1. 6/8 decks in shoe
    // 2. 3:2 blackjack payout
    // 3. Hit/Stand on a soft 17
    // 4. Re-splitting Aces (exceptionally rare)
    // 5. Surrender

    DECKS_IN_SHOE: Decks_in_Shoe,
    FORCE_RESHUFFLE: ((52 * Decks_in_Shoe) * 3) / 4,

    // True => Must stand after the Ace split (stand on the Ace plus the one card dealt after split)
    // True => no double down after the Ace split, no splitting Aces after the Ace split
    NO_MORE_CARDS_AFTER_SPLITTING_ACES: true,

    // [9, 10, 11] aka range(9, 12) => "Reno Rules"
    DOUBLE_DOWN_ON_TOTAL: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},

    // Does not apply tp Aces if NO_MORE_CARDS_AFTER_SPLITTING_ACES is true
    DOUBLE_DOWN_AFTER_SPLIT: true,

    // 3 => turn one hand into no more than 4 hands
    SPLITS_PER_HAND: 3,

    // rank match like K-K always can split, values match allows K-10 split
    SPLIT_ON_VALUE_MATCH: true,

    // Hit on soft 17 (6/8 decks) is more common on low bet tables.
    DEALER_HITS_HARD_ON: 16,  // or less
    DEALER_HITS_SOFT_ON: 17,  // or less

    // 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
    // 6 to 5 is more common in two deck games
    NATURAL_BLACKJACK_PAYOUT: 1.5,

    // Usually 8 deck game, no Ace re-splitting, 50-100 minimum bet ...
    // Setting True here since I am a high roller ;) and want to shake out the code.
    SURRENDER_ALLOWED: true,
}
