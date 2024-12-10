package game

import (
	"fmt"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
	house_rules "github.com/bnwest/GoBlackjackSimulation/go/blackjack/rules"
	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/strategy"
)

type BlackJackResults struct {
	HandsPlayed int
	HandsWon    int
	HandsLost   int
	HandsPushed int
	Proceeds    int
}

type BlackJackStats struct {
	DoubleDownCount int
	SurrenderCount  int
	AcesSplit       int
}

func CreateBlackJackStats() BlackJackStats {
	return BlackJackStats{
		DoubleDownCount: 0,
		SurrenderCount:  0,
		AcesSplit:       0,
	}
}

type BlackJack struct {
	Shoe    []cards.Card
	ShoeTop int
	Players []*Player
	Results map[string]*BlackJackResults
	Stats   BlackJackStats
}

func CreateBlackJack() *BlackJack {
	blackjack := BlackJack{
		Shoe:    cards.CreateShoe(),
		ShoeTop: 0,
		Players: []*Player{},
		Results: make(map[string]*BlackJackResults),
		Stats:   CreateBlackJackStats(),
	}
	return &blackjack
}

func (self *BlackJack) NumPlayers() int {
	return len(self.Players)
}

func (self *BlackJack) ReshuffleShoe() {
	cards.ShuffleShoe(self.Shoe)
	self.ShoeTop = 0
}

func (self *BlackJack) GetCardFromShoe() cards.Card {
	card := self.Shoe[self.ShoeTop]
	self.ShoeTop++
	return card
}

