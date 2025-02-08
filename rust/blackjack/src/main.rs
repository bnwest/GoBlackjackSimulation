// Lessons learned:
//
// 1. Coing rust is an ongoing stryggle.
// 2. Ownership and borrowing happen implicily.  This is alomost a deal breaker.
// 3. Ownership and borrowing errors are hard to grok.
// 4. Traits to add to types are determined via compiler errors.
// 5. basic enums are integer based.
// 6. to get a working integer enum, you need a lot code.

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
