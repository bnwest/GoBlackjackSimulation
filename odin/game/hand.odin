package game

import "core:fmt"

import "../cards"
import house_rules "../rules"

HandOutcome  :: enum {
	STAND,   // == 0
    BUST,
    SURRENDER,
    DEALER_BLACKJACK,
    IN_PLAY
}

hand_outcome_string := [HandOutcome]string {
    .STAND            = "stand",
    .BUST             = "bust",
    .SURRENDER        = "surrender",
    .DEALER_BLACKJACK = "dealer-blackjack",
    .IN_PLAY          = "in-play",
}

to_hand_outcome_string :: proc(hand_outcome: HandOutcome) -> string {
    return hand_outcome_string[hand_outcome]
}

//
// PlayerHand
//

PlayerHand :: struct {
    cards: [dynamic]cards.Card, // dynamic => stored on the heap
    from_split: bool,
    bet: uint,
    outcome: HandOutcome,
}

// factory
create_player_hand :: proc(
    from_split: bool,
    bet: uint,
) -> PlayerHand {
    player_hand := PlayerHand{
        cards=[dynamic]cards.Card{},  // NO heap pointer, just zero-ed memory
        from_split=from_split,
        bet=bet,
        outcome=HandOutcome.IN_PLAY,
    }
    // create a copy of the stack variable hand
    // and return the copy
    return player_hand
}

player_add_card :: proc(self: ^PlayerHand, card: cards.Card) {
    append(&self.cards, card)
}

player_reset_cards :: proc(self: ^PlayerHand) {
    clear(&self.cards)
    // self.cards still points to the same heap location 
    // with the same capacity, but the length has been reset to 0
}

player_free_cards :: proc(self: ^PlayerHand) {
    delete(self.cards)
    // self.cards now has a pointer to freed memory aka a dangling pointer
    // since we collectively learn no lessons over time
    self.cards = [dynamic]cards.Card{}
    // self.cards is now zero-ed memory
}

player_num_cards :: proc(self: ^PlayerHand) -> uint {
    return len(self.cards)
}

is_from_split :: proc(self: ^PlayerHand) -> bool {
    return self.from_split
}

get_card :: proc(self: ^PlayerHand, card_index: uint) -> cards.Card {
    return self.cards[card_index]
}

player_aces_count :: proc(self: ^PlayerHand) -> uint {
    count: uint = 0
    for card in self.cards {
        if card.rank == cards.CardRank.ACE {
            count += 1
        }
    }
    return count
}

player_hard_count :: proc(self: ^PlayerHand) -> uint {
    count: uint = 0
    for card in self.cards {
        count += cards.to_int(card.rank)
    }
    return count
}

player_soft_count :: proc(self: ^PlayerHand) -> uint {
	// if the soft count is a bust, we convert the Ace values
	// back to the value of 1, one at a time, until the soft count
	// is no longer a bust or until there are no more Aces
	// and the soft count has become the hard count.
    count: uint = 0
    aces_count: uint = 0
    for card in self.cards {
        if card.rank == cards.CardRank.ACE {
            count += 11
            aces_count += 1
        } else {
            count += cards.to_int(card.rank)
        }
    }
    if count > 21 {
        for i: uint = 0; i < aces_count; i += 1 {
            count -= 10
            if count <= 21 {
                break
            }
        }
    }
    return count
}

player_count :: proc(self: ^PlayerHand) -> uint {
	// return the highest count for hand,
	// which is always the soft count.
    return soft_count(self)
}

player_is_natural :: proc(self: ^PlayerHand) -> bool {
    if !self.from_split {
        if num_cards(self) == 2 {
            if soft_count(self) == 21 {
                return true
            }
        }
    }
    return false
}

player_is_bust :: proc(self: ^PlayerHand) -> bool {
    return count(self) > 21
}

player_can_split :: proc(self: ^PlayerHand) -> bool {
	// there are other split house rules that will be applied
	// at a higher abstraction level ... like splitting aces
	// after a split ...like limiting the number of splits
	// from the original (aka "master") hand.
    if num_cards(self) == 2 {
        card1: cards.Card = self.cards[0]
        card2: cards.Card = self.cards[1]
        if house_rules.SPLIT_ON_VALUE_MATCH {
            if cards.to_int(card1.rank) == cards.to_int(card2.rank) {
                return true
            }
        } else {
            if card1.rank == card2.rank {
                return true
            }
        }
    }
    return false
}

