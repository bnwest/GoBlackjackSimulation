// file srchandcards/tests.rs defines project module "hand::tests".

use super::HandOutcome;

#[test]
fn test_hand_outcomes() {
    let mut num_hand_outcomes: u32 = 0;
    for _suite in HandOutcome::iterator() {
        num_hand_outcomes += 1
    }

    assert_eq!(num_hand_outcomes, 5);
    assert_eq!(HandOutcome::STAND as usize, 0);
    assert_eq!(HandOutcome::BUST as usize, 1);
    assert_eq!(HandOutcome::SURRENDER as usize, 2);
    assert_eq!(HandOutcome::DEALER_BLACKJACK as usize, 3);
    assert_eq!(HandOutcome::IN_PLAY as usize, 4);

    let mut suite: HandOutcome = HandOutcome::STAND;
    assert_eq!(suite.discriminant(), 0);
    suite = HandOutcome::BUST;
    assert_eq!(suite.discriminant(), 1);
    suite = HandOutcome::SURRENDER;
    assert_eq!(suite.discriminant(), 2);
    suite = HandOutcome::DEALER_BLACKJACK;
    assert_eq!(suite.discriminant(), 3);
    suite = HandOutcome::IN_PLAY;
    assert_eq!(suite.discriminant(), 4);

    for hand_outcome in HandOutcome::iterator() {
        // hand_outcome: &HandOutcome
        let roundtrip_hand_outcome: HandOutcome =
            HandOutcome::transmute(hand_outcome.discriminant());
        assert_eq!(hand_outcome, &roundtrip_hand_outcome);

        println!("{}", hand_outcome.to_string());
        println!("{:?}", hand_outcome);
        println!("{:#?}", hand_outcome);
        println!("{hand_outcome:?}");
        println!("{hand_outcome:#?}");
    }
}

use super::PlayerHand;
use crate::cards;

#[test]
fn test_create_player_hand() {
    let from_split: bool = false;
    let bet: u32 = 100;
    let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
    assert_eq!(player_hand.cards.len(), 0);
    assert_eq!(player_hand.from_split, from_split);
    assert_eq!(player_hand.bet, bet);
    assert_eq!(player_hand.outcome, HandOutcome::IN_PLAY);

    assert_eq!(player_hand.num_cards(), 0);
    assert_eq!(player_hand.is_from_split(), from_split);
}

#[test]
fn test_player_hand_add_card() {
    let from_split: bool = false;
    let bet: u32 = 100;
    let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);

    let card: cards::Card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    player_hand.add_card(&card);
    assert_eq!(player_hand.num_cards(), 1);
    assert_eq!(player_hand.cards[0].suite, cards::CardSuite::SPADES);
    assert_eq!(player_hand.cards[0].rank, cards::CardRank::ACE);

    let card: cards::Card = player_hand.get_card(0);
    assert_eq!(card.suite, cards::CardSuite::SPADES);
    assert_eq!(card.rank, cards::CardRank::ACE);

    let aces_count: usize = player_hand.aces_count();
    assert_eq!(aces_count, 1);
}

