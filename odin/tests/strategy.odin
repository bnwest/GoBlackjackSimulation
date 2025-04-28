package tests

import "core:testing"
import "core:fmt"

import "../cards"
import "../game"
import "../strategy"

@(test)
test_decisons :: proc(t: ^testing.T) {
    for decision in strategy.Decision {
        decision_str := strategy.to_string(decision)
        testing.expectf(
            t,
            len(decision_str) > 0,
            "map decision {} to string? {}",
            decision,
            decision_str,
        )
    }
}

@(test)
test_player_decisons :: proc(t: ^testing.T) {
    for player_decision in strategy.PlayerDecision {
        player_decision_str := strategy.to_string(player_decision)
        testing.expectf(
            t,
            len(player_decision_str) > 0,
            "map player decision {} to string? {}",
            player_decision,
            player_decision_str,
        )
    }
}

@(test)
test_get_hard_total_decision :: proc(t: ^testing.T) {
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

                player_total: uint
                player_total = game.hard_count(&player_hand)

                dealer_top_card: cards.Card = cards.Card{
                    rank=dealer_rank, suite=cards.CardSuite.DIAMONDS,
                }

                decision: strategy.Decision
                decision = strategy.get_hard_total_decision(
                    player_total=player_total,
                    dealer_top_card=dealer_top_card.rank,
                )
            }
        }
    }
}

@(test)
test_get_soft_total_decision :: proc(t: ^testing.T) {
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

                player_total: uint
                player_total = game.soft_count(&player_hand)

                dealer_top_card: cards.Card = cards.Card{
                    rank=dealer_rank, suite=cards.CardSuite.DIAMONDS,
                }

                decision: strategy.Decision
                decision = strategy.get_soft_total_decision(
                    player_total=player_total,
                    dealer_top_card=dealer_top_card.rank,
                )
            }
        }
    }
}

@(test)
test_get_pairs_split_decision :: proc(t: ^testing.T) {
    for pair_rank in cards.CardRank {
        for dealer_rank in cards.CardRank {
            player_hand: game.PlayerHand
            player_hand = game.create_player_hand(
                from_split=false, bet=100,
            )
            defer game.free_cards(&player_hand)

            card1: cards.Card = cards.Card{
                rank=pair_rank, suite=cards.CardSuite.SPADES,
            }
            game.add_card(&player_hand, card1)
            card2: cards.Card = cards.Card{
                rank=pair_rank, suite=cards.CardSuite.HEARTS,
            }
            game.add_card(&player_hand, card1)

            dealer_top_card: cards.Card = cards.Card{
                rank=dealer_rank, suite=cards.CardSuite.DIAMONDS,
            }

            decision: strategy.Decision
            decision = strategy.get_pairs_split_decision(
                player_rank_for_pair=pair_rank,
                dealer_top_card=dealer_top_card.rank,
            )
        }
    }
}

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
                player_decision = strategy.determine_basic_strategy_play(
                    dealer_top_card=dealer_top_card,
                    player_hand=&player_hand,
                    hand_allows_more_splits=true,
                )
            }
        }
    }
}
