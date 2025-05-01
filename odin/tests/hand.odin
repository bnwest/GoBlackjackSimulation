package tests

import "core:testing"
import "core:fmt"

import "../game"
import "../cards"
import house_rules "../rules"

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
        "HandOutcome to_string()) maps to hand_outcome_string[]"
    )
}

//
// PlayerHand
//

@(test)
test_create_player_hand :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_player_add_card :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_player_free_cards :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_player_hard_count :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_soft_count :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_player_count :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_player_is_natural :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_player_is_bust :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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
test_can_split :: proc(t: ^testing.T) {
    hand: game.PlayerHand
    hand = game.create_player_hand(
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

//
// DealerHand
//

@(test)
test_create_dealer_hand :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
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
test_dealer_hand_add_card :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
    defer game.free_cards(&hand)

    testing.expect(
        t,
        game.num_cards(&hand) == 0,
        "create_dealer_hand() returns hand with no cards"
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
test_dealer_hand_free_cards :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
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
test_dealer_hand_hard_count :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
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
test_dealer_hand_count :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
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
test_dealer_hand_is_natural :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
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
test_dealer_hand_is_bust :: proc(t: ^testing.T) {
    hand: game.DealerHand
    hand = game.create_dealer_hand()
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

//
// PlayerMasterHand
//

@(test)
test_create_player_master_hand :: proc(t: ^testing.T) {
    master_hand := game.create_player_master_hand()
    testing.expect(
        t,
        len(master_hand.hands) == 0,
        "create_player_master_hand() returns empty master hand"
    )
    testing.expect(
        t,
        master_hand.HANDS_LIMIT == house_rules.SPLITS_PER_HAND + 1,
        "create_player_master_hand() returns master hand with correct hands limit"
    )
    game.log_hands(&master_hand, "testing 1 2 3")
}

@(test)
test_player_master_hand_add_start_hand :: proc(t: ^testing.T) {
    master_hand := game.create_player_master_hand()

    game.add_start_hand(&master_hand, bet=100)
    defer game.free_hands(&master_hand)
    testing.expect(
        t,
        game.num_hands(&master_hand) == 1,
        "add_start_hand() returns master hand with one hand"
    )

    card: cards.Card
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&master_hand.hands[0], card)
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&master_hand.hands[0], card)
    defer game.free_cards(&master_hand.hands[0])

    game.log_hands(&master_hand, "testing 1 2 3")
}

@(test)
test_player_master_hand_can_split :: proc(t: ^testing.T) {
    master_hand := game.create_player_master_hand()

    game.add_start_hand(&master_hand, bet=100)
    defer game.free_hands(&master_hand)

    card: cards.Card

    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&master_hand.hands[0], card)
    card = cards.Card{
        rank=cards.CardRank.TEN,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&master_hand.hands[0], card)
    defer game.free_cards(&master_hand.hands[0])

    hand_index: uint

    hand_index = 0
    testing.expect(
        t,
        game.can_split(&master_hand, hand_index),
        "can_split() returns true as expected"
    )

    master_hand.hands[0].cards[0].rank = cards.CardRank.NINE
    hand_index = 0
    testing.expect(
        t,
        game.can_split(&master_hand, hand_index) == false,
        "can_split() returns false as expected"
    )
}

@(test)
test_player_master_hand_split_hand :: proc(t: ^testing.T) {
    master_hand: game.PlayerMasterHand
    master_hand = game.create_player_master_hand()

    bet: uint = 2

    card1: cards.Card
    card2: cards.Card
    new_card1: cards.Card
    new_card2: cards.Card
    cards_to_add: [2]cards.Card
    hand_index: uint
    new_hand_index: uint

    for rank in cards.CardRank {
		// reset master hand back to have no hands
        game.reset_hands(&master_hand)

        // add pair to start hand in the master hand
        game.add_start_hand(&master_hand, bet)
        defer game.free_hands(&master_hand)

        card1 = cards.Card{suite=cards.CardSuite.HEARTS, rank=rank}
        card2 = cards.Card{suite=cards.CardSuite.SPADES, rank=rank}

		// have a card pair to split to the first hand
        append(&master_hand.hands[0].cards, card1, card2)

		//
		// SPLIT #1:
		// hand[0] [A♥️, A♠️]
		// splits into
		// hand[0] [A♥️, A♦️] and hand[1] [A♠️, A♣️]
		//

		// create two new cards to add second to each split hand
        new_card1 = cards.Card{suite=cards.CardSuite.DIAMONDS, rank=rank}
        new_card2 = cards.Card{suite=cards.CardSuite.CLUBS, rank=rank}
        cards_to_add = [2]cards.Card{new_card1, new_card2}

        hand_index = 0
        testing.expect(
            t,
            game.can_split(&master_hand, hand_index),
            "Hand should be split-able"
        )

        new_hand_index = game.split_hand(&master_hand, hand_index, cards_to_add)
        testing.expect(
            t,
            game.num_hands(&master_hand) == 2,
            "Master Hand should now have 2 hands"
        )
        testing.expect(
            t,
            new_hand_index == 1,
            "New split hand got added as expected"
        )

		//
		// SPLIT #2:
		// hand[0] [A♥️, A♦️]
		// splits into
		// hand[0] [A♥️, A♣️] and hand[2] [A♦️, A♠️]
		//

		// create two new cards to add second to each split hand
        new_card1 = cards.Card{suite=cards.CardSuite.CLUBS, rank=rank}
        new_card2 = cards.Card{suite=cards.CardSuite.SPADES, rank=rank}
        cards_to_add = [2]cards.Card{new_card1, new_card2}

        // game.log_hands(&master_hand, "after first split")
        hand_index = 0
        testing.expect(
            t,
            game.can_split(&master_hand, hand_index),
            "Hand should be split-able"
        )

        new_hand_index = game.split_hand(&master_hand, hand_index, cards_to_add)
        testing.expect(
            t,
            game.num_hands(&master_hand) == 3,
            "Master Hand should now have 2 hands"
        )
        testing.expect(
            t,
            new_hand_index == 2,
            "New split hand got added as expected"
        )

		//
		// SPLIT #3:
		// hand[1] [A♠️, A♣️]
		// splits into
		// hand[1] [A♠️, A♦️] and hand[3] [A♣️, A♥️]
		//

		// create two new cards to add second to each split hand
        new_card1 = cards.Card{suite=cards.CardSuite.DIAMONDS, rank=rank}
        new_card2 = cards.Card{suite=cards.CardSuite.HEARTS, rank=rank}
        cards_to_add = [2]cards.Card{new_card1, new_card2}

        hand_index = 1
        testing.expect(
            t,
            game.can_split(&master_hand, hand_index),
            "Hand should be split-able"
        )

        new_hand_index = game.split_hand(&master_hand, hand_index, cards_to_add)
        testing.expect(
            t,
            game.num_hands(&master_hand) == 4,
            "Master Hand should now have 2 hands"
        )
        testing.expect(
            t,
            new_hand_index == 3,
            "New split hand got added as expected"
        )

		//
		// A master should only be able to be split 3 times
		// => a single master hand can turn into no more than four hands
		//

        testing.expect(
            t,
            game.can_split(&master_hand, hand_index=0) == false,
            "Hand should not be split-able"
        )
        testing.expect(
            t,
            game.can_split(&master_hand, hand_index=1) == false,
            "Hand should not be split-able"
        )
        testing.expect(
            t,
            game.can_split(&master_hand, hand_index=2) == false,
            "Hand should not be split-able"
        )
        testing.expect(
            t,
            game.can_split(&master_hand, hand_index=3) == false,
            "Hand should not be split-able"
        )

		//
		// Verify that thee four hand in the master hands have the expected cards.
		//

		// masterHand.Hands[0]  [A♥️, A♣️]
        {
            expected_card1 := cards.Card{cards.CardSuite.HEARTS, rank}
            expected_card2 := cards.Card{cards.CardSuite.CLUBS, rank}
            expected_card1_str := cards.to_string(expected_card1)
            expected_card2_str := cards.to_string(expected_card2)
            got_card1_str := cards.to_string(master_hand.hands[0].cards[0])
            got_card2_str := cards.to_string(master_hand.hands[0].cards[1])
            testing.expectf(
                t,
                master_hand.hands[0].cards[0] == expected_card1,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
            testing.expectf(
                t,
                master_hand.hands[0].cards[1] == expected_card2,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
        }

        // masterHand.Hands[1]  [A♠️, A♦️]
        {
            expected_card1 := cards.Card{cards.CardSuite.SPADES, rank}
            expected_card2 := cards.Card{cards.CardSuite.DIAMONDS, rank}
            expected_card1_str := cards.to_string(expected_card1)
            expected_card2_str := cards.to_string(expected_card2)
            got_card1_str := cards.to_string(master_hand.hands[1].cards[0])
            got_card2_str := cards.to_string(master_hand.hands[1].cards[1])
            testing.expectf(
                t,
                master_hand.hands[1].cards[0] == expected_card1,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
            testing.expectf(
                t,
                master_hand.hands[1].cards[1] == expected_card2,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
        }

        // masterHand.Hands[2]  [A♦️, A♠️]
        {
            expected_card1 := cards.Card{cards.CardSuite.DIAMONDS, rank}
            expected_card2 := cards.Card{cards.CardSuite.SPADES, rank}
            expected_card1_str := cards.to_string(expected_card1)
            expected_card2_str := cards.to_string(expected_card2)
            got_card1_str := cards.to_string(master_hand.hands[2].cards[0])
            got_card2_str := cards.to_string(master_hand.hands[2].cards[1])
            testing.expectf(
                t,
                master_hand.hands[2].cards[0] == expected_card1,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
            testing.expectf(
                t,
                master_hand.hands[2].cards[1] == expected_card2,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
        }

        // masterHand.Hands[3]  [A♣️, A♥️]
        {
            expected_card1 := cards.Card{cards.CardSuite.CLUBS, rank}
            expected_card2 := cards.Card{cards.CardSuite.HEARTS, rank}
            expected_card1_str := cards.to_string(expected_card1)
            expected_card2_str := cards.to_string(expected_card2)
            got_card1_str := cards.to_string(master_hand.hands[3].cards[0])
            got_card2_str := cards.to_string(master_hand.hands[3].cards[1])
            testing.expectf(
                t,
                master_hand.hands[3].cards[0] == expected_card1,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
            testing.expectf(
                t,
                master_hand.hands[3].cards[1] == expected_card2,
                "Hand[0] expected [ {0}, {1} ] got [ {2}, {3} ].",
                expected_card1_str,
                expected_card2_str,
                got_card1_str,
                got_card2_str,
            )
        }
    }
}
