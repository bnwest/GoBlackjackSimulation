use rand::prelude::*;
use rand_chacha::ChaCha8Rng;

mod cards;
mod rules;

fn main() {
    println!("Hello, world!");

    println!("cards: CardSuiteValue: {:#?}", cards::CardSuiteValue);

    let deck: Vec<cards::Card> = cards::create_unshuffle_deck();
    for card in deck.iter() {
        println!("{card:#?}");
    }

    let mut rng: ChaCha8Rng = ChaCha8Rng::seed_from_u64(42_u64);
    let mut shoe: Vec<cards::Card> = cards::create_shoe(rules::DECKS_IN_SHOE);
    cards::shuffle_shoe(&mut shoe, &mut rng);
    cards::display_shoe(&shoe);
}
