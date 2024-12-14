// Lessons Learned:
// 1. TypeScript files are trans-compiled into JavaScript
// 1.1 and run locally in my case in the Node.js environment
// 1.2 Runtime errors map to the JS code :(
// 1.3 Trans-complie JavaScript a real "piece of work", JavaScript library worthy
// 2. string1 in [string1, string2] did not work, FTW.

import Rand, { PRNG } from 'rand-seed';

import { PlayerHandInterface, determineBasicStrategyPlay, PlayerDecision } from "./basic";
import { CardSuite, CardRank, Card, getCardValue } from "./card";
import { HouseRules } from "./house";

const seededRand = new Rand('42');
// log(`rand-seed: ${seededRand.next()}`) returns 0 <= float < 1

function log(msg: string) {
    console.log(msg);
}

const UNSHUFFLED_DECK: Card[] = [
    // HEARTS
    new Card(CardSuite.HEARTS, CardRank.ACE),
    new Card(CardSuite.HEARTS, CardRank.TWO),
    new Card(CardSuite.HEARTS, CardRank.THREE),
    new Card(CardSuite.HEARTS, CardRank.FOUR),
    new Card(CardSuite.HEARTS, CardRank.FIVE),
    new Card(CardSuite.HEARTS, CardRank.SIX),
    new Card(CardSuite.HEARTS, CardRank.SEVEN),
    new Card(CardSuite.HEARTS, CardRank.EIGHT),
    new Card(CardSuite.HEARTS, CardRank.NINE),
    new Card(CardSuite.HEARTS, CardRank.TEN),
    new Card(CardSuite.HEARTS, CardRank.JACK),
    new Card(CardSuite.HEARTS, CardRank.QUEEN),
    new Card(CardSuite.HEARTS, CardRank.KING),
    // DIAMONDS
    new Card(CardSuite.DIAMONDS, CardRank.ACE),
    new Card(CardSuite.DIAMONDS, CardRank.TWO),
    new Card(CardSuite.DIAMONDS, CardRank.THREE),
    new Card(CardSuite.DIAMONDS, CardRank.FOUR),
    new Card(CardSuite.DIAMONDS, CardRank.FIVE),
    new Card(CardSuite.DIAMONDS, CardRank.SIX),
    new Card(CardSuite.DIAMONDS, CardRank.SEVEN),
    new Card(CardSuite.DIAMONDS, CardRank.EIGHT),
    new Card(CardSuite.DIAMONDS, CardRank.NINE),
    new Card(CardSuite.DIAMONDS, CardRank.TEN),
    new Card(CardSuite.DIAMONDS, CardRank.JACK),
    new Card(CardSuite.DIAMONDS, CardRank.QUEEN),
    new Card(CardSuite.DIAMONDS, CardRank.KING),
    // SPADES
    new Card(CardSuite.SPADES, CardRank.ACE),
    new Card(CardSuite.SPADES, CardRank.TWO),
    new Card(CardSuite.SPADES, CardRank.THREE),
    new Card(CardSuite.SPADES, CardRank.FOUR),
    new Card(CardSuite.SPADES, CardRank.FIVE),
    new Card(CardSuite.SPADES, CardRank.SIX),
    new Card(CardSuite.SPADES, CardRank.SEVEN),
    new Card(CardSuite.SPADES, CardRank.EIGHT),
    new Card(CardSuite.SPADES, CardRank.NINE),
    new Card(CardSuite.SPADES, CardRank.TEN),
    new Card(CardSuite.SPADES, CardRank.JACK),
    new Card(CardSuite.SPADES, CardRank.QUEEN),
    new Card(CardSuite.SPADES, CardRank.KING),
    // CLUBS
    new Card(CardSuite.CLUBS, CardRank.ACE),
    new Card(CardSuite.CLUBS, CardRank.TWO),
    new Card(CardSuite.CLUBS, CardRank.THREE),
    new Card(CardSuite.CLUBS, CardRank.FOUR),
    new Card(CardSuite.CLUBS, CardRank.FIVE),
    new Card(CardSuite.CLUBS, CardRank.SIX),
    new Card(CardSuite.CLUBS, CardRank.SEVEN),
    new Card(CardSuite.CLUBS, CardRank.EIGHT),
    new Card(CardSuite.CLUBS, CardRank.NINE),
    new Card(CardSuite.CLUBS, CardRank.TEN),
    new Card(CardSuite.CLUBS, CardRank.JACK),
    new Card(CardSuite.CLUBS, CardRank.QUEEN),
    new Card(CardSuite.CLUBS, CardRank.KING),
];