#[test]
fn test_player_hand_count() {
    let from_split: bool = false;
    let bet: u32 = 100;
    let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);

    let card1: cards::Card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    let card2: cards::Card = cards::Card {
        suite: cards::CardSuite::HEARTS,
        rank: cards::CardRank::NINE,
    };
    player_hand.add_card(&card1);
    player_hand.add_card(&card2);
    assert_eq!(player_hand.num_cards(), 2);

    let aces_count: usize = player_hand.aces_count();
    assert_eq!(aces_count, 1);

    let hard_count: usize = player_hand.hard_count();
    assert_eq!(hard_count, 10);

    let soft_count: usize = player_hand.soft_count();
    assert_eq!(soft_count, 20);

    let count: usize = player_hand.count();
    assert_eq!(count, 20);

    let is_bust: bool = player_hand.is_bust();
    assert_eq!(is_bust, false);

    let is_natural: bool = player_hand.is_natural();
    assert_eq!(is_natural, false);

    let card3: cards::Card = cards::Card {
        suite: cards::CardSuite::DIAMONDS,
        rank: cards::CardRank::TEN,
    };
    player_hand.add_card(&card3);
    assert_eq!(player_hand.num_cards(), 3);

    let aces_count: usize = player_hand.aces_count();
    assert_eq!(aces_count, 1);

    let hard_count: usize = player_hand.hard_count();
    assert_eq!(hard_count, 20);

    let soft_count: usize = player_hand.soft_count();
    assert_eq!(soft_count, 20);

    let count: usize = player_hand.count();
    assert_eq!(count, 20);

    let is_bust: bool = player_hand.is_bust();
    assert_eq!(is_bust, false);

    let is_natural: bool = player_hand.is_natural();
    assert_eq!(is_natural, false);

    let card4: cards::Card = cards::Card {
        suite: cards::CardSuite::CLUBS,
        rank: cards::CardRank::EIGHT,
    };
    player_hand.add_card(&card4);
    assert_eq!(player_hand.num_cards(), 4);

    let hard_count: usize = player_hand.hard_count();
    assert_eq!(hard_count, 28);

    let soft_count: usize = player_hand.soft_count();
    assert_eq!(soft_count, 28);

    let count: usize = player_hand.count();
    assert_eq!(count, 28);

    let is_bust: bool = player_hand.is_bust();
    assert_eq!(is_bust, true);

    let is_natural: bool = player_hand.is_natural();
    assert_eq!(is_natural, false);
}

#[test]
fn test_player_is_natural() {
    for last_card_dealt_rank in [
        cards::CardRank::TEN,
        cards::CardRank::JACK,
        cards::CardRank::QUEEN,
        cards::CardRank::KING,
    ]
    .iter()
    {
        // last_card_dealt_rank: &cards::CardRank:
        let from_split: bool = false;
        let bet: u32 = 100;
        let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);

        let card1: cards::Card = cards::Card {
            suite: cards::CardSuite::SPADES,
            rank: cards::CardRank::ACE,
        };
        let card2: cards::Card = cards::Card {
            suite: cards::CardSuite::HEARTS,
            rank: *last_card_dealt_rank,
        };
        player_hand.add_card(&card1);
        player_hand.add_card(&card2);
        assert_eq!(player_hand.num_cards(), 2);

        let aces_count: usize = player_hand.aces_count();
        assert_eq!(aces_count, 1);

        let hard_count: usize = player_hand.hard_count();
        assert_eq!(hard_count, 11);

        let soft_count: usize = player_hand.soft_count();
        assert_eq!(soft_count, 21);

        let count: usize = player_hand.count();
        assert_eq!(count, 21);

        let is_bust: bool = player_hand.is_bust();
        assert_eq!(is_bust, false);

        let is_natural: bool = player_hand.is_natural();
        assert_eq!(is_natural, true);
    }
}

#[test]
fn test_player_is_natura_1() {
    for last_card_dealt_rank in [
        cards::CardRank::TEN,
        cards::CardRank::JACK,
        cards::CardRank::QUEEN,
        cards::CardRank::KING,
    ]
    .iter()
    {
        // last_card_dealt_rank: &cards::CardRank:
        let from_split: bool = true;
        let bet: u32 = 100;
        let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);

        let card1: cards::Card = cards::Card {
            suite: cards::CardSuite::SPADES,
            rank: cards::CardRank::ACE,
        };
        let card2: cards::Card = cards::Card {
            suite: cards::CardSuite::HEARTS,
            rank: *last_card_dealt_rank,
        };
        player_hand.add_card(&card1);
        player_hand.add_card(&card2);
        assert_eq!(player_hand.num_cards(), 2);

        let aces_count: usize = player_hand.aces_count();
        assert_eq!(aces_count, 1);

        let hard_count: usize = player_hand.hard_count();
        assert_eq!(hard_count, 11);

        let soft_count: usize = player_hand.soft_count();
        assert_eq!(soft_count, 21);

        let count: usize = player_hand.count();
        assert_eq!(count, 21);

        let is_bust: bool = player_hand.is_bust();
        assert_eq!(is_bust, false);

        let is_natural: bool = player_hand.is_natural();
        assert_eq!(is_natural, false);
    }
}

