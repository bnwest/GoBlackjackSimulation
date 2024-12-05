package game

import (
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"
)

type BlackJack struct {
	Shoe []cards.Card
	ShoeTop int
	Players []*Player
}

func CreateBlackJack() *BlackJack {
	blackjack := BlackJack{
		Shoe: cards.CreateShoe(),
		ShoeTop: 0,
		Players: []*Player{},
	}
	return &blackjack
}

func (self *BlackJack) ReshuffleShoe() {
	cards.ShuffleShoe(self.Shoe)
}

func (self *BlackJack) GetCardFromShoe() cards.Card {
	card := self.Shoe[self.ShoeTop]
	self.ShoeTop++
	return card
}

func (self *BlackJack) SetPlayersForGame(players []*Player) {
	self.Players = players
}

func (self *BlackJack) PlayGame() {
	if self.ShoeTop > house_rules.FORCE_RESHUFFLE {
		self.ReshuffleShoe()
	}

	//var dealer *Dealer = CreateDealer()

	var player1 *Player = CreatePlayer("John")
	var player2 *Player = CreatePlayer("Jane")

	self.SetPlayersForGame([]*Player{player1, player2})

	player1.SetGameBets([]int{2})
	player2.SetGameBets([]int{2, 2})
}
