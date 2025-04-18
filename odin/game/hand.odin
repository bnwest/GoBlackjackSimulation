package game

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
        cards=[dynamic]cards.Card{},  // heap pointer
        from_split=from_split,
        bet=bet,
        outcome=HandOutcome.IN_PLAY,
    }
    // create a copy of the stack variable hand
    // and return the copy
    return player_hand
}

add_card :: proc(self: ^PlayerHand, card: cards.Card) {
    append(&self.cards, card)
}

free_cards :: proc(self: ^PlayerHand) {
    delete(self.cards)
    // self.cards now has a pointer to free memory
    // since we collectively learn no lessons over time
    self.cards = [dynamic]cards.Card{}
}

num_cards :: proc(self: ^PlayerHand) -> uint {
    return len(self.cards)
}

is_from_split :: proc(self: ^PlayerHand) -> bool {
    return self.from_split
}

get_card :: proc(self: ^PlayerHand, card_index: uint) -> cards.Card {
    return self.cards[card_index]
}

aces_count :: proc(self: ^PlayerHand) -> uint {
    count: uint = 0
    for card in self.cards {
        if card.rank == cards.CardRank.ACE {
            count += 1
        }
    }
    return count
}

hard_count :: proc(self: ^PlayerHand) -> uint {
    count: uint = 0
    for card in self.cards {
        count += cards.to_int(card.rank)
    }
    return count
}

soft_count :: proc(self: ^PlayerHand) -> uint {
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

count :: proc(self: ^PlayerHand) -> uint {
	// return the highest count for hand,
	// which is always the soft count.
    return soft_count(self)
}

is_natural :: proc(self: ^PlayerHand) -> bool {
    if !self.from_split {
        if num_cards(self) == 2 {
            if soft_count(self) == 21 {
                return true
            }
        }
    }
    return false
}

is_bust :: proc(self: ^PlayerHand) -> bool {
    return count(self) > 21
}

can_split :: proc(self: ^PlayerHand) -> bool {
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

is_hand_over :: proc(self: ^PlayerHand) -> bool {
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
            // aka default when case hs no expression
            // compiler not smart enough to now this was not needed
            return false
    }
}

//
// Explicit Procedure Overloading
//

to_string :: proc {
    to_hand_outcome_string,
}
