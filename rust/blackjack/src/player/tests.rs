// file src/player/tests.rs defines project module "player::tests".

use crate::cards;

use super::Player;

#[test]
fn test_player_create() {
    let mut player: Player = Player::create("Jack");

    let bets: Vec<u32> = [100, 100, 100].to_vec();
    player.set_game_bets(&bets);
    assert_eq!(bets.len(), 3);
    assert_eq!(player.num_master_hands(), 3);

    player.game_reset();
    assert_eq!(player.num_master_hands(), 0);
}

use super::Dealer;

#[test]
fn test_dealer_create() {
    let mut dealer: Dealer = Dealer::create("Riverboat Dealer");
    assert_eq!(dealer.hand.num_cards(), 0);

    let card1: cards::Card = cards::Card {
        suite: cards::CardSuite::SPADES,
        rank: cards::CardRank::ACE,
    };
    let card2: cards::Card = cards::Card {
        suite: cards::CardSuite::HEARTS,
        rank: cards::CardRank::TEN,
    };
    dealer.hand.add_card(&card1);
    dealer.hand.add_card(&card2);
    assert_eq!(dealer.hand.num_cards(), 2);

    assert_eq!(dealer.top_card(), card1);
    assert_eq!(dealer.hole_card(), card2);

    dealer.game_reset();
    assert_eq!(dealer.hand.num_cards(), 0);
}
