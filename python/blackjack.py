"""
Blackjack simulation.

Learn to play blackjack:
https://bicyclecards.com/how-to-play/blackjack

The House Edge and the Basic Strategy of Blackjack 
https://www.shs-conferences.org/articles/shsconf/pdf/2022/18/shsconf_icprss2022_03038.pdf

Blackjack - Wikipedia
https://en.wikipedia.org/wiki/Blackjack
https://en.wikipedia.org/wiki/Blackjack#Basic_strategy

6 decks, Dealer hits on soft 17, BJ pays 3 to 2
https://blog.mbitcasino.io/wp-content/uploads/2024/04/chart01_6decks.jpg
6 decks, Dealer stands on soft 17, BJ pays 3 to 2
https://blog.mbitcasino.io/wp-content/uploads/2024/04/chart03_6decks.jpg

Central Las Vegas Strip Blackjack Survey 2024
https://vegasadvantage.com/las-vegas-blackjack/central-las-vegas-strip-blackjack/

Las Vegas Blackjack Guide
https://vegasadvantage.com/las-vegas-blackjack-guide/
"""

from enum import StrEnum
import random
import pdb

# seed the random number generator to make game play 100% reproducible:
random.seed(0xDEADBEEF)
random.seed(42)  # hit, stand, split
random.seed(43)  # double, hit, hit
random.seed(45)  # double, stand, surrender
random.seed(48)  # stand, hit , surrender
random.seed(56)  # dealer and player blackjack
random.seed(59)  # double, stand, stand (player blackjack)
random.seed(60)  # surrender, hit, double
random.seed(66)  # hit, split (double, hit), stand

random.seed(0xDEADBEEF)

#
# House Rules
#


class HouseRules:
    # Hit/Stand on a soft 17 and 3:2 black jack payouts
    # are what casinos advetize wrt their BJ tables:
    # 1. 6/8 decks in shoe
    # 2. 3:2 blackjack payout
    # 3. Hit/Stand on a soft 17
    # 4. Re-splitting Aces (exceptionally rare)
    # 5. Surrender

    DECKS_IN_SHOE: int = 6
    FORCE_RESHUFFLE: int = ((52 * DECKS_IN_SHOE) * 3) / 4

    # True => Must stand after the Ace split (stand on the Ace plus the one card dealt after split)
    # True => no double down after the Ace split, no splitting Aces after the Ace split
    NO_MORE_CARDS_AFTER_SPLITTING_ACES: bool = True

    # [9, 10, 11] aka range(9, 12) => "Reno Rules"
    DOUBLE_DOWN_ON_TOTAL: list[int] = [i for i in range(1, 22)]

    # Does not apply tp Aces if NO_MORE_CARDS_AFTER_SPLITTING_ACES is true
    DOUBLE_DOWN_AFTER_SPLIT: bool = True

    # 3 => turn one hand into no more than 4 hands
    SPLITS_PER_HAND: int = 3

    # rank match like K-K always can split, values match allows K-10 split
    SPLIT_ON_VALUE_MATCH: bool = True

    # Hit on soft 17 (6/8 decks) is more common on low bet tables.
    DEALER_HITS_HARD_ON: int = 16  # or less
    DEALER_HITS_SOFT_ON: int = 17  # or less

    # 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
    # 6 to 5 is more common in two deck games
    NATURAL_BLACKJACK_PAYOUT: float = 1.5

    # Usually 8 deck game, no Ace re-splitting, 50-100 minimum bet ...
    # Setting True here since I am a high roller ;) and want to shake out the code.
    SURRENDER_ALLOWED: bool = True


#
# Card and Deck and Shoe
#


# fmt: off
class CardSuit(StrEnum):
    # these are emoji symbols which require two unicode characters
    HEARTS   = "♥️"  # aka U+2665 + U+fe0f
    DIAMONDS = "♦️"  # aka U+2666 + U+fe0f
    SPADES   = "♠️"  # aka U+2660 + U+fe0f
    CLUBS    = "♣️"  # aka U+2663 + U+fe0f
# fmt: on

# print(f"{ord('♥️'[0]): #04x}") # 0x02665
# print(f"{ord('♥️'[1]): #04x}") # 0x0fe0f
# print("\u2665\ufe0f")


class CardRank(StrEnum):
    ACE = "A"
    TWO = "2"
    THREE = "3"
    FOUR = "4"
    FIVE = "5"
    SIX = "6"
    SEVEN = "7"
    EIGHT = "8"
    NINE = "9"
    TEN = "10"
    JACK = "J"
    QUEEN = "Q"
    KING = "K"