player_is_hand_over :: proc(self: ^PlayerHand) -> bool {
    switch self.outcome {
        case .STAND:
            return true
        case .BUST:
            return true
        case .SURRENDER:
            return true
        case .DEALER_BLACKJACK:
            return true
        case .IN_PLAY:
            return false
        case: 
            // aka default when case has no expression
            // compiler not smart enough to know this was not needed
            return false
    }
}

//
// DealerHand
//

DealerHand :: struct {
    cards: [dynamic]cards.Card, // dynamic => stored on the heap
    outcome: HandOutcome,
}

// factory
create_dealer_hand :: proc() -> DealerHand {
    dealer_hand := DealerHand{
        cards=[dynamic]cards.Card{},  // heap pointer
        outcome=HandOutcome.IN_PLAY,
    }
    // create a copy of the stack variable hand
    // and return the copy
    return dealer_hand
}

dealer_add_card :: proc(self: ^DealerHand, card: cards.Card) {
    append(&self.cards, card)
}

dealer_reset_cards :: proc(self: ^DealerHand) {
    clear(&self.cards)
    // self.cards still points to the same heap location 
    // with the same capacity, but the length has been reset to 0
}

dealer_free_cards :: proc(self: ^DealerHand) {
    delete(self.cards)
    // self.cards now has a pointer to freed memory aka a dangling pointer
    // since we collectively learn no lessons over time
    self.cards = [dynamic]cards.Card{}
    // self.cards is now zero-ed memory
}

dealer_num_cards :: proc(self: ^DealerHand) -> uint {
    return len(self.cards)
}

dealer_aces_count :: proc(self: ^DealerHand) -> uint {
    count: uint = 0
    for card in self.cards {
        if card.rank == cards.CardRank.ACE {
            count += 1
        }
    }
    return count
}

dealer_hard_count :: proc(self: ^DealerHand) -> uint {
    count: uint = 0
    for card in self.cards {
        count += cards.to_int(card.rank)
    }
    return count
}

dealer_soft_count :: proc(self: ^DealerHand) -> uint {
	// if the soft count is a bust, we convert the Ace values
	// back to the value of 1, one at a time, until the soft count
	// is no longer a bust or until there are no more Aces
	// and the soft count has become the hard count.
    count: uint = 0
    aces_count: uint = 0
    for card in self.cards {
        if card.rank == cards.CardRank.ACE {
            count += 11
            aces_count += 1
        } else {
            count += cards.to_int(card.rank)
        }
    }
    if count > 21 {
        for i: uint = 0; i < aces_count; i += 1 {
            count -= 10
            if count <= 21 {
                break
            }
        }
    }
    return count
}

dealer_count :: proc(self: ^DealerHand) -> uint {
	// return the highest count for hand,
	// which is always the soft count.
    return soft_count(self)
}

dealer_is_natural :: proc(self: ^DealerHand) -> bool {
    if num_cards(self) == 2 {
        if soft_count(self) == 21 {
            return true
        }
    }
    return false
}

dealer_is_bust :: proc(self: ^DealerHand) -> bool {
    return count(self) > 21
}

dealer_is_hand_over :: proc(self: ^DealerHand) -> bool {
    switch self.outcome {
        case .STAND:
            return true
        case .BUST:
            return true
        case .SURRENDER:
            return true
        case .DEALER_BLACKJACK:
            return true
        case .IN_PLAY:
            return false
        case: 
            // aka default when case has no expression
            // compiler not smart enough to know this was not needed
            return false
    }
}

//
// PlayerMasterHand
//

PlayerMasterHand :: struct {
    HANDS_LIMIT: uint,
    hands: [dynamic]PlayerHand,
}

// factory
create_player_master_hand :: proc() -> PlayerMasterHand {
    master_hand := PlayerMasterHand{
        HANDS_LIMIT=house_rules.SPLITS_PER_HAND + 1,
        hands=[dynamic]PlayerHand{},
    }
    return master_hand
}

