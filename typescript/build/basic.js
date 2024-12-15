"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PlayerDecision = void 0;
exports.determineBasicStrategyPlay = determineBasicStrategyPlay;
const card_1 = require("./card");
const house_1 = require("./house");
//
// Basic Strategy:
//     https://en.wikipedia.org/wiki/Blackjack#Basic_strategy
//     All Hail The Wikipedia!
//
// to make decision tables readable:
const S = "stand";
const H = "hit";
const Dh = "double-down-if-allowed-or-hit";
const Ds = "double-down-if-allowed-or-stand";
const SP = "split";
// U => Surrender, in a world of too many S words
const Uh = "surrender-if-allowed-or-hit";
const Us = "surrender-if-allowed-or-stand";
const Usp = "surrender-if-allowed-or-split";
const NO = "no-decision";
function isDecisionValid(decision) {
    switch (decision) {
        case S:
        case H:
        case Dh:
        case Ds:
        case SP:
        case Uh:
        case Us:
        case Usp:
        case NO:
            return true;
        default:
            return false;
    }
}
var PlayerDecision;
(function (PlayerDecision) {
    PlayerDecision["STAND"] = "stand";
    PlayerDecision["HIT"] = "hit";
    PlayerDecision["DOUBLE"] = "double-down";
    PlayerDecision["SPLIT"] = "split";
    PlayerDecision["SURRENDER"] = "surrender";
})(PlayerDecision || (exports.PlayerDecision = PlayerDecision = {}));
function convertToPlayerDecision(decision, playerHand) {
    // Decision sometimes return Xy, which translates to do X if allowed else do y.
    // Determine the X or the y here.
    const isFirstDecision = (playerHand.numCards == 2);
    const isFirstPostSplitDecision = (isFirstDecision && playerHand.isFromSplit);
    var playerDecision = PlayerDecision.STAND;
    const hardCount = playerHand.hardCount;
    const softCount = playerHand.softCount;
    if (decision == S) {
        return PlayerDecision.STAND;
    }
    else if (decision == H) {
        return PlayerDecision.HIT;
    }
    else if (decision == Dh || decision == Ds) {
        // may be only allow to down on hand totals [9, 10, 11] or some such
        // basic stratgey wants to double down on
        //     hand hard totals [9, 10, 11]
        //     hand soft totals [12, 13,14, 15, 16, 17, 18, 19]
        var nondoubleDownDecision;
        if (decision == Dh) {
            nondoubleDownDecision = PlayerDecision.HIT;
        }
        else {
            nondoubleDownDecision = PlayerDecision.STAND;
        }
        var doubleDownAllowed;
        if (isFirstDecision) {
            if (isFirstPostSplitDecision) {
                if (house_1.HouseRules.DOUBLE_DOWN_AFTER_SPLIT) {
                    doubleDownAllowed = true;
                }
                else {
                    doubleDownAllowed = false;
                }
            }
            else {
                doubleDownAllowed = true;
            }
        }
        else {
            doubleDownAllowed = false;
        }
        if (doubleDownAllowed) {
            var doubleDown;
            if ((0, house_1.canDoubleDown)(hardCount)) {
                doubleDown = true;
            }
            else if ((0, house_1.canDoubleDown)(softCount)) {
                doubleDown = true;
            }
            else {
                doubleDown = false;
            }
            playerDecision = doubleDown ? PlayerDecision.DOUBLE : nondoubleDownDecision;
        }
        else {
            playerDecision = nondoubleDownDecision;
        }
    }
    else if (decision == SP) {
        playerDecision = PlayerDecision.SPLIT;
    }
    else if (decision == Uh || decision == Us || decision == Usp) {
        // surrent decision must be allowed in the House Rules and
        // must be a first decision (before splitting)
        var nonsurrenderDecision;
        switch (decision) {
            case Uh:
                nonsurrenderDecision = PlayerDecision.HIT;
                break;
            case Us:
                nonsurrenderDecision = PlayerDecision.STAND;
                break;
            case Usp:
                nonsurrenderDecision = PlayerDecision.SPLIT;
                break;
            default:
                throw new Error("convertToPlayerDecision() ran into a little trouble in town");
        }
        var surrenderAllowed = (isFirstDecision && !isFirstPostSplitDecision && house_1.HouseRules.SURRENDER_ALLOWED);
        playerDecision = surrenderAllowed ? PlayerDecision.SURRENDER : nonsurrenderDecision;
    }
    // console.log(`convertToPlayerDecision: decision "${decision}" -> playerDecision "${playerDecision}"`);
    return playerDecision;
}
// Use NO since we live in a zero index world.
// Only use the SP aka SPLIT decision from this table.  The other decisions 
// mirror exactly what the hard/soft total decision tables yield.
const _PAIRS_DECISION = [
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // player pair card: Ace  x  dealer top card
    [NO, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],
    // player pair card: 2  x  dealer top card
    [NO, H, SP, SP, SP, SP, SP, SP, H, H, H, H, H, H],
    // player pair card: 3  x  dealer top card
    [NO, H, SP, SP, SP, SP, SP, SP, H, H, H, H, H, H],
    // player pair card: 4  x  dealer top card
    [NO, H, H, H, H, SP, SP, H, H, H, H, H, H, H],
    // player pair card: 5  x  dealer top card
    [NO, H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, H, H, H, H],
    // player pair card: 6 x  dealer top card
    [NO, H, SP, SP, SP, SP, SP, H, H, H, H, H, H, H],
    // player pair card: 7  x  dealer top card
    [NO, H, SP, SP, SP, SP, SP, SP, H, H, H, H, H, H],
    // player pair card: 8  x  dealer top card
    [NO, Usp, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],
    // player pair card: 9  x  dealer top card
    [NO, S, SP, SP, SP, SP, SP, S, SP, SP, S, S, S, S],
    // player pair card: 10  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // player pair card: J  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // player pair card: Q  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // player pair card: K  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
];
function createPairsDecison() {
    /*
    Turn the _PAIRS_DECISION table into a rank based dictiopnary:
        PAIRS_DECISION: dict[CardRank, dict[CardRank, str]] = {
            CardRank.ACE: {
                CardRank.ACE: SP, CardRank.TWO: SP, ..., CardRank.KING: SP,
            }, ...
        }
    */
    var decisions = new Map();
    const ranks = Object.values(card_1.CardRank);
    ranks.forEach((playerPairRank) => {
        decisions.set(playerPairRank, new Map());
    });
    ranks.forEach((playerPairRank) => {
        var playerIndex = (0, card_1.getCardIndex)(playerPairRank);
        ranks.forEach((dealerTopCardRank) => {
            var _a;
            var dealerIndex = (0, card_1.getCardIndex)(dealerTopCardRank);
            var decision = _PAIRS_DECISION[playerIndex][dealerIndex];
            (_a = decisions.get(playerPairRank)) === null || _a === void 0 ? void 0 : _a.set(dealerTopCardRank, decision);
        });
    });
    return decisions;
}
const PAIRS_DECISION = createPairsDecison();
// Expect to use the soft total decision table for: (A,A) and (A,2),
// which is the only way to get to hard totals 2 and 3.
const _HARD_TOTAL_DECISION = [
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // hard total: 1  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // hard total: 2  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // hard total: 3  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // hard total: 4  x  dealer top card
    [NO, H, H, H, H, H, H, H, H, H, H, H, H, H],
    // hard total: 5  x  dealer top card
    [NO, H, H, H, H, H, H, H, H, H, H, H, H, H],
    // hard total: 6  x  dealer top card
    [NO, H, H, H, H, H, H, H, H, H, H, H, H, H],
    // hard total: 7  x  dealer top card
    [NO, H, H, H, H, H, H, H, H, H, H, H, H, H],
    // hard total: 8  x  dealer top card
    [NO, H, H, H, H, H, H, H, H, H, H, H, H, H],
    // hard total: 9  x  dealer top card
    [NO, H, H, Dh, Dh, Dh, Dh, H, H, H, H, H, H, H],
    // hard total: 10  x  dealer top card
    [NO, H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, H, H, H, H],
    // hard total: 11  x  dealer top card
    [NO, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh],
    // hard total: 12  x  dealer top card
    [NO, H, H, H, S, S, S, H, H, H, H, H, H, H],
    // hard total: 13  x  dealer top card
    [NO, H, S, S, S, S, S, H, H, H, H, H, H, H],
    // hard total: 14  x  dealer top card
    [NO, H, S, S, S, S, S, H, H, H, H, H, H, H],
    // hard total: 15  x  dealer top card
    [NO, Uh, S, S, S, S, S, H, H, H, Uh, Uh, Uh, Uh],
    // hard total: 16  x  dealer top card
    [NO, Uh, S, S, S, S, S, H, H, Uh, Uh, Uh, Uh, Uh],
    // hard total: 17  x  dealer top card
    [NO, Us, S, S, S, S, S, S, S, S, S, S, S, S],
    // hard total: 18  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // hard total: 19  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // hard total: 20  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // hard total: 21  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
];
function createHardTotalDecision() {
    /*
    Turn _HARD_TOTAL_DECISION table into
        HARD_TOTAL_DECISION = [
            ...
            {CardRank.ACE: H, CardRank.TWO: H, ..., CardRank.KING: H}, # [5]
            ...
        ]
    */
    var decisions = [];
    for (let hardCount = 0; hardCount <= 21; hardCount++) {
        decisions.push(new Map());
    }
    const ranks = Object.values(card_1.CardRank);
    for (let hardCount = 0; hardCount <= 21; hardCount++) {
        ranks.forEach((dealerTopCardRank) => {
            var dealerIndex = (0, card_1.getCardIndex)(dealerTopCardRank);
            var decision = _HARD_TOTAL_DECISION[hardCount][dealerIndex];
            decisions[hardCount].set(dealerTopCardRank, decision);
        });
    }
    return decisions;
}
const HARD_TOTAL_DECISION = createHardTotalDecision();
const _SOFT_TOTAL_DECISION = [
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 1  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 2  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 3  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 4  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 5  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 6  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 7  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 8  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 9  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 10  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 11  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],
    // soft total: 12 (A, A)  x  dealer top card
    [NO, H, H, H, H, H, Dh, H, H, H, H, H, H, H],
    // soft total: 13 (A, 2)  x  dealer top card
    [NO, H, H, H, H, Dh, Dh, H, H, H, H, H, H, H],
    // soft total: 14 (A, 3)  x  dealer top card
    [NO, H, H, H, H, Dh, Dh, H, H, H, H, H, H, H],
    // soft total: 15 (A, 4)  x  dealer top card
    [NO, H, H, H, Dh, Dh, Dh, H, H, H, H, H, H, H],
    // soft total: 16 (A, 5)  x  dealer top card
    [NO, H, H, H, Dh, Dh, Dh, H, H, H, H, H, H, H],
    // soft total: 17 (A, 6)  x  dealer top card
    [NO, H, H, Dh, Dh, Dh, Dh, H, H, H, H, H, H, H],
    // soft total: 18 (A, 7)  x  dealer top card
    [NO, H, Ds, Ds, Ds, Ds, Ds, S, S, H, H, H, H, H],
    // soft total: 19 (A, 8)  x  dealer top card
    [NO, S, S, S, S, S, Ds, S, S, S, S, S, S, S],
    // soft total: 20 (A, 9)  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    // soft total: 21 (A, 10)  x  dealer top card
    [NO, S, S, S, S, S, S, S, S, S, S, S, S, S],
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
];
function createSoftTotalDecision() {
    /*
    Turn _SOFT_TOTAL_DECISION table into
        HARD_TOTAL_DECISION = [
            ...
            {CardRank.ACE: H, CardRank.TWO: H, ..., CardRank.KING: H}, # [5]
            ...
        ]
    */
    var decisions = [];
    for (let hardCount = 0; hardCount <= 21; hardCount++) {
        decisions.push(new Map());
    }
    const ranks = Object.values(card_1.CardRank);
    for (let hardCount = 0; hardCount <= 21; hardCount++) {
        ranks.forEach((dealerTopCardRank) => {
            var dealerIndex = (0, card_1.getCardIndex)(dealerTopCardRank);
            var decision = _SOFT_TOTAL_DECISION[hardCount][dealerIndex];
            decisions[hardCount].set(dealerTopCardRank, decision);
        });
    }
    return decisions;
}
const SOFT_TOTAL_DECISION = createSoftTotalDecision();
function determineBasicStrategyPlay(dealerTopCard, playerHand, handsAllowMoreSplits) {
    var _a;
    var isFirstDecision = (playerHand.numCards == 2);
    var isFirstPostSplitDecision = (isFirstDecision && playerHand.isFromSplit);
    var playerCard1 = playerHand.getCard(0);
    var playerCard2 = playerHand.getCard(1);
    var decision;
    var playerDecision;
    var gotPairs;
    if (isFirstDecision) {
        if (house_1.HouseRules.SPLIT_ON_VALUE_MATCH) {
            gotPairs = (0, card_1.getCardValue)(playerCard1.rank) == (0, card_1.getCardValue)(playerCard2.rank);
        }
        else {
            gotPairs = playerCard1.rank == playerCard2.rank;
        }
    }
    else {
        gotPairs = false;
    }
    if (gotPairs && handsAllowMoreSplits) {
        // Determine if the pairs can be split.
        // Note all of the non-split decisions that are ignored below
        // (which will not contradict the hard/soft total decision).
        var pairRank;
        if ((0, card_1.getCardValue)(playerCard1.rank) == 10) {
            pairRank = card_1.CardRank.TEN;
        }
        else {
            pairRank = playerCard1.rank;
        }
        decision = (_a = PAIRS_DECISION.get(pairRank)) === null || _a === void 0 ? void 0 : _a.get(dealerTopCard.rank);
        playerDecision = convertToPlayerDecision(decision, playerHand);
        if (playerDecision == PlayerDecision.SPLIT) {
            return PlayerDecision.SPLIT;
        }
    }
    var hardCount = playerHand.hardCount;
    var softCount = playerHand.softCount;
    var useSoftTotal = (hardCount < softCount && softCount <= 21);
    if (useSoftTotal) {
        decision = SOFT_TOTAL_DECISION[softCount].get(dealerTopCard.rank);
        playerDecision = convertToPlayerDecision(decision, playerHand);
        return playerDecision;
    }
    else {
        decision = HARD_TOTAL_DECISION[softCount].get(dealerTopCard.rank);
        playerDecision = convertToPlayerDecision(decision, playerHand);
        return playerDecision;
    }
    console.log("determineBasicStrategyPlay() ran into a little trouble in town.");
    throw new Error("determineBasicStrategyPlay() ran into a little trouble in town.");
}