#[test]
fn test_player_is_natura_2() {
    let from_split: bool = false;
    let bet: u32 = 100;
    let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);

    let card1: cards::Card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    let card2: cards::Card = cards::Card {
        suite: cards::CardSuite::HEARTS,
        rank: cards::CardRank::NINE,
    };
    let card3: cards::Card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    player_hand.add_card(&card1);
    player_hand.add_card(&card2);
    player_hand.add_card(&card3);
    assert_eq!(player_hand.num_cards(), 3);

    let hard_count: usize = player_hand.hard_count();
    assert_eq!(hard_count, 11);

    let soft_count: usize = player_hand.soft_count();
    assert_eq!(soft_count, 21);

    let count: usize = player_hand.count();
    assert_eq!(count, 21);

    let is_bust: bool = player_hand.is_bust();
    assert_eq!(is_bust, false);

    let is_natural: bool = player_hand.is_natural();
    assert_eq!(is_natural, false);
}

#[test]
fn test_player_can_split() {
    let from_split: bool = false;
    let bet: u32 = 100;

    for rank in cards::CardRank::iterator() {
        // rank: &cards::CardRank
        let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
        let card1: cards::Card = cards::Card {
            suite: cards::CardSuite::SPADES,
            rank: *rank,
        };
        let card2: cards::Card = cards::Card {
            suite: cards::CardSuite::HEARTS,
            rank: *rank,
        };
        player_hand.add_card(&card1);
        player_hand.add_card(&card2);
        assert_eq!(player_hand.num_cards(), 2);

        let can_split: bool = player_hand.can_split();
        assert_eq!(can_split, true)
    }

    for rank1 in cards::CardRank::iterator() {
        // rank1: &cards::CardRank
        if *rank1 == cards::CardRank::JACK {
            break;
        }
        for rank2 in cards::CardRank::iterator() {
            // rank2: &cards::CardRank
            if *rank2 == cards::CardRank::JACK {
                break;
            }
            if *rank1 == *rank2 {
                continue;
            }
            let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
            let card1: cards::Card = cards::Card {
                suite: cards::CardSuite::SPADES,
                rank: *rank1,
            };
            let card2: cards::Card = cards::Card {
                suite: cards::CardSuite::HEARTS,
                rank: *rank2,
            };
            player_hand.add_card(&card1);
            player_hand.add_card(&card2);
            assert_eq!(player_hand.num_cards(), 2);

            let can_split: bool = player_hand.can_split();
            assert_eq!(can_split, false);
        }
    }

    for rank1 in [
        cards::CardRank::TEN,
        cards::CardRank::JACK,
        cards::CardRank::QUEEN,
        cards::CardRank::KING,
    ]
    .iter()
    {
        for rank2 in [
            cards::CardRank::TEN,
            cards::CardRank::JACK,
            cards::CardRank::QUEEN,
            cards::CardRank::KING,
        ]
        .iter()
        {
            let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
            let card1: cards::Card = cards::Card {
                suite: cards::CardSuite::SPADES,
                rank: *rank1,
            };
            let card2: cards::Card = cards::Card {
                suite: cards::CardSuite::HEARTS,
                rank: *rank2,
            };
            player_hand.add_card(&card1);
            player_hand.add_card(&card2);
            assert_eq!(player_hand.num_cards(), 2);

            let can_split: bool = player_hand.can_split();
            assert_eq!(can_split, true);
        }
    }
}
#[test]
fn test_player_is_hand_over() {
    let from_split: bool = false;
    let bet: u32 = 100;
    let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
    let card1: cards::Card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    let card2: cards::Card = cards::Card {
        suite: cards::CardSuite::HEARTS,
        rank: cards::CardRank::ACE,
    };
    player_hand.add_card(&card1);
    player_hand.add_card(&card2);

    player_hand.outcome = HandOutcome::STAND;
    let is_hand_over: bool = player_hand.is_hand_over();
    assert_eq!(is_hand_over, true);

    player_hand.outcome = HandOutcome::BUST;
    let is_hand_over: bool = player_hand.is_hand_over();
    assert_eq!(is_hand_over, true);

    player_hand.outcome = HandOutcome::SURRENDER;
    let is_hand_over: bool = player_hand.is_hand_over();
    assert_eq!(is_hand_over, true);

    player_hand.outcome = HandOutcome::DEALER_BLACKJACK;
    let is_hand_over: bool = player_hand.is_hand_over();
    assert_eq!(is_hand_over, true);

    player_hand.outcome = HandOutcome::IN_PLAY;
    let is_hand_over: bool = player_hand.is_hand_over();
    assert_eq!(is_hand_over, false);
}

