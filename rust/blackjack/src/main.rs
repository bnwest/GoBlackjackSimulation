mod cards;

fn main() {
    println!("Hello, world!");

    // println!("{sudoku_game:#?}");
    // println!("cards: CardSuiteValue: {cards::CardSuiteValue:#?}");

    // println!("cards: CardSuites::HEARTS {}", cards::CardSuites::HEARTS);
    println!("cards: CardSuiteValue: {:#?}", cards::CardSuiteValue);

    let deck: Vec<cards::Card> = cards::create_unshuffle_deck();
    for card in deck.iter() {
        println!("{:#?}", card);
    }
}
