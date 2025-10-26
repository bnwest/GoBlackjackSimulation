// file src/hand.rs defines project module "hand".

use std::fmt;
use std::mem::transmute;
use std::slice::Iter;

// use cards::Card;
use crate::cards;
use crate::rules;
use std::array;

// #[derive(Eq, PartialEq, Hash, Copy, Clone)]
#[derive(Copy, Clone, PartialEq)]
#[repr(u8)]
pub enum HandOutcome {
    STAND = 0,
    BUST = 1,
    SURRENDER = 2,
    DEALER_BLACKJACK = 3,
    IN_PLAY = 4,
}

impl HandOutcome {
    pub fn iterator() -> array::IntoIter<HandOutcome, 5> {
        [
            HandOutcome::STAND,
            HandOutcome::BUST,
            HandOutcome::SURRENDER,
            HandOutcome::DEALER_BLACKJACK,
            HandOutcome::IN_PLAY,
        ].into_iter()
    }
    pub fn discriminant(&self) -> u8 {
        // fn returns the integer discriminat for the enum
        // some may see "as" type casts as a red flag
        *self as u8
    }
    pub fn transmute(discrim: u8) -> HandOutcome {
        // FAILS: rank = CardRank(2);
        // FAILS: rank = 2 as CardRank;
        // WORKS POORLY: rank = unsafe { transmute(2 as u8) };
        match discrim {
            0 => HandOutcome::STAND,
            1 => HandOutcome::BUST,
            2 => HandOutcome::SURRENDER,
            3 => HandOutcome::DEALER_BLACKJACK,
            4 => HandOutcome::IN_PLAY,
            _ => HandOutcome::STAND, // Default fallback
        }
    }
    pub fn _to_string(&self) -> &str {
        static STRINGS: [&str; 5] = [
            "stand",            // HandOutcome::STAND
            "bust",             // HandOutcome::BUST
            "surrender",        // HandOutcome::SURRENDER
            "dealer-blackjack", // HandOutcome::DEALER_BLACKJACK
            "in-play",          // HandOutcome::IN_PLAY
        ];
        STRINGS[self.discriminant() as usize]
    }
}

impl fmt::Debug for HandOutcome {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        return write!(f, "{}", self._to_string());
    }
}

impl fmt::Display for HandOutcome {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        return write!(f, "{}", self._to_string());
    }
}

//
// PlayerHand
//

pub struct PlayerHand {
    pub cards: Vec<cards::Card>,
    pub from_split: bool,
    pub bet: u32,
    pub outcome: HandOutcome,
}

impl PlayerHand {
    pub fn create(from_split: bool, bet: u32) -> PlayerHand {
        let mut player_hand: PlayerHand = PlayerHand {
            cards: vec![],
            from_split: from_split,
            bet: bet,
            outcome: HandOutcome::IN_PLAY,
        };
        player_hand
    }
    pub fn num_cards(&self) -> usize {
        return self.cards.len();
    }
    pub fn add_card(&mut self, card: &cards::Card) {
        self.cards.push(*card);
    }
    pub fn is_from_split(&self) -> bool {
        return self.from_split;
    }
    pub fn get_card(&self, hand_index: usize) -> cards::Card {
        let card: cards::Card = self.cards[hand_index];
        return card;
    }
    pub fn aces_count(&self) -> usize {
        let mut count: usize = 0;
        for card in self.cards.iter() {
            // card: &cards::Card
            if card.rank == cards::CardRank::ACE {
                count += 1;
            }
        }
        return count;
    }
    pub fn hard_count(&self) -> usize {
        let mut hard_count: usize = 0;
        for card in self.cards.iter() {
            hard_count += card.rank.value();
        }
        return hard_count;
    }
    pub fn soft_count(&self) -> usize {
        // if the soft count is a bust, we convert the Ace values
        // back to the value of 1, one at a time, until the soft count
        // is no longer a bust or until there are no more Aces
        // and the soft count has become the hard count.
        let mut soft_count: usize = 0;
        let mut aces_count: usize = 0;
        for card in self.cards.iter() {
            if card.rank == cards::CardRank::ACE {
                soft_count += 11;
                aces_count += 1;
            } else {
                soft_count += card.rank.value();
            }
        }
        if soft_count > 21 {
            for _i in 0..aces_count {
                soft_count -= 10;
                if soft_count <= 21 {
                    break;
                }
            }
        }
        return soft_count;
    }
    pub fn count(&self) -> usize {
        // return the highest count for hand,
        // which is always the soft count.
        return self.soft_count();
    }
    pub fn is_bust(&self) -> bool {
        return self.soft_count() > 21;
    }
    pub fn is_surrender(&self) -> bool {
        // need outcome to be finalized ...
        return self.outcome == HandOutcome::SURRENDER;
    }
    pub fn is_natural(&self) -> bool {
        if !self.from_split {
            if self.num_cards() == 2 {
                if self.soft_count() == 21 {
                    return true;
                }
            }
        }
        return false;
    }
    pub fn can_split(&self) -> bool {
        // there are other split house rules that will be applied
        // at a higher abstraction level ... like splitting aces
        // after a split ...like limiting the number of splits
        // from the original (aka "master") hand.
        if self.num_cards() == 2 {
            let card1: cards::Card = self.get_card(0);
            let card2: cards::Card = self.get_card(1);
            if rules::SPLIT_ON_VALUE_MATCH {
                if card1.rank.value() == card2.rank.value() {
                    return true;
                }
            } else {
                if card1.rank == card2.rank {
                    return true;
                }
            }
        }
        return false;
    }
    pub fn is_hand_over(&self) -> bool {
        match self.outcome {
            HandOutcome::STAND => return true,
            HandOutcome::BUST => return true,
            HandOutcome::SURRENDER => return true,
            HandOutcome::DEALER_BLACKJACK => return true,
            HandOutcome::IN_PLAY => return false,
        }
        // return self.outcome != HandOutcome::IN_PLAY;
    }
}

