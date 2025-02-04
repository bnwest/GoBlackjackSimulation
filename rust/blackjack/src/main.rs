use rand::prelude::*;
use rand_chacha::ChaCha8Rng;

mod cards;

fn main() {
    println!("Hello, world!");

    println!("cards: CardSuiteValue: {:#?}", cards::CardSuiteValue);

    let deck: Vec<cards::Card> = cards::create_unshuffle_deck();
    for card in deck.iter() {
        println!("{card:#?}");
    }

    let num_decks_in_shoe: usize = 6;
    let mut rng: ChaCha8Rng = ChaCha8Rng::seed_from_u64(42);
    let mut shoe: Vec<cards::Card> = cards::create_shoe(num_decks_in_shoe);
    cards::shuffle_shoe(&mut shoe, &mut rng);
    cards::display_shoe(&shoe);
}
