package strategy

type Decision string

const (
	S  Decision = "stand"
	H  Decision = "hit"
	Dh Decision = "double-down-if-allowed-or-hit"
	Ds Decision = "double-down-if-allowed-or-stand"
	SP Decision = "split"
	// U => Surrender, in a world of too many S words
	Uh  Decision = "surrender-if-allowed-or-hit"
	Us  Decision = "surrender-if-allowed-or-stand"
	Usp Decision = "surrender-if-allowed-or-split"
	NO  Decision = "no-decision"
)

func IsValidDecision(decision Decision) bool {
	// Decision("foobar") will not generate a run time exception
	switch decision {
	case Decision(S), Decision(H), Decision(Dh), Decision(Ds), Decision(SP), Decision(Uh), Decision(Us), Decision(Usp):
		return true
	}
	return false
}

type PlayerDecision string

const (
	STAND     PlayerDecision = "stand"
	HIT       PlayerDecision = "hit"
	DOUBLE    PlayerDecision = "double-down"
	SPLIT     PlayerDecision = "split"
	SURRENDER PlayerDecision = "surrender"
)

func IsValidPlayerDecision(player_decision PlayerDecision) bool {
	switch player_decision {
	case PlayerDecision(STAND), PlayerDecision(HIT), PlayerDecision(DOUBLE), PlayerDecision(SPLIT), PlayerDecision(SURRENDER):
		return true
	}
	return false
}