//
// DealerHand
//

pub struct DealerHand {
    pub cards: Vec<cards::Card>,
    pub outcome: HandOutcome,
}

impl DealerHand {
    pub fn create() -> DealerHand {
        let dealer_hand: DealerHand = DealerHand {
            cards: vec![],
            outcome: HandOutcome::IN_PLAY,
        };
        dealer_hand
    }
    pub fn num_cards(&self) -> usize {
        return self.cards.len();
    }
    pub fn add_card(&mut self, card: &cards::Card) {
        self.cards.push(*card);
    }
    pub fn hard_count(&self) -> usize {
        let mut hard_count: usize = 0;
        for card in self.cards.iter() {
            hard_count += card.rank.value();
        }
        return hard_count;
    }
    pub fn soft_count(&self) -> usize {
        // if the soft count is a bust, we convert the Ace values
        // back to the value of 1, one at a time, until the soft count
        // is no longer a bust or until there are no more Aces
        // and the soft count has become the hard count.
        let mut soft_count: usize = 0;
        let mut aces_count: usize = 0;
        for card in self.cards.iter() {
            if card.rank == cards::CardRank::ACE {
                soft_count += 11;
                aces_count += 1;
            } else {
                soft_count += card.rank.value();
            }
        }
        if soft_count > 21 {
            for _i in 0..aces_count {
                soft_count -= 10;
                if soft_count <= 21 {
                    break;
                }
            }
        }
        return soft_count;
    }
    pub fn count(&self) -> usize {
        // return the highest count for hand,
        // which is always the soft count.
        return self.soft_count();
    }
    pub fn is_bust(&self) -> bool {
        return self.soft_count() > 21;
    }
    pub fn is_natural(&self) -> bool {
        if self.num_cards() == 2 {
            if self.soft_count() == 21 {
                return true;
            }
        }
        return false;
    }
    pub fn is_hand_over(&self) -> bool {
        match self.outcome {
            HandOutcome::STAND => return true,
            HandOutcome::BUST => return true,
            HandOutcome::SURRENDER => return true,
            HandOutcome::DEALER_BLACKJACK => return true,
            HandOutcome::IN_PLAY => return false,
        }
        // return self.outcome != HandOutcome::IN_PLAY;
    }
}

//
// PlayerMasterHand
//

pub struct PlayerMasterHand {
    pub HANDS_LIMIT: usize,
    pub hands: Vec<PlayerHand>,
}

impl PlayerMasterHand {
    pub fn create() -> PlayerMasterHand {
        let master_hand: PlayerMasterHand = PlayerMasterHand {
            hands: vec![],
            HANDS_LIMIT: rules::SPLITS_PER_HAND + 1,
        };
        master_hand
    }
    pub fn num_hands(&self) -> usize {
        return self.hands.len();
    }
    pub fn add_start_hand(&mut self, bet: u32) {
        let from_split: bool = false;
        let player_hand: PlayerHand = PlayerHand::create(from_split, bet);
        self.hands.push(player_hand);
    }
    pub fn can_split(&self, hand_index: usize) -> bool {
        if self.num_hands() < self.HANDS_LIMIT {
            // master hand allows
            let player_hand: &PlayerHand = &self.hands[hand_index];
            if player_hand.can_split() {
                // individual hand allows
                return true;
            }
        }
        return false;
    }
    pub fn split_hand(&mut self, hand_index: usize, cards_to_add: [cards::Card; 2]) -> usize {
        let hand = &mut self.hands[hand_index];
        let [card1, card2] = [hand.cards[0], hand.cards[1]];

        // Update existing hand
        hand.cards.clear();
        hand.cards.extend([card1, cards_to_add[0]]);
        hand.from_split = true;
        hand.outcome = HandOutcome::IN_PLAY;

        // Create new hand using functional style
        let new_hand = PlayerHand {
            cards: [card2, cards_to_add[1]].into(),
            from_split: true,
            bet: hand.bet,
            outcome: HandOutcome::IN_PLAY,
        };

        let new_hand_index = self.num_hands();
        self.hands.push(new_hand);
        new_hand_index
    }
    pub fn log_hands(&self, preface: &str) {
        println!("{preface}: MasterHand");
        for i in 0..self.hands.len() {
            let player_hand: &PlayerHand = &self.hands[i];
            println!("    Hand {}", i + 1);
            for j in 0..player_hand.num_cards() {
                let card: &cards::Card = &player_hand.cards[j];
                println!("        Card {}: {:#?}", j + 1, card);
            }
        }
    }
}

#[cfg(test)]
mod tests;
