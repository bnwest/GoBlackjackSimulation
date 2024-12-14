import Rand, { PRNG } from 'rand-seed';

import { PlayerHandInterface, determineBasicStrategyPlay } from "./basic";
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
    get numMasterHands() {
        return this.hands.length;
    }
    addStartHand(bet: number) {
        let playerHand = new PlayerHand(bet);
        this.hands.push(playerHand);
    }
    canSplit(handIndex: number): boolean {
        if ( this.numMasterHands < PlayerMasterHand.HAND_LIMIT ) {
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

        let newHandIndex: number = this.numMasterHands;
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

        log("\n\nDEAL HANDS")

        var card: Card;
        var player: Player;
        var masterHand: PlayerMasterHand;

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
    }
}

var blackjack: BlackJack = new BlackJack();

for ( let i = 0; i < 1; i++ ) {
    blackjack.playGame();
}
