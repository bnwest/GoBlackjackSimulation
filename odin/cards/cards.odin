package cards

import house_rules "../rules"

import "core:strings"
import "core:math/rand"

CardSuite :: enum {
    HEARTS,    // == 0
    DIAMONDS,  // == 1
    SPADES,    // == 2
    CLUBS,     // == 3
}

/*
// Error: Compound literals of dynamic types are disabled by default
CardSuiteValue := map[CardSuite]string{
    .HEARTS   = "♥️", // aka U+2665 + U+fe0f
    .DIAMONDS = "♦️", // aka U+2666 + U+fe0f
    .SPADES   = "♠️", // aka U+2660 + U+fe0f
    .CLUBS    = "♣️", // aka U+2663 + U+fe0f
}
*/

card_suite_string := [CardSuite]string {
    .HEARTS   = "♥️", // aka U+2665 + U+fe0f, 2 odin runes
    .DIAMONDS = "♦️", // aka U+2666 + U+fe0f, 2 odin runes
    .SPADES   = "♠️", // aka U+2660 + U+fe0f, 2 odin runes
    .CLUBS    = "♣️", // aka U+2663 + U+fe0f, 2 odin runes
}

suite_to_string :: proc(
    suite: CardSuite
) -> string {
    return strings.concatenate({card_suite_string[suite], " "})
}

CardRank :: enum {
    ACE,   // == 0
    TWO,   // == 1
    THREE, // == 2
    FOUR,  // == 3
    FIVE,  // == 4
    SIX,   // == 5
    SEVEN, // == 6
    EIGHT, // == 7
    NINE,  // == 8
    TEN,   // == 9
    JACK,  // == 10
    QUEEN, // == 11
    KING,  // == 12
}

card_rank_string := [CardRank]string {
    .ACE   = "A",
    .TWO   = "2",
    .THREE = "3",
    .FOUR  = "4",
    .FIVE  = "5",
    .SIX   = "6",
    .SEVEN = "7",
    .EIGHT = "8",
    .NINE  = "9",
    .TEN   = "10",
    .JACK  = "J",
    .QUEEN = "Q",
    .KING  = "K",
}

rank_to_string :: proc(
    rank: CardRank
) -> string {
    return card_rank_string[rank]
}

card_rank_integer := [CardRank]uint {
    .ACE   = 1,
    .TWO   = 2,
    .THREE = 3,
    .FOUR  = 4,
    .FIVE  = 5,
    .SIX   = 6,
    .SEVEN = 7,
    .EIGHT = 8,
    .NINE  = 9,
    .TEN   = 10,
    .JACK  = 10,
    .QUEEN = 10,
    .KING  = 10,
}

rank_to_int :: proc(
    rank: CardRank
) -> uint {
    return card_rank_integer[rank]
}

Card :: struct {
    suite: CardSuite,
    rank: CardRank,
}

cards_strings := [CardRank][CardSuite]string{
    .ACE  ={.HEARTS="A♥️ ",  .DIAMONDS="A♦️ ",  .SPADES="A♠️ ",  .CLUBS="A♣️ "},
    .TWO  ={.HEARTS="2♥️ ",  .DIAMONDS="2♦️ ",  .SPADES="2♠️ ",  .CLUBS="2♣️ "},
    .THREE={.HEARTS="3♥️ ",  .DIAMONDS="3♦️ ",  .SPADES="3♠️ ",  .CLUBS="3♣️ "},
    .FOUR ={.HEARTS="4♥️ ",  .DIAMONDS="4♦️ ",  .SPADES="4♠️ ",  .CLUBS="4♣️ "},
    .FIVE ={.HEARTS="5♥️ ",  .DIAMONDS="5♦️ ",  .SPADES="5♠️ ",  .CLUBS="5♣️ "},
    .SIX  ={.HEARTS="6♥️ ",  .DIAMONDS="6♦️ ",  .SPADES="6♠️ ",  .CLUBS="6♣️ "},
    .SEVEN={.HEARTS="7♥️ ",  .DIAMONDS="7♦️ ",  .SPADES="7♠️ ",  .CLUBS="7♣️ "},
    .EIGHT={.HEARTS="8♥️ ",  .DIAMONDS="8♦️ ",  .SPADES="8♠️ ",  .CLUBS="8♣️ "},
    .NINE ={.HEARTS="9♥️ ",  .DIAMONDS="9♦️ ",  .SPADES="9♠️ ",  .CLUBS="9♣️ "},
    .TEN  ={.HEARTS="10♥️ ", .DIAMONDS="10♦️ ", .SPADES="10♠️ ", .CLUBS="10♣️ "},
    .JACK ={.HEARTS="J♥️ ",  .DIAMONDS="J♦️ ",  .SPADES="J♠️ ",  .CLUBS="J♣️ "},
    .QUEEN={.HEARTS="Q♥️ ",  .DIAMONDS="Q♦️ ",  .SPADES="Q♠️ ",  .CLUBS="Q♣️ "},
    .KING ={.HEARTS="K♥️ ",  .DIAMONDS="K♦️ ",  .SPADES="K♠️ ",  .CLUBS="K♣️ "},
}

