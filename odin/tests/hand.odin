package tests

import "core:testing"
import "core:fmt"

import "../game"
import "../cards"

@(test)
test_hand_outcome :: proc(t: ^testing.T) {
    testing.expect(
        t,
        game.HandOutcome.STAND == game.HandOutcome(0),
        "HandOutcome enum starts at 0"
    )
    testing.expect(
        t,
        game.to_hand_outcome_string(game.HandOutcome.STAND) == game.hand_outcome_string[game.HandOutcome.STAND],
        "HandOutcome to_hand_outcome_string() maps to hand_outcome_string[]"
    )
    testing.expect(
        t,
        game.to_string(game.HandOutcome.STAND) == game.hand_outcome_string[game.HandOutcome.STAND],
        "HandOutcome to_strin()) maps to hand_outcome_string[]"
    )
}

@(test)
test_create_player_hand :: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    testing.expect(
        t,
        hand.from_split == false,
        "create_player_hand() returns correct from_split"
    )
    testing.expect(
        t,
        hand.bet == 100,
        "create_player_hand() returns correct bet"
    )
    testing.expect(
        t,
        hand.outcome == game.HandOutcome.IN_PLAY,
        "create_player_hand() returns correct outcome"
    )
    testing.expect(
        t,
        len(hand.cards) == 0,
        "create_player_hand() returns hand with no cards"
    )
}

@(test)
test_add_card :: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    testing.expect(
        t,
        game.num_cards(&hand) == 0,
        "create_player_hand() returns hand with no cards"
    )
    card1: cards.Card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card1)
    testing.expect(
        t,
        game.num_cards(&hand) == 1,
        "add_card() adds one card"
    )
    card2: cards.Card = cards.Card{
        rank=cards.CardRank.JACK,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card2)
    testing.expect(
        t,
        game.num_cards(&hand) == 2,
        "add_card() adds one card"
    )
}

@(test)
test_free_cards :: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    testing.expect(
        t,
        game.num_cards(&hand) == 0,
        "free_cards() returns hand with no cards"
    )
    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.num_cards(&hand) == 1,
        "free_cards() adding one card adds one card"
    )
    game.free_cards(&hand)
    testing.expect(
        t,
        game.num_cards(&hand) == 0,
        "free_cards() freeing cards frees all cards"
    )
}

@(test)
test_hard_count:: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.hard_count(&hand) == 9,
        "hard_count()"
    )
    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.HEARTS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.hard_count(&hand) == 18,
        "hard_count() returns 9 + 9 = 18"
    )
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.hard_count(&hand) == 19,
        "hard_count() returns 9 + 9 + A = 19"
    )
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.CLUBS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.hard_count(&hand) == 20,
        "hard_count() returns 9 + 9 + A + A = 20"
    )
}

@(test)
test_soft_count:: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.soft_count(&hand) == 9,
        "soft_count() returns 9"
    )
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.soft_count(&hand) == 20,
        "soft_count() returns 9 + A = 20"
    )
    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.HEARTS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.soft_count(&hand) == 19,
        "soft_count() returns 9 + A + 9 = 19"
    )
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.CLUBS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.soft_count(&hand) == 20,
        "soft_count() returns 9 + A + 9 + A = 20"
    )
}

@(test)
test_count:: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.count(&hand) == 9,
        "count() returns 9"
    )
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.count(&hand) == 20,
        "count() returns 9 + A = 20"
    )
    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.HEARTS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.count(&hand) == 19,
        "count() returns 9 + A + 9 = 19"
    )
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.CLUBS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.count(&hand) == 20,
        "count() returns 9 + A + 9 + A = 20"
    )
}

@(test)
test_is_natural:: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.is_natural(&hand),
        "is_natural() returns true for 10 + A"
    )

    game.free_cards(&hand)

    card = cards.Card{
        rank=cards.CardRank.NINE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.is_natural(&hand) == false,
        "is_natural() returns false for 9 + A"
    )
}

@(test)
test_is_bust:: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.CLUBS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.is_bust(&hand) == false,
        "is_bust() returns false for 10 + 10 + A"
    )

    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.HEARTS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.is_bust(&hand),
        "is_bust() returns true for 10 + 10 + A + A"
    )
}

@(test)
test_can_split:: proc(t: ^testing.T) {
    hand := game.create_player_hand(
        from_split=false,
        bet=100,
    )
    defer game.free_cards(&hand)

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    card = cards.Card{
        rank=cards.CardRank.KING,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.can_split(&hand),
        "can_split() returns true for 10 + K"
    )

    game.free_cards(&hand)

    card = cards.Card{
        rank=cards.CardRank.KING,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&hand, card)
    card = cards.Card{
        rank=cards.CardRank.KING,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&hand, card)
    testing.expect(
        t,
        game.can_split(&hand),
        "can_split() returns true for K + K"
    )
}
