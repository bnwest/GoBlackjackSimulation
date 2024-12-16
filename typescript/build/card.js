"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Card = exports.CardRank = exports.CardSuite = void 0;
exports.getCardValue = getCardValue;
exports.getCardIndex = getCardIndex;
var CardSuite;
(function (CardSuite) {
    CardSuite["HEARTS"] = "\u2665\uFE0F";
    CardSuite["DIAMONDS"] = "\u2666\uFE0F";
    CardSuite["SPADES"] = "\u2660\uFE0F";
    CardSuite["CLUBS"] = "\u2663\uFE0F";
})(CardSuite || (exports.CardSuite = CardSuite = {}));
;
var CardRank;
(function (CardRank) {
    CardRank["ACE"] = "A";
    CardRank["TWO"] = "2";
    CardRank["THREE"] = "3";
    CardRank["FOUR"] = "4";
    CardRank["FIVE"] = "5";
    CardRank["SIX"] = "6";
    CardRank["SEVEN"] = "7";
    CardRank["EIGHT"] = "8";
    CardRank["NINE"] = "9";
    CardRank["TEN"] = "10";
    CardRank["JACK"] = "J";
    CardRank["QUEEN"] = "Q";
    CardRank["KING"] = "K";
})(CardRank || (exports.CardRank = CardRank = {}));
;
const CardValue = new Map([
    [CardRank.ACE, 1],
    [CardRank.TWO, 2],
    [CardRank.THREE, 3],
    [CardRank.FOUR, 4],
    [CardRank.FIVE, 5],
    [CardRank.SIX, 6],
    [CardRank.SEVEN, 7],
    [CardRank.EIGHT, 8],
    [CardRank.NINE, 9],
    [CardRank.TEN, 10],
    [CardRank.JACK, 10],
    [CardRank.QUEEN, 10],
    [CardRank.KING, 10],
]);
Object.freeze(CardValue);
function getCardValue(rank) {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const value = CardValue.get(rank);
    return value;
}
const CardIndex = new Map([
    [CardRank.ACE, 1],
    [CardRank.TWO, 2],
    [CardRank.THREE, 3],
    [CardRank.FOUR, 4],
    [CardRank.FIVE, 5],
    [CardRank.SIX, 6],
    [CardRank.SEVEN, 7],
    [CardRank.EIGHT, 8],
    [CardRank.NINE, 9],
    [CardRank.TEN, 10],
    [CardRank.JACK, 11],
    [CardRank.QUEEN, 12],
    [CardRank.KING, 13],
]);
Object.freeze(CardIndex);
function getCardIndex(rank) {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const index = CardIndex.get(rank); // as number;
    return index;
}
class Card {
    constructor(suite, rank) {
        this.suite = suite;
        this.rank = rank;
    }
    str() {
        return `${this.rank}${this.suite}`;
    }
}
exports.Card = Card;
