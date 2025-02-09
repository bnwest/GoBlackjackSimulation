// Lessons learned:
//
// 1. Coing rust is an ongoing stryggle.
// 2. Ownership and borrowing happen implicily.  This is almost a deal breaker.
// 2a. Variables are located either or the stack or on the hap.
// 2b. Gentle coder will need to sort where variable get located on their own.
// 2c. Ownership and borrowing apply only to heap variables.
// 2d. Ownership and borrowing errors are hard to grok.
// 3. Traits to add to types are determined via compiler errors.
// 4. basic enums are integer based.
// 5. to get a working integer enum, you need a lot code.
// 6. Global variable are very hard to initialize at run time (easy at compile time).
// 6a. Had to use the "lazy_static!"" macro, which appears to thread safe.

use rand::prelude::*;
use rand_chacha::ChaCha8Rng;

mod cards;
mod rules;
mod strategy;

fn main() {
    println!("Hello, world!");

    println!("cards: CardSuiteValue: {:#?}", cards::CardSuiteValue);

    let mut rng: ChaCha8Rng = ChaCha8Rng::seed_from_u64(42_u64);
    let mut shoe: Vec<cards::Card> = cards::create_shoe(rules::DECKS_IN_SHOE);
    cards::shuffle_shoe(&mut shoe, &mut rng);
    // cards::display_shoe(&shoe);
}
