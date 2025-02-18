// file src/game/tests.rs defines project module "game::tests".

use crate::game::BlackJack;
use crate::rules;
use crate::player;

#[test]
fn test_blackjack_create() {
    let mut blackjack: BlackJack = BlackJack::create();
    assert_eq!(blackjack.num_players(), 0);
    assert_eq!(blackjack.shoe.len(), 52 * rules::DECKS_IN_SHOE);
    assert_eq!(blackjack.shoe_top, 0);
    // println!("blackjack: {:#?}", blackjack);
}

#[test]
fn test_blackjack_set_players_for_game() {
    let mut blackjack: BlackJack = BlackJack::create();

    let players: Vec<player::Player> = vec![
        player::Player::create("Jack"),
        player::Player::create("Jill"),
    ];
    blackjack.set_players_for_game(players);
    assert_eq!(blackjack.num_players(), 2);
}

#[test]
fn test_blackjack_set_game_bets() {
    let mut blackjack: BlackJack = BlackJack::create();

    let players: Vec<player::Player> = vec![
        player::Player::create("Jack"),
        player::Player::create("Jill"),
    ];
    blackjack.set_players_for_game(players);

    let initial_bet: u32 = 100;
    blackjack.players[0].set_game_bets(&vec![initial_bet, ]);
    blackjack.players[1].set_game_bets(&vec![initial_bet, initial_bet]);
    assert_eq!(blackjack.players[0].num_master_hands(), 1);
    assert_eq!(blackjack.players[1].num_master_hands(), 2);
}
