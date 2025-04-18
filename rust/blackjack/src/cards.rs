// file src/cards.rs defines project module "cards".

// Had to add the following to Cargo.toml
//    [dependencies]
//    lazy_static = "1.5.0"
//    rand = "0.9.0"
//    rand_chacha = "0.9.0"

use lazy_static::lazy_static;

// two peas in a pod
use rand::prelude::*;
use rand_chacha::ChaCha8Rng;

use std::collections::HashMap;
use std::fmt;
use std::mem::transmute;
use std::slice::Iter;

#[derive(Eq, PartialEq, Hash, Copy, Clone)]
#[repr(u8)]
pub enum CardSuite {
    HEARTS = 1,
    DIAMONDS,
    SPADES,
    CLUBS,
}

impl CardSuite {
    pub fn iterator() -> Iter<'static, CardSuite> {
        static SUITES: [CardSuite; 4] = [
            CardSuite::HEARTS,
            CardSuite::DIAMONDS,
            CardSuite::SPADES,
            CardSuite::CLUBS,
        ];
        SUITES.iter()
    }
    pub fn discriminant(&self) -> u8 {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as u8
    }
    pub fn transmute(discrim: u8) -> CardSuite {
        // FAILS: rank = CardRank(2);
        // FAILS: rank = 2 as CardRank;
        // WORKS: rank = unsafe { transmute(2 as u8) };
        unsafe { transmute(discrim) }
        // why is this not recursive?
    }
    pub fn to_string(&self) -> String {
        static STRINGS: [&str; 5] = [
            "¯\\_(ツ)_/¯", // this is not a valid CardRank
            "♥️",          // CardSuite::HEARTS
            "♦️",          // CardSuite::DIAMONDS
            "♠️",          // CardSuite::SPADES
            "♣️",          // CardSuite::CLUBS
        ];
        STRINGS[self.discriminant() as usize].to_string()
    }
}

impl fmt::Debug for CardSuite {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        return write!(f, "{}", self.to_string());
    }
}

/*
wanting:
var CardSuiteValue = map[CardSuite]string{
    HEARTS:   "♥️", // aka U+2665 + U+fe0f
    DIAMONDS: "♦️", // aka U+2666 + U+fe0f
    SPADES:   "♠️", // aka U+2660 + U+fe0f
    CLUBS:    "♣️", // aka U+2663 + U+fe0f
}

a cautionary tale or hilarity ensues ...
*/

// https://rust-lang-nursery.github.io/rust-cookbook/mem/global_static.html
// https://stackoverflow.com/questions/19605132/is-it-possible-to-use-global-variables-in-rust

// Declares a lazily evaluated constant HashMap. The HashMap will be evaluated once and
// stored behind a global static reference => thread safe

lazy_static! {
    // HashMap documentation:
    // https://doc.rust-lang.org/std/collections/struct.HashMap.html

    // &'static str
    // https://users.rust-lang.org/t/quick-question-static-str/35940/2
    //
    // "'" is a lifetime notation prefix.
    // "'static" represents the lifetime of the process itself
    // "&'static" means a reference that cannot be dangling
    //
    // &'static str is a reference to the UTF-8 encoded variable length of byte sequence,
    // which is valid for the entire lifetime of the process

    pub static ref CardSuiteValue: HashMap<CardSuite, &'static str> = {
        let map = HashMap::from([
            (CardSuite::HEARTS,   "♥️"), // aka "\u{2665}\u{fe0f}" ('\u{fe0f' is the combining mark char)
            (CardSuite::DIAMONDS, "♦️"), // aka "\u{2666}\u{fe0f}"
            (CardSuite::SPADES,   "♠️"), // aka "\u{2660}\u{fe0f}"
            (CardSuite::CLUBS,    "♣️"), // aka "\u{2663}\u{fe0f}"
        ]);
        map
    };
}

// #[derive(Debug)] // adding so pretty print will work ... {:#?} for pretty-print
// manually implement the Debug trait
// https://stackoverflow.com/questions/22243527/how-to-implement-a-custom-fmtdebug-trait

impl fmt::Debug for CardSuiteValue {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        return write!(
            f,
            "{{ HEARTS: {}, DIAMONDS: {}, SPADES: {}, CLUBS: {} }}",
            CardSuiteValue[&CardSuite::HEARTS],
            CardSuiteValue[&CardSuite::DIAMONDS],
            CardSuiteValue[&CardSuite::SPADES],
            CardSuiteValue[&CardSuite::CLUBS],
        );
    }
}

#[derive(Copy, Clone, PartialEq, Eq, Hash)]
#[repr(usize)]
pub enum CardRank {
    ACE = 1,
    TWO = 2,
    THREE = 3,
    FOUR = 4,
    FIVE = 5,
    SIX = 6,
    SEVEN = 7,
    EIGHT = 8,
    NINE = 9,
    TEN = 10,
    JACK = 11,
    QUEEN = 12,
    KING = 13,
}

