package tests

import "core:testing"
import "core:fmt"

import "../game"
import "../cards"
import house_rules "../rules"

@(test)
test_memory_cycle_player_hand :: proc(t: ^testing.T) {
    player_hand: game.PlayerHand
    player_hand = game.create_player_hand(
        from_split=false, bet=100,
    )
    defer game.free_cards(&player_hand)

    // player_hand.cards in initialized with
    //     [dynamic]cards.Card{}
    // which does not allocate space on the heap
    // player_hand.cards should be given zero-ed memory
    testing.expectf(
        t,
        len(player_hand.cards) == 0,
        "player_hand.cards is zero length? {}",
        len(player_hand.cards),
    )
    testing.expectf(
        t,
        cap(player_hand.cards) == 0,
        "player_hand.cards has zero capacity? {}",
        cap(player_hand.cards)
    )

    // add a card to the player hand
    // player_hand.cards will get a heap allocation that will need to be freed
    card1: cards.Card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&player_hand, card1)
    testing.expectf(
        t,
        len(player_hand.cards) == 1,
        "player_hand.cards has length of 1? {}",
        len(player_hand.cards),
    )
    testing.expectf(
        t,
        cap(player_hand.cards) > 0,
        "player_hand.cards has nonzero capacity? {}",
        cap(player_hand.cards),
    )

    // clear player_hand.cards
    // player_hand.cards retains heap allocation but length is reset to 0
    game.player_reset_cards(&player_hand)
    testing.expectf(
        t,
        len(player_hand.cards) == 0,
        "player_hand.cards is zero length? {}",
        len(player_hand.cards),
    )
    testing.expectf(
        t,
        cap(player_hand.cards) > 0,
        "player_hand.cards has nonzero capacity? {}",
        cap(player_hand.cards),
    )

    game.add_card(&player_hand, card1)

    // free player_hand.cards
    // player_hand.cards frees heap allocation
    game.free_cards(&player_hand)
    testing.expectf(
        t,
        len(player_hand.cards) == 0,
        "player_hand.cards is zero length? {}",
        len(player_hand.cards),
    )
    testing.expectf(
        t,
        cap(player_hand.cards) == 0,
        "player_hand.cards has zero capacity? {}",
        cap(player_hand.cards)
    )

    // the deferred call to game.free_cards() is run
    // with no heap memory to free ... without issue
}

@(test)
test_memory_cycle_master_hand :: proc(t: ^testing.T) {
    master_hand: game.PlayerMasterHand
    player_hand1: game.PlayerHand
    player_hand2: game.PlayerHand
    card: cards.Card

    master_hand = game.create_player_master_hand()
    testing.expectf(
        t,
        len(master_hand.hands) == 0,
        "master_hand.hands is zero length? {}",
        len(master_hand.hands),
    )
    testing.expectf(
        t,
        cap(master_hand.hands) == 0,
        "master_hand.hands is zero capacity? {}",
        cap(master_hand.hands),
    )

    // add two hands to the master hand
    player_hand1 = game.create_player_hand(
        from_split=false, bet=100,
    )

    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&player_hand1, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.CLUBS,
    }
    game.add_card(&player_hand1, card)

    player_hand2 = game.create_player_hand(
        from_split=false, bet=100,
    )

    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.HEARTS,
    }
    game.add_card(&player_hand2, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&player_hand2, card)

    // not how we add hand to the master hand
    // (which happens when the player splits his hand)
    // but we fake it here
    append(&master_hand.hands, player_hand1)
    append(&master_hand.hands, player_hand2)
    testing.expectf(
        t,
        len(master_hand.hands) == 2,
        "master_hand.hands has nonzero length? {}",
        len(master_hand.hands),
    )
    testing.expectf(
        t,
        cap(master_hand.hands) > 0,
        "master_hand.hands has nonzero capacity? {}",
        cap(master_hand.hands),
    )

    // reset the master hand
    game.reset_hands(&master_hand)
    testing.expectf(
        t,
        len(master_hand.hands) == 0,
        "master_hand.hands has zero length? {}",
        len(master_hand.hands),
    )
    testing.expectf(
        t,
        cap(master_hand.hands) > 0,
        "master_hand.hands has nonzero capacity? {}",
        cap(master_hand.hands),
    )

    // add two hands to the master hand
    player_hand1 = game.create_player_hand(
        from_split=false, bet=100,
    )

    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.SPADES,
    }
    game.add_card(&player_hand1, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.CLUBS,
    }
    game.add_card(&player_hand1, card)

    player_hand2 = game.create_player_hand(
        from_split=false, bet=100,
    )

    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.HEARTS,
    }
    game.add_card(&player_hand2, card)
    card = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&player_hand2, card)

    append(&master_hand.hands, player_hand1)
    append(&master_hand.hands, player_hand2)
    testing.expectf(
        t,
        len(master_hand.hands) == 2,
        "master_hand.hands has nonzero length? {}",
        len(master_hand.hands),
    )
    testing.expectf(
        t,
        cap(master_hand.hands) > 0,
        "master_hand.hands has nonzero capacity? {}",
        cap(master_hand.hands),
    )

    // free master_hand.hands and exit
    game.free_hands(&master_hand)
    testing.expectf(
        t,
        len(master_hand.hands) == 0,
        "master_hand.hands has zero length? {}",
        len(master_hand.hands),
    )
    testing.expectf(
        t,
        cap(master_hand.hands) == 0,
        "master_hand.hands has zero capacity? {}",
        cap(master_hand.hands),
    )
}

//
// Player
//

