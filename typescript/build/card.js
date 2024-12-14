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
const CardValue = new Map();
CardValue.set(CardRank.ACE, 1);
CardValue.set(CardRank.TWO, 2);
CardValue.set(CardRank.THREE, 3);
CardValue.set(CardRank.FOUR, 4);
CardValue.set(CardRank.FIVE, 5);
CardValue.set(CardRank.SIX, 6);
CardValue.set(CardRank.SEVEN, 7);
CardValue.set(CardRank.EIGHT, 8);
CardValue.set(CardRank.NINE, 9);
CardValue.set(CardRank.TEN, 10);
CardValue.set(CardRank.JACK, 10);
CardValue.set(CardRank.QUEEN, 10);
CardValue.set(CardRank.KING, 10);
Object.freeze(CardValue);
function getCardValue(rank) {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const value = CardValue.get(rank);
    return value;
}
const CardIndex = new Map();
CardIndex.set(CardRank.ACE, 1);
CardIndex.set(CardRank.TWO, 2);
CardIndex.set(CardRank.THREE, 3);
CardIndex.set(CardRank.FOUR, 4);
CardIndex.set(CardRank.FIVE, 5);
CardIndex.set(CardRank.SIX, 6);
CardIndex.set(CardRank.SEVEN, 7);
CardIndex.set(CardRank.EIGHT, 8);
CardIndex.set(CardRank.NINE, 9);
CardIndex.set(CardRank.TEN, 10);
CardIndex.set(CardRank.JACK, 11);
CardIndex.set(CardRank.QUEEN, 12);
CardIndex.set(CardRank.KING, 13);
Object.freeze(CardIndex);
function getCardIndex(rank) {
    // typescript does not know that the CardValue map has ALL 
    // of the CardRank values but I do.
    const value = CardIndex.get(rank);
    return value;
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