CARD_VALUE = {
    CardRank.ACE: 1,
    CardRank.TWO: 2,
    CardRank.THREE: 3,
    CardRank.FOUR: 4,
    CardRank.FIVE: 5,
    CardRank.SIX: 6,
    CardRank.SEVEN: 7,
    CardRank.EIGHT: 8,
    CardRank.NINE: 9,
    CardRank.TEN: 10,
    CardRank.JACK: 10,
    CardRank.QUEEN: 10,
    CardRank.KING: 10,
}


class Card:
    def __init__(self, suite: CardSuit, rank: CardRank):
        self.suite = suite
        self.rank = rank

    suite: CardSuit
    rank: CardRank


UNSHUFFLED_DECK: list[Card] = [
    # HEARTS
    Card(suite=CardSuit.HEARTS, rank=CardRank.ACE),
    Card(suite=CardSuit.HEARTS, rank=CardRank.TWO),
    Card(suite=CardSuit.HEARTS, rank=CardRank.THREE),
    Card(suite=CardSuit.HEARTS, rank=CardRank.FOUR),
    Card(suite=CardSuit.HEARTS, rank=CardRank.FIVE),
    Card(suite=CardSuit.HEARTS, rank=CardRank.SIX),
    Card(suite=CardSuit.HEARTS, rank=CardRank.SEVEN),
    Card(suite=CardSuit.HEARTS, rank=CardRank.EIGHT),
    Card(suite=CardSuit.HEARTS, rank=CardRank.NINE),
    Card(suite=CardSuit.HEARTS, rank=CardRank.TEN),
    Card(suite=CardSuit.HEARTS, rank=CardRank.JACK),
    Card(suite=CardSuit.HEARTS, rank=CardRank.QUEEN),
    Card(suite=CardSuit.HEARTS, rank=CardRank.KING),
    # DIAMONDS
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.ACE),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.TWO),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.THREE),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.FOUR),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.FIVE),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.SIX),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.SEVEN),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.EIGHT),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.NINE),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.TEN),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.JACK),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.QUEEN),
    Card(suite=CardSuit.DIAMONDS, rank=CardRank.KING),
    # SPADES
    Card(suite=CardSuit.SPADES, rank=CardRank.ACE),
    Card(suite=CardSuit.SPADES, rank=CardRank.TWO),
    Card(suite=CardSuit.SPADES, rank=CardRank.THREE),
    Card(suite=CardSuit.SPADES, rank=CardRank.FOUR),
    Card(suite=CardSuit.SPADES, rank=CardRank.FIVE),
    Card(suite=CardSuit.SPADES, rank=CardRank.SIX),
    Card(suite=CardSuit.SPADES, rank=CardRank.SEVEN),
    Card(suite=CardSuit.SPADES, rank=CardRank.EIGHT),
    Card(suite=CardSuit.SPADES, rank=CardRank.NINE),
    Card(suite=CardSuit.SPADES, rank=CardRank.TEN),
    Card(suite=CardSuit.SPADES, rank=CardRank.JACK),
    Card(suite=CardSuit.SPADES, rank=CardRank.QUEEN),
    Card(suite=CardSuit.SPADES, rank=CardRank.KING),
    # CLUBS
    Card(suite=CardSuit.CLUBS, rank=CardRank.ACE),
    Card(suite=CardSuit.CLUBS, rank=CardRank.TWO),
    Card(suite=CardSuit.CLUBS, rank=CardRank.THREE),
    Card(suite=CardSuit.CLUBS, rank=CardRank.FOUR),
    Card(suite=CardSuit.CLUBS, rank=CardRank.FIVE),
    Card(suite=CardSuit.CLUBS, rank=CardRank.SIX),
    Card(suite=CardSuit.CLUBS, rank=CardRank.SEVEN),
    Card(suite=CardSuit.CLUBS, rank=CardRank.EIGHT),
    Card(suite=CardSuit.CLUBS, rank=CardRank.NINE),
    Card(suite=CardSuit.CLUBS, rank=CardRank.TEN),
    Card(suite=CardSuit.CLUBS, rank=CardRank.JACK),
    Card(suite=CardSuit.CLUBS, rank=CardRank.QUEEN),
    Card(suite=CardSuit.CLUBS, rank=CardRank.KING),
]


def shuffle_shoe(shoe: list[Card]) -> None:
    """Shuffle shoe in place."""
    random.shuffle(shoe)


def create_shoe() -> list[Card]:
    "Create a shoe at the start of the match.  Returned shoe is shuffled."
    shoe: list[Card]
    shoe = [card for i in range(HouseRules.DECKS_IN_SHOE) for card in UNSHUFFLED_DECK]
    shuffle_shoe(shoe)
    return shoe


def display_shoe(shoe: list[Card]) -> None:
    for card in shoe:
        print(f"""{card.rank.value}{card.suite.value}""")


