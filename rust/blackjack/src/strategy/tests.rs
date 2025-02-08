// file src/strategy/tests.rs defines project module "strategy::tests".

use super::*;

#[test]
fn test_decisions() {
    const TOTAL_NUM_DECISION: u32 = 9;
    let mut num_decisions: u32 = 0;
    for _decision in Decision::iterator() {
        num_decisions += 1;
    }
    assert_eq!(num_decisions, TOTAL_NUM_DECISION);

    for decision in Decision::iterator() {
        // decision: &Decision
        let discrim = decision.discriminant();
        let decision_roundtrip: Decision = Decision::transmute(discrim);
        assert_eq!(*decision, decision_roundtrip);
        // binary operation `==` cannot be applied to type `strategy::Decision`
        // need: #[derive(PartialEq)]
        // `strategy::Decision` cannot be formatted using `{:?}`
        // need: #[derive(Debug)]
    }

    for decision in Decision::iterator() {
        // println!("{}", decision);
        // the trait `std::fmt::Display` is not implemented for `strategy::Decision`
        println!("{}", decision.to_string());
        // `strategy::Decision` cannot be formatted using `{:?}`
        // need: #[derive(Debug)]
        println!("{:?}", decision.to_string());
        println!("{:#?}", decision.to_string());

        // println!("{decision}");
        // the trait `std::fmt::Display` is not implemented for `strategy::Decision`
        println!("{decision:?}");
        println!("{decision:#?}");
    }
}

#[test]
fn test_player_decisions() {
    const TOTAL_NUM_PLAYER_DECISION: u32 = 5;
    let mut num_player_decisions: u32 = 0;
    for _decision in PlayerDecision::iterator() {
        num_player_decisions += 1;
    }
    assert_eq!(num_player_decisions, TOTAL_NUM_PLAYER_DECISION);

    for decision in PlayerDecision::iterator() {
        // decision: &PlayerDecision
        let discrim = decision.discriminant();
        let decision_roundtrip: PlayerDecision = PlayerDecision::transmute(discrim);
        assert_eq!(*decision, decision_roundtrip);
    }

    for decision in PlayerDecision::iterator() {
        // println!("{}", decision);
        println!("{}", decision.to_string());
        println!("{:?}", decision.to_string());
        println!("{:#?}", decision.to_string());

        // println!("{decision}");
        println!("{decision:?}");
        println!("{decision:#?}");
    }
}
