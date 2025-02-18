// file src/game.rs defines project module "game".

// two peas in a pod
use rand::prelude::*;
use rand_chacha::ChaCha8Rng;

use std::collections::HashMap;
use std::fmt::format;
use std::num;

use crate::cards::Card;
use crate::cards::CardRank;
use crate::hand::HandOutcome;
use crate::strategy;
use crate::strategy::PlayerDecision;

use super::cards;
use super::hand;
use super::player;
use super::rules;

pub struct BlackJackPlayerResults {
    pub hands_played: u64,
    pub hands_won: u64,
    pub hands_lost: u64,
    pub hands_pushed: u64,
    pub proceeds: i64,
}

impl BlackJackPlayerResults {
    pub fn create() -> BlackJackPlayerResults {
        BlackJackPlayerResults {
            hands_played: 0,
            hands_won: 0,
            hands_lost: 0,
            hands_pushed: 0,
            proceeds: 0,
        }
    }
}

pub struct BlackJackStats {
    pub double_down_count: u64,
    pub surrender_count: u64,
    pub split_count: u64,
    pub aces_split: u64,
}

impl BlackJackStats {
    pub fn create() -> BlackJackStats {
        BlackJackStats {
            double_down_count: 0,
            surrender_count: 0,
            split_count: 0,
            aces_split: 0,
        }
    }
}

const RNG_SEED: u64 = 42_u64;

pub struct BlackJack {
    shoe: Vec<cards::Card>,
    shoe_top: usize,
    players: Vec<player::Player>,
    results: HashMap<String, BlackJackPlayerResults>,
    stats: BlackJackStats,
    rng: ChaCha8Rng,
}