class Shoe {
    cards: Card[];
    top: number;
    constructor (numDecks: number) {
        this.cards = []
        for (let i = 0; i < numDecks; i++) {
            this.cards = this.cards.concat(UNSHUFFLED_DECK);
        }
        this.shuffle();
        this.top = 0;  // redundant, but the TS compiler need to be appeas3d
    }
    shuffle() {
        // need third party support for a seeded random number generator
        for (let i = 0; i < this.cards.length; i++) {
            const j = Math.floor(seededRand.next() * this.cards.length);
            if (i != j) {
                [this.cards[i], this.cards[j]] = [this.cards[j], this.cards[i]]
            }
        }
        this.top = 0;
    }
    display() {
        for (let i = 0; i < this.cards.length; i++) {
            const card: Card = this.cards[i];
            log(`(${i+1}): ${card.rank}${card.suite}`)
        }
    }
    getCard(): Card {
        var card: Card = this.cards[this.top];
        this.top++;
        return card;
    }
}

enum HandOutcome {
    STAND            = "stand",
    BUST             = "bust",
    SURRENDER        = "surrender",
    DEALER_BLACKJACK = "dealer-blackjack",
    IN_PLAY          = "in-play",
}

class BaseHand {
    cards: Card[];
    outcome: HandOutcome;
    constructor() {
        this.cards = [];
        this.outcome = HandOutcome.IN_PLAY;
    }
    get acesCount(): number {
        var count: number = 0;
        for ( let i = 0; i< this.numCards; i++ ) {
            if ( this.cards[i].rank == CardRank.ACE ) {
                count++;
            }
        }
        return count;
    }
    get hardCount(): number {
        var count: number = 0;
        for ( let i = 0; i< this.numCards; i++ ) {
            count += getCardValue(this.cards[i].rank);
        }
        return count;
    }
    get softCount(): number {
        var count: number = 0;
        for (let i = 0; i< this.numCards; i++) {
            var card: Card = this.cards[i];
            if ( card.rank == CardRank.ACE ) {
                count += 11;
            } else {
                count += getCardValue(this.cards[i].rank);
            }
        }

        // case of Ace + 5, where the count can be 6 or 16
        // case of Ace + Ace + 5, where count can be 7 or 17 or 27

        if ( count > 21 ) {
            for ( let i = 0; i < this.acesCount; i++ ){
                count -= 10;
                if ( count <= 21 ) {
                    return count;
                }
            }
        }

        // count is now the hardCount.

        return count;
    }
    get count(): number {
        return this.softCount;
    }
    get isBust(): boolean {
        return this.count > 21;
    }
    get numCards(): number {
        return this.cards.length;
    }
    get isHandOver(): boolean {
        if ( this.outcome == HandOutcome.STAND ) {
            return true;
        } else if ( this.outcome == HandOutcome.BUST ) {
            return true;
        } else if ( this.outcome == HandOutcome.SURRENDER ) {
            return true;
        } else if ( this.outcome == HandOutcome.DEALER_BLACKJACK ) {
            return true;
        } else { // this.outcome == HandOutcome.IN_PLAY
            return false;
        }
    }
    addCard(card: Card) {
        this.cards.push(card);
    }
    getCard(cardIndex: number): Card {
        return this.cards[cardIndex];
    }
}

class PlayerHand extends BaseHand {
    // cards: Card[];
    // outcome: HandOutcome;
    fromSplit: boolean;
    bet: number;
    constructor(bet: number, fromSplit: boolean = false) {
        super();
        this.bet = bet;
        this.fromSplit = fromSplit;
    }
    get isNatural(): boolean {
        if ( !this.fromSplit ) {
            if ( this.numCards == 2 ) {
                if ( this.count == 21 ) {
                    return true;
                }
            }
        }
        return false;
    }
    get canSplit(): boolean {
        if ( this.numCards == 2 ) {
            var card1: Card = this.cards[0];
            var card2: Card = this.cards[1];
            // check card value or rank, depending on house rules
            if ( HouseRules.SPLIT_ON_VALUE_MATCH ){
                if ( getCardValue(card1.rank) == getCardValue(card2.rank) ) {
                    return true;
                }
            } else {
                if ( card1.rank == card2.rank ) {
                    return true;
                }
            }
        }
        return false;
    }
    get isFromSplit(): boolean {
        return this.fromSplit;
    }
}

