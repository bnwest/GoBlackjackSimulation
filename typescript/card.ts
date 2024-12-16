export enum CardSuite {
    HEARTS   = "♥️",  // aka U+2665 + U+fe0f aka "\u2665\uFE0F"
    DIAMONDS = "♦️",  // aka U+2666 + U+fe0f aka "\u2666\uFE0F"
    SPADES   = "♠️",  // aka U+2660 + U+fe0f aka "\u2660\uFE0F"
    CLUBS    = "♣️",  // aka U+2663 + U+fe0f aka "\u2663\uFE0F"
};

export enum CardRank {
    ACE   = "A",
    TWO   = "2",
    THREE = "3",
    FOUR  = "4",
    FIVE  = "5",
    SIX   = "6",
    SEVEN = "7",
    EIGHT = "8",
    NINE  = "9",
    TEN   ="10",
    JACK  = "J",
    QUEEN = "Q",
    KING  = "K",
};

const CardValue: Map<CardRank, number> = new Map<CardRank, number>([
    [CardRank.ACE,    1],
    [CardRank.TWO,    2],
    [CardRank.THREE,  3],
    [CardRank.FOUR,   4],
    [CardRank.FIVE,   5],
    [CardRank.SIX,    6],
    [CardRank.SEVEN,  7],
    [CardRank.EIGHT,  8],
    [CardRank.NINE,   9],
    [CardRank.TEN,   10],
    [CardRank.JACK,  10],
    [CardRank.QUEEN, 10],
    [CardRank.KING,  10],
]);
Object.freeze(CardValue);

export function getCardValue(rank: CardRank): number {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const value: number = CardValue.get(rank) as number;
    return value;
}

const CardIndex: Map<CardRank, number> = new  Map<CardRank, number>([
    [CardRank.ACE,    1],
    [CardRank.TWO,    2],
    [CardRank.THREE,  3],
    [CardRank.FOUR,   4],
    [CardRank.FIVE,   5],
    [CardRank.SIX,    6],
    [CardRank.SEVEN,  7],
    [CardRank.EIGHT,  8],
    [CardRank.NINE,   9],
    [CardRank.TEN,   10],
    [CardRank.JACK,  11],
    [CardRank.QUEEN, 12],
    [CardRank.KING,  13],
])
Object.freeze(CardIndex);

export function getCardIndex(rank: CardRank): number {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const index: number = CardIndex.get(rank)!; // as number;
    return index;
}

export class Card {
    suite: CardSuite;
    rank: CardRank;
    constructor(suite: CardSuite, rank: CardRank) {
        this.suite = suite;
        this.rank = rank;
    }
    str(): string {
        return `${this.rank}${this.suite}`;
    }
}
