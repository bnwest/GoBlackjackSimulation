package tests

import "core:testing"
import "core:fmt"

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
            "map decision {} to string? {}",
            player_decision,
            player_decision_str,
        )
    }
}
