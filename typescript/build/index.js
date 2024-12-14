"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const rand_seed_1 = __importDefault(require("rand-seed"));
const card_1 = require("./card");
const house_1 = require("./house");
const seededRand = new rand_seed_1.default('42');
// log(`rand-seed: ${seededRand.next()}`) returns 0 <= float < 1
function log(msg) {
    console.log(msg);
}
const UNSHUFFLED_DECK = [
    // HEARTS
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.ACE),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.TWO),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.THREE),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.FOUR),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.FIVE),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.SIX),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.SEVEN),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.EIGHT),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.NINE),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.TEN),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.JACK),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.QUEEN),
    new card_1.Card(card_1.CardSuite.HEARTS, card_1.CardRank.KING),
    // DIAMONDS
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.ACE),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.TWO),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.THREE),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.FOUR),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.FIVE),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.SIX),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.SEVEN),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.EIGHT),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.NINE),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.TEN),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.JACK),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.QUEEN),
    new card_1.Card(card_1.CardSuite.DIAMONDS, card_1.CardRank.KING),
    // SPADES
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.ACE),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.TWO),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.THREE),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.FOUR),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.FIVE),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.SIX),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.SEVEN),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.EIGHT),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.NINE),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.TEN),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.JACK),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.QUEEN),
    new card_1.Card(card_1.CardSuite.SPADES, card_1.CardRank.KING),
    // CLUBS
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.ACE),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.TWO),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.THREE),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.FOUR),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.FIVE),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.SIX),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.SEVEN),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.EIGHT),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.NINE),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.TEN),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.JACK),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.QUEEN),
    new card_1.Card(card_1.CardSuite.CLUBS, card_1.CardRank.KING),
];
class Shoe {
    constructor(numDecks) {
        this.cards = [];
        for (let i = 0; i < numDecks; i++) {
            this.cards = this.cards.concat(UNSHUFFLED_DECK);
        }
        this.shuffle();
        this.top = 0; // redundant, but the TS compiler need to be appeas3d
    }
    shuffle() {
        // need third party support for a seeded random number generator
        for (let i = 0; i < this.cards.length; i++) {
            const j = Math.floor(seededRand.next() * this.cards.length);
            if (i != j) {
                [this.cards[i], this.cards[j]] = [this.cards[j], this.cards[i]];
            }
        }
        this.top = 0;
    }
    display() {
        for (let i = 0; i < this.cards.length; i++) {
            const card = this.cards[i];
            log(`(${i + 1}): ${card.rank}${card.suite}`);
        }
    }
    getCard() {
        var card = this.cards[this.top];
        this.top++;
        return card;
    }
}
var HandOutcome;
(function (HandOutcome) {
    HandOutcome["STAND"] = "stand";
    HandOutcome["BUST"] = "bust";
    HandOutcome["SURRENDER"] = "surrender";
    HandOutcome["DEALER_BLACKJACK"] = "dealer-blackjack";
    HandOutcome["IN_PLAY"] = "in-play";
})(HandOutcome || (HandOutcome = {}));
class BaseHand {
    constructor() {
        this.cards = [];
        this.outcome = HandOutcome.IN_PLAY;
    }
    get acesCount() {
        var count = 0;
        for (let i = 0; i < this.numCards; i++) {
            if (this.cards[i].rank == card_1.CardRank.ACE) {
                count++;
            }
        }
        return count;
    }
    get hardCount() {
        var count = 0;
        for (let i = 0; i < this.numCards; i++) {
            count += (0, card_1.getCardValue)(this.cards[i].rank);
        }
        return count;
    }
    get softCount() {
        var count = 0;
        for (let i = 0; i < this.numCards; i++) {
            var card = this.cards[i];
            if (card.rank == card_1.CardRank.ACE) {
                count += 11;
            }
            else {
                count += (0, card_1.getCardValue)(this.cards[i].rank);
            }
        }
        // case of Ace + 5, where the count can be 6 or 16
        // case of Ace + Ace + 5, where count can be 7 or 17 or 27
        if (count > 21) {
            for (let i = 0; i < this.acesCount; i++) {
                count -= 10;
                if (count <= 21) {
                    return count;
                }
            }
        }
        // count is now the hardCount.
        return count;
    }
    get count() {
        return this.softCount;
    }
    get isBust() {
        return this.count > 21;
    }
    get numCards() {
        return this.cards.length;
    }
    get isHandOver() {
        if (this.outcome == HandOutcome.STAND) {
            return true;
        }
        else if (this.outcome == HandOutcome.BUST) {
            return true;
        }
        else if (this.outcome == HandOutcome.SURRENDER) {
            return true;
        }
        else if (this.outcome == HandOutcome.DEALER_BLACKJACK) {
            return true;
        }
        else { // this.outcome == HandOutcome.IN_PLAY
            return false;
        }
    }
    addCard(card) {
        this.cards.push(card);
    }
    getCard(cardIndex) {
        return this.cards[cardIndex];
    }
}
class PlayerHand extends BaseHand {
    constructor(bet, fromSplit = false) {
        super();
        this.bet = bet;
        this.fromSplit = fromSplit;
    }
    get isNatural() {
        if (!this.fromSplit) {
            if (this.numCards == 2) {
                if (this.count == 21) {
                    return true;
                }
            }
        }
        return false;
    }
    get canSplit() {
        if (this.numCards == 2) {
            var card1 = this.cards[0];
            var card2 = this.cards[1];
            // check card value or rank, depending on house rules
            if (house_1.HouseRules.SPLIT_ON_VALUE_MATCH) {
                if ((0, card_1.getCardValue)(card1.rank) == (0, card_1.getCardValue)(card2.rank)) {
                    return true;
                }
            }
            else {
                if (card1.rank == card2.rank) {
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
    get isNatural() {
        if (this.numCards == 2) {
            if (this.count == 21) {
                return true;
            }
        }
        return false;
    }
}
class PlayerMasterHand {
    constructor() {
        this.hands = [];
    }
    get numMasterHands() {
        return this.hands.length;
    }
    addStartHand(bet) {
        let playerHand = new PlayerHand(bet);
        this.hands.push(playerHand);
    }
    canSplit(handIndex) {
        if (this.numMasterHands < PlayerMasterHand.HAND_LIMIT) {
            // master hand allows
            let playerHand = this.hands[handIndex];
            if (playerHand.canSplit) {
                // individual hand allows
                return true;
            }
        }
        return false;
    }
    splitHand(handIndex, cardsToAdd) {
        let card1 = this.hands[handIndex].cards[0];
        let card2 = this.hands[handIndex].cards[1];
        let oldPlayerHand = this.hands[handIndex];
        // oldPlayerHand is a reference to this.hands[handIndex]
        // => the below changes this.hands[handIndex] "in place"
        oldPlayerHand.cards = [card1, cardsToAdd[0]];
        oldPlayerHand.fromSplit = true;
        oldPlayerHand.outcome = HandOutcome.IN_PLAY;
        let newPlayerHand = new PlayerHand(oldPlayerHand.bet, true);
        newPlayerHand.cards = [card2, cardsToAdd[1]];
        newPlayerHand.outcome = HandOutcome.IN_PLAY;
        let newHandIndex = this.numMasterHands;
        this.hands.push(newPlayerHand);
        return newHandIndex;
    }
}
PlayerMasterHand.HAND_LIMIT = house_1.HouseRules.SPLITS_PER_HAND + 1;
class Player {
    constructor(name) {
        this.name = name;
        this.masterHands = [];
    }
    get numMasterHands() {
        return this.masterHands.length;
    }
    setGameBets(bets) {
        this.masterHands = [];
        for (let i = 0; i < bets.length; i++) {
            let bet = bets[i];
            let playerMasterHand = new PlayerMasterHand();
            playerMasterHand.addStartHand(bet);
            this.masterHands.push(playerMasterHand);
        }
    }
}
class Dealer {
    constructor() {
        this.hand = new DealerHand();
    }
    get topCard() {
        return this.hand.cards[0];
    }
    get holeCard() {
        return this.hand.cards[1];
    }
}
class BlackJackPlayerResults {
    constructor() {
        this.handsPlayed = 0;
        this.handsWon = 0;
        this.handsLost = 0;
        this.handsPushed = 0;
        this.proceeds = 0;
    }
}
class BlackJackStats {
    constructor() {
        this.doubleDownCount = 0;
        this.surrenderCount = 0;
        this.splitCount = 0;
        this.acesSplit = 0;
    }
}
class BlackJack {
    constructor() {
        this.shoe = new Shoe(house_1.HouseRules.DECKS_IN_SHOE);
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
    getCardFromShoe() {
        var card = this.shoe.getCard();
        return card;
    }
    setPlayersForGame(players) {
        this.players = players;
        for (let i = 0; i < this.numPlayers; i++) {
            let player = this.players[i];
            if (!this.playerResults.has(player.name)) {
                this.playerResults.set(player.name, new BlackJackPlayerResults());
            }
        }
    }
    playGame() {
        if (this.shoe.top > house_1.HouseRules.FORCE_RESHUFFLE) {
            this.reshuffleShoe();
        }
        // setting up the dealer and the player(s) could be done
        // by the caller and pass here via parameters.  ¯\_(ツ)_/¯
        var dealer = new Dealer();
        var player1 = new Player("Jack");
        var player2 = new Player("Jill");
        this.setPlayersForGame([player1, player2]);
        var initialBet = 2;
        player1.setGameBets([initialBet]);
        player2.setGameBets([initialBet, initialBet]);
        //
        // DEAL HANDS
        //
        log("\n\nDEAL HANDS");
        var card;
        var player;
        var masterHand;
        // deal two cards to players and dealer, 
        // all face up except for dealer's second card.
        for (let i = 0; i < 2; i++) {
            for (let j = 0; j < this.numPlayers; j++) {
                player = this.players[j];
                for (let k = 0; k < player.numMasterHands; k++) {
                    card = this.getCardFromShoe();
                    masterHand = player.masterHands[k];
                    let firstHand = masterHand.hands[0];
                    firstHand.addCard(card);
                }
            }
            card = this.getCardFromShoe();
            dealer.hand.addCard(card);
        }
        var dealerTopCard = dealer.topCard;
        log(`dealer top card: ${dealerTopCard.str()}`);
    }
}
var blackjack = new BlackJack();
for (let i = 0; i < 1; i++) {
    blackjack.playGame();
}
