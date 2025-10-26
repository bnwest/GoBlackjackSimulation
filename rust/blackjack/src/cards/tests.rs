// file src/cards/tests.rs defines project module "cards::tests".

use super::*;

#[test]
fn test_card_suites() {
    let mut num_suite: u32 = 0;
    for _suite in CardSuite::iterator() {
        num_suite += 1
    }
    assert_eq!(num_suite, 4);

    assert_eq!(CardSuite::HEARTS as usize, 1);
    assert_eq!(CardSuite::DIAMONDS as usize, 2);
    assert_eq!(CardSuite::SPADES as usize, 3);
    assert_eq!(CardSuite::CLUBS as usize, 4);

    let mut suite: CardSuite = CardSuite::HEARTS;
    assert_eq!(suite.discriminant(), 1);
    suite = CardSuite::DIAMONDS;
    assert_eq!(suite.discriminant(), 2);
    suite = CardSuite::SPADES;
    assert_eq!(suite.discriminant(), 3);
    suite = CardSuite::CLUBS;
    assert_eq!(suite.discriminant(), 4);

    for suite in CardSuite::iterator() {
        let roundtrip_suite: CardSuite = CardSuite::transmute(suite.discriminant());
        assert_eq!(suite, &roundtrip_suite);
        println!("{}", suite.to_string());
        println!("{:?}", suite);
        println!("{:#?}", suite);
        println!("{suite:?}");
        println!("{suite:#?}");
    }
}

#[test]
fn test_card_suite_value() {
    assert_eq!(CardSuiteValue[&CardSuite::HEARTS], "♥️");
    assert_eq!(CardSuiteValue[&CardSuite::DIAMONDS], "♦️");
    assert_eq!(CardSuiteValue[&CardSuite::SPADES], "♠️");
    assert_eq!(CardSuiteValue[&CardSuite::CLUBS], "♣️");

    println!("cards: CardSuiteValue: {:#?}", CardSuiteValue);
    println!("cards: CardSuiteValue: {CardSuiteValue:#?}");
}

#[test]
fn test_card_rank() {
    let mut num_rank = 0;
    for _rank in CardRank::iterator() {
        num_rank += 1;
    }
    assert_eq!(num_rank, 13);

    assert_eq!(CardRank::ACE as usize, 1);
    assert_eq!(CardRank::KING as usize, 13);

    let mut rank: CardRank = CardRank::KING;
    assert_eq!(rank.discriminant(), 13);

    rank = CardRank::ACE;
    assert_eq!(rank.discriminant(), 1);

    assert_eq!(CardRank::transmute(2), CardRank::TWO);
    for rank in CardRank::iterator() {
        let roundtrip_rank: CardRank = CardRank::transmute(rank.discriminant());
        assert_eq!(rank, &roundtrip_rank);
    }

    for rank in CardRank::iterator() {
        println!("{}", rank.to_string());
        println!("{:?}", rank);
        println!("{:#?}", rank);
        println!("{rank:?}");
        println!("{rank:#?}");
    }

    for rank in CardRank::iterator() {
        let value: usize = rank.value();
        let discrim: u8 = rank.discriminant();
        if 1 <= discrim && discrim <= 10 {
            assert_eq!(value, discrim as usize);
        } else {
            assert_eq!(value, 10);
        }
    }
}

#[test]
fn test_card() {
    let card = Card {
        suite: CardSuite::SPADES,
        rank: CardRank::ACE,
    };
    println!("{:#?}", card);
    println!("{card:#?}");

    let card_str = card.to_string();
    println!("{}", card.to_string());
    assert_eq!(card_str, "♠️ A");

    let mut card2 = card;
    card2.suite = CardSuite::HEARTS;
    assert_eq!(card.suite, CardSuite::SPADES);
    assert_eq!(card2.suite, CardSuite::HEARTS);
}