#
# Players and Dealer
#


class HandOutcome(StrEnum):
    STAND = "stand"
    BUST = "bust"
    SURRENDER = "surrender"
    DEALER_BLACKJACK = "dealer-blackjack"
    IN_PLAY = "in-play"


class PlayerHand:
    cards: list[Card]
    from_split: bool
    bet: int
    outcome: HandOutcome

    def __init__(self, bet: int = 2, from_split: bool = False):
        self.cards = []
        self.from_split = from_split
        self.bet = bet
        self.outcome = HandOutcome.IN_PLAY

    @property
    def aces_count(self) -> int:
        return sum([1 for card in self.cards if card.rank == CardRank.ACE])

    @property
    def hard_count(self) -> int:
        hard_count: int = 0
        for card in self.cards:
            hard_count += CARD_VALUE[card.rank]
        return hard_count

    @property
    def soft_count(self) -> int:
        """Return the highest soft count possible."""
        hard_count: int = self.hard_count
        aces_count: int = self.aces_count
        if aces_count == 0:
            return hard_count

        soft_count: int = 0
        for card in self.cards:
            if card.rank == CardRank.ACE:
                soft_count += 11
            else:
                soft_count += CARD_VALUE[card.rank]

        # case of Ace + 5, where the count can be 6 or 16
        # case of Ace + Ace + 5, where count can be 7 or 17 or 27

        if soft_count > 21:
            for i in range(aces_count):
                soft_count -= 10
                if soft_count <= 21:
                    return soft_count

        return soft_count

    @property
    def count(self) -> int:
        return self.soft_count

    @property
    def is_natural(self) -> bool:
        if not self.from_split:
            if len(self.cards) == 2:
                if self.count == 21:
                    return True
        return False

    @property
    def is_bust(self) -> bool:
        return self.count > 21

    @property
    def card_count(self) -> int:
        return len(self.cards)

    @property
    def can_split(self) -> bool:
        if len(self.cards) == 2:
            card1 = self.cards[0]
            card2 = self.cards[1]
            # check card value or rank, depending on hiuse rules
            if HouseRules.SPLIT_ON_VALUE_MATCH:
                if CARD_VALUE[card1.rank] == CARD_VALUE[card2.rank]:
                    return True
            else:
                if card1.rank == card2.rank:
                    return True
        return False

    def add_card(self, card: Card):
        self.cards.append(card)

    @property
    def is_hand_over(self) -> bool:
        if self.outcome == HandOutcome.STAND:
            return True
        elif self.outcome == HandOutcome.BUST:
            return True
        elif self.outcome == HandOutcome.SURRENDER:
            return True
        elif self.outcome == HandOutcome.DEALER_BLACKJACK:
            return True
        else:  # HandOutcome.IN_PLAY
            return False


class DealerHand(PlayerHand):
    def __init__(self):
        self.cards = []
        self.outcome = HandOutcome.IN_PLAY

    @property
    def is_natural(self) -> bool:
        if len(self.cards) == 2:
            if self.count == 21:
                return True
        return False


class PlayerMasterHand:
    HAND_LIMIT: int = HouseRules.SPLITS_PER_HAND + 1

    hands: list[PlayerHand]
    num_hands: int

    def __init__(self):
        self.hands = []
        self.num_hands = 0

    def add_start_hand(self, bet: int):
        self.hands.append(PlayerHand(bet=bet))
        self.num_hands += 1

    def split_hand(self, hand_index: int, cards_to_add: tuple):
        # there are two in the hand of the same value
        # or rank depending of the house rules.
        card1 = self.hands[hand_index].cards[0]
        card2 = self.hands[hand_index].cards[1]

        old_player_hand = self.hands[hand_index]
        old_player_hand.cards = [card1, cards_to_add[0]]
        old_player_hand.from_split = True
        old_player_hand.outcome = HandOutcome.IN_PLAY

        new_player_hand = PlayerHand(bet=old_player_hand.bet, from_split=True)
        new_player_hand.cards = [card2, cards_to_add[1]]
        new_player_hand.outcome = HandOutcome.IN_PLAY
        self.hands.append(new_player_hand)

        new_hand_index: int = self.num_hands
        self.num_hands += 1
        return new_hand_index

    def can_split(self, hand_index) -> bool:
        hand: PlayerHand = self.hands[hand_index]
        if hand.card_count == 2:
            if hand.can_split:
                return True
        return False


# ye olde factory
def create_player_master_hand(hand_limit: int = 4) -> PlayerMasterHand:
    player_hands: PlayerMasterHand = PlayerMasterHand()
    player_hands.hands = []
    player_hands.HAND_LIMIT = hand_limit
    return player_hands


