package game

import (
	"fmt"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/strategy"
)

type BlackJack struct {
	Shoe    []cards.Card
	ShoeTop int
	Players []*Player
}

func CreateBlackJack() *BlackJack {
	blackjack := BlackJack{
		Shoe:    cards.CreateShoe(),
		ShoeTop: 0,
		Players: []*Player{},
	}
	return &blackjack
}

func (self *BlackJack) Num_Players() int {
	return len(self.Players)
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

func (self *BlackJack) log(msg string) {
	fmt.Println(msg)
}

func (self *BlackJack) PlayGame() {
	if self.ShoeTop > house_rules.FORCE_RESHUFFLE {
		self.ReshuffleShoe()
	}

	// setting up the dealer and the player(s) could be done
	// by the caller and pass here via parameters.

	var dealer *Dealer = CreateDealer()

	var player1 *Player = CreatePlayer("John")
	var player2 *Player = CreatePlayer("Jane")

	self.SetPlayersForGame([]*Player{player1, player2})

	player1.SetGameBets([]int{2})
	player2.SetGameBets([]int{2, 2})

	//
	// DEAL HANDS
	//

	self.log("\n\nDEAL HANDS")

	var card cards.Card

	for i := 0; i < 2; i++ {
		for j := 0; j < self.Num_Players(); j++ {
			var player *Player = self.Players[j]
			for k := 0; k < player.Num_Master_Hands(); k++ {
				card = self.GetCardFromShoe()
				var master_hand *PlayerMasterHand = player.PlayerMasterHands[k]
				var first_hand *PlayerHand = master_hand.Hands[0]
				first_hand.AddCard(card)
			}
		}

		card = self.GetCardFromShoe()
		dealer.DealerHand.AddCard(card)
	}

	var dealer_top_card cards.Card = *dealer.TopCard()
	self.log(fmt.Sprintf("dealer top card: %v", dealer_top_card.Str()))

	var dealer_hole_card cards.Card = *dealer.HoleCard()

	//
	// PLAY HANDS
	//

	self.log("PLAY HANDS")

	if dealer.DealerHand.IsNatural() {
		// a real simulation would have to take care of Insurance, which is a sucker's bet,
		// so we just assume that no player will ask for insurance.
		// two cases:
		//     1. player has a natural and their bet is pushed
		//     2. player loses

		dealer.DealerHand.OutCome = HandOutcome(DEALER_BLACKJACK)

		for i := 0; i < self.Num_Players(); i++ {
			var player *Player = self.Players[i]
			for j := 0; j < player.Num_Master_Hands(); j++ {
				var master_hand *PlayerMasterHand = player.PlayerMasterHands[j]
				for k := 0; k < master_hand.Num_Hands(); k++ {
					// really should only be one hand in the master hand at this point
					var hand *PlayerHand = master_hand.Hands[k]
					// standing will do the right thing for both cases
					hand.OutCome = HandOutcome(STAND)
				}
			}
		}

	} else {
		// dealer does not have a natural
		for i := 0; i < self.Num_Players(); i++ {
			var player *Player = self.Players[i]
			self.log((fmt.Sprintf("player %v - %v", i+1, player.Name)))
			for j := 0; j < player.Num_Master_Hands(); j++ {
				var master_hand *PlayerMasterHand = player.PlayerMasterHands[j]
				for k := 0; k < master_hand.Num_Hands(); k++ {
					var hand *PlayerHand = master_hand.Hands[k]
					self.log(fmt.Sprintf("    hand %v.%v:", j+1, k+1))
					for l := 0; l < hand.Num_Cards(); l++ {
						var card cards.Card = hand.Cards[l]
						self.log(fmt.Sprintf("    card %v: %v", l+1, card.Str()))
					}

					var is_split_possible bool = master_hand.Num_Hands() < master_hand.HANDS_LIMIT

					// Need to make decisions pr player hand ...
					for {
						var decision strategy.PlayerDecision = strategy.DetermineBasicStrategyPlay(
							dealer_top_card, hand, is_split_possible,
						)
						self.log(fmt.Sprintf("    basic strategy: %v", decision))

						if decision == strategy.STAND {
							hand.OutCome = HandOutcome(STAND)
							self.log(fmt.Sprintf("    stand total H%v S%v", hand.HardCount(), hand.SoftCount()))
							break

						} else if decision == strategy.SURRENDER {
							hand.OutCome = HandOutcome(strategy.SURRENDER)
							hand.Bet = int(hand.Bet / 2)
							break

						} else if decision == strategy.HIT {
							card = self.GetCardFromShoe()
							hand.AddCard(card)
							hand_total := hand.Count()
							self.log(fmt.Sprintf("    hit: %v, total H%v S%v", card.Str(), hand.HardCount(), hand.SoftCount()))
							if hand_total > 21 {
								hand.OutCome = HandOutcome(BUST)
								self.log(fmt.Sprintf("    %v", hand.OutCome))
								break
							} else {
								hand.OutCome = HandOutcome(IN_PLAY)
							}

						} else if decision == strategy.SPLIT {
							var card1 cards.Card = self.GetCardFromShoe()
							var card2 cards.Card = self.GetCardFromShoe()
							var handIndex int = j
							var newHandIndex int = master_hand.SplitHand(handIndex, [2]cards.Card{card1, card2})
							self.log(fmt.Sprintf("   split, new hand index %v, adding cards %v, %v", newHandIndex, card1.Str(), card2.Str()))

						} else {
							self.log("FTW")
							hand.OutCome = HandOutcome(STAND)
							break
						}

					}
				}
			}
		}

		//
		// DEALER HANDS
		//

		self.log("DEALER HAND")
		self.log(fmt.Sprintf("dealer top card: %v", dealer_top_card.Str()))
		self.log(fmt.Sprintf("dealer hole card: %v", dealer_hole_card.Str()))
		var dealerDone bool = false
		for !dealerDone {
			hardCount := dealer.DealerHand.HardCount()
			softCount := dealer.DealerHand.SoftCount()
			var useSoftCount bool = hardCount < softCount && softCount < 21
			if useSoftCount && softCount < house_rules.DEALER_HITS_SOFT_ON {
				card = self.GetCardFromShoe()
				dealer.DealerHand.AddCard(card)
				self.log(fmt.Sprintf("    add: %v", card.Str()))
			} else if !useSoftCount && hardCount < house_rules.DEALER_HITS_HARD_ON {
				card = self.GetCardFromShoe()
				dealer.DealerHand.AddCard(card)
				self.log(fmt.Sprintf("    add: %v", card.Str()))
			} else {
				dealer.DealerHand.OutCome = HandOutcome(STAND)
				dealerDone = true
			}

			if dealer.DealerHand.Count() > 21 {
				dealer.DealerHand.OutCome = HandOutcome(BUST)
				dealerDone = true
			}
		}
	}

	//
	// SETTLE HANDS
	//

	self.log("SETTLE HAND")
}