func (self *BlackJack) SetPlayersForGame(players []*Player) {
	self.Players = players
	for i := 0; i < self.NumPlayers(); i++ {
		var player *Player = self.Players[i]
		_, ok := self.Results[player.Name]
		if !ok {
			self.Results[player.Name] = &BlackJackResults{
				HandsPlayed: 0,
				HandsWon:    0,
				HandsLost:   0,
				HandsPushed: 0,
				Proceeds:    0,
			}
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (self *BlackJack) AddResult(
	player *Player,
	playerHand *PlayerHand,
	initialBet int,
	result int,
) {
	self.Results[player.Name].HandsPlayed++
	if result > 0 {
		self.Results[player.Name].HandsWon++
	} else if result < 0 {
		self.Results[player.Name].HandsLost++
	} else {
		self.Results[player.Name].HandsPushed++
	}
	self.Results[player.Name].Proceeds += result

	isDoubleDown := playerHand.NumCards() == 3 && abs(initialBet)*2 == abs(result)
	if isDoubleDown {
		self.Stats.DoubleDownCount++
	}

	if playerHand.OutCome == HandOutcome(SURRENDER) {
		self.Stats.SurrenderCount++
	}

	splittingAces := playerHand.FromSplit && playerHand.Cards[0].Rank == cards.ACE
	if splittingAces {
		self.Stats.AcesSplit++
	}
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

	initialBet := 2
	player1.SetGameBets([]int{initialBet})
	player2.SetGameBets([]int{initialBet, initialBet})

	//
	// DEAL HANDS
	//

	self.log("\n\nDEAL HANDS")

	var card cards.Card

	for i := 0; i < 2; i++ {
		for j := 0; j < self.NumPlayers(); j++ {
			var player *Player = self.Players[j]
			for k := 0; k < player.NumMasterHands(); k++ {
				card = self.GetCardFromShoe()
				var masterHand *PlayerMasterHand = player.PlayerMasterHands[k]
				var firstHand *PlayerHand = masterHand.Hands[0]
				firstHand.AddCard(card)
			}
		}

		card = self.GetCardFromShoe()
		dealer.DealerHand.AddCard(card)
	}

	var dealerTopCard cards.Card = dealer.TopCard()
	self.log(fmt.Sprintf("dealer top card: %v", dealerTopCard.Str()))

	var dealerHoleCard cards.Card = dealer.HoleCard()

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

		for i := 0; i < self.NumPlayers(); i++ {
			var player *Player = self.Players[i]
			for j := 0; j < player.NumMasterHands(); j++ {
				var masterHand *PlayerMasterHand = player.PlayerMasterHands[j]
				for k := 0; k < masterHand.NumHands(); k++ {
					// really should only be one hand in the master hand at this point
					var hand *PlayerHand = masterHand.Hands[k]
					// standing will do the right thing for both cases
					hand.OutCome = HandOutcome(STAND)
				}
			}
		}

	} else {
		// dealer does not have a natural
		for i := 0; i < self.NumPlayers(); i++ {
			var player *Player = self.Players[i]
			self.log((fmt.Sprintf("player %v - %v", i+1, player.Name)))
			for j := 0; j < player.NumMasterHands(); j++ {
				var masterHand *PlayerMasterHand = player.PlayerMasterHands[j]
				for k := 0; k < masterHand.NumHands(); k++ {
					var hand *PlayerHand = masterHand.Hands[k]
					self.log(fmt.Sprintf("    hand %v.%v:", j+1, k+1))
					for l := 0; l < hand.NumCards(); l++ {
						var card cards.Card = hand.Cards[l]
						self.log(fmt.Sprintf("    card %v: %v", l+1, card.Str()))
					}

					var isSplitPossible bool = masterHand.NumHands() < masterHand.HANDS_LIMIT

					// Need to make decisions pr player hand ...
					for {
						if hand.OutCome == HandOutcome(STAND) {
							// product of a prior ace split, outcome has already been determined.
							self.log(fmt.Sprintf("    prior aces split: %v", strategy.PlayerDecision(strategy.STAND)))
							break
						}

						var decision strategy.PlayerDecision = strategy.DetermineBasicStrategyPlay(
							dealerTopCard, hand, isSplitPossible,
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

						} else if decision == strategy.DOUBLE {
							card = self.GetCardFromShoe()
							hand.AddCard(card)
							hand.Bet *= 2
							self.log(fmt.Sprintf("    hit: %v, total H%v S%v", card.Str(), hand.HardCount(), hand.SoftCount()))
							hand.OutCome = HandOutcome(STAND)
							self.log(fmt.Sprintf("    stand total H%v S%v", hand.HardCount(), hand.SoftCount()))
							break

						} else if decision == strategy.HIT {
							card = self.GetCardFromShoe()
							hand.AddCard(card)
							handTotal := hand.Count()
							self.log(fmt.Sprintf("    hit: %v, total H%v S%v", card.Str(), hand.HardCount(), hand.SoftCount()))
							if handTotal > 21 {
								hand.OutCome = HandOutcome(BUST)
								self.log(fmt.Sprintf("    %v", hand.OutCome))
								break
							} else {
								hand.OutCome = HandOutcome(IN_PLAY)
							}

						} else if decision == strategy.SPLIT {
							var card1 cards.Card = self.GetCardFromShoe()
							var card2 cards.Card = self.GetCardFromShoe()
							var handIndex int = k
							var newHandIndex int = masterHand.SplitHand(handIndex, [2]cards.Card{card1, card2})
							self.log(fmt.Sprintf("    split, new hand index %v, adding cards %v, %v", newHandIndex+1, card1.Str(), card2.Str()))
							self.log(fmt.Sprintf("    new card 2: %v", card1.Str()))
							splittingAces := hand.Cards[0].Rank == cards.ACE
							if splittingAces && house_rules.NO_MORE_CARDS_AFTER_SPLITTING_ACES {
								hand.OutCome = HandOutcome(STAND)
								self.log(fmt.Sprintf("    aces split: %v", strategy.PlayerDecision(strategy.STAND)))
								masterHand.Hands[newHandIndex].OutCome = HandOutcome(STAND)
								break
							}

						} else {
							self.log("FTW")
							self.log(fmt.Sprintf("FTW: dealerTopCard: %v, isSplitPossible %v", dealerTopCard.Str(), isSplitPossible))
							self.log(fmt.Sprintf("FTW: player hand count: H%v S%v", hand.HardCount(), hand.SoftCount()))
							self.log(fmt.Sprintf("FTW: decision: %v", decision))
							hand.OutCome = HandOutcome(STAND)
							break
						}
					}
				}
			}
		}

		//
		// DEALER HAND
		//

		self.log("DEALER HAND")
		self.log(fmt.Sprintf("dealer top card: %v", dealerTopCard.Str()))
		self.log(fmt.Sprintf("dealer hole card: %v", dealerHoleCard.Str()))
		var dealerDone bool = false
		for !dealerDone {
			hardCount := dealer.DealerHand.HardCount()
			softCount := dealer.DealerHand.SoftCount()
			var useSoftCount bool = hardCount < softCount && softCount < 21
			if useSoftCount && softCount <= house_rules.DEALER_HITS_SOFT_ON {
				card = self.GetCardFromShoe()
				dealer.DealerHand.AddCard(card)
				self.log(fmt.Sprintf("    add: %v", card.Str()))
			} else if !useSoftCount && hardCount <= house_rules.DEALER_HITS_HARD_ON {
				card = self.GetCardFromShoe()
				dealer.DealerHand.AddCard(card)
				self.log(fmt.Sprintf("    add: %v", card.Str()))
			} else {
				dealer.DealerHand.OutCome = HandOutcome(STAND)
				dealerDone = true
				self.log(fmt.Sprintf("    stand: total H%v S%v", dealer.DealerHand.HardCount(), dealer.DealerHand.SoftCount()))
			}

			if dealer.DealerHand.Count() > 21 {
				dealer.DealerHand.OutCome = HandOutcome(BUST)
				dealerDone = true
				self.log("    bust")
			}
		}
	}

	//
	// SETTLE HANDS
	//

	self.log("SETTLE HAND")
	if dealer.DealerHand.OutCome == HandOutcome(DEALER_BLACKJACK) {
		for i := 0; i < self.NumPlayers(); i++ {
			var player *Player = self.Players[i]
			self.log(fmt.Sprintf("Player %v - %v", i+1, player.Name))
			for j := 0; j < player.NumMasterHands(); j++ {
				var masterHand *PlayerMasterHand = player.PlayerMasterHands[j]
				for k := 0; k < masterHand.NumHands(); k++ {
					var hand *PlayerHand = masterHand.Hands[k]
					if hand.IsNatural() {
						self.AddResult(player, hand, initialBet, 0)
						self.log(fmt.Sprintf("    hand %v.%v: push both player and dealer had naturals", j+1, k+1))
					} else {
						self.AddResult(player, hand, initialBet, -hand.Bet)
						self.log(fmt.Sprintf("    hand %v.%v: lost $%v", j+1, k+1, hand.Bet))
					}
				}
			}
		}

	} else {
		// dealer does not have a natural
		for i := 0; i < self.NumPlayers(); i++ {
			var player *Player = self.Players[i]
			self.log(fmt.Sprintf("Player %v - %v", i+1, player.Name))
			for j := 0; j < player.NumMasterHands(); j++ {
				var masterHand *PlayerMasterHand = player.PlayerMasterHands[j]
				for k := 0; k < masterHand.NumHands(); k++ {
					var hand *PlayerHand = masterHand.Hands[k]
					if hand.OutCome == HandOutcome(BUST) {
						self.AddResult(player, hand, initialBet, -hand.Bet)
						self.log(fmt.Sprintf("    hand %v.%v: bust: lost $%v", j+1, k+1, hand.Bet))

					} else if hand.OutCome == HandOutcome(SURRENDER) {
						self.AddResult(player, hand, initialBet, -hand.Bet)
						self.log(fmt.Sprintf("    hand %v.%v: surrender: lost $%v", j+1, k+1, hand.Bet))

					} else {
						// player has a non-bust, non-surrender hand
						if hand.IsNatural() {
							var payout int = int(float32(hand.Bet) * house_rules.NATURAL_BLACKJACK_PAYOUT)
							self.AddResult(player, hand, initialBet, payout)
							self.log(fmt.Sprintf("    hand %v.%v: natural: won $%v", j+1, k+1, payout))

						} else if dealer.DealerHand.OutCome == HandOutcome(BUST) {
							self.AddResult(player, hand, initialBet, hand.Bet)
							self.log(fmt.Sprintf("    hand %v.%v: won: dealer bust: won $%v", j+1, k+1, hand.Bet))

						} else {
							if hand.Count() < dealer.DealerHand.Count() {
								self.AddResult(player, hand, initialBet, -hand.Bet)
								self.log(fmt.Sprintf("    hand %v.%v: lost $%v", j+1, k+1, hand.Bet))

							} else if hand.Count() > dealer.DealerHand.Count() {
								self.AddResult(player, hand, initialBet, hand.Bet)
								self.log(fmt.Sprintf("    hand %v.%v: won $%v", j+1, k+1, hand.Bet))

							} else {
								self.AddResult(player, hand, initialBet, 0)
								self.log(fmt.Sprintf("    hand %v.%v: push", j+1, k+1))
							}
						}
					}
				}
			}
		}
	}
}