use super::PlayerMasterHand;

#[test]
fn test_player_player_master_hand_create() {
    let mut player_master_hand: PlayerMasterHand = PlayerMasterHand::create();

    // let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
    let bet: u32 = 100;
    player_master_hand.add_start_hand(bet);
    assert_eq!(player_master_hand.num_hands(), 1);

    let mut deal_card: cards::Card;

    deal_card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    player_master_hand.hands[0].add_card(&deal_card);
    println!("deal card: {deal_card:#?}");

    deal_card = cards::Card {
        suite: cards::CardSuite::HEARTS,
        rank: cards::CardRank::ACE,
    };
    player_master_hand.hands[0].add_card(&deal_card);

    player_master_hand.log_hands("TEST");
}

#[test]
fn test_player_player_master_hand_split() {
    let mut player_master_hand: PlayerMasterHand = PlayerMasterHand::create();

    // let mut player_hand: PlayerHand = PlayerHand::create(from_split, bet);
    let bet: u32 = 100;
    player_master_hand.add_start_hand(bet);
    assert_eq!(player_master_hand.num_hands(), 1);

    let mut deal_card: cards::Card;

    deal_card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    player_master_hand.hands[0].add_card(&deal_card);

    deal_card = cards::Card {
        suite: cards::CardSuite::HEARTS,
        rank: cards::CardRank::ACE,
    };
    player_master_hand.hands[0].add_card(&deal_card);

    assert_eq!(player_master_hand.num_hands(), 1);

    player_master_hand.log_hands("PRE-SPLIT");

    let can_split: bool = player_master_hand.can_split(0);
    assert_eq!(can_split, true);

    let cards_to_add: [cards::Card; 2] = [
        cards::Card {
            suite: cards::CardSuite::DIAMONDS,
            rank: cards::CardRank::ACE,
        },
        cards::Card {
            suite: cards::CardSuite::CLUBS,
            rank: cards::CardRank::ACE,
        },
    ];

    let new_hand_index: usize = player_master_hand.split_hand(0, cards_to_add);
    assert_eq!(new_hand_index, 1);
    assert_eq!(player_master_hand.num_hands(), 2);

    player_master_hand.log_hands("POST-SPLIT");

    let cards_to_add: [cards::Card; 2] = [
        cards::Card {
            suite: cards::CardSuite::SPADES,
            rank: cards::CardRank::TEN,
        },
        cards::Card {
            suite: cards::CardSuite::DIAMONDS,
            rank: cards::CardRank::TEN,
        },
    ];

    let new_hand_index: usize = player_master_hand.split_hand(0, cards_to_add);
    assert_eq!(new_hand_index, 2);
    assert_eq!(player_master_hand.num_hands(), 3);

    player_master_hand.log_hands("POST-SPLIT");

    let cards_to_add: [cards::Card; 2] = [
        cards::Card {
            suite: cards::CardSuite::HEARTS,
            rank: cards::CardRank::TEN,
        },
        cards::Card {
            suite: cards::CardSuite::CLUBS,
            rank: cards::CardRank::TEN,
        },
    ];

    let new_hand_index: usize = player_master_hand.split_hand(1, cards_to_add);
    assert_eq!(new_hand_index, 3);
    assert_eq!(player_master_hand.num_hands(), 4);

    player_master_hand.log_hands("POST-SPLIT");
    /*
    assert!(false);

    POST-SPLIT: MasterHand
        Hand 1
            Card 1: ♠️ A
            Card 2: ♠️ 10
        Hand 2
            Card 1: ♥️ A
            Card 2: ♥️ 10
        Hand 3
            Card 1: ♦️ A
            Card 2: ♦️ 10
        Hand 4
            Card 1: ♣️ A
            Card 2: ♣️ 10
    */
}