impl CardRank {
    pub fn iterator() -> Iter<'static, CardRank> {
        static SUITES: [CardRank; 13] = [
            CardRank::ACE,
            CardRank::TWO,
            CardRank::THREE,
            CardRank::FOUR,
            CardRank::FIVE,
            CardRank::SIX,
            CardRank::SEVEN,
            CardRank::EIGHT,
            CardRank::NINE,
            CardRank::TEN,
            CardRank::JACK,
            CardRank::QUEEN,
            CardRank::KING,
        ];
        SUITES.iter()
    }
    // https://doc.rust-lang.org/reference/items/enumerations.html
    // Each enum instance has a discriminant: an integer logically associated to it
    // that is used to determine which variant it holds.
    pub fn discriminant(&self) -> usize {
        // https://doc.rust-lang.org/std/mem/fn.discriminant.html
        // fn returns the integer discriminat for the enum
        // *self as u8
        unsafe { *<*const _>::from(self).cast::<usize>() }
    }
    pub fn transmute(discrim: usize) -> CardRank {
        unsafe { transmute(discrim) }
    }
    pub fn to_string(&self) -> String {
        static STRINGS: [&str; 14] = [
            "bad dog", // this is not a valid CardRank
            "A",       // CardRank::ACE
            "2",       // CardRank::TWO
            "3",       // CardRank::THREE
            "4",       // CardRank::FOUR
            "5",       // CardRank::FIVE
            "6",       // CardRank::SIX
            "7",       // CardRank::SEVEN
            "8",       // CardRank::EIGHT
            "9",       // CardRank::NINE
            "10",      // CardRank::TEN
            "J",       // CardRank::JACK
            "Q",       // CardRank::QUEEN
            "K",       // CardRank::KING
        ];
        STRINGS[self.discriminant()].to_string()
    }
    pub fn value(&self) -> usize {
        // for counting purposes, the value of the card.
        static VALUES: [usize; 14] = [
            0,  // this is not a valid CardRank
            1,  // CardRank::ACE
            2,  // CardRank::TWO
            3,  // CardRank::THREE
            4,  // CardRank::FOUR
            5,  // CardRank::FIVE
            6,  // CardRank::SIX
            7,  // CardRank::SEVEN
            8,  // CardRank::EIGHT
            9,  // CardRank::NINE
            10, // CardRank::TEN
            10, // CardRank::JACK
            10, // CardRank::QUEEN
            10, // CardRank::KING
        ];
        VALUES[self.discriminant()]
    }
}

// this is no bueno. fmt::Result aint what you think?
impl fmt::Debug for CardRank {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        return write!(f, "{}", self.to_string());
    }
}

#[derive(Copy, Clone, PartialEq)]
pub struct Card {
    pub suite: CardSuite,
    pub rank: CardRank,
}

impl Card {
    pub fn to_string(&self) -> String {
        let suite_str: String = self.suite.to_string();
        let rank_str: String = self.rank.to_string();
        let card_str = format!("{suite_str} {rank_str}");
        return card_str;
    }
}

impl fmt::Debug for Card {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        return write!(f, "{}", self.to_string());
    }
}

pub fn create_unshuffle_deck() -> Vec<Card> {
    // should have been: static deck ...
    // but static declard variables can not have run time allocations
    let deck: Vec<Card> = vec![
        // HEARTS
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::ACE,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::TWO,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::THREE,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::FOUR,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::FIVE,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::SIX,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::SEVEN,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::EIGHT,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::NINE,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::TEN,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::JACK,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::QUEEN,
        },
        Card {
            suite: CardSuite::HEARTS,
            rank: CardRank::KING,
        },
        // DIAMONDS
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::ACE,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::TWO,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::THREE,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::FOUR,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::FIVE,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::SIX,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::SEVEN,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::EIGHT,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::NINE,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::TEN,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::JACK,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::QUEEN,
        },
        Card {
            suite: CardSuite::DIAMONDS,
            rank: CardRank::KING,
        },
        // SPADES
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::ACE,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::TWO,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::THREE,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::FOUR,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::FIVE,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::SIX,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::SEVEN,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::EIGHT,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::NINE,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::TEN,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::JACK,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::QUEEN,
        },
        Card {
            suite: CardSuite::SPADES,
            rank: CardRank::KING,
        },
        // CLUBS
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::ACE,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::TWO,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::THREE,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::FOUR,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::FIVE,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::SIX,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::SEVEN,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::EIGHT,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::NINE,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::TEN,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::JACK,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::QUEEN,
        },
        Card {
            suite: CardSuite::CLUBS,
            rank: CardRank::KING,
        },
    ];
    return deck;
}

pub fn create_shoe(decks_in_shoe: usize /*= 6*/) -> Vec<Card> {
    let mut shoe = vec![];
    let unshuffle_deck: Vec<Card> = create_unshuffle_deck();
    for _i in 0..decks_in_shoe {
        for card in unshuffle_deck.iter() {
            // card: &Card
            // *card deferences and copies it.
            shoe.push(*card);
        }
    }
    return shoe;
}

/*
Note to future self:

Given (Vec<Card> is allocated on the heap, not the stack)
pub fn shuffle_shoe(cards: Vec<Card>) {...}
and the call
    let mut shoe : Vec<Card> = create_shoe(6);
    shuffle_shoe(shoe);
the variable "shoe" has its ownership transfered to fn paramater "cards", for the call.
the parameter "cards" goes out of scope at the end of the fn and its heap is dropped/freed.
shuffle_shoe() could return the "cards" paramter to the caller, but this does not scale.

what we want to do is send a reference to "shoes" to the fn.  the fn will then "borrow" the
value for its duration and will not drop the value at the end of scope,
since it does not have ownership.
this looks like:
pub fn shuffle_shoe(cards: &Vec<Card>) {...}
and
    shuffle_shoe(&shoe);

by default, borrowing does not allow mutation (like "let").  we need mutation for the fn.

to allow the value to be borrowe and mutated, we need the following:
pub fn shuffle_shoe(cards: &mut Vec<Card>) {...}
and
    shuffle_shoe(&mut shoe);
the "mut" is associated with the reference "&" (like "let").
rust only allows one mutible reference per value a time, to prevent race condition.
*/

pub fn shuffle_shoe(shoe: &mut Vec<Card>, rng: &mut ChaCha8Rng) {
    shoe.shuffle(rng);
}

pub fn display_shoe(shoe: &Vec<Card>) {
    for card in shoe.iter() {
        // card: &Card
        println!("{card:#?}");
    }
}

#[cfg(test)]
mod tests;