num_hands :: proc(self: ^PlayerMasterHand) -> uint {
    return len(self.hands)
}

add_start_hand :: proc(self: ^PlayerMasterHand, bet: uint) {
    from_split : bool : false

    player_hand: PlayerHand
    player_hand = create_player_hand(
        from_split=from_split, bet=bet
    )

    append(&self.hands, player_hand)
}

reset_hands :: proc(self: ^PlayerMasterHand) {
    for &hand in self.hands {
        // when we append to self.hands the next time,
        // self.hands[i].cards will be overrwritten and 
        // the old heap pointer will be potentially leaked,
        // so self.hands[i].cards will be set free
        free_cards(&hand)
    }
    clear(&self.hands)
}

free_hands :: proc(self: ^PlayerMasterHand) {
    // for hand, hand_index in self.hands {
    //     free_cards(&self.hands[hand_index])
    // }
    for &hand in self.hands {
        free_cards(&hand)
    }
    delete(self.hands)
    // self.hands now has a pointer to free memory
    // since we collectively learn no lessons over time
    self.hands = [dynamic]PlayerHand{}
}

log_hands :: proc(self: ^PlayerMasterHand, preface: string) {
    // log to stderr => using eprintfln()
    fmt.eprintfln("{0}: MasterHand", preface)
    for hand, i in self.hands {
        fmt.eprintfln("    Hand {0}", i+1)
        for card, j in hand.cards {
            card_string := cards.to_string(card)
            fmt.eprintfln("        Card {0}: {1}", j+1, card_string)
            delete(card_string)
        }
    }
}

master_can_split :: proc(self: ^PlayerMasterHand, hand_index: uint) -> bool {
    if num_hands(self) < self.HANDS_LIMIT {
        // master hand allows
        hand := self.hands[hand_index]
        if can_split(&hand) {
            // individual hand allows
            return true
        }
    }
    return false
}

split_hand :: proc(
    self: ^PlayerMasterHand, 
    hand_index: uint,
    cards_to_add: [2]cards.Card,
) -> uint {
	// there are two cards in the hand of the same value
	// or rank depending of the house rules.
	card1: cards.Card = self.hands[hand_index].cards[0]
	card2: cards.Card = self.hands[hand_index].cards[1]

    old_player_hand: ^PlayerHand
    old_player_hand = &self.hands[hand_index]
    delete(old_player_hand.cards)
    old_player_hand.cards = [dynamic]cards.Card{}
    append(&old_player_hand.cards, card1, cards_to_add[0])
    old_player_hand.from_split = true
    old_player_hand.outcome = HandOutcome.IN_PLAY

    new_player_hand: PlayerHand
    new_player_hand = create_player_hand(from_split=true, bet=old_player_hand.bet)
    delete(new_player_hand.cards)
    new_player_hand.cards = [dynamic]cards.Card{}
    append(&new_player_hand.cards, card2, cards_to_add[1])
    new_player_hand.outcome = HandOutcome.IN_PLAY

    new_hand_index := num_hands(self)
    append(&self.hands, new_player_hand)

    return new_hand_index
}

//
// Explicit Procedure Overloading
//

to_string :: proc {
    to_hand_outcome_string,
}

add_card :: proc {
    dealer_add_card,
    player_add_card,
}

reset_cards :: proc {
    dealer_reset_cards,
    player_reset_cards,
}

free_cards :: proc {
    dealer_free_cards,
    player_free_cards,
}

num_cards :: proc {
    dealer_num_cards,
    player_num_cards,
}

aces_count :: proc {
    dealer_aces_count,
    player_aces_count,
}

hard_count :: proc {
    dealer_hard_count,
    player_hard_count,
}

soft_count :: proc {
    dealer_soft_count,
    player_soft_count,
}

count :: proc {
    dealer_count,
    player_count,
}

is_natural :: proc {
    dealer_is_natural,
    player_is_natural,
}

is_bust :: proc {
    dealer_is_bust,
    player_is_bust,
}

is_hand_over :: proc {
    dealer_is_hand_over,
    player_is_hand_over,

}

can_split :: proc {
    player_can_split,
    master_can_split,
}