#[test]
fn test_create_unshuffle_deck() {
    let mut deck: Vec<Card> = create_unshuffle_deck();

    for card in deck.iter() {
        // card: &Card
        println!("{}", card.to_string());
        println!("{:#?}", card);
        println!("{:#?}", *card);
        println!("{card:#?}");

        let card_copy: Card = *card;
        println!("{}", card_copy.to_string());
        println!("{:#?}", card_copy);
        println!("{:#?}", &card_copy);
        println!("{card_copy:#?}");
    }

    for suite in CardSuite::iterator() {
        let mut card_in_suite_count: u32 = 0;
        for card in deck.iter() {
            if card.suite == *suite {
                card_in_suite_count += 1;
            }
        }
        // are there 13 hearts in one deck?
        assert_eq!(card_in_suite_count, 13);
    }

    for rank in CardRank::iterator() {
        let mut num_cards = 0;
        for card in deck.iter() {
            if card.rank == *rank {
                num_cards += 1;
            }
        }
        // are there 4 aces in one deck?
        assert_eq!(num_cards, 4);
    }

    let card1 = deck[42];
    let card2 = deck[21];
    assert_eq!(card1, deck[42]);
    assert_eq!(card2, deck[21]);

    deck[42] = card2;
    deck[21] = card1;
    assert_eq!(card2, deck[42]);
    assert_eq!(card1, deck[21]);
}

#[test]
fn test_create_shoe() {
    let num_decks_in_shoe: usize = 6;
    let shoe: Vec<Card> = create_shoe(num_decks_in_shoe);

    for suite in CardSuite::iterator() {
        let mut card_in_suite_count: usize = 0;
        for card in shoe.iter() {
            if card.suite == *suite {
                card_in_suite_count += 1;
            }
        }
        // are there 13 hearts in each deck in the shoe?
        assert_eq!(card_in_suite_count, 13 * num_decks_in_shoe);
    }

    for rank in CardRank::iterator() {
        let mut num_cards: usize = 0;
        for card in shoe.iter() {
            if card.rank == *rank {
                num_cards += 1;
            }
        }
        // are there 4 aces in each deck in the shoe?
        assert_eq!(num_cards, 4 * num_decks_in_shoe);
    }

    display_shoe(&shoe);
}

#[test]
fn test_shuffle_shoe() {
    let num_decks_in_shoe: usize = 6;
    let mut rng: ChaCha8Rng = ChaCha8Rng::seed_from_u64(42);
    let mut shoe: Vec<Card> = create_shoe(num_decks_in_shoe);

    let mut card1: Card = shoe[41];
    let mut card2: Card = shoe[42];
    assert_eq!(card1, shoe[41]);
    assert_eq!(card2, shoe[42]);

    shoe[41] = card2;
    shoe[42] = card1;
    assert_eq!(card1, shoe[42]);
    assert_eq!(card2, shoe[41]);

    shuffle_shoe(&mut shoe, &mut rng);

    card1 = shoe[41];
    card2 = shoe[42];
    assert_eq!(card1, shoe[41]);
    assert_eq!(card2, shoe[42]);

    shoe[41] = card2;
    shoe[42] = card1;
    assert_eq!(card1, shoe[42]);
    assert_eq!(card2, shoe[41]);

    for suite in CardSuite::iterator() {
        let mut card_in_suite_count: usize = 0;
        for card in shoe.iter() {
            if card.suite == *suite {
                card_in_suite_count += 1;
            }
        }
        // are there 13 hearts in each deck in the shoe?
        assert_eq!(card_in_suite_count, 13 * num_decks_in_shoe);
    }

    for rank in CardRank::iterator() {
        let mut num_cards: usize = 0;
        for card in shoe.iter() {
            if card.rank == *rank {
                num_cards += 1;
            }
        }
        // are there 4 aces in each deck in the shoe?
        assert_eq!(num_cards, 4 * num_decks_in_shoe);
    }
}
