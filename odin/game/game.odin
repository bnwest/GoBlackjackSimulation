package game

import "core:fmt"

import "../cards"
import house_rules "../rules"
import "../strategy"

BlackJackPlayerResults :: struct {
    hands_played: uint,
    hands_won: uint,
    hands_lost: uint,
    hands_pushed: uint,
    proceeds: int,
}

BlackJackStats :: struct {
    double_down_count: uint,
    surrender_count: uint,
    split_count: uint,
    aces_split_count: uint,
}

create_blackjack_stats :: proc() -> BlackJackStats {
    // default behavior is to zero out menory
    // => all of the BlackJackStats fields are initialized to zero.
    return BlackJackStats{}
}

BlackJack :: struct {
    shoe: [dynamic]cards.Card,
    shoe_top: uint,
    // players and dealer are external entities
    players: [dynamic]^Player,
    dealer: ^Dealer,
    // collect stats across multiple games
    results: map[string]BlackJackPlayerResults,
    stats: BlackJackStats,
}

create_blackjack :: proc() -> BlackJack {
    blackjack := BlackJack{
        shoe=cards.create_shoe(),
        shoe_top=0,
        players=[dynamic]^Player{},
        results=map[string]BlackJackPlayerResults{},
        stats=create_blackjack_stats(),
    }
    return blackjack
}

free_blackjack :: proc(self: ^BlackJack) {
    delete(self.shoe)
    self.shoe = [dynamic]cards.Card{}
    // players are external objects
    // which are not owned/managed by BlackJack
    // for player in self.players {
    //     free_player(player)
    // }
    delete(self.players)
    self.players = [dynamic]^Player{}
    delete(self.results)
    self.results = map[string]BlackJackPlayerResults{}
}

num_players :: proc(self: ^BlackJack) -> uint {
    return len(self.players)
}

reshuffle_shoe :: proc(self: ^BlackJack) {
    cards.shuffle_shoe(self.shoe)
    self.shoe_top = 0
}

get_card_from_shoe :: proc(self: ^BlackJack) -> cards.Card {
    card := self.shoe[self.shoe_top]
    self.shoe_top += 1
    return card
}

set_dealer :: proc(self: ^BlackJack, dealer: ^Dealer) {
    self.dealer = dealer
}

set_players_for_game :: proc(
    self: ^BlackJack, 
    players: [dynamic]^Player
) {
    reset_players(self)
    self.players = [dynamic]^Player{}
    for player in players {
        append(&self.players, player)
        if player.name not_in self.results {
            // default behavior is to zero out menory
            // => all of the BlackJackPlayerResults fields are initialized to zero.
            self.results[player.name] = BlackJackPlayerResults{}
        }
    }
}

reset_players :: proc(self: ^BlackJack) {
    // players space is not managed by BlackJack
    // => do not free the player space
    delete(self.players)
    self.players = [dynamic]^Player{}
    // self.players is now zeroed memory
}

abs :: proc(i: int) -> uint {
    // math.abs(T) -> T which is so very wrong
    return i < 0 ? uint(-i) : uint(i)
}

add_result :: proc(
    self: ^BlackJack,
    player: ^Player,
    hand_index: uint,
    player_hand: ^PlayerHand,
    initial_bet: uint,
    result: int,
) {
    // Cannot assign to struct field in map
    // self.results[player.name].hands_played += 1
    // workaround: 
    //     get struct copy, modify copy and reassign copy back
    new_results := self.results[player.name]
    new_results.hands_played += 1
    if result > 0 {
        new_results.hands_won += 1
    } else if result < 0 {
        new_results.hands_lost += 1
    } else {
        new_results.hands_pushed += 1
    }
    new_results.proceeds += result
    self.results[player.name] = new_results

    is_double_down: bool
    is_double_down = (
        num_cards(player_hand) == 3
        && (initial_bet * 2) == abs(result)
    )
    if is_double_down {
        self.stats.double_down_count += 1
    }

    if player_hand.outcome == HandOutcome.SURRENDER {
        self.stats.surrender_count += 1
    }

    if hand_index > 0 {
        self.stats.split_count += 1
    }

    splitting_aces: bool
    splitting_aces = (
        player_hand.from_split
        && player_hand.cards[0].rank == cards.CardRank.ACE
    )
    if splitting_aces {
        self.stats.aces_split_count += 1
    }
}

