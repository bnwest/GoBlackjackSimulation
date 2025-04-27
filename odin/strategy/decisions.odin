package strategy

Decision :: enum {
    S,
    H,
    Dh,
    Ds,
    SP,
	// U => Surrender, in a world of too many S words
    Uh,
    Us,
    Usp,
    NO,
}

decision_string := [Decision]string {
    .S   = "stand",
    .H   = "hit",
    .Dh  = "double-down-if-allowed-or-hit",
    .Ds  = "double-down-if-allowed-or-stand",
    .SP  = "split",
    .Uh  = "surrender-if-allowed-or-hit",
    .Us  = "surrender-if-allowed-or-stand",
    .Usp = "surrender-if-allowed-or-split",
    .NO  = "no-decision",
}

decision_to_string :: proc(
    decision: Decision
) -> string {
    return decision_string[decision]
}

PlayerDecision :: enum {
    STAND,
    HIT,
    DOUBLE,
    SPLIT,
    SURRENDER,
}

player_decision_string := [PlayerDecision]string {
    .STAND     = "stand",
    .HIT       = "hit",
    .DOUBLE    = "double",
    .SPLIT     = "split",
    .SURRENDER = "surrender",
}

player_decision_to_string :: proc(
    player_decision: PlayerDecision
) -> string {
    return player_decision_string[player_decision]
}

to_string :: proc {
    decision_to_string,
    player_decision_to_string,
}
