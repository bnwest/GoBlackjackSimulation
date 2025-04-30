package game

import "core:fmt"

import "../cards"
import house_rules "../rules"

//
// Player
//

// a player can play a set of hands.
// each indivdual master hand can be split into more hands,
// for which there is hard limit.  each master hand can typically be split
// up to three times, for a total of four hands starting from the master hand.

Player :: struct {
    master_hands: [dynamic]PlayerMasterHand,
    name: string,
}

create_player :: proc(name: string) -> Player {
    player := Player{
        master_hands=[dynamic]PlayerMasterHand{},
        name=name,
    }
    return player
    // match with:
    //     defer free_player(&player)
}

free_player :: proc(self: ^Player) {
    delete(self.master_hands)
    self.master_hands = [dynamic]PlayerMasterHand{}
}

num_master_hands :: proc(self: ^Player) -> uint {
    return len(self.master_hands)
}

//
// Reset paradigm for nested dynamic array:
//    1. iterate over the allocate heap space, freeing all pointers therein
//    2. "clear" the allocate heap space => reset the "top" of the allocated space
//    3. the allocated heap space can now be re-used without leaking nested allocations
//

player_game_reset :: proc(self: ^Player) {
    for &master_hand in self.master_hands {
        free_hands(&master_hand)
    }
    clear(&self.master_hands)
}

set_game_bets :: proc(self: ^Player, bets: [dynamic]uint) {
    game_reset(self)
    for bet, i in bets {
        master_hand: PlayerMasterHand
        master_hand = create_player_master_hand()
        add_start_hand(&master_hand, bet)
        append(&self.master_hands, master_hand)
    }
}

//
// Dealer
//

Dealer :: struct {
    dealer_hand: DealerHand,
    name: string,
}

create_dealer :: proc(name: string = "Riverboat Dealer") -> Dealer {
    dealer := Dealer{
        dealer_hand=create_dealer_hand(),
        name=name,
    }
    return dealer
    // match with:
    //     defer free_dealer(&dealer)
}

free_dealer :: proc(self: ^Dealer) {
    dealer_free_cards(&self.dealer_hand)
}

dealer_game_reset :: proc(self: ^Dealer) {
    dealer_reset_cards(&self.dealer_hand)
}

top_card :: proc(self: ^Dealer) -> cards.Card {
    return self.dealer_hand.cards[0]
}

hole_card :: proc(self: ^Dealer) -> cards.Card {
    return self.dealer_hand.cards[1]
}

//
// Explicit Procedure Overloading
//

game_reset :: proc {
    dealer_game_reset,
    player_game_reset,
}
