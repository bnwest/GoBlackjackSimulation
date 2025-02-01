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
}

#[test]
fn test_card_suite_value() {
    assert_eq!(CardSuiteValue[&CardSuite::HEARTS], "♥️");
    assert_eq!(CardSuiteValue[&CardSuite::DIAMONDS], "♦️");
    assert_eq!(CardSuiteValue[&CardSuite::SPADES], "♠️");
    assert_eq!(CardSuiteValue[&CardSuite::CLUBS], "♣️");

    println!("cards: CardSuiteValue: {:#?}", CardSuiteValue);
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

    assert_eq!(CardRank::transmuate(2), CardRank::TWO);
    for rank in CardRank::iterator() {
        let roundtrip_rank: CardRank = CardRank::transmuate(rank.discriminant());
        assert_eq!(rank, &roundtrip_rank);
    }

    for rank in CardRank::iterator() {
        println!("{}", rank.to_string());
        println!("{:#?}", rank);
    }

    for rank in CardRank::iterator() {
        let value: u8 = rank.value();
        let discrim: u8 = rank.discriminant();
        if 1 <= discrim && discrim <= 10 {
            assert_eq!(value, discrim);
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

    let card_str = card.to_string();
    println!("{}", card.to_string());
    assert_eq!(card_str, "♠️ A");

    let mut card2 = card;
    card2.suite = CardSuite::HEARTS;
    assert_eq!(card.suite, CardSuite::SPADES);
    assert_eq!(card2.suite, CardSuite::HEARTS);
}