card_to_string :: proc(card: Card) -> string {
    // really, really, really do not want to return a heap allocated string
    return cards_strings[card.rank][card.suite]
}

UNSHUFFLED_DECK :: []Card {
    // HEARTS
    Card{suite=CardSuite.HEARTS, rank=CardRank.ACE},
    Card{suite=CardSuite.HEARTS, rank=CardRank.TWO},
    Card{suite=CardSuite.HEARTS, rank=CardRank.THREE},
    Card{suite=CardSuite.HEARTS, rank=CardRank.FOUR},
    Card{suite=CardSuite.HEARTS, rank=CardRank.FIVE},
    Card{suite=CardSuite.HEARTS, rank=CardRank.SIX},
    Card{suite=CardSuite.HEARTS, rank=CardRank.SEVEN},
    Card{suite=CardSuite.HEARTS, rank=CardRank.EIGHT},
    Card{suite=CardSuite.HEARTS, rank=CardRank.NINE},
    Card{suite=CardSuite.HEARTS, rank=CardRank.TEN},
    Card{suite=CardSuite.HEARTS, rank=CardRank.JACK},
    Card{suite=CardSuite.HEARTS, rank=CardRank.QUEEN},
    Card{suite=CardSuite.HEARTS, rank=CardRank.KING},
    // DIAMONDS
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.ACE},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.TWO},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.THREE},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.FOUR},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.FIVE},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.SIX},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.SEVEN},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.EIGHT},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.NINE},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.TEN},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.JACK},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.QUEEN},
    Card{suite=CardSuite.DIAMONDS, rank=CardRank.KING},
    // SPADES
    Card{suite=CardSuite.SPADES, rank=CardRank.ACE},
    Card{suite=CardSuite.SPADES, rank=CardRank.TWO},
    Card{suite=CardSuite.SPADES, rank=CardRank.THREE},
    Card{suite=CardSuite.SPADES, rank=CardRank.FOUR},
    Card{suite=CardSuite.SPADES, rank=CardRank.FIVE},
    Card{suite=CardSuite.SPADES, rank=CardRank.SIX},
    Card{suite=CardSuite.SPADES, rank=CardRank.SEVEN},
    Card{suite=CardSuite.SPADES, rank=CardRank.EIGHT},
    Card{suite=CardSuite.SPADES, rank=CardRank.NINE},
    Card{suite=CardSuite.SPADES, rank=CardRank.TEN},
    Card{suite=CardSuite.SPADES, rank=CardRank.JACK},
    Card{suite=CardSuite.SPADES, rank=CardRank.QUEEN},
    Card{suite=CardSuite.SPADES, rank=CardRank.KING},
    // CLUBS
    Card{suite=CardSuite.CLUBS, rank=CardRank.ACE},
    Card{suite=CardSuite.CLUBS, rank=CardRank.TWO},
    Card{suite=CardSuite.CLUBS, rank=CardRank.THREE},
    Card{suite=CardSuite.CLUBS, rank=CardRank.FOUR},
    Card{suite=CardSuite.CLUBS, rank=CardRank.FIVE},
    Card{suite=CardSuite.CLUBS, rank=CardRank.SIX},
    Card{suite=CardSuite.CLUBS, rank=CardRank.SEVEN},
    Card{suite=CardSuite.CLUBS, rank=CardRank.EIGHT},
    Card{suite=CardSuite.CLUBS, rank=CardRank.NINE},
    Card{suite=CardSuite.CLUBS, rank=CardRank.TEN},
    Card{suite=CardSuite.CLUBS, rank=CardRank.JACK},
    Card{suite=CardSuite.CLUBS, rank=CardRank.QUEEN},
    Card{suite=CardSuite.CLUBS, rank=CardRank.KING},
}

RNG_SEED :: 314159

create_shoe :: proc (
    num_shoes: u32 = house_rules.DECKS_IN_SHOE
) -> [dynamic]Card {
    // updates context.random_generator with a RNG seed
    rand.reset(RNG_SEED)

    shoe: [dynamic]Card
    for i in 0..<num_shoes {
        for card in UNSHUFFLED_DECK {
            append(&shoe, card)
        }
    }

    shuffle_shoe(shoe)

    return shoe
}

shuffle_shoe :: proc(shoe: [dynamic]Card) {
    rand.shuffle(shoe[:])
}

//
// Explicit Procedure Overloading
// really just aliasing a set of existing functions to the same name,
// the existing functions can live within the C ABI, while the code gets
// to act like that there is function overloading.
//

to_string :: proc {
    suite_to_string,
    rank_to_string,
    card_to_string
}

to_int :: proc {
    rank_to_int
}