class DealerHand extends BaseHand {
    // cards: Card[];
    // outcome: HandOutcome;
    constructor() {
        super();
    }
    get isNatural(): boolean {
        if ( this.numCards == 2 ) {
            if ( this.count == 21 ) {
                return true;
            }
        }
        return false;
    }
}

class PlayerMasterHand {
    static HAND_LIMIT: number = HouseRules.SPLITS_PER_HAND + 1;
    hands: PlayerHand[];
    constructor() {
        this.hands = [];
    }
    get numHands() {
        return this.hands.length;
    }
    addStartHand(bet: number) {
        let playerHand = new PlayerHand(bet);
        this.hands.push(playerHand);
    }
    canSplit(handIndex: number): boolean {
        if ( this.numHands < PlayerMasterHand.HAND_LIMIT ) {
            // master hand allows
            let playerHand = this.hands[handIndex];
            if ( playerHand.canSplit ) {
                // individual hand allows
                return true;
            }
        }
        return false;
    }
    splitHand(handIndex: number, cardsToAdd: Card[]) {
        let card1 = this.hands[handIndex].cards[0];
        let card2 = this.hands[handIndex].cards[1];

        let oldPlayerHand: PlayerHand = this.hands[handIndex];
        // oldPlayerHand is a reference to this.hands[handIndex]
        // => the below changes this.hands[handIndex] "in place"
        oldPlayerHand.cards = [card1, cardsToAdd[0]];
        oldPlayerHand.fromSplit = true;
        oldPlayerHand.outcome = HandOutcome.IN_PLAY;

        let newPlayerHand: PlayerHand = new PlayerHand(oldPlayerHand.bet, true);
        newPlayerHand.cards = [card2, cardsToAdd[1]];
        newPlayerHand.outcome = HandOutcome.IN_PLAY;

        let newHandIndex: number = this.numHands;
        this.hands.push(newPlayerHand);

        return newHandIndex;
    }
}

class Player {
    masterHands: PlayerMasterHand[];
    name: string;
    constructor(name:string) {
        this.name = name;
        this.masterHands = [];
    }
    get numMasterHands(): number {
        return this.masterHands.length;
    }
    setGameBets(bets: number[]) {
        this.masterHands = [];
        for ( let i = 0; i < bets.length; i++ ) {
            let bet: number = bets[i];
            let playerMasterHand: PlayerMasterHand = new PlayerMasterHand();
            playerMasterHand.addStartHand(bet);
            this.masterHands.push(playerMasterHand);
        }
    }
}

class Dealer {
    hand: DealerHand;
    constructor() {
        this.hand = new DealerHand();
    }
    get topCard(): Card {
        return this.hand.cards[0];
    }
    get holeCard(): Card {
        return this.hand.cards[1];
    }
}


class BlackJackPlayerResults {
    handsPlayed: number;
    handsWon: number;
    handsLost: number;
    handsPushed: number;
    proceeds: number;
    constructor() {
        this.handsPlayed = 0;
        this.handsWon = 0;
        this.handsLost = 0;
        this.handsPushed = 0;
        this.proceeds = 0;
    }
}

class BlackJackStats {
    doubleDownCount: number;
    surrenderCount: number;
    splitCount: number;
    acesSplit: number;
    constructor() {
        this.doubleDownCount = 0;
        this.surrenderCount = 0;
        this.splitCount = 0;
        this.acesSplit = 0;
    }
}

