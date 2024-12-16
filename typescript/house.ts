export class HouseRules {
    // Hit/Stand on a soft 17 and 3:2 black jack payouts
    // are what casinos advetize wrt their BJ tables:
    // 1. 6/8 decks in shoe
    // 2. 3:2 blackjack payout
    // 3. Dealer Hit/Stand on a soft 17
    // 4. Re-splitting Aces (exceptionally rare)
    // 5. Surrender

    static DECKS_IN_SHOE = 6;
    static FORCE_RESHUFFLE: number = Math.floor(
        (52 * HouseRules.DECKS_IN_SHOE * 3) / 4,
    );

    // True => Must stand after the Ace split (stand on the Ace plus the one card dealt after split)
    // True => no double down after the Ace split, no splitting Aces after the Ace split
    static NO_MORE_CARDS_AFTER_SPLITTING_ACES = true;

    // [9, 10, 11] aka range(9, 12) => "Reno Rules"
    // static DOUBLE_DOWN_ON_TOTAL: number[] = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21]
    // python list comprehension can be faked with lambdas, unsure how TS-ic that is.
    static DOUBLE_DOWN_ON_TOTAL: number[] = ((): number[] => {
        const doubleOn: number[] = [];
        for (let i = 2; i <= 21; i++) {
            doubleOn.push(i);
        }
        return doubleOn;
    })();

    // Does not apply tp Aces if NO_MORE_CARDS_AFTER_SPLITTING_ACES is true
    static DOUBLE_DOWN_AFTER_SPLIT = true;

    // 3 => turn one hand into no more than 4 hands
    static SPLITS_PER_HAND = 3;

    // rank match like K-K always can split, values match allows K-10 split
    static SPLIT_ON_VALUE_MATCH = true;

    // Hit on soft 17 (6/8 decks) is more common on low bet tables.
    static DEALER_HITS_HARD_ON = 16; // or less
    static DEALER_HITS_SOFT_ON = 17; // or less

    // 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
    // 6 to 5 is more common in two deck games
    static NATURAL_BLACKJACK_PAYOUT = 1.5;

    // Usually 8 deck game, no Ace re-splitting, 50-100 minimum bet ...
    // Setting True here since I am a high roller ;) and want to shake out the code.
    static SURRENDER_ALLOWED = true;
}
Object.freeze(HouseRules);

export function canDoubleDown(total: number): boolean {
    return total in HouseRules.DOUBLE_DOWN_ON_TOTAL;
}
