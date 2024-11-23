"""Blsckjsck simulation."""

from enum import StrEnum
import random
import pdb

# seed the random number generator to make game play 100% reproducible:
random.seed(0xDEADBEEF)

#
# Card and Deck and Shoe
#


class CardSuit(StrEnum):
    # these are emoji symbols which require two unicode characters
    HEARTS = "♥️"  #   aka U+2665 + U+fe0f
    DIAMONDS = "♦️"  # aka U+2666 + U+fe0f
    SPADES = "♠️"  # aka U+2660 + U+fe0f
    CLUBS = "♣️"  #    aka U+2663 + U+fe0f


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
# House Rules
#


class HouseRules:
    DECKS_IN_SHOE: int = 6
    FORCE_RESHUFFLE: int = ((52 * DECKS_IN_SHOE) * 3) / 4

    # True => Must stand after Ace split -- stand on the Ace plus the one card dealt after split
    NO_MORE_CARDS_AFTER_SPLITTING_ACES: bool = True

    # [9, 10, 11] aka range(9, 12) => "Reno Rules"
    DOUBLE_DOWN_ON_TOTAL: list[int] = [i for i in range(1, 22)]

    DOUBLE_DOWN_AFTER_SPLIT: bool = True

    # 3 => turn one hand into no more than 4 hands
    SPLITS_PER_HAND: int = 3

    # rank match like K-K always can split, values match allows K-10 split
    SPLIT_ON_VALUE_MATCH: bool = True

    DEALER_HITS_HARD_ON: bool = 16  # or less
    DEALER_HITS_SOFT_ON: bool = 17  # or less

    # 1.5 => 3 to 2 payout, 1.2 => 6 to 5 payout
    NATURAL_BLACKJACK_PAYOUT: float = 1.5

    SURRENDER_ALLOWED: bool = True


#
# Players and Dealer
#


class PlayerHand:
    cards: list[Card]
    from_split: bool
    bet: int

    def __init__(self, bet: int = 2, from_split: bool = False):
        self.cards = []
        self.from_split = from_split
        self.bet = bet

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
                if CARD_VALUE(card1.rank) == CARD_VALUE(card2.rank):
                    return True
            else:
                if card1.rank == card2.rank:
                    return True
        return False

    def add_card(self, card: Card):
        self.cards.append(card)


class DealerHand(PlayerHand):
    def __init__(self):
        self.cards = []


class PlayerHands:
    hands: list[PlayerHand]
    HAND_LIMIT: int = HouseRules.SPLITS_PER_HAND

    def add_hand(self, bet: int):
        self.hands.append(PlayerHand(bet=bet))

    def split_hand(self, hand_index):
        # there are two in the hand of the same value
        # or rank depending of the house rules.
        card1 = self.hands[hand_index].cards[0:1]
        card2 = self.hands[hand_index].cards[1:2]

        old_player_hand = self.hands[hand_index]
        old_player_hand.cards = [card1]
        old_player_hand.from_split = True

        new_player_hand = PlayerHand(from_split=True)
        new_player_hand.cards = [card2]
        self.hands.append(PlayerHand)

    def can_split(self, hand_index) -> bool:
        hand: PlayerHand = self.hands[hand_index]
        if hand.card_count == 2:
            if hand.can_split:
                return True
        return False


# ye olde factory
def create_player_hands(hand_limit: int = 4) -> PlayerHands:
    player_hands: PlayerHands = PlayerHands()
    player_hands.hands = []
    player_hands.HAND_LIMIT = hand_limit
    return player_hands


class Player:
    # a player can play a set of hands
    # each inidivdual hand can be split into more hands, but is hard limit
    hands: list[PlayerHands]
    num_of_hands: int

    def set_game_bets(self, bets: list[int] = [2]):
        self.hands = []
        for bet in bets:
            # starts with one hand per bet
            player_hands: PlayerHands = create_player_hands()
            player_hands.add_hand(bet=bet)
            self.hands.append(player_hands)
        self.num_of_hands = len(bets)


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
# U => Surrender, in a world of too many Ss
Uh = "surrender-if-allowed-or-hit"
Us = "surrender-if-allowed-or-stand"
Usp = "surrender-if-allowed-or-split"
NO = "no-decision"


