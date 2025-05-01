package tests

import "core:testing"
import "core:fmt"

import "../cards"
import "../game"
import "../strategy"

@(test)
test_determine_basic_strategy_play :: proc(t: ^testing.T) {
    for player_rank1 in cards.CardRank {
        for player_rank2 in cards.CardRank {
            for dealer_rank in cards.CardRank {
                player_hand: game.PlayerHand
                player_hand = game.create_player_hand(
                    from_split=false, bet=100,
                )
                defer game.free_cards(&player_hand)

                card1: cards.Card = cards.Card{
                    rank=player_rank1, suite=cards.CardSuite.SPADES,
                }
                game.add_card(&player_hand, card1)
                card2: cards.Card = cards.Card{
                    rank=player_rank2, suite=cards.CardSuite.HEARTS,
                }
                game.add_card(&player_hand, card1)

                dealer_top_card: cards.Card = cards.Card{
                    rank=dealer_rank, suite=cards.CardSuite.DIAMONDS,
                }

                player_decision: strategy.PlayerDecision
                player_decision = game.determine_basic_strategy_play(
                    dealer_top_card=dealer_top_card,
                    player_hand=&player_hand,
                    hand_allows_more_splits=true,
                )
            }
        }
    }
}

@(test)
test_play_game :: proc(t: ^testing.T) {
    blackjack := game.create_blackjack()
    defer game.free_blackjack(&blackjack)

    game.play_game(&blackjack)
}