log :: proc(msg: string) {
    // log to stderr
    fmt.eprintln(msg)
}

play_game :: proc(self: ^BlackJack) {
    if self.shoe_top > house_rules.FORCE_RESHUFFLE {
        reshuffle_shoe(self)
    }

	// setting up the dealer and the player(s) could be done
	// by the caller and pass here via parameters.

    dealer: Dealer = create_dealer()
    set_dealer(self, &dealer)
    defer free_dealer(&dealer)

    player1: Player = create_player(name="Jack")
    defer free_player(&player1)

    player2: Player = create_player(name="Jill")
    defer free_player(&player2)

    players := [dynamic]^Player{}
    append(&players, &player1, &player2)
    defer {
        delete(players)
        players := [dynamic]^Player{}
    }

    set_players_for_game(self, players=players)

    // create one player master hand per bet

    initial_bet: uint = 2
    bets: [dynamic]uint

    clear(&bets)
    append(&bets, initial_bet)
    set_game_bets(self.players[0], bets=bets)

    clear(&bets)
    append(&bets, initial_bet, initial_bet)
    set_game_bets(self.players[1], bets=bets)

    defer {
        delete(bets)
        bets = [dynamic]uint{}
    }
 
    //
	// DEAL HANDS
	//

	log("\n\nDEAL HANDS")

    card: cards.Card

    for i in 0..<2 {
        // log(fmt.tprintf("deal round: {}", i+1))
        for &player in self.players {
            for &master_hand in player.master_hands {
                card = get_card_from_shoe(self)
                add_card(&master_hand.hands[0], card)
            }
        }

        card = get_card_from_shoe(self)
        add_card(&self.dealer.dealer_hand, card)
    }

    dealer_top_card: cards.Card = top_card(self.dealer)
    dealer_hole_card: cards.Card = hole_card(self.dealer)
    log(
        fmt.tprintf("dealer top  card: {}", cards.to_string(dealer_top_card))
    )

    //
	// PLAY HANDS
	//

	log("PLAY HANDS")

    if is_natural(&self.dealer.dealer_hand) {
		// a real simulation would have to take care of Insurance, which is a sucker's bet,
		// so we just assume that no player will ask for insurance.
		// two cases:
		//     1. player has a natural and their bet is pushed
		//     2. player loses

        self.dealer.dealer_hand.outcome = HandOutcome.DEALER_BLACKJACK
        for &player in self.players {
            for &master_hand in player.master_hands {
                for &hand in master_hand.hands {
					// standing will do the right thing in the settlement logic below
                    hand.outcome = HandOutcome.STAND
                }
            }
        }

    } else {
		// dealer does not have a natural

        for &player, i in self.players {
            log(fmt.tprintf("player - {} - {}", i+1, player.name))
            for &master_hand, j in player.master_hands {
                for &hand, k in master_hand.hands {
                    log(fmt.tprintf("    hand {}.{}:", j+1, k+1))
                    for &card, l in hand.cards {
                        log(fmt.tprintf("        card {}: {}", l+1, cards.to_string(card)))
                    }

                    is_split_possible: bool = num_hands(&master_hand) < master_hand.HANDS_LIMIT

                    // Need to make decisions per player hand ...
                    for {
                        if hand.outcome == HandOutcome.STAND {
							// product of a prior ace split, outcome has already been determined.
                            log(
                                fmt.tprintf(
                                    "        prior aces split: {}, total H{} S{}",
                                    strategy.PlayerDecision.STAND, hard_count(&hand), soft_count(&hand)
                                )
                            )
							break
                        }

                        decision: strategy.PlayerDecision
                        decision = determine_basic_strategy_play(
                            dealer_top_card, &hand, is_split_possible
                        )
                        log(fmt.tprintf("        basic strategy: {}", strategy.to_string(decision)))

                        if decision == strategy.PlayerDecision.STAND {
                            hand.outcome = HandOutcome.STAND
                            log(fmt.tprintf("        stand total H{} S{}", hard_count(&hand), soft_count(&hand)))
                            break

                        } else if decision == strategy.PlayerDecision.SURRENDER {
                            hand.outcome = HandOutcome.SURRENDER
                            hand.bet /= 2
                            break

                        } else if decision == strategy.PlayerDecision.DOUBLE {
                            card = get_card_from_shoe(self)
                            add_card(&hand, card)
                            hand.bet *= 2
                            log(fmt.tprintf("        hit: {}, total H{} S{}", cards.to_string(card), hard_count(&hand), soft_count(&hand)))
                            hand.outcome = HandOutcome.STAND
                            log(fmt.tprintf("        stand total H{} S{}", hard_count(&hand), soft_count(&hand)))
                            break

                        } else if decision == strategy.PlayerDecision.HIT {
                            card = get_card_from_shoe(self)
                            add_card(&hand, card)
                            hand_total := count(&hand)
                            log(fmt.tprintf("        hit: {}, total H{} S{}", cards.to_string(card), hard_count(&hand), soft_count(&hand)))
                            if hand_total > 21 {
                                hand.outcome = HandOutcome.BUST
                                log(fmt.tprintf("        {}", to_string(hand.outcome)))
								break
                            } else {
                                hand.outcome = HandOutcome.IN_PLAY
                            }

                        } else if decision == strategy.PlayerDecision.SPLIT {
                            card1: cards.Card = get_card_from_shoe(self)
                            card2: cards.Card = get_card_from_shoe(self)
                            split_cards := [2]cards.Card{card1, card2}
                            hand_index: uint = uint(k)
                            new_hand_index := split_hand(&master_hand, hand_index, split_cards)
                            log(fmt.tprintf("        split, new hand index {}, adding cards {}, {}", new_hand_index+1, cards.to_string(split_cards[0]), cards.to_string(split_cards[1])))
                            log(fmt.tprintf("        card 1: {}", cards.to_string(hand.cards[0])))
                            log(fmt.tprintf("        card 2: {}", cards.to_string(hand.cards[1])))
                            splitting_aces: bool = hand.cards[0].rank == cards.CardRank.ACE
                            if splitting_aces && house_rules.NO_MORE_CARDS_AFTER_SPLITTING_ACES {
                                hand.outcome = HandOutcome.STAND
                                log(fmt.tprintf("        aces split: {}, total H{} S{}", to_string(hand.outcome), hard_count(&hand), soft_count(&hand)))
                                master_hand.hands[new_hand_index].outcome = HandOutcome.STAND
                                break
                            }

                        } else {
                            log(fmt.tprintf("FTW"))
                            log(fmt.tprintf("FTW: dealer_top_card: {}, is_split_possible: {}", cards.to_string(dealer_top_card), is_split_possible))
                            log(fmt.tprintf("FTW: player hand count: H{}, S{}", hard_count(&hand), soft_count(&hand)))
                            log(fmt.tprintf("FTW: decision: {}", strategy.to_string(decision)))
                            hand.outcome = HandOutcome.STAND
                            break
                        }
                    }
                }
            }
        }

        //
		// DEALER HAND
		//

        log("DEALER HAND")
        log(
            fmt.tprintf("dealer top  card: {}", cards.to_string(dealer_top_card))
        )
        log(
            fmt.tprintf("dealer hole card: {}", cards.to_string(dealer_hole_card))
        )

        dealer_done: bool = false
        for !dealer_done {
            hard_count := hard_count(&dealer.dealer_hand)
            soft_count := soft_count(&dealer.dealer_hand)
            use_soft_count: bool = hard_count < soft_count && soft_count <= 21
            if use_soft_count && soft_count < house_rules.DEALER_HITS_SOFT_ON {
                card = get_card_from_shoe(self)
                add_card(&dealer.dealer_hand, card)
                log(fmt.tprintf("    add: {}", cards.to_string(card)))
            } else if !use_soft_count && hard_count < house_rules.DEALER_HITS_HARD_ON {
                card = get_card_from_shoe(self)
                add_card(&dealer.dealer_hand, card)
                log(fmt.tprintf("    add: {}", cards.to_string(card)))
            } else {
                dealer.dealer_hand.outcome = HandOutcome.STAND
                dealer_done = true
                log(fmt.tprintf("   stand total H{} S{}", hard_count, soft_count))
            }

            if count(&dealer.dealer_hand) > 21 {
                dealer.dealer_hand.outcome = HandOutcome.BUST
                dealer_done = true
                log(fmt.tprintf("    bust"))
            }
        }
    }

    //
	// SETTLE HANDS
	//

	log("SETTLE HAND")

    if dealer.dealer_hand.outcome == HandOutcome.DEALER_BLACKJACK {
        for &player, i in self.players {
            log(fmt.tprintf("player - {} - {}", i+1, player.name))
            for &master_hand, j in player.master_hands {
                for &hand, k in master_hand.hands {
                    if is_natural(&hand) {
                        add_result(self, player, uint(k), &hand, initial_bet, 0)
                        log(fmt.tprintf("    hand {}.{}: push both player and dealer had naturals", j+1, k+1))
                    } else {
                        add_result(self, player, uint(k), &hand, initial_bet, -int(hand.bet))
                        log(fmt.tprintf("    hand {}.{}: dealer natural: lost ${}", j+1, k+1, hand.bet))
                    }
                }
            }
        }

    } else {
		// dealer does not have a natural
        for &player, i in self.players {
            log(fmt.tprintf("player - {} - {}", i+1, player.name))
            for &master_hand, j in player.master_hands {
                for &hand, k in master_hand.hands {
                    if hand.outcome == HandOutcome.BUST {
                        add_result(self, player, uint(k), &hand, initial_bet, -int(hand.bet))
                        log(fmt.tprintf("    hand {}.{}: bust: lost ${}", j+1, k+1, hand.bet))
                    } else if hand.outcome == HandOutcome.SURRENDER {
                        add_result(self, player, uint(k), &hand, initial_bet, -int(hand.bet))
                        log(fmt.tprintf("    hand {}.{}: surrender: lost ${}", j+1, k+1, hand.bet))
                    } else {
						// player has a non-bust, non-surrender hand
                        if is_natural(&hand) {
                            payout: int = int(f32(hand.bet) * house_rules.NATURAL_BLACKJACK_PAYOUT)
                            add_result(self, player, uint(k), &hand, initial_bet, payout)
                            log(fmt.tprintf("    hand {}.{}: won ${}", j+1, k+1, payout))
                        } else if dealer.dealer_hand.outcome == HandOutcome.BUST {
                            add_result(self, player, uint(k), &hand, initial_bet, int(hand.bet))
                            log(fmt.tprintf("    hand {}.{}: dealer bust: won ${}", j+1, k+1, hand.bet))
                        } else {
                            if count(&hand) < count(&dealer.dealer_hand) {
                                add_result(self, player, uint(k), &hand, initial_bet, -int(hand.bet))
                                log(fmt.tprintf("    hand {}.{}: lost ${}", j+1, k+1, hand.bet))
                            } else if count(&hand) > count(&dealer.dealer_hand) {
                                add_result(self, player, uint(k), &hand, initial_bet, int(hand.bet))
                                log(fmt.tprintf("    hand {}.{}: won ${}", j+1, k+1, hand.bet))
                            } else {
                                add_result(self, player, uint(k), &hand, initial_bet, 0)
                                log(fmt.tprintf("    hand {}.{}: push", j+1, k+1))
                            }
                        }
                    }
                }
            }
        }
    }

    //
    // CLEAN THIS MESS UP
    //

    free_dealer(self.dealer)
    for player in self.players {
        free_player(player)
    }
}
