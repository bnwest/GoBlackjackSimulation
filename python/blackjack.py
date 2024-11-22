"""Blsckjsck simulation."""

from enum import StrEnum
import random
import pdb

# seed the random number generator to make game play reproducible:
random.seed(0xDEADBEEF)


class CardSuit(StrEnum):
    # these are emoji symbols which require two unicode characters
    HEARTS = "♥️"  #   aka U+2665 + U+fe0f
    DIAMONDS = "♦️"  # aka U+2666 + U+fe0f
    SPADES = "♠️"  # aka U+2660 + U+fe0f
    CLUBS = "♣️"  #    aka U+2663 + U+fe0f


# print(f"{ord('♥️'[0]): #04x}") # 0x02665
# print(f"{ord('♥️'[1]): #04x}") # 0x0fe0f
# print("\u2665\ufe0f")

# convert enum to string:
# print(f"""Black spade: {CardSuit.SPADES.value}.""")


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

DECKS_IN_SHOE: int = 6
FORCE_RESHUFFLE: int = (52 * DECKS_IN_SHOE) * 3 / 4

import pdb


def shuffle_shoe(shoe: list[Card]) -> None:
    """Shuffle shoe in place."""
    random.shuffle(shoe)


def create_shoe() -> list[Card]:
    "Create a shoe at the start of the match.  Returned shoe is shuffled."
    shoe: list[Card]
    shoe = [card for i in range(DECKS_IN_SHOE) for card in UNSHUFFLED_DECK]
    shuffle_shoe(shoe)
    return shoe


def display_shoe(shoe: list[Card]) -> None:
    for card in shoe:
        print(f"""{card.rank.value}{card.suite.value}""")


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
    HAND_LIMIT: int = 4

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
        if self.shoe_top > FORCE_RESHUFFLE:
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


# pdb.set_trace()

bj = BlackJack()
bj.play_game()

# exit(0xDEADBEEF)

print("Done.")