class Player:
    # a player can play a set of hands.
    # each indivdual master hand can be split into more hands,
    # for which there is hard limit.  each master hand can typically be split
    # up to three times, for a total of four hands starting from the master hand.
    master_hands: list[PlayerMasterHand]
    num_of_hands: int
    name: str

    def __init__(self, name: str):
        self.master_hands = []
        self.num_of_hands = 0
        self.name = name

    def set_game_bets(self, bets: list[int] = [2]):
        """
        At start of the game, the player will place separate bets
        for each hand that they want. Each of these original hands
        will be considered a "master" hand.
        """
        self.master_hands = []
        self.num_of_hands = 0
        for bet in bets:
            # starts with one hand per bet
            player_master_hand: PlayerMasterHand = create_player_master_hand()
            player_master_hand.add_start_hand(bet=bet)
            self.master_hands.append(player_master_hand)
            self.num_of_hands += 1


class Dealer:
    hand: DealerHand

    def __init__(self):
        self.hand = DealerHand()

    @property
    def top_card(self) -> Card:
        return self.hand.cards[0]

    @property
    def hole_card(self) -> Card:
        return self.hand.cards[1]


#
# Basic Strategy:
#     https://en.wikipedia.org/wiki/Blackjack#Basic_strategy
#     All Hail The Wikipedia!
#


class PlayerDecision(StrEnum):
    STAND = "stand"
    HIT = "hit"
    DOUBLE = "double-down"
    SPLIT = "split"
    SURRENDER = "surrender"


# to make decision tables readable:
S = "stand"
H = "hit"
Dh = "double-down-if-allowed-or-hit"
Ds = "double-down-if-allowed-or-stand"
SP = "split"
# U => Surrender, in a world of too many S words
Uh = "surrender-if-allowed-or-hit"
Us = "surrender-if-allowed-or-stand"
Usp = "surrender-if-allowed-or-split"
NO = "no-decision"


def convert_to_player_decision(
    decision: str, player_hand: PlayerHand
) -> PlayerDecision:
    """
    Decision sometimes return Xy, which translates to do X if allowed else do y.
    Determine the X or the y here.
    """
    is_first_decision: bool = len(player_hand.cards) == 2
    is_first_postsplit_decision: bool = is_first_decision and player_hand.from_split

    surrender_can_be_played: bool = is_first_decision and HouseRules.SURRENDER_ALLOWED

    if decision == S:
        return PlayerDecision.STAND

    elif decision == H:
        return PlayerDecision.HIT

    elif decision == Dh or decision == Ds:
        # may be only allow to down on hand totals [9, 10, 11] or some such
        # basic stratgey wants to double down on
        #     hand hard totals [9, 10, 11]
        #     hand soft totals [12, 13,14, 15, 16, 17, 18, 19]
        player_decision: PlayerDecision
        nondouble_down_decision: PlayerDecision = (
            PlayerDecision.HIT if decision == Dh else PlayerDecision.STAND
        )

        can_double_down: bool
        if is_first_decision:
            if is_first_postsplit_decision:
                if HouseRules.DOUBLE_DOWN_AFTER_SPLIT:
                    can_double_down = True
                else:
                    can_double_down = False
            else:
                can_double_down = True
        else:
            can_double_down = False

        if can_double_down:
            double_down: bool
            hard_total: int = player_hand.hard_count
            soft_total: int = player_hand.soft_count
            totals_allowed: list[int] = HouseRules.DOUBLE_DOWN_ON_TOTAL
            if hard_total in totals_allowed:
                double_down = True
            elif soft_total in totals_allowed:
                double_down = True
            else:
                double_down = False
            player_decision = (
                PlayerDecision.DOUBLE if double_down else nondouble_down_decision
            )
        else:
            player_decision = nondouble_down_decision

        return player_decision

    elif decision == SP:
        return PlayerDecision.SPLIT

    elif decision == Uh or decision == Us or decision == Usp:
        # surrent decision must be allowed in the House Rules and
        # must be a first decision (before splitting)
        player_decision: PlayerDecision
        nonsurrender_decision: PlayerDecision = (
            PlayerDecision.HIT
            if decision == Uh
            else PlayerDecision.STAND
            if decision == "Us"
            else PlayerDecision.SPLIT
        )

        can_surrender: bool
        if is_first_decision:
            if not is_first_postsplit_decision:
                if HouseRules.SURRENDER_ALLOWED:
                    can_surrender = True
                else:
                    can_surrender = False
            else:
                can_surrender = False
        else:
            can_surrender = False

        if can_surrender:
            player_decision = PlayerDecision.SURRENDER
        else:
            player_decision = nonsurrender_decision

        return player_decision

    raise Exception("convert_to_player_decision() ran into a little trouble in town.")