class BlackJack {
    shoe: Shoe;
    players: Player[];
    playerResults: Map<string, BlackJackPlayerResults>;
    stats: BlackJackStats;
    constructor() {
        this.shoe = new Shoe(HouseRules.DECKS_IN_SHOE);
        this.players = [];
        this.playerResults = new Map();
        this.stats = new BlackJackStats();
    }
    get numPlayers() {
        return this.players.length;
    }
    reshuffleShoe() {
        this.shoe.shuffle();
    }
    getCardFromShoe(): Card {
        var card: Card = this.shoe.getCard();
        return card;
    }
    setPlayersForGame(players: Player[]) {
        this.players = players;
        for ( let i = 0; i < this.numPlayers; i++ ) {
            let player: Player = this.players[i];
            if ( !this.playerResults.has(player.name) ) {
                this.playerResults.set(player.name, new BlackJackPlayerResults());
            }
        }
    }
    playGame() {
        if ( this.shoe.top > HouseRules.FORCE_RESHUFFLE ) {
            this.reshuffleShoe();
        }
        // setting up the dealer and the player(s) could be done
        // by the caller and pass here via parameters.  ¯\_(ツ)_/¯

        var dealer: Dealer = new Dealer();

        var player1: Player = new Player("Jack");
        var player2: Player = new Player("Jill");

        this.setPlayersForGame([player1, player2]);

        var initialBet: number = 2;
        player1.setGameBets([initialBet]);
        player2.setGameBets([initialBet, initialBet]);

        //
        // DEAL HANDS
        //

        log("\n\nDEAL HANDS");

        var card: Card;
        var player: Player;
        var masterHand: PlayerMasterHand;
        var hand: PlayerHand;
        var playerDecision: PlayerDecision;

        // deal two cards to players and dealer, 
        // all face up except for dealer's second card.
        for ( let i = 0; i < 2; i++ ) {
            for ( let j = 0; j < this.numPlayers; j++ ) {
                player = this.players[j];
                for ( let k = 0; k < player.numMasterHands; k++ ) {
                    card = this.getCardFromShoe();
                    masterHand = player.masterHands[k];
                    let firstHand: PlayerHand = masterHand.hands[0];
                    firstHand.addCard(card);
                }
            }
            card = this.getCardFromShoe();
            dealer.hand.addCard(card);
        }

        var dealerTopCard: Card = dealer.topCard;
        log(`dealer top card: ${dealerTopCard.str()}`);

        var dealerHoleCard: Card = dealer.holeCard;

        //
        // PLAY HANDS
        //

        log("PLAY HANDS");

        if ( dealer.hand.isNatural ) {
            // a real simulation would have to take care of Insurance, which is a sucker's bet,
            // so we just assume that no player will ask for insurance.
            // two cases:
            //     1. player has a natural and their bet is pushed
            //     2. player loses

            dealer.hand.outcome = HandOutcome.DEALER_BLACKJACK;

            for ( let i = 0; i < this.numPlayers; i++ ) {
                player = this.players[i];
                for ( let j = 0; j < player.numMasterHands; j++ ) {
                    masterHand = player.masterHands[j];
                    for ( let k = 0; k < masterHand.numHands; k++ ) {
                        // really should only be one hand in the master hand at this point
                        hand = masterHand.hands[k];
                        // standing will do the right thing in the settlement logic below
                        hand.outcome = HandOutcome.STAND;
                    }
                }
            }

        } else {
            // dealer does not have a natural
            for ( let i = 0; i < this.numPlayers; i++ ) {
                player = this.players[i];
                log(`player ${i+1} - ${player.name}`);
                for ( let j = 0; j < player.numMasterHands; j++ ) {
                    masterHand = player.masterHands[j];
                    for ( let k = 0; k < masterHand.numHands; k++ ) {
                        hand = masterHand.hands[k];

                        log(`    hand ${j+1}.${k+1}`);
                        for ( let l = 0; l < hand.numCards; l++ ) {
                            card = hand.cards[l];
                            log(`        card ${l+1}: ${card.str()}`);
                        }

                        var isSplitPossible: boolean = ( masterHand.numHands < PlayerMasterHand.HAND_LIMIT );

                        // resolve the current hand ...
                        while ( true ) {
                            if ( hand.outcome == HandOutcome.STAND ) {
                                // product of a prior ace split, outcome has already been determined.
                                log(`        prior aces split; ${PlayerDecision.STAND}`);
                                break;
                            }

                            var handInterface: PlayerHandInterface = hand;
                            playerDecision = determineBasicStrategyPlay(
                                dealerTopCard, 
                                handInterface, 
                                isSplitPossible
                            );
                            log(`        basic strategy: ${playerDecision}`);

                            if ( playerDecision == PlayerDecision.STAND ) {
                                hand.outcome = HandOutcome.STAND;
                                log(`        stand total H${hand.hardCount} S${hand.softCount}`)
                                break;
                            }

                            else if ( playerDecision == PlayerDecision.SURRENDER ) {
                                hand.outcome = HandOutcome.SURRENDER;
                                hand.bet = Math.floor(hand.bet / 2);
                                break;
                            }

                            else if ( playerDecision == PlayerDecision.DOUBLE ) {
                                card = this.getCardFromShoe();
                                hand.addCard(card);
                                hand.bet *= 2;
                                log(`        hit: ${card.str()}}`);
                                hand.outcome = HandOutcome.STAND;
                                log(`        stand total H${hand.hardCount} S${hand.softCount}`);
                                break;
                            }

                            else if ( playerDecision == PlayerDecision.HIT ) {
                                card = this.getCardFromShoe();
                                hand.addCard(card);
                                let handTotal: number = hand.count;
                                log(`        hit: ${card.str()}}, H${hand.hardCount} S${hand.softCount}`);
                                if ( handTotal > 21 ) {
                                    hand.outcome = HandOutcome.BUST;
                                    log(`        bust`);
                                    break;
                                } else {
                                    hand.outcome = HandOutcome.IN_PLAY;
                                }
                            }

                            else if ( playerDecision == PlayerDecision.SPLIT ) {
                                let card1: Card = this.getCardFromShoe();
                                let card2: Card = this.getCardFromShoe();
                                let cardsToAdd: Card[] = [card1, card2];
                                let handIndex: number = k;
                                let newHandIndex: number = masterHand.splitHand(handIndex, cardsToAdd);
                                log(`        split, new hand index ${newHandIndex+1}, adding cards ${card1.str()}, ${card2.str()}`);
                                log(`        new card 2: ${card1.str()}`);

                                let splittingAces: boolean = ( card1.rank == CardRank.ACE );
                                if ( splittingAces && HouseRules.NO_MORE_CARDS_AFTER_SPLITTING_ACES ) {
                                    hand.outcome = HandOutcome.STAND;
                                    log(`        aces split: stand`);
                                    masterHand.hands[newHandIndex].outcome = HandOutcome.STAND;
                                    break;
                                }
                            }
                            else {
                                log(`FTW`);
                                log(`FTW: dealer top card: ${dealerTopCard.str()}`);
                                log(`FTW: player hand count: H${hand.hardCount} S${hand.softCount}`);
                                log(`FTW: player decision: ${playerDecision}`);
                                break;
                            }
                        }
                    }
                }
            }

            //
            // DEALER HAND
            //

            log("DEALER HAND");
            log(`dealer top card:  ${dealerTopCard.str()}`);
            log(`dealer hole card: ${dealerHoleCard.str()}`);
    
            var dealerDone: boolean = false;
            while ( !dealerDone ) {
                let hardCount: number = dealer.hand.hardCount;
                let softCount: number = dealer.hand.softCount;
                let useSoftCount: boolean = ( hardCount < softCount && softCount <= 21 );
                if ( useSoftCount && softCount <= HouseRules.DEALER_HITS_SOFT_ON ) {
                    card = this.getCardFromShoe();
                    dealer.hand.addCard(card);
                    log(`    add: ${card.str()}, total H${dealer.hand.hardCount} S${dealer.hand.softCount}`);
                }
                else if ( !useSoftCount && hardCount <= HouseRules.DEALER_HITS_HARD_ON ) {
                    card = this.getCardFromShoe();
                    dealer.hand.addCard(card);
                    log(`    add: ${card.str()}, total H${dealer.hand.hardCount} S${dealer.hand.softCount}`);
                } else {
                    dealer.hand.outcome = HandOutcome.STAND;
                    dealerDone = true;
                    log(`    stand`);
                }

                if ( dealer.hand.count > 21 ) {
                    dealer.hand.outcome = HandOutcome.BUST;
                    dealerDone = true;
                    log(`    bust`);
                }
            }
    
        }

        //
        // SETTLE HANDS
        //

        log("SETTLE HANDS");

    }
}

var blackjack: BlackJack = new BlackJack();

for ( let i = 0; i < 8; i++ ) {
    blackjack.playGame();
}
