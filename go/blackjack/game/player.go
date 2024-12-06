package game

import (
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
)

//
// Player
//

// a player can play a set of hands.
// each indivdual master hand can be split into more hands,
// for which there is hard limit.  each master hand can typically be split
// up to three times, for a total of four hands starting from the master hand.

type Player struct {
	PlayerMasterHands []*PlayerMasterHand
	Name              string
}

func CreatePlayer(name string) *Player {
	player := Player{
		PlayerMasterHands: []*PlayerMasterHand{},
		Name:              name,
	}
	return &player
}

func (self *Player) NumMasterHands() int {
	return len(self.PlayerMasterHands)
}

func (self *Player) GameReset() {
	self.PlayerMasterHands = []*PlayerMasterHand{}
}

func (self *Player) SetGameBets(bets []int) {
	// At start of the game, the player will place separate bets,
	// for each hand that they want. Each of these original hands
	// will be considered a "master" hand.
	// the number of bets is the number of master hands the player
	// wants for this game.
	self.GameReset()

	for i := 0; i < len(bets); i++ {
		bet := bets[i]
		var player_master_hand *PlayerMasterHand
		player_master_hand = CreatePlayerMasterHand()
		player_master_hand.AddStartHand(bet)

		self.PlayerMasterHands = append(self.PlayerMasterHands, player_master_hand)
	}
}

//
// Dealer
//

type Dealer struct {
	DealerHand DealerHand
	Name       string
}

func CreateDealer() *Dealer {
	dealer := Dealer{
		DealerHand: *CreateDealerHand(),
		Name:       "Riverboat Dealer",
	}
	return &dealer
}

func (self *Dealer) GameReset() {
	self.DealerHand = *CreateDealerHand()
}

func (self *Dealer) TopCard() *cards.Card {
	return &self.DealerHand.Cards[0]
}

func (self *Dealer) HoleCard() *cards.Card {
	return &self.DealerHand.Cards[1]
}
