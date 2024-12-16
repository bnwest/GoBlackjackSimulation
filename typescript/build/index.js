"use strict";
// Lessons Learned:
// 1. TypeScript files are trans-compiled into JavaScript
// 1.1 and run locally in my case in the Node.js environment
// 1.2 Runtime errors map to the JS code :(
// 1.3 Trans-complie JavaScript a real "piece of work", JavaScript library worthy
// 1.4 compile and run is a 2 step process: 1. npx tsc, 2. node build/index.js
// 2. string1 in [string1, string2] did not work, FTW.
// 3. python-like dictionaries are a pain
// 3.1 A JS idion object.field is the same as object["field"]
// 3.2 dictionary keys are dynamically created fields
// 3.3 getting a dynamic field is problematic, since the TS compiler does not know
// if that field exists
// 3.4 TS/JS has a generic Map type with a set of methods: has(), get(), set(), delete()
// which are verbose
// 3.5 Map.get() does not work with a little help since it returns "type | undefined"
// so the undefined has to be worked around via "as type" or "!" (tells the compiler
// that undefined is not possible, double pink promise)
// 4. TS/JS OO is solid.  Supports methods, inheritance and interfaces.
// 5. "===" versus "=="
// 6. gts lint-er had a pretty heavy hand when "correcting" white space "errors"
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const rand_seed_1 = __importDefault(require("rand-seed"));
const basic_1 = require("./basic");
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
        this.top = 0; // redundant, but the TS compiler needs to be appeased
    }
    shuffle() {
        // need third party support for a seeded random number generator
        for (let i = 0; i < this.cards.length; i++) {
            const j = Math.floor(seededRand.next() * this.cards.length);
            if (i !== j) {
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
        const card = this.cards[this.top];
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
        let count = 0;
        for (let i = 0; i < this.numCards; i++) {
            if (this.cards[i].rank === card_1.CardRank.ACE) {
                count++;
            }
        }
        return count;
    }
    get hardCount() {
        let count = 0;
        for (let i = 0; i < this.numCards; i++) {
            count += (0, card_1.getCardValue)(this.cards[i].rank);
        }
        return count;
    }
    get softCount() {
        let count = 0;
        for (let i = 0; i < this.numCards; i++) {
            const card = this.cards[i];
            if (card.rank === card_1.CardRank.ACE) {
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
        if (this.outcome === HandOutcome.STAND) {
            return true;
        }
        else if (this.outcome === HandOutcome.BUST) {
            return true;
        }
        else if (this.outcome === HandOutcome.SURRENDER) {
            return true;
        }
        else if (this.outcome === HandOutcome.DEALER_BLACKJACK) {
            return true;
        }
        else {
            // this.outcome === HandOutcome.IN_PLAY
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
            if (this.numCards === 2) {
                if (this.count === 21) {
                    return true;
                }
            }
        }
        return false;
    }
    get canSplit() {
        if (this.numCards === 2) {
            const card1 = this.cards[0];
            const card2 = this.cards[1];
            // check card value or rank, depending on house rules
            if (house_1.HouseRules.SPLIT_ON_VALUE_MATCH) {
                if ((0, card_1.getCardValue)(card1.rank) === (0, card_1.getCardValue)(card2.rank)) {
                    return true;
                }
            }
            else {
                if (card1.rank === card2.rank) {
                    return true;
                }
            }
        }
        return false;
    }
    get isFromSplit() {
        return this.fromSplit;
    }
}
class DealerHand extends BaseHand {
    // cards: Card[];
    // outcome: HandOutcome;
    constructor() {
        super();
    }
    get isNatural() {
        if (this.numCards === 2) {
            if (this.count === 21) {
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
    get numHands() {
        return this.hands.length;
    }
    addStartHand(bet) {
        const playerHand = new PlayerHand(bet);
        this.hands.push(playerHand);
    }
    canSplit(handIndex) {
        if (this.numHands < PlayerMasterHand.HAND_LIMIT) {
            // master hand allows
            const playerHand = this.hands[handIndex];
            if (playerHand.canSplit) {
                // individual hand allows
                return true;
            }
        }
        return false;
    }
    splitHand(handIndex, cardsToAdd) {
        const card1 = this.hands[handIndex].cards[0];
        const card2 = this.hands[handIndex].cards[1];
        const oldPlayerHand = this.hands[handIndex];
        // oldPlayerHand is a reference to this.hands[handIndex]
        // => the below changes this.hands[handIndex] "in place"
        oldPlayerHand.cards = [card1, cardsToAdd[0]];
        oldPlayerHand.fromSplit = true;
        oldPlayerHand.outcome = HandOutcome.IN_PLAY;
        const newPlayerHand = new PlayerHand(oldPlayerHand.bet, true);
        newPlayerHand.cards = [card2, cardsToAdd[1]];
        newPlayerHand.outcome = HandOutcome.IN_PLAY;
        const newHandIndex = this.numHands;
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
            const bet = bets[i];
            const playerMasterHand = new PlayerMasterHand();
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
        const card = this.shoe.getCard();
        return card;
    }
    setPlayersForGame(players) {
        this.players = players;
        for (let i = 0; i < this.numPlayers; i++) {
            const player = this.players[i];
            if (!this.playerResults.has(player.name)) {
                this.playerResults.set(player.name, new BlackJackPlayerResults());
            }
        }
    }
    addResult(player, handIndex, playerHand, initialBet, result) {
        // typescript map get method can return undefined
        // "!" asserts that the key is present, double pinky promise.
        this.playerResults.get(player.name).handsPlayed++;
        if (result > 0) {
            this.playerResults.get(player.name).handsWon++;
        }
        else if (result < 0) {
            this.playerResults.get(player.name).handsLost++;
        }
        else {
            this.playerResults.get(player.name).handsPushed++;
        }
        this.playerResults.get(player.name).proceeds += result;
        const isDoubleDown = playerHand.numCards === 3 &&
            Math.abs(initialBet) * 2 === Math.abs(result);
        if (isDoubleDown) {
            this.stats.doubleDownCount++;
        }
        if (playerHand.outcome === HandOutcome.SURRENDER) {
            this.stats.surrenderCount++;
        }
        if (handIndex > 0) {
            this.stats.splitCount++;
        }
        const splittingAces = playerHand.fromSplit && playerHand.cards[0].rank === card_1.CardRank.ACE;
        if (splittingAces) {
            this.stats.acesSplit++;
        }
    }
    playGame() {
        if (this.shoe.top > house_1.HouseRules.FORCE_RESHUFFLE) {
            this.reshuffleShoe();
        }
        // setting up the dealer and the player(s) could be done
        // by the caller and pass here via parameters.  ¯\_(ツ)_/¯
        const dealer = new Dealer();
        const player1 = new Player('Jack');
        const player2 = new Player('Jill');
        this.setPlayersForGame([player1, player2]);
        const initialBet = 2;
        player1.setGameBets([initialBet]);
        player2.setGameBets([initialBet, initialBet]);
        //
        // DEAL HANDS
        //
        log('\n\nDEAL HANDS');
        let card;
        let player;
        let masterHand;
        let hand;
        let playerDecision;
        // deal two cards to players and dealer,
        // all face up except for dealer's second card.
        for (let i = 0; i < 2; i++) {
            for (let j = 0; j < this.numPlayers; j++) {
                player = this.players[j];
                for (let k = 0; k < player.numMasterHands; k++) {
                    card = this.getCardFromShoe();
                    masterHand = player.masterHands[k];
                    const firstHand = masterHand.hands[0];
                    firstHand.addCard(card);
                }
            }
            card = this.getCardFromShoe();
            dealer.hand.addCard(card);
        }
        const dealerTopCard = dealer.topCard;
        log(`dealer top card: ${dealerTopCard.str()}`);
        const dealerHoleCard = dealer.holeCard;
        //
        // PLAY HANDS
        //
        log('PLAY HANDS');
        if (dealer.hand.isNatural) {
            // a real simulation would have to take care of Insurance, which is a sucker's bet,
            // so we just assume that no player will ask for insurance.
            // two cases:
            //     1. player has a natural and their bet is pushed
            //     2. player loses
            dealer.hand.outcome = HandOutcome.DEALER_BLACKJACK;
            for (let i = 0; i < this.numPlayers; i++) {
                player = this.players[i];
                for (let j = 0; j < player.numMasterHands; j++) {
                    masterHand = player.masterHands[j];
                    for (let k = 0; k < masterHand.numHands; k++) {
                        // really should only be one hand in the master hand at this point
                        hand = masterHand.hands[k];
                        // standing will do the right thing in the settlement logic below
                        hand.outcome = HandOutcome.STAND;
                    }
                }
            }
        }
        else {
            // dealer does not have a natural
            for (let i = 0; i < this.numPlayers; i++) {
                player = this.players[i];
                log(`player ${i + 1} - ${player.name}`);
                for (let j = 0; j < player.numMasterHands; j++) {
                    masterHand = player.masterHands[j];
                    for (let k = 0; k < masterHand.numHands; k++) {
                        hand = masterHand.hands[k];
                        log(`    hand ${j + 1}.${k + 1}`);
                        for (let l = 0; l < hand.numCards; l++) {
                            card = hand.cards[l];
                            log(`        card ${l + 1}: ${card.str()}`);
                        }
                        const isSplitPossible = masterHand.numHands < PlayerMasterHand.HAND_LIMIT;
                        // resolve the current hand ...
                        while (true === true) {
                            if (hand.outcome === HandOutcome.STAND) {
                                // product of a prior ace split, outcome has already been determined.
                                log(`        prior aces split; ${basic_1.PlayerDecision.STAND}, total H${hand.hardCount} S${hand.softCount}`);
                                break;
                            }
                            const handInterface = hand;
                            playerDecision = (0, basic_1.determineBasicStrategyPlay)(dealerTopCard, handInterface, isSplitPossible);
                            log(`        basic strategy: ${playerDecision}`);
                            if (playerDecision === basic_1.PlayerDecision.STAND) {
                                hand.outcome = HandOutcome.STAND;
                                log(`        stand total H${hand.hardCount} S${hand.softCount}`);
                                break;
                            }
                            else if (playerDecision === basic_1.PlayerDecision.SURRENDER) {
                                hand.outcome = HandOutcome.SURRENDER;
                                hand.bet = Math.floor(hand.bet / 2);
                                break;
                            }
                            else if (playerDecision === basic_1.PlayerDecision.DOUBLE) {
                                card = this.getCardFromShoe();
                                hand.addCard(card);
                                hand.bet *= 2;
                                log(`        hit: ${card.str()}}`);
                                hand.outcome = HandOutcome.STAND;
                                log(`        stand total H${hand.hardCount} S${hand.softCount}`);
                                break;
                            }
                            else if (playerDecision === basic_1.PlayerDecision.HIT) {
                                card = this.getCardFromShoe();
                                hand.addCard(card);
                                const handTotal = hand.count;
                                log(`        hit: ${card.str()}}, H${hand.hardCount} S${hand.softCount}`);
                                if (handTotal > 21) {
                                    hand.outcome = HandOutcome.BUST;
                                    log('        bust');
                                    break;
                                }
                                else {
                                    hand.outcome = HandOutcome.IN_PLAY;
                                }
                            }
                            else if (playerDecision === basic_1.PlayerDecision.SPLIT) {
                                const card1 = this.getCardFromShoe();
                                const card2 = this.getCardFromShoe();
                                const cardsToAdd = [card1, card2];
                                const handIndex = k;
                                const newHandIndex = masterHand.splitHand(handIndex, cardsToAdd);
                                log(`        split, new hand index ${newHandIndex + 1}, adding cards ${card1.str()}, ${card2.str()}`);
                                log(`        card 1: ${hand.cards[0].str()}`);
                                log(`        card 2: ${hand.cards[1].str()}`);
                                const splittingAces = card1.rank === card_1.CardRank.ACE;
                                if (splittingAces &&
                                    house_1.HouseRules.NO_MORE_CARDS_AFTER_SPLITTING_ACES) {
                                    hand.outcome = HandOutcome.STAND;
                                    log(`        aces split: stand, total H${hand.hardCount} S${hand.softCount}`);
                                    masterHand.hands[newHandIndex].outcome = HandOutcome.STAND;
                                    break;
                                }
                            }
                            else {
                                log('FTW');
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
            log('DEALER HAND');
            log(`dealer top card:  ${dealerTopCard.str()}`);
            log(`dealer hole card: ${dealerHoleCard.str()}`);
            let dealerDone = false;
            while (!dealerDone) {
                const hardCount = dealer.hand.hardCount;
                const softCount = dealer.hand.softCount;
                const useSoftCount = hardCount < softCount && softCount <= 21;
                if (useSoftCount && softCount <= house_1.HouseRules.DEALER_HITS_SOFT_ON) {
                    card = this.getCardFromShoe();
                    dealer.hand.addCard(card);
                    log(`    add: ${card.str()}, total H${dealer.hand.hardCount} S${dealer.hand.softCount}`);
                }
                else if (!useSoftCount &&
                    hardCount <= house_1.HouseRules.DEALER_HITS_HARD_ON) {
                    card = this.getCardFromShoe();
                    dealer.hand.addCard(card);
                    log(`    add: ${card.str()}, total H${dealer.hand.hardCount} S${dealer.hand.softCount}`);
                }
                else {
                    dealer.hand.outcome = HandOutcome.STAND;
                    dealerDone = true;
                    log('    stand');
                }
                if (dealer.hand.count > 21) {
                    dealer.hand.outcome = HandOutcome.BUST;
                    dealerDone = true;
                    log('    bust');
                }
            }
        }
        //
        // SETTLE HANDS
        //
        log('SETTLE HANDS');
        if (dealer.hand.outcome === HandOutcome.DEALER_BLACKJACK) {
            for (let i = 0; i < this.numPlayers; i++) {
                player = this.players[i];
                log(`player ${i + 1} - ${player.name}`);
                for (let j = 0; j < player.numMasterHands; j++) {
                    masterHand = player.masterHands[j];
                    for (let k = 0; k < masterHand.numHands; k++) {
                        hand = masterHand.hands[k];
                        if (hand.isNatural) {
                            this.addResult(player, k, hand, initialBet, 0);
                            log(`    hand ${j + 1}.${k + 1}: push both player and dealer had naturals`);
                        }
                        else {
                            this.addResult(player, k, hand, initialBet, -hand.bet);
                            log(`    hand ${j + 1}.${k + 1}: dealer natural: lost ${hand.bet}`);
                        }
                    }
                }
            }
        }
        else {
            // dealer does not have a natural
            for (let i = 0; i < this.numPlayers; i++) {
                player = this.players[i];
                log(`player ${i + 1} - ${player.name}`);
                for (let j = 0; j < player.numMasterHands; j++) {
                    masterHand = player.masterHands[j];
                    for (let k = 0; k < masterHand.numHands; k++) {
                        hand = masterHand.hands[k];
                        if (hand.outcome === HandOutcome.BUST) {
                            this.addResult(player, k, hand, initialBet, -hand.bet);
                            log(`    hand ${j + 1}.${k + 1}: bust: lost ${hand.bet}`);
                        }
                        else if (hand.outcome === HandOutcome.SURRENDER) {
                            this.addResult(player, k, hand, initialBet, -hand.bet);
                            log(`    hand ${j + 1}.${k + 1}: surrender: lost ${hand.bet}`);
                        }
                        else {
                            // player has a non-bust, non-surrender hand
                            if (hand.isNatural) {
                                const payout = Math.floor(hand.bet * house_1.HouseRules.NATURAL_BLACKJACK_PAYOUT);
                                this.addResult(player, k, hand, initialBet, payout);
                                log(`    hand ${j + 1}.${k + 1}: natural: won ${payout}`);
                            }
                            else if (dealer.hand.outcome === HandOutcome.BUST) {
                                this.addResult(player, k, hand, initialBet, hand.bet);
                                log(`    hand ${j + 1}.${k + 1}: dealer bust: won ${hand.bet}`);
                            }
                            else {
                                // cards in the player and dealer hands now matter
                                if (hand.count < dealer.hand.count) {
                                    this.addResult(player, k, hand, initialBet, -hand.bet);
                                    log(`    hand ${j + 1}.${k + 1}: lost ${hand.bet}`);
                                }
                                else if (hand.count > dealer.hand.count) {
                                    this.addResult(player, k, hand, initialBet, hand.bet);
                                    log(`    hand ${j + 1}.${k + 1}: won ${hand.bet}`);
                                }
                                else {
                                    this.addResult(player, k, hand, initialBet, 0);
                                    log(`    hand ${j + 1}.${k + 1}: push`);
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
const blackjack = new BlackJack();
for (let i = 0; i < 100; i++) {
    blackjack.playGame();
}
// for the above loop, playing a million games takes about 2 minutes on the laptop.
// for 1,000,000 games, 3 master hands per game and $2 bets per hand => about $6,000,000 bet
// player Jack: {"handsPlayed":1027656,"handsWon":436271,"handsLost":507237,"handsPushed":84148,"proceeds":-7612}
// player Jill: {"handsPlayed":2055362,"handsWon":870914,"handsLost":1015242,"handsPushed":169206,"proceeds":-21542}
// 42% hands won, 49% hands lost, 8% hands pushed
blackjack.playerResults.forEach((value, key) => {
    const playerName = key;
    const playerResult = value;
    log(`\nplayer ${playerName}: ${JSON.stringify(playerResult)}`);
});
// for 1,000,000 games with 3 master hands per game => about 3,000,000 hands
// game stats: {"doubleDownCount":303854,"surrenderCount":155797,"splitCount":83018,"acesSplit":33595}
// roughly 10% hands double down, 5% surrender, 2.5% split, 1% split Aces
log(`\ngame stats: ${JSON.stringify(blackjack.stats)}`);
