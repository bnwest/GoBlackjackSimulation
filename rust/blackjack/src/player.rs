// file src/hand.rs defines project module "hand".

use crate::cards;
use crate::cards::Card;
use crate::hand;

//
// Player
//

// a player can play a set of hands.
// each indivdual master hand can be split into more hands,
// for which there is hard limit.  each master hand can typically be split
// up to three times, for a total of four hands starting from the master hand.

pub struct Player {
    pub master_hands: Vec<hand::PlayerMasterHand>,
    pub name: String,
}

impl Player {
    pub fn create(name: &str) -> Player {
        Player {
            master_hands: vec![],
            name: name.to_string(),
        }
    }
    pub fn num_master_hands(&self) -> usize {
        self.master_hands.len()
    }
    pub fn game_reset(&mut self) {
        self.master_hands = vec![];
    }
    pub fn set_game_bets(&mut self, bets: &Vec<u32>) {
        self.game_reset();

        for bet in bets.iter() {
            // bet: &u32
            let mut master_hand: hand::PlayerMasterHand = hand::PlayerMasterHand::create();
            master_hand.add_start_hand(*bet);
            self.master_hands.push(master_hand);
        }
    }
}

pub struct Dealer {
    pub hand: hand::DealerHand,
    pub name: String,
}

impl Dealer {
    pub fn create(name: &str) -> Dealer {
        Dealer {
            hand: hand::DealerHand::create(),
            name: name.to_string(),
        }
    }
    pub fn game_reset(&mut self) {
        self.hand = hand::DealerHand::create();
    }
    pub fn top_card(&self) -> cards::Card {
        let card: &cards::Card = &self.hand.cards[0];
        return *card;
    }
    pub fn hole_card(&self) -> cards::Card {
        let card: &cards::Card = &self.hand.cards[1];
        return *card;
    }
}

#[cfg(test)]
mod tests;
