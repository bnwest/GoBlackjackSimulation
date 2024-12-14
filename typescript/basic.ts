import { CardSuite, CardRank, Card, getCardValue, getCardIndex } from "./card";
import { HouseRules, canDoubleDown } from "./house";

//
// Basic Strategy:
//     https://en.wikipedia.org/wiki/Blackjack#Basic_strategy
//     All Hail The Wikipedia!
//

// to make decision tables readable:
const S   = "stand";
const H   = "hit";
const Dh  = "double-down-if-allowed-or-hit";
const Ds  = "double-down-if-allowed-or-stand";
const SP  = "split";
// U => Surrender, in a world of too many S words
const Uh  = "surrender-if-allowed-or-hit";
const Us  = "surrender-if-allowed-or-stand";
const Usp = "surrender-if-allowed-or-split";
const NO  = "no-decision";

function isDecisionValid(decision: string): boolean {
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

export enum PlayerDecision {
    STAND     = "stand",
    HIT       = "hit",
    DOUBLE    = "double-down",
    SPLIT     = "split",
    SURRENDER = "surrender",
}

export interface PlayerHandInterface {
    get hardCount(): number;
    get softCount(): number;
    get numCards(): number;
    get isFromSplit(): boolean;
    getCard(cardIndex:number): Card;
}

function convertToPlayerDecision(
    decision: string, 
    playerHand: PlayerHandInterface
): PlayerDecision {
	// Decision sometimes return Xy, which translates to do X if allowed else do y.
	// Determine the X or the y here.

	const isFirstDecision: boolean = ( playerHand.numCards == 2 );
	const isFirstPostSplitDecision: boolean = ( isFirstDecision && playerHand.isFromSplit );

    var playerDecision: PlayerDecision = PlayerDecision.STAND;

    const hardCount: number = playerHand.hardCount;
    const softCount: number = playerHand.softCount;

    if ( decision == S ) {
        return PlayerDecision.STAND;

    } else if ( decision == H ) {
        return PlayerDecision.HIT;

    } else if ( decision == Dh || decision == Ds ) {
		// may be only allow to down on hand totals [9, 10, 11] or some such
		// basic stratgey wants to double down on
		//     hand hard totals [9, 10, 11]
		//     hand soft totals [12, 13,14, 15, 16, 17, 18, 19]
		var nondoubleDownDecision: PlayerDecision;
        if ( decision == Dh ) {
            nondoubleDownDecision = PlayerDecision.HIT;
        } else {
            nondoubleDownDecision = PlayerDecision.STAND;
        }

        var doubleDownAllowed: boolean;
        if ( isFirstDecision ) {
            if ( isFirstPostSplitDecision ) {
                if ( HouseRules.DOUBLE_DOWN_AFTER_SPLIT ) {
                    doubleDownAllowed = true;
                } else {
                    doubleDownAllowed = false;
                }
            } else {
                doubleDownAllowed = true;
            }
        } else {
            doubleDownAllowed = false;
        }

        if ( doubleDownAllowed ) {
            var doubleDown: boolean;
            if ( canDoubleDown(hardCount) ) {
                doubleDown = true;
            } else if ( canDoubleDown(softCount) ) {
                doubleDown = true;
            } else {
                doubleDown = false;
            }

            playerDecision = doubleDown ? PlayerDecision.DOUBLE : nondoubleDownDecision;

        } else {
            playerDecision = nondoubleDownDecision;
        }

    } else if ( decision == SP ) {
        playerDecision = PlayerDecision.SPLIT;

    } else if ( decision == Uh || decision == Us || decision == Usp ) {
		// surrent decision must be allowed in the House Rules and
		// must be a first decision (before splitting)
		var nonsurrenderDecision: PlayerDecision;
        switch ( decision ) {
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
        var surrenderAllowed: boolean = ( isFirstDecision && !isFirstPostSplitDecision && HouseRules.SURRENDER_ALLOWED );
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
    [NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],
    // player pair card: 3  x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],
    // player pair card: 4  x  dealer top card
    [NO,  H,  H,  H, H, SP, SP,  H,  H,  H,  H,   H,  H,  H],
    // player pair card: 5  x  dealer top card
    [NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H],
    // player pair card: 6 x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H,  H],
    // player pair card: 7  x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],
    // player pair card: 8  x  dealer top card
    [NO, Usp, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],
    // player pair card: 9  x  dealer top card
    [NO,  S, SP, SP, SP, SP, SP,  S, SP, SP,  S,  S,  S,  S],
    // player pair card: 10  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // player pair card: J  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // player pair card: Q  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // player pair card: K  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
]

function createPairsDecison(): Map<CardRank, Map<CardRank, string>> {
    /*
    Turn the _PAIRS_DECISION table into a rank based dictiopnary:
        PAIRS_DECISION: dict[CardRank, dict[CardRank, str]] = {
            CardRank.ACE: {
                CardRank.ACE: SP, CardRank.TWO: SP, ..., CardRank.KING: SP,
            }, ...
        }
    */
    var decisions: Map<CardRank, Map<CardRank, string>> = new Map();

    const ranks: CardRank[] = Object.values(CardRank);

    ranks.forEach((playerPairRank) => {
        decisions.set(playerPairRank, new Map())
    });

    ranks.forEach((playerPairRank) => {
        var playerIndex: number = getCardIndex(playerPairRank);
        ranks.forEach((dealerTopCardRank) => {
            var dealerIndex: number = getCardIndex(dealerTopCardRank);
            var decision: string = _PAIRS_DECISION[playerIndex][dealerIndex];
            decisions.get(playerPairRank)?.set(dealerTopCardRank, decision);
        });
    });

    return decisions;
}

const PAIRS_DECISION:  Map<CardRank, Map<CardRank, string>> = createPairsDecison();

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
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 5  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 6  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 7  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 8  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 9  x  dealer top card
    [NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 10  x  dealer top card
    [NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H],
    // hard total: 11  x  dealer top card
    [NO, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh],
    // hard total: 12  x  dealer top card
    [NO,  H,  H,  H,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 13  x  dealer top card
    [NO,  H,  S,  S,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 14  x  dealer top card
    [NO,  H,  S,  S,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H],
    // hard total: 15  x  dealer top card
    [NO, Uh,  S,  S,  S,  S,  S,  H,  H,  H, Uh, Uh, Uh, Uh],
    // hard total: 16  x  dealer top card
    [NO, Uh,  S,  S,  S,  S,  S,  H,  H, Uh, Uh, Uh, Uh, Uh],
    // hard total: 17  x  dealer top card
    [NO, Us,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // hard total: 18  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // hard total: 19  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // hard total: 20  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // hard total: 21  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
];

function createHardTotalDecision(): Map<CardRank, string>[] {
    /*
    Turn _HARD_TOTAL_DECISION table into
        HARD_TOTAL_DECISION = [
            ...
            {CardRank.ACE: H, CardRank.TWO: H, ..., CardRank.KING: H}, # [5]
            ...
        ]
    */
    var decisions: Map<CardRank, string>[] = [];
    for ( let hardCount = 0; hardCount <= 21; hardCount++ ) {
        decisions.push(new Map());
    }

    const ranks: CardRank[] = Object.values(CardRank);

    for ( let hardCount = 0; hardCount <= 21; hardCount++ ) {
        ranks.forEach((dealerTopCardRank) => {
            var dealerIndex: number = getCardIndex(dealerTopCardRank);
            var decision: string = _HARD_TOTAL_DECISION[hardCount][dealerIndex];
            decisions[hardCount].set(dealerTopCardRank, decision);
        })
    }

    return decisions;
}

const HARD_TOTAL_DECISION: Map<CardRank, string>[] = createHardTotalDecision();

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
    [NO,  H,  H,  H,  H,  H, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft total: 13 (A, 2)  x  dealer top card
    [NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft total: 14 (A, 3)  x  dealer top card
    [NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft total: 15 (A, 4)  x  dealer top card
    [NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft total: 16 (A, 5)  x  dealer top card
    [NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft total: 17 (A, 6)  x  dealer top card
    [NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],
    // soft total: 18 (A, 7)  x  dealer top card
    [NO,  H, Ds, Ds, Ds, Ds, Ds,  S,  S,  H,  H,  H,  H,  H],
    // soft total: 19 (A, 8)  x  dealer top card
    [NO,  S,  S,  S,  S,  S, Ds,  S,  S,  S,  S,  S,  S,  S],
    // soft total: 20 (A, 9)  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    // soft total: 21 (A, 10)  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],
    //0   A   2   3   4   5   6   7   8   9  10   J   Q   K
];

function createSoftTotalDecision(): Map<CardRank, string>[] {
    /*
    Turn _SOFT_TOTAL_DECISION table into
        HARD_TOTAL_DECISION = [
            ...
            {CardRank.ACE: H, CardRank.TWO: H, ..., CardRank.KING: H}, # [5]
            ...
        ]
    */
    var decisions: Map<CardRank, string>[] = [];
    for ( let hardCount = 0; hardCount <= 21; hardCount++ ) {
        decisions.push(new Map());
    }

    const ranks: CardRank[] = Object.values(CardRank);

    for ( let hardCount = 0; hardCount <= 21; hardCount++ ) {
        ranks.forEach((dealerTopCardRank) => {
            var dealerIndex: number = getCardIndex(dealerTopCardRank);
            var decision: string = _SOFT_TOTAL_DECISION[hardCount][dealerIndex];
            decisions[hardCount].set(dealerTopCardRank, decision);
        })
    }

    return decisions;
}

const SOFT_TOTAL_DECISION: Map<CardRank, string>[] = createSoftTotalDecision();

export function determineBasicStrategyPlay(
    dealerTopCard: Card,
    playerHand: PlayerHandInterface,
    handsAllowMoreSplits: boolean,
): PlayerDecision {
    var isFirstDecision: boolean = ( playerHand.numCards == 2 );
    var isFirstPostSplitDecision: boolean = ( isFirstDecision && playerHand.isFromSplit );

    var playerCard1: Card = playerHand.getCard(0);
    var playerCard2: Card = playerHand.getCard(1);

    var decision: string;
    var playerDecision: PlayerDecision;

    var gotPairs: boolean;
    if ( isFirstDecision ) {
        if ( HouseRules.SPLIT_ON_VALUE_MATCH ) {
            gotPairs = getCardValue(playerCard1.rank) == getCardValue(playerCard2.rank);
        } else {
            gotPairs = playerCard1.rank == playerCard2.rank;
        }
    } else {
        gotPairs = false;
    }

    if ( gotPairs && handsAllowMoreSplits ) {
		// Determine if the pairs can be split.
		// Note all of the non-split decisions that are ignored below
		// will not contradict the hard/soft total decision.
		var pairRank: CardRank;
        if ( getCardValue(playerCard1.rank) == 10 ) {
            pairRank = CardRank.TEN;
        } else {
            pairRank = playerCard1.rank;
        }

        decision = PAIRS_DECISION.get(pairRank)?.get(dealerTopCard.rank) as string;
        playerDecision = convertToPlayerDecision(decision, playerHand);
        if ( playerDecision == PlayerDecision.SPLIT ) {
            return PlayerDecision.SPLIT;
        }
    }

    var hardCount: number = playerHand.hardCount;
    var softCount: number = playerHand.softCount;
    var useSoftTotal: boolean = ( hardCount < softCount && softCount <= 21 );
    if ( useSoftTotal ) {
        decision = SOFT_TOTAL_DECISION[softCount].get(dealerTopCard.rank) as string;
        playerDecision = convertToPlayerDecision(decision, playerHand);
        return playerDecision;

    } else {
        decision = HARD_TOTAL_DECISION[softCount].get(dealerTopCard.rank) as string;
        playerDecision = convertToPlayerDecision(decision, playerHand);
        return playerDecision;
    }

    console.log("determineBasicStrategyPlay() ran into a little trouble in town.");
    throw new Error("determineBasicStrategyPlay() ran into a little trouble in town.");
}
