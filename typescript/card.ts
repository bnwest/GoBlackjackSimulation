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

const CardValue: Map<CardRank, number> = new Map();
CardValue.set(CardRank.ACE,    1);
CardValue.set(CardRank.TWO,    2);
CardValue.set(CardRank.THREE,  3);
CardValue.set(CardRank.FOUR,   4);
CardValue.set(CardRank.FIVE,   5);
CardValue.set(CardRank.SIX,    6);
CardValue.set(CardRank.SEVEN,  7);
CardValue.set(CardRank.EIGHT,  8);
CardValue.set(CardRank.NINE,   9);
CardValue.set(CardRank.TEN,   10);
CardValue.set(CardRank.JACK,  10);
CardValue.set(CardRank.QUEEN, 10);
CardValue.set(CardRank.KING,  10);
Object.freeze(CardValue);

export function getCardValue(rank: CardRank): number {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const value: number = CardValue.get(rank) as number;
    return value;
}

const CardIndex: Map<CardRank, number> = new Map();
CardIndex.set(CardRank.ACE,    1);
CardIndex.set(CardRank.TWO,    2);
CardIndex.set(CardRank.THREE,  3);
CardIndex.set(CardRank.FOUR,   4);
CardIndex.set(CardRank.FIVE,   5);
CardIndex.set(CardRank.SIX,    6);
CardIndex.set(CardRank.SEVEN,  7);
CardIndex.set(CardRank.EIGHT,  8);
CardIndex.set(CardRank.NINE,   9);
CardIndex.set(CardRank.TEN,   10);
CardIndex.set(CardRank.JACK,  11);
CardIndex.set(CardRank.QUEEN, 12);
CardIndex.set(CardRank.KING,  13);
Object.freeze(CardIndex);

export function getCardIndex(rank: CardRank): number {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const value: number = CardIndex.get(rank) as number;
    return value;
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