@(test)
test_create_player :: proc(t: ^testing.T) {
    player: game.Player
    player = game.create_player(name="Scott Free")
    testing.expectf(
        t,
        game.num_master_hands(&player) == 0,
        "player.master_hands has zero length? {}",
        game.num_master_hands(&player),
    )
    testing.expectf(
        t,
        cap(player.master_hands) == 0,
        "player.master_hand has zero capacity? {}",
        cap(player.master_hands),
    )
}

@(test)
test_set_game_bets :: proc(t: ^testing.T) {
    player: game.Player
    player = game.create_player(name="Scott Free")
    defer game.free_player(&player)

    testing.expectf(
        t,
        game.num_master_hands(&player) == 0,
        "player.master_hands has zero length? {}",
        game.num_master_hands(&player),
    )
    testing.expectf(
        t,
        cap(player.master_hands) == 0,
        "player.master_hand has zero capacity? {}",
        cap(player.master_hands),
    )

    //
    // Set the best for the first game
    //

    bets: [dynamic]uint
    append(&bets, 2, 4, 6)
    defer delete(bets)

    game.set_game_bets(&player, bets)
    testing.expectf(
        t,
        game.num_master_hands(&player) == 3,
        "player.master_hands has nonzero length? {}",
        game.num_master_hands(&player),
    )
    testing.expectf(
        t,
        cap(player.master_hands) > 0,
        "player.master_hand has nonzero capacity? {}",
        cap(player.master_hands),
    )

    for &master_hand in player.master_hands {
        testing.expectf(
            t,
            game.num_hands(&master_hand) == 1,
            "master_hand.hands has nonzero length? {}",
            game.num_hands(&master_hand),
        )            
        testing.expectf(
            t,
            cap(master_hand.hands) > 0,
            "master_hand.hands has nonzero capacity? {}",
            cap(master_hand.hands),
        )            
    }

    game.game_reset(&player)
    testing.expectf(
        t,
        game.num_master_hands(&player) == 0,
        "player.master_hands has zero length? {}",
        game.num_master_hands(&player),
    )
    testing.expectf(
        t,
        cap(player.master_hands) > 0,
        "player.master_hand has zero capacity? {}",
        cap(player.master_hands),
    )

    //
    // Set the best for the second game
    //

    bets2: [dynamic]uint
    append(&bets2, 3, 5, 7)
    defer delete(bets2)

    game.set_game_bets(&player, bets2)
    testing.expectf(
        t,
        game.num_master_hands(&player) == 3,
        "player.master_hands has nonzero length? {}",
        game.num_master_hands(&player),
    )
    testing.expectf(
        t,
        cap(player.master_hands) > 0,
        "player.master_hand has nonzero capacity? {}",
        cap(player.master_hands),
    )

    for &master_hand in player.master_hands {
        testing.expectf(
            t,
            game.num_hands(&master_hand) == 1,
            "master_hand.hands has nonzero length? {}",
            game.num_hands(&master_hand),
        )            
        testing.expectf(
            t,
            cap(master_hand.hands) > 0,
            "master_hand.hands has nonzero capacity? {}",
            cap(master_hand.hands),
        )            
    }

    game.game_reset(&player)
    testing.expectf(
        t,
        game.num_master_hands(&player) == 0,
        "player.master_hands has zero length? {}",
        game.num_master_hands(&player),
    )
    testing.expectf(
        t,
        cap(player.master_hands) > 0,
        "player.master_hand has zero capacity? {}",
        cap(player.master_hands),
    )
}

//
// Dealer
//

@(test)
test_create_dealer :: proc(t: ^testing.T) {
    dealer := game.create_dealer()
    defer game.free_dealer(&dealer)

    testing.expectf(
        t,
        len(dealer.dealer_hand.cards) == 0,
        "dealer.dealer_hand.cards has zero length? {}",
        len(dealer.dealer_hand.cards),
    )
    testing.expectf(
        t,
        cap(dealer.dealer_hand.cards) == 0,
        "dealer.dealer_hand.cards has zero capacity? {}",
        cap(dealer.dealer_hand.cards),
    )
}

@(test)
test_dealer_add_card :: proc(t: ^testing.T) {
    dealer := game.create_dealer()
    defer game.free_dealer(&dealer)

    card1: cards.Card
    card1 = cards.Card{
        rank=cards.CardRank.ACE,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&dealer.dealer_hand, card1)

    card2: cards.Card
    card2 = cards.Card{
        rank=cards.CardRank.KING,
        suite=cards.CardSuite.DIAMONDS,
    }
    game.add_card(&dealer.dealer_hand, card2)

    testing.expectf(
        t,
        game.num_cards(&dealer.dealer_hand) == 2,
        "dealer.dealer_hand.cards has 2 cards? {}",
        game.num_cards(&dealer.dealer_hand),
    )

    {
        expected_card1_str := cards.to_string(card1)
        expected_card2_str := cards.to_string(card2)
        got_top_card_str := cards.to_string(game.top_card(&dealer))
        got_hole_card_str := cards.to_string(game.hole_card(&dealer))
        testing.expectf(
            t,
            game.top_card(&dealer) == card1,
            "dealer.top_card is {}? {}",
            expected_card1_str,
            got_top_card_str,
        )
        testing.expectf(
            t,
            game.hole_card(&dealer) == card2,
            "dealer.top_card is {}? {}",
            expected_card2_str,
            got_hole_card_str,
        )
    }

    game.game_reset(&dealer)
    testing.expectf(
        t,
        len(dealer.dealer_hand.cards) == 0,
        "dealer.dealer_hand.cards has zero length? {}",
        len(dealer.dealer_hand.cards),
    )
    testing.expectf(
        t,
        cap(dealer.dealer_hand.cards) > 0,
        "dealer.dealer_hand.cards has nonzero capacity? {}",
        cap(dealer.dealer_hand.cards),
    )
}