RANK_TO_INDEX = {
    CardRank.ACE: 1,
    CardRank.TWO: 2,
    CardRank.THREE: 3,
    CardRank.FOUR: 4,
    CardRank.FIVE: 5,
    CardRank.SIX: 6,
    CardRank.SEVEN: 7,
    CardRank.EIGHT: 8,
    CardRank.NINE: 9,
    CardRank.TEN: 10,
    CardRank.JACK: 11,
    CardRank.QUEEN: 12,
    CardRank.KING: 13,
}

# Use "fmt: skip" to force black to leave these table alone.
# Use NO since we live in a zero index world.
# Only use the SP decision from this table.  The other decisions mirror exactly
# what the hard/soft total decision tables yield.
_PAIRS_DECISION = [
    # 0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # player pair card: Ace  x  dealer top card
    [NO, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],  # fmt: skip
    # player pair card: 2  x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # player pair card: 3  x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # player pair card: 4  x  dealer top card
    [NO,  H,  H,  H, H, SP, SP,  H,  H,  H,  H,   H,  H,  H],  # fmt: skip
    # player pair card: 5  x  dealer top card
    [NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H],  # fmt: skip
    # player pair card: 6 x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # player pair card: 7  x  dealer top card
    [NO,  H, SP, SP, SP, SP, SP, SP,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # player pair card: 8  x  dealer top card
    [NO, Usp, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP, SP],  # fmt: skip
    # player pair card: 9  x  dealer top card
    [NO,  S, SP, SP, SP, SP, SP,  S, SP, SP,  S,  S,  S,  S],  # fmt: skip
    # player pair card: 10  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # player pair card: J  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # player pair card: Q  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # player pair card: K  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
]


def create_pairs_decision() -> dict[CardRank, dict[CardRank, str]]:
    """
    Turn the _PAIRS_DECISION table into a rank based dictiopnary:
        PAIRS_DECISION: dict[CardRank, str] = {
            CardRank.ACE: {
                CardRank.ACE: SP, CardRank.TWO: SP, ..., CardRank.KING: SP,
            }, ...
        }
    """
    decisions: dict[CardRank, dict[CardRank, str]] = {}

    for player_pair_card_rank in CardRank:
        decisions[player_pair_card_rank] = {}

    for player_pair_card_rank in CardRank:
        player_index: int = RANK_TO_INDEX[player_pair_card_rank]
        for dealer_top_card_rank in CardRank:
            dealer_index: int = RANK_TO_INDEX[dealer_top_card_rank]
            decision: str = _PAIRS_DECISION[player_index][dealer_index]
            decisions[player_pair_card_rank][dealer_top_card_rank] = decision

    return decisions


PAIRS_DECISION: dict[CardRank, str] = create_pairs_decision()

# Expect to use the soft total decision table for: (A,A) and (A,2),
# which is the only way to get to hard totals 2 and 3.
_HARD_TOTAL_DECISION = [
    # 0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # hard total: 1  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # hard total: 2  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # hard total: 3  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # hard total: 4  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 5  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 6  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 7  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 8  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 9  x  dealer top card
    [NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 10  x  dealer top card
    [NO,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H],  # fmt: skip
    # hard total: 11  x  dealer top card
    [NO, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh, Dh],  # fmt: skip
    # hard total: 12  x  dealer top card
    [NO,  H,  H,  H,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 13  x  dealer top card
    [NO,  H,  S,  S,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 14  x  dealer top card
    [NO,  H,  S,  S,  S,  S,  S,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 15  x  dealer top card
    [NO, Uh,  S,  S,  S,  S,  S,  H,  H,  H, Uh, Uh, Uh, Uh],  # fmt: skip
    # hard total: 16  x  dealer top card
    [NO, Uh,  S,  S,  S,  S,  S,  H,  H, Uh, Uh, Uh, Uh, Uh],  # fmt: skip
    # hard total: 17  x  dealer top card
    [NO, Us,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # hard total: 18  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # hard total: 19  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # hard total: 20  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # hard total: 21  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
]


def create_hard_total_decision() -> list[dict[CardRank, str]]:
    """
    Turn _HARD_TOTAL_DECISION table into
        HARD_TOTAL_DECISION = [
            ...
            {CardRank.ACE: H, CardRank.TWO: H, ..., CardRank.KING: H}, # [5]
            ...
        ]
    """
    decisions = [{} for i in range(22)]
    for hard_total in range(22):
        for rank in CardRank:
            decision: str = _HARD_TOTAL_DECISION[hard_total][CARD_VALUE[rank]]
            decisions[hard_total][rank] = decision
    return decisions


HARD_TOTAL_DECISION: list[dict[CardRank, str]] = create_hard_total_decision()


_SOFT_TOTAL_DECISION = [
    # 0   A   2   3   4   5   6   7   8   9  10   J   Q   K
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 1  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 2  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 3  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 4  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 5  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 6  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 7  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 8  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 9  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 10  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 11  x  dealer top card
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # soft total: 12 (A, A)  x  dealer top card
    [NO,  H,  H,  H,  H,  H, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 13 (A, 2)  x  dealer top card
    [NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 14 (A, 3)  x  dealer top card
    [NO,  H,  H,  H,  H, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 15 (A, 4)  x  dealer top card
    [NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 16 (A, 5)  x  dealer top card
    [NO,  H,  H,  H, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 17 (A, 6)  x  dealer top card
    [NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 18 (A, 7)  x  dealer top card
    [NO,  H, Ds, Ds, Ds, Ds, Ds,  S,  S,  H,  H,  H,  H,  H],  # fmt: skip
    # soft total: 19 (A, 8)  x  dealer top card
    [NO,  S,  S,  S,  S,  S, Ds,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # soft total: 20 (A, 9)  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
    # soft total: 21 (A, 10)  x  dealer top card
    [NO,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S,  S],  # fmt: skip
]


def create_soft_total_decision() -> list[dict[CardRank, str]]:
    """
    Turn _SOFT_TOTAL_DECISION table into
        SOFT_TOTAL_DECISION = [
            ...
            {CardRank.ACE: H, CardRank.TWO: H, ..., CardRank.KING: H}, # [5]
            ...
        ]
    """
    decisions = [{} for i in range(22)]
    for soft_total in range(22):
        for rank in CardRank:
            decision: str = _SOFT_TOTAL_DECISION[soft_total][CARD_VALUE[rank]]
            decisions[soft_total][rank] = decision
    return decisions


SOFT_TOTAL_DECISION: list[dict[CardRank, str]] = create_soft_total_decision()


class BasicStrategy:
    @staticmethod
    def determine_play(
        dealer_top_card: Card,
        player_hand: PlayerHand,
        hand_allows_more_splits: bool,
    ) -> PlayerDecision:
        is_first_decision: bool = len(player_hand.cards) == 2
        is_first_postsplit_decision: bool = is_first_decision and player_hand.from_split

        player_decision: PlayerDecision

        player_card1: Card = player_hand.cards[0]
        player_card2: Card = player_hand.cards[1]

        got_pairs: bool = False
        if HouseRules.SPLIT_ON_VALUE_MATCH:
            # recognize pair match for all card with a value of 10, plus ordinary pairs
            got_pairs = CARD_VALUE[player_card1.rank] == CARD_VALUE[player_card2.rank]
        else:
            got_pairs = player_card1.rank == player_card2.rank

        if got_pairs and hand_allows_more_splits:
            # Determine if the pairs can be split.
            # Note all of the non-split decisions that are ignored below
            # will not contradict the hard/soft total decision.
            pair_rank: CardRank
            if CARD_VALUE[player_card1.rank] == 10:
                pair_rank = CardRank.TEN
            else:
                pair_rank = player_card1.rank

            decision: str = PAIRS_DECISION[player_card1.rank][dealer_top_card.rank]

            player_decision = convert_to_player_decision(
                decision=decision, player_hand=player_hand
            )
            if player_decision == PlayerDecision.SPLIT:
                return PlayerDecision.SPLIT

        use_soft_total: bool = player_hand.hard_count < player_hand.soft_count <= 21

        if use_soft_total:
            soft_total: int = player_hand.soft_count
            decision: str = SOFT_TOTAL_DECISION[soft_total][dealer_top_card.rank]
            player_decision = convert_to_player_decision(
                decision=decision, player_hand=player_hand
            )
            return player_decision

        else:
            hard_total: int = player_hand.hard_count
            decision: str = HARD_TOTAL_DECISION[hard_total][dealer_top_card.rank]
            player_decision = convert_to_player_decision(
                decision=decision, player_hand=player_hand
            )
            return player_decision

        raise Exception(
            "BasicStrategy.determine_play() ran into a little trouble in town."
        )


#
# Game
#

player1: Player = Player(name="Jack")
player2: Player = Player(name="Jill")


class BlackJack:
    shoe: list[Card]
    shoe_top: int
    players: list[Player]

    def __init__(self):
        self.shoe = create_shoe()
        self.shoe_top = 0

    def reshuffle_shoe(self):
        shuffle_shoe(self.shoe)
        self.shoe_top = 0

    def get_card_from_shoe(self) -> Card:
        card: Card = self.shoe[self.shoe_top]
        self.shoe_top += 1
        return card

    def set_players(self, players: list[Player]) -> None:
        self.players = players

    def play_game(self):
        if self.shoe_top > HouseRules.FORCE_RESHUFFLE:
            self.reshuffle_shoe()

        # who is playing?
        self.set_players(players=[player1, player2])

        # how many hands==bets for each player?
        player1.set_game_bets(bets=[2])
        player2.set_game_bets(bets=[2, 2])

        # dealer gets a hand
        dealer: Dealer = Dealer()

        card: Card

        print("\nDEAL HANDS")

        # deal Player1, Player2, Dealer
        # where players may have more than one hand to start
        for _ in range(2):
            for i in range(player1.num_of_hands):
                card = self.get_card_from_shoe()
                master_hand: PlayerMasterHand = player1.master_hands[i]
                hand: PlayerHand = master_hand.hands[0]
                hand.add_card(card)

            for i in range(player2.num_of_hands):
                card = self.get_card_from_shoe()
                master_hand: PlayerMasterHand = player2.master_hands[i]
                hand: PlayerHand = master_hand.hands[0]
                hand.add_card(card)

            card = self.get_card_from_shoe()
            dealer.hand.add_card(card)

        dealer_top_card: Card = dealer.top_card
        dealer_hole_card: Card = dealer.hole_card

        print(
            f"dealer top card: {dealer_top_card.rank.value}{dealer_top_card.suite.value}."
        )

        print("PLAY HANDS")
        #
        # 1. Does Dealer have a natural?
        # 1.1. Is the delaer top card A, 10, J, Q, K?
        # 1.2. Do players want Insurance?
        #

        if dealer.hand.is_natural:
            # a real simulation would have to take care of Insurance, which is a sucker's bet,
            # so we just assume that no player will ask for insurance.
            # two cases:
            #     1. player has a natural and their bet is pushed
            #     2. player loses

            dealer.hand.outcome = HandOutcome.DEALER_BLACKJACK

            for p, player in enumerate(self.players):
                for mh, master_hand in enumerate(player.master_hands):
                    hand_index: int = 0
                    while hand_index < master_hand.num_hands:
                        hand: PlayerHand = master_hand.hands[hand_index]
                        hand.outcome = HandOutcome.STAND
                        hand_index += 1

        else:
            for p, player in enumerate(self.players):
                print(f"player {p+1} - {player.name}:")
                for mh, master_hand in enumerate(player.master_hands):
                    hand_index: int = 0
                    while hand_index < master_hand.num_hands:
                        hand: PlayerHand = master_hand.hands[hand_index]

                        print(f"    hand {mh+1}.{hand_index+1}:")
                        for card_index, card in enumerate(hand.cards):
                            print(f"""        Card {card_index+1}: {card.rank.value}{card.suite.value}""")

                        num_hands: int = len(master_hand.hands)
                        is_split_possible: bool = (
                            num_hands < PlayerMasterHand.HAND_LIMIT
                        )

                        #
                        # 2. Player decides each hand via Basic Strategy
                        #

                        while True:
                            decision: PlayerDecision = BasicStrategy.determine_play(
                                dealer_top_card=dealer_top_card,
                                player_hand=hand,
                                hand_allows_more_splits=is_split_possible,
                            )
                            print(
                                f"""        decision: {decision.value}"""
                            )

                            if decision == PlayerDecision.STAND:
                                hand.outcome = HandOutcome.STAND
                                print(
                                    f"""        stand total H{hand.hard_count} S{hand.soft_count}"""
                                )
                                break

                            elif decision == PlayerDecision.SURRENDER:
                                hand.outcome = HandOutcome.SURRENDER
                                hand.bet = int(hand.bet / 2)
                                break

                            elif decision == PlayerDecision.DOUBLE:
                                card = self.get_card_from_shoe()
                                hand.add_card(card)
                                hand.bet *= 2
                                hand.outcome = HandOutcome.STAND
                                print(
                                    f"""        double down: {card.rank.value}{card.suite.value}, total H{hand.hard_count} S{hand.soft_count}"""
                                )
                                break

                            elif decision == PlayerDecision.HIT:
                                card = self.get_card_from_shoe()
                                hand.add_card(card)
                                hand_total: int = hand.count
                                print(
                                    f"""        hit: {card.rank.value}{card.suite.value}, total H{hand.hard_count} S{hand.soft_count}"""
                                )
                                if hand_total > 21:
                                    hand.outcome = HandOutcome.BUST
                                    print(f"""        bust""")
                                    break
                                else:
                                    hand.outcome = HandOutcome.IN_PLAY

                            elif decision == PlayerDecision.SPLIT:
                                card1: Card = self.get_card_from_shoe()
                                card2: Card = self.get_card_from_shoe()
                                new_hand_index: int
                                new_hand_index = master_hand.split_hand(
                                    hand_index, cards_to_add=(card1, card2)
                                )
                                print(
                                    f"""        split, adding cards: {card1.rank.value}{card1.suite.value}, {card2.rank.value}{card2.suite.value}"""
                                )

                            else:
                                print("FTW")
                                break

                        hand_index += 1

        #
        # Dealer finishes their hand
        #

        dealer_done: bool = dealer.hand.outcome == HandOutcome.DEALER_BLACKJACK
        print("dealer:")
        print(
            f"    top  card: {dealer_top_card.rank.value}{dealer_top_card.suite.value}."
        )
        print(
            f"    hole card: {dealer_hole_card.rank.value}{dealer_hole_card.suite.value}."
        )
        while not dealer_done:
            hard_count: int = dealer.hand.hard_count
            soft_count: int = dealer.hand.soft_count

            use_soft_count: bool
            if hard_count < soft_count <= 21:
                use_soft_count = True
            else:
                use_soft_count = False

            if use_soft_count and soft_count <= HouseRules.DEALER_HITS_SOFT_ON:
                card = self.get_card_from_shoe()
                dealer.hand.add_card(card)
                print(
                    f"""    add: {card.rank.value}{card.suite.value}, total {dealer.hand.count}"""
                )
            elif not use_soft_count and hard_count <= HouseRules.DEALER_HITS_HARD_ON:
                card = self.get_card_from_shoe()
                dealer.hand.add_card(card)
                print(
                    f"""    add: {card.rank.value}{card.suite.value}, total {dealer.hand.count}"""
                )
            else:
                dealer.hand.outcome = HandOutcome.STAND
                dealer_done = True
                continue

            if dealer.hand.count > 21:
                dealer.hand.outcome = HandOutcome.BUST
                dealer_done = True
            else:
                dealer_done = False

        print(f"""    outcome: {dealer.hand.outcome}, total {dealer.hand.count}""")

        #
        # Settle hands
        #

        print("SETTLE HANDS")
        if dealer.hand.outcome == HandOutcome.DEALER_BLACKJACK:
            for p, player in enumerate(self.players):
                print(f"player {p+1} - {player.name}:")
                for mh, master_hand in enumerate(player.master_hands):
                    hand_index: int = 0
                    while hand_index < master_hand.num_hands:
                        hand: PlayerHand = master_hand.hands[hand_index]
                        if hand.is_natural:
                            # player neither wins or loses, bet is pushed.
                            print(
                                f"""    hand {mh+1}.{hand_index+1}: push, both player and dealer had naturals"""
                            )
                        else:
                            print(
                                f"""    hand {mh+1}.{hand_index+1}: lost {hand.bet}"""
                            )

                        hand_index += 1

        else:
            for p, player in enumerate(self.players):
                print(f"player {p+1} - {player.name}:")
                for mh, master_hand in enumerate(player.master_hands):
                    hand_index: int = 0
                    while hand_index < master_hand.num_hands:
                        hand: PlayerHand = master_hand.hands[hand_index]
                        if hand.outcome == HandOutcome.BUST:
                            print(
                                f"""    hand {mh+1}.{hand_index+1}: lost {hand.bet}, player bust"""
                            )
                        elif hand.outcome == HandOutcome.SURRENDER:
                            print(
                                f"""    hand {mh+1}.{hand_index+1}: lost {hand.bet}, player surrender"""
                            )
                        else:
                            # player has a non-bust, non-surrender hand
                            if hand.is_natural:
                                print(
                                    f"""    hand {mh+1}.{hand_index+1}: won {int(hand.bet * HouseRules.NATURAL_BLACKJACK_PAYOUT)} with natural"""
                                )
                            elif dealer.hand.outcome == HandOutcome.BUST:
                                print(
                                    f"""    hand {mh+1}.{hand_index+1}: won {hand.bet}, dealer bust"""
                                )
                            else:
                                if hand.count < dealer.hand.count:
                                    print(
                                        f"""    hand {mh+1}.{hand_index+1}: lost {hand.bet}, dealer total { dealer.hand.count}, player total {hand.count}."""
                                    )
                                elif hand.count > dealer.hand.count:
                                    print(
                                        f"""    hand {mh+1}.{hand_index+1}: won {hand.bet}, dealer total { dealer.hand.count},  player total {hand.count}."""
                                    )
                                else:
                                    print(f"""    hand {mh+1}.{hand_index+1}: push""")

                        hand_index += 1


# pdb.set_trace()

bj = BlackJack()

for i in range(10):
    bj.play_game()

# exit(0xDEADBEEF)

print("Done.")