def convert_to_player_decision(decision: str) -> PlayerDecision:
    if decision == S:
        return PlayerDecision.STAND
    elif decision == H:
        return PlayerDecision.HIT
    elif decision == Dh:
        return PlayerDecision.DOUBLE
    elif decision == Ds:
        return PlayerDecision.HIT
    elif decision == SP:
        return PlayerDecision.HIT
    elif decision == Uh:
        return PlayerDecision.HIT
    elif decision == Us:
        return PlayerDecision.HIT
    elif decision == Usp:
        return PlayerDecision.HIT
    else:
        return PlayerDecision.STAND


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


def create_pairs_decision() -> dict[CardRank, str]:
    """
    Turn the _PAIRS_DECISION table into a rank based dictiopnary:
        PAIRS_DECISION: dict[CardRank, str] = {
            CardRank.ACE: {
                CardRank.ACE: SP, CardRank.TWO: SP, ..., CardRank.KING: SP,
            }, ...
        }
    """
    decisions: dict[CardRank, str] = {}

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
    [NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO, NO],  # fmt: skip
    # hard total: 5  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 6  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 7  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 8  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 9  x  dealer top card
    [NO,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 10  x  dealer top card
    [NO,  H,  H, Dh, Dh, Dh, Dh,  H,  H,  H,  H,  H,  H,  H],  # fmt: skip
    # hard total: 10  x  dealer top card
    [NO,  H,  H, Dh, Dh, Dh, Dh, Dh, Dh, Dh,  H,  H,  H,  H],  # fmt: skip
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
    # soft total: 20 (A, 9)  x  dealer top card
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
        dealer_top_card: Card, player_hand: PlayerHand
    ) -> PlayerDecision:
        is_first_decision: bool = len(player_hand) == 2 and not player_hand.from_split
        if is_first_decision:
            if HouseRules.SURRENDER_ALLOWED:
                pass


#
# Game
#

player1: Player = Player()
player2: Player = Player()


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

        # deal Player1, Player2, Dealer
        # where players may have more than one hand to start
        for _ in range(2):
            for i in range(player1.num_of_hands):
                card = self.get_card_from_shoe()
                initial_hands: PlayerHands = player1.hands[i]
                hand: PlayerHand = initial_hands.hands[0]
                hand.add_card(card)

            for i in range(player2.num_of_hands):
                card = self.get_card_from_shoe()
                initial_hands: PlayerHands = player2.hands[i]
                hand: PlayerHand = initial_hands.hands[0]
                hand.add_card(card)

            card = self.get_card_from_shoe()
            dealer.hand.add_card(card)

        dealer_top_card: Card = dealer.top_card
        dealer_hole_card: Card = dealer.hole_card

        # ready to rumble
        print("player 1:")
        for i in range(player1.num_of_hands):
            print(f"    hand {i+1}:")
            player_hands: PlayerHands = player1.hands[i]
            for j in range(len(player_hands.hands)):
                print(f"    hand {i+1}.{j+1}:")
                for k in range(len(player_hands.hands[j].cards)):
                    card = player_hands.hands[j].cards[k]
                    print(f"        card {k+1}: {card.rank.value}{card.suite.value}")

        print("player 2:")
        for i in range(player2.num_of_hands):
            print(f"    hand {i+1}:")
            player_hands: PlayerHands = player2.hands[i]
            for j in range(len(player_hands.hands)):
                print(f"    hand {i+1}.{j+1}:")
                for k in range(len(player_hands.hands[j].cards)):
                    card = player_hands.hands[j].cards[k]
                    print(f"        card {k+1}: {card.rank.value}{card.suite.value}")

        print("dealer:")
        print(
            f"    top  card: {dealer_top_card.rank.value}{dealer_top_card.suite.value}."
        )
        print(
            f"    hole card: {dealer_hole_card.rank.value}{dealer_hole_card.suite.value}."
        )


pdb.set_trace()

bj = BlackJack()
bj.play_game()

# exit(0xDEADBEEF)

print("Done.")
