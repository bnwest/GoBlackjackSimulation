package game

import (
	"fmt"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"
)

type HandOutcome string

const (
	STAND            HandOutcome = "stand"
	BUST             HandOutcome = "bust"
	SURRENDER        HandOutcome = "surrender"
	DEALER_BLACKJACK HandOutcome = "dealer-blackjack"
	IN_PLAY          HandOutcome = "in-play"
)

//
// PlayerHand
//

type PlayerHand struct {
	Cards     []cards.Card
	FromSplit bool
	Bet       int
	OutCome   HandOutcome
}

// factory
func CreatePlayerHand(from_split bool, bet int) *PlayerHand {
	hand := PlayerHand{
		Cards:     []cards.Card{},
		FromSplit: from_split,
		Bet:       bet,
		OutCome:   HandOutcome(IN_PLAY),
	}
	return &hand
}

// AddCard() is a receiver function which takes a special object parameter
// like class methods do.  go statis-checker dislikes my naming convention,
// while I dislike "adjacent" functions standing in for class methods.

// (self *PlayerHand) allows the referenced object to be updated,
//     I suspect that a copy of the reference/pointer is being sent
// (self PlayerHand) uses a copy of the object,
//     all updates are made to the copy

func (self *PlayerHand) NumCards() int {
	return len(self.Cards)
}

func (self *PlayerHand) IsFromSplit() bool {
	return self.FromSplit
}

func (self *PlayerHand) GetCard(cardIndex int) cards.Card {
	return self.Cards[cardIndex]
}

func (self *PlayerHand) AddCard(card cards.Card) {
	self.Cards = append(self.Cards, card)
}

func (self *PlayerHand) AcesCount() int {
	count := 0
	for i := 0; i < self.NumCards(); i++ {
		var card cards.Card = self.Cards[i]
		if card.Rank == cards.ACE {
			count++
		}
	}
	return count
}

func (self *PlayerHand) HardCount() int {
	hard_count := 0
	for i := 0; i < self.NumCards(); i++ {
		var card cards.Card = self.Cards[i]
		hard_count += cards.CardRankValue[card.Rank]
	}
	return hard_count
}

func (self *PlayerHand) SoftCount() int {
	// if the soft count is a bust, we convert the Ace values
	// back to the value of 1, one at a time, until the soft count
	// is no longer a bust or until there are no more Aces
	// and the soft count has become the hard count.
	soft_count := 0
	aces_count := 0
	for i := 0; i < self.NumCards(); i++ {
		var card cards.Card = self.Cards[i]
		if card.Rank == cards.ACE {
			soft_count += 11
			aces_count++
		} else {
			soft_count += cards.CardRankValue[card.Rank]
		}
	}
	if soft_count > 21 {
		for i := 0; i < aces_count; i++ {
			soft_count -= 10
			if soft_count <= 21 {
				break
			}
		}
	}
	return soft_count
}

func (self *PlayerHand) Count() int {
	// return the highest count for hand,
	// which is always the soft count.
	return self.SoftCount()
}

func (self *PlayerHand) IsNatural() bool {
	if !self.FromSplit {
		if self.NumCards() == 2 {
			if self.SoftCount() == 21 {
				return true
			}
		}
	}
	return false
}

func (self *PlayerHand) IsBust() bool {
	return self.Count() > 21
}

func (self *PlayerHand) CanSplit() bool {
	// there are other split house rules that will be applied
	// at a higher abstraction level ... like splitting aces
	// after a split ...like limiting the number of splits
	// from the original (aka "master") hand.
	if self.NumCards() == 2 {
		var card1 cards.Card = self.Cards[0]
		var card2 cards.Card = self.Cards[1]
		if house_rules.SPLIT_ON_VALUE_MATCH {
			if cards.CardRankValue[card1.Rank] == cards.CardRankValue[card2.Rank] {
				return true
			}
		} else {
			if card1.Rank == card2.Rank {
				return true
			}
		}
	}
	return false
}

func (self *PlayerHand) IsHandOver() bool {
	switch self.OutCome {
	case HandOutcome(STAND):
		return true
	case HandOutcome(BUST):
		return true
	case HandOutcome(SURRENDER):
		return true
	case HandOutcome(DEALER_BLACKJACK):
		return true
	case HandOutcome(IN_PLAY):
		return false
	default:
		return false
	}
}

//
// DealerHand
//

type DealerHand struct {
	Cards   []cards.Card
	OutCome HandOutcome
}

// factory
func CreateDealerHand() *DealerHand {
	var hand DealerHand = DealerHand{
		Cards:   []cards.Card{},
		OutCome: HandOutcome(IN_PLAY),
	}
	return &hand
}

func (self *DealerHand) AddCard(card cards.Card) {
	self.Cards = append(self.Cards, card)
}