impl BlackJack {
    pub fn create() -> BlackJack {
        let mut players = vec![];

        let mut results: HashMap<String, BlackJackPlayerResults> = HashMap::new();

        let mut rng: ChaCha8Rng = ChaCha8Rng::seed_from_u64(RNG_SEED);

        let mut blackjack = BlackJack {
            shoe: cards::create_shoe(rules::DECKS_IN_SHOE),
            shoe_top: 0,
            players: players,
            results: results,
            stats: BlackJackStats::create(),
            rng: rng,
        };

        blackjack.reshuffle_shoe();

        return blackjack;
    }
    pub fn num_players(&self) -> usize {
        self.players.len()
    }
    pub fn reshuffle_shoe(&mut self) {
        cards::shuffle_shoe(&mut self.shoe, &mut self.rng);
        self.shoe_top = 0;
    }
    pub fn get_card_from_shoe(&mut self) -> cards::Card {
        let card: &cards::Card = &self.shoe[self.shoe_top];
        self.shoe_top += 1;
        return *card;
    }
    pub fn set_players_for_game(&mut self, players: Vec<player::Player>) {
        self.players = players;

        for p in 0..self.players.len() {
            let seen_player: bool = self.results.contains_key(&self.players[p].name);
            if !seen_player {
                self.results.insert(
                    self.players[p].name.clone(),
                    BlackJackPlayerResults::create(),
                );
            }
        }
    }
    pub fn add_result(
        &mut self,
        player_index: usize,
        master_hand_index: usize,
        hand_index: usize,
        initial_bet: u32,
        result: i32,
    ) {
        let player = &self.players[player_index];

        if let Some(player_results) = self.results.get_mut(&player.name) {
            player_results.hands_played += 1;
            if result > 0 {
                player_results.hands_won += 1;
            } else if result < 0 {
                player_results.hands_lost += 1;
            } else {
                player_results.hands_pushed += 1;
            }
            player_results.proceeds += result as i64;
        }
        // help: trait `IndexMut` is required to modify indexed content, but it is not implemented for `HashMap<String, BlackJackPlayerResults>`
        //  help: to modify a `HashMap<String, BlackJackPlayerResults>`, use `.get_mut()`, `.insert()` or the entry API

        let player_hand =
            &self.players[player_index].master_hands[master_hand_index].hands[hand_index];

        let double_down: bool =
            player_hand.num_cards() == 3 && initial_bet * 2 == result.abs() as u32;
        if double_down {
            self.stats.double_down_count += 1;
        }

        if player_hand.outcome == HandOutcome::SURRENDER {
            self.stats.surrender_count += 1;
        }

        if hand_index > 0 {
            self.stats.split_count += 1;
        }

        let splitting_aces: bool =
            player_hand.from_split && player_hand.cards[0].rank == CardRank::ACE;
        if splitting_aces {
            self.stats.aces_split += 1;
        }
    }
    pub fn log(msg: String) {
        println!("{msg}");
    }
    pub fn play_game(&mut self) {
        if self.shoe_top > rules::FORCE_RESHUFFLE {
            self.reshuffle_shoe();
        }

        let mut dealer: player::Dealer = player::Dealer::create("Steamboat Dealer");

        //
        // Players put down their bets for the next game,
        // one bet per hand.
        //

        let players: Vec<player::Player> = vec![
            player::Player::create("Jack"),
            player::Player::create("Jill"),
        ];

        self.set_players_for_game(players);
        // variable 'players' is dropped.

        let initial_bet: u32 = 2;
        self.players[0].set_game_bets(&vec![initial_bet]);
        self.players[1].set_game_bets(&vec![initial_bet, initial_bet]);

        //
        // DEAL HANDS
        //

        BlackJack::log("\n\nDEAL HANDS".to_string());

        for _i in 0..2 {
            for p in 0..self.num_players() {
                for mh in 0..self.players[p].master_hands.len() {
                    let card: cards::Card = self.get_card_from_shoe();
                    self.players[p].master_hands[mh].hands[0].add_card(&card);
                }
            }

            let card: cards::Card = self.get_card_from_shoe();
            dealer.hand.add_card(&card);
        }

        let dealer_top_card: cards::Card = dealer.top_card();
        let dealer_hole_card: cards::Card = dealer.hole_card();

        BlackJack::log(format!("dealer top card: {dealer_top_card:#?}"));

        //
        // PLAY HANDS
        //

        BlackJack::log("PLAY HANDS".to_string());
        if dealer.hand.is_natural() {
            // a real simulation would have to take care of Insurance, which is a sucker's bet,
            // so we just assume that no player will ask for insurance.
            // two cases:
            //     1. player has a natural and their bet is pushed
            //     2. player loses
            // either way the player will not get any more cards dealt and effectively stands.

            dealer.hand.outcome = HandOutcome::DEALER_BLACKJACK;
            for p in 0..self.num_players() {
                for mh in 0..self.players[p].num_master_hands() {
                    for h in 0..self.players[p].master_hands[mh].num_hands() {
                        // standing will do the right thing in the settlement logic below
                        self.players[p].master_hands[mh].hands[h].outcome = HandOutcome::STAND;
                    }
                }
            }
        } else {
            // dealer does not have a natural
            for p in 0..self.num_players() {
                BlackJack::log(format!("player {} - {}", p + 1, self.players[p].name));
                for mh in 0..self.players[p].num_master_hands() {
                    for h in 0..self.players[p].master_hands[mh].num_hands() {
                        BlackJack::log(format!("    hand {}.{}", mh + 1, h + 1));

                        for c in 0..self.players[p].master_hands[mh].hands[h].num_cards() {
                            let card: cards::Card =
                                self.players[p].master_hands[mh].hands[h].cards[c].clone();
                            BlackJack::log(format!("        card {}: {:#?}", c + 1, card));
                        }

                        let is_split_possible: bool = self.players[p].master_hands[mh].num_hands()
                            < self.players[p].master_hands[mh].HANDS_LIMIT;

                        loop {
                            if self.players[p].master_hands[mh].hands[h].outcome
                                == HandOutcome::STAND
                            {
                                // product of a prior ace split, outcome has already been determined.
                                BlackJack::log(format!(
                                    "        prior aces split: {:#?}: H{} S{}",
                                    strategy::PlayerDecision::STAND,
                                    self.players[p].master_hands[mh].hands[h].hard_count(),
                                    self.players[p].master_hands[mh].hands[h].soft_count(),
                                ));
                                break;
                            }

                            let hand_allow_more_splits: bool = is_split_possible;
                            let decision: strategy::PlayerDecision;
                            decision = strategy::determine_basic_strategy(
                                &dealer_top_card,
                                &self.players[p].master_hands[mh].hands[h],
                                hand_allow_more_splits,
                            );
                            BlackJack::log(format!("        basic strategy: {:#?}", decision));

                            if decision == PlayerDecision::STAND {
                                self.players[p].master_hands[mh].hands[h].outcome =
                                    HandOutcome::STAND;
                                BlackJack::log(format!(
                                    "        stand total H{} S{}",
                                    self.players[p].master_hands[mh].hands[h].hard_count(),
                                    self.players[p].master_hands[mh].hands[h].soft_count()
                                ));
                                break;
                            } else if decision == PlayerDecision::SURRENDER {
                                self.players[p].master_hands[mh].hands[h].outcome =
                                    HandOutcome::SURRENDER;
                                self.players[p].master_hands[mh].hands[h].bet /= 2;
                                break;
                            } else if decision == PlayerDecision::DOUBLE {
                                let card: cards::Card = self.get_card_from_shoe();
                                self.players[p].master_hands[mh].hands[h].add_card(&card);
                                self.players[p].master_hands[mh].hands[h].bet *= 2;
                                BlackJack::log(format!(
                                    "        hit: {:#?}, total H{} S{}",
                                    card,
                                    self.players[p].master_hands[mh].hands[h].hard_count(),
                                    self.players[p].master_hands[mh].hands[h].soft_count()
                                ));
                                self.players[p].master_hands[mh].hands[h].outcome =
                                    HandOutcome::STAND;
                                BlackJack::log(format!(
                                    "        stand total H{} S{}",
                                    self.players[p].master_hands[mh].hands[h].hard_count(),
                                    self.players[p].master_hands[mh].hands[h].soft_count()
                                ));
                                break;
                            } else if decision == PlayerDecision::HIT {
                                let card: cards::Card = self.get_card_from_shoe();
                                self.players[p].master_hands[mh].hands[h].add_card(&card);
                                BlackJack::log(format!(
                                    "        hit: {:#?}, total H{} S{}",
                                    card,
                                    self.players[p].master_hands[mh].hands[h].hard_count(),
                                    self.players[p].master_hands[mh].hands[h].soft_count()
                                ));
                                let hand_total: usize =
                                    self.players[p].master_hands[mh].hands[h].count();
                                if hand_total > 21 {
                                    self.players[p].master_hands[mh].hands[h].outcome =
                                        HandOutcome::BUST;
                                    BlackJack::log(format!(
                                        "        {:#?}",
                                        self.players[p].master_hands[mh].hands[h].outcome
                                    ));
                                    break;
                                } else {
                                    self.players[p].master_hands[mh].hands[h].outcome =
                                        HandOutcome::IN_PLAY;
                                }
                            } else if decision == PlayerDecision::SPLIT {
                                let card1: cards::Card = self.get_card_from_shoe();
                                let card2: cards::Card = self.get_card_from_shoe();
                                let hand_index: usize = h;
                                let new_hand_index: usize;
                                new_hand_index = self.players[p].master_hands[mh]
                                    .split_hand(hand_index, [card1, card2]);
                                BlackJack::log(format!(
                                    "        split, new hand index {}, adding cards {:#?}, {:#?}",
                                    new_hand_index, card1, card2
                                ));
                                BlackJack::log(format!(
                                    "        card 1: {:#?}",
                                    self.players[p].master_hands[mh].hands[h].cards[0]
                                ));
                                BlackJack::log(format!(
                                    "        card 2: {:#?}",
                                    self.players[p].master_hands[mh].hands[h].cards[1]
                                ));
                                let splitting_aces: bool =
                                    self.players[p].master_hands[mh].hands[h].cards[0].rank
                                        == CardRank::ACE;
                                if splitting_aces && rules::NO_MORE_CARDS_AFTER_SPLITTING_ACES {
                                    self.players[p].master_hands[mh].hands[h].outcome =
                                        HandOutcome::STAND;
                                    BlackJack::log(format!(
                                        "        aces split: {:#?},  total H{} S{}",
                                        self.players[p].master_hands[mh].hands[h].outcome,
                                        self.players[p].master_hands[mh].hands[h].hard_count(),
                                        self.players[p].master_hands[mh].hands[h].soft_count()
                                    ));
                                    self.players[p].master_hands[mh].hands[new_hand_index]
                                        .outcome = HandOutcome::STAND;
                                    break;
                                }
                            } else {
                                BlackJack::log("FTW".to_string());
                                // clues at the scene of the crime ...
                                BlackJack::log(format!(
                                    "FTW: dealer top card: {:#?}, is split possible? {}",
                                    dealer_top_card, is_split_possible
                                ));
                                BlackJack::log(format!(
                                    "FTW: player hand count: H{} S{}",
                                    self.players[p].master_hands[mh].hands[h].hard_count(),
                                    self.players[p].master_hands[mh].hands[h].soft_count()
                                ));
                                BlackJack::log(format!("FTW: decision: {:#?}", decision));
                                self.players[p].master_hands[mh].hands[h].outcome =
                                    HandOutcome::STAND;
                                break;
                            }
                            // break;
                        }
                    }
                }
            }
        }

        //
        // DEALER HAND
        //

        BlackJack::log("DEALER HAND".to_string());
        BlackJack::log(format!("dealter top  card: {:#?}", dealer_top_card));
        BlackJack::log(format!("dealter hole card: {:#?}", dealer_hole_card));
        let mut dealer_done: bool = false;
        while !dealer_done {
            let hard_count: usize = dealer.hand.hard_count();
            let soft_count: usize = dealer.hand.soft_count();
            let use_soft_count: bool = hard_count < soft_count && soft_count <= 21;
            if use_soft_count && soft_count < rules::DEALER_HITS_SOFT_ON {
                let card: cards::Card = self.get_card_from_shoe();
                dealer.hand.add_card(&card);
                BlackJack::log(format!("    add: {:#?}", card));
            } else if !use_soft_count && hard_count <= rules::DEALER_HITS_HARD_ON {
                let card: cards::Card = self.get_card_from_shoe();
                dealer.hand.add_card(&card);
                BlackJack::log(format!("    add: {:#?}", card));
            } else {
                dealer.hand.outcome = HandOutcome::STAND;
                dealer_done = true;
                BlackJack::log(format!(
                    "    stand: total H{} S{}",
                    dealer.hand.hard_count(),
                    dealer.hand.soft_count()
                ));
            }

            if dealer.hand.count() > 21 {
                dealer.hand.outcome = HandOutcome::BUST;
                dealer_done = true;
                BlackJack::log(format!("    bust"));
            }
        }

        //
        // SETTlE HANDS
        //

        BlackJack::log("SETTLE HANDS".to_string());
        if dealer.hand.outcome == HandOutcome::DEALER_BLACKJACK {
            for p in 0..self.num_players() {
                BlackJack::log(format!("player {} - {}", p + 1, self.players[p].name));
                for mh in 0..self.players[p].num_master_hands() {
                    for h in 0..self.players[p].master_hands[mh].num_hands() {
                        if self.players[p].master_hands[mh].hands[h].is_natural() {
                            self.add_result(p, mh, h, initial_bet, 0_i32);
                            BlackJack::log(format!(
                                "    hand {}.{}: push both player and dealer had naturals",
                                mh + 1,
                                h + 1
                            ));
                        } else {
                            self.add_result(
                                p,
                                mh,
                                h,
                                initial_bet,
                                -(self.players[p].master_hands[mh].hands[h].bet as i32),
                            );
                            BlackJack::log(format!(
                                "    hand {}.{}: dealer natural: lost {}",
                                mh + 1,
                                h + 1,
                                self.players[p].master_hands[mh].hands[h].bet
                            ));
                        }
                    }
                }
            }
        } else {
            // dealer does not have a natural
            for p in 0..self.num_players() {
                BlackJack::log(format!("player {} - {}", p + 1, self.players[p].name));
                for mh in 0..self.players[p].num_master_hands() {
                    // need the traditional for loop
                    // for h in 0..self.players[p].master_hands[mh].num_hands() {
                    //    fails since the loop range is only calculated once and not
                    //    evertime through the loop
                    let mut h: usize = 0;
                    while h < self.players[p].master_hands[mh].num_hands() {
                        if self.players[p].master_hands[mh].hands[h].is_bust() {
                            self.add_result(
                                p,
                                mh,
                                h,
                                initial_bet,
                                -(self.players[p].master_hands[mh].hands[h].bet as i32),
                            );
                            BlackJack::log(format!(
                                "    hand {}.{}: bust: lost {}",
                                mh + 1,
                                h + 1,
                                self.players[p].master_hands[mh].hands[h].bet
                            ));
                        } else if self.players[p].master_hands[mh].hands[h].is_surrender() {
                            self.add_result(
                                p,
                                mh,
                                h,
                                initial_bet,
                                -(self.players[p].master_hands[mh].hands[h].bet as i32),
                            );
                            BlackJack::log(format!(
                                "    hand {}.{}: surrender: lost {}",
                                mh + 1,
                                h + 1,
                                self.players[p].master_hands[mh].hands[h].bet
                            ));
                        } else {
                            // player has a non-bust, non-surrender hand
                            if self.players[p].master_hands[mh].hands[h].is_natural() {
                                let payout: i32 = ((self.players[p].master_hands[mh].hands[h].bet as f32) * rules::NATURAL_BLACKJACK_PAYOUT) as i32;
                                self.add_result(
                                    p,
                                    mh,
                                    h,
                                    initial_bet,
                                    payout,
                                );
                                BlackJack::log(format!(
                                    "    hand {}.{}: natural: won {}",
                                    mh + 1,
                                    h + 1,
                                    payout
                                ));
                            } else if dealer.hand.is_bust() {
                                self.add_result(
                                    p,
                                    mh,
                                    h,
                                    initial_bet,
                                    self.players[p].master_hands[mh].hands[h].bet as i32,
                                );
                                BlackJack::log(format!(
                                    "    hand {}.{}: dealer bust: won {}",
                                    mh + 1,
                                    h + 1,
                                    self.players[p].master_hands[mh].hands[h].bet
                                ));    
                            } else {
                                // determine who wins by solely checking the hand totals
                                if self.players[p].master_hands[mh].hands[h].count() < dealer.hand.count() {
                                    self.add_result(
                                        p,
                                        mh,
                                        h,
                                        initial_bet,
                                        -(self.players[p].master_hands[mh].hands[h].bet as i32),
                                    );
                                    BlackJack::log(format!(
                                        "    hand {}.{}: lost {}",
                                        mh + 1,
                                        h + 1,
                                        self.players[p].master_hands[mh].hands[h].bet
                                    ));        
                                } else if self.players[p].master_hands[mh].hands[h].count() > dealer.hand.count() {
                                    self.add_result(
                                        p,
                                        mh,
                                        h,
                                        initial_bet,
                                        self.players[p].master_hands[mh].hands[h].bet as i32,
                                    );
                                    BlackJack::log(format!(
                                        "    hand {}.{}: won {}",
                                        mh + 1,
                                        h + 1,
                                        self.players[p].master_hands[mh].hands[h].bet
                                    ));        
                                } else {
                                    self.add_result(
                                        p,
                                        mh,
                                        h,
                                        initial_bet,
                                        0_i32,
                                    );
                                    BlackJack::log(format!(
                                        "    hand {}.{}: push",
                                        mh + 1,
                                        h + 1,
                                    ));        
                                }
                            }
                        }

                        h += 1;
                    }
                }
            }
        }
    }
}