func (self *DealerHand) HardCount() int {
	hard_count := 0
	for i := 0; i < self.NumCards(); i++ {
		var card cards.Card = self.Cards[i]
		hard_count += cards.CardRankValue[card.Rank]
	}
	return hard_count
}

func (self *DealerHand) SoftCount() int {
	// if the soft count is a bust, we convert the Ace values
	// back to the value of 1, one at a time, until the soft count
	// is no longer a bust or until there are no more Aces
	// and the soft count has become the hard count.
	soft_count := 0
	aces_count := 0
	for i := 0; i < self.NumCards(); i++ {
		var card cards.Card = self.Cards[i]
		if card.Rank == cards.ACE {
			soft_count += 11
			aces_count++
		} else {
			soft_count += cards.CardRankValue[card.Rank]
		}
	}
	if soft_count > 21 {
		for i := 0; i < aces_count; i++ {
			soft_count -= 10
			if soft_count <= 21 {
				break
			}
		}
	}
	return soft_count
}

func (self *DealerHand) Count() int {
	// return the highest count for hand,
	// which is always the soft count.
	return self.SoftCount()
}

func (self *DealerHand) IsNatural() bool {
	if self.NumCards() == 2 {
		if self.SoftCount() == 21 {
			return true
		}
	}
	return false
}

func (self *DealerHand) IsBust() bool {
	return self.Count() > 21
}

func (self *DealerHand) NumCards() int {
	return len(self.Cards)
}

func (self *DealerHand) IsHandOver() bool {
	switch self.OutCome {
	case HandOutcome(STAND):
		return true
	case HandOutcome(BUST):
		return true
	case HandOutcome(DEALER_BLACKJACK):
		return true
	case HandOutcome(IN_PLAY):
		return false
	default:
		return false
	}
}

//
// PlayerMasterHand
//

type PlayerMasterHand struct {
	HANDS_LIMIT int
	Hands       []*PlayerHand
}

// factory
func CreatePlayerMasterHand() *PlayerMasterHand {
	var master_hand PlayerMasterHand = PlayerMasterHand{
		Hands:       []*PlayerHand{},
		HANDS_LIMIT: house_rules.SPLITS_PER_HAND + 1,
	}
	return &master_hand
}

func (self *PlayerMasterHand) NumHands() int {
	return len(self.Hands)
}

func (self *PlayerMasterHand) AddStartHand(bet int) {
	const from_split bool = false

	var player_hand *PlayerHand
	player_hand = CreatePlayerHand(from_split, bet)

	self.Hands = append(self.Hands, player_hand)
}

func (self *PlayerMasterHand) logHands(preface string) {
	fmt.Printf("%v: MasterHand\n", preface)
	for i := 0; i < self.NumHands(); i++ {
		var hand *PlayerHand = self.Hands[i]
		fmt.Printf("    Hand %v\n", i+1)
		for j := 0; j < hand.NumCards(); j++ {
			var card cards.Card = hand.Cards[j]
			fmt.Printf(
				"        Card %v: %v%v\n",
				j+1,
				cards.CardRankString[card.Rank],
				cards.CardSuiteValue[card.Suite],
			)
		}
	}
}

func (self *PlayerMasterHand) SplitHand(
	handIndex int,
	cardsToAdd [2]cards.Card,
) int {
	// fmt.Printf("SplitHand: hand index: %v\n", handIndex)
	// fmt.Printf("SplitHand: cards to add: %v  %v", cardsToAdd[0].Str(), cardsToAdd[1].Str())
	// self.logHands("before")

	// there are two in the hand of the same value
	// or rank depending of the house rules.
	var card1 cards.Card = self.Hands[handIndex].Cards[0]
	var card2 cards.Card = self.Hands[handIndex].Cards[1]

	var oldPlayerHand *PlayerHand
	oldPlayerHand = self.Hands[handIndex]
	oldPlayerHand.Cards = []cards.Card{card1, cardsToAdd[0]}
	oldPlayerHand.FromSplit = true
	oldPlayerHand.OutCome = HandOutcome(IN_PLAY)

	var newPlayerHand *PlayerHand
	newPlayerHand = CreatePlayerHand(true, oldPlayerHand.Bet)
	newPlayerHand.Cards = []cards.Card{card2, cardsToAdd[1]}
	newPlayerHand.OutCome = HandOutcome(IN_PLAY)

	newHandIndex := self.NumHands()
	self.Hands = append(self.Hands, newPlayerHand)

	// self.logHands("after")

	return newHandIndex
}

func (self *PlayerMasterHand) CanSplit(handIndex int) bool {
	if self.NumHands() < self.HANDS_LIMIT {
		// master hand allows
		var hand *PlayerHand = self.Hands[handIndex]
		if hand.CanSplit() {
			// individual hand allows
			return true
		}
	}
	return false
}
