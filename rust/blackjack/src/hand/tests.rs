// file srchandcards/tests.rs defines project module "hand::tests".

use super::HandOutcome;

#[test]
fn test_hand_outcomes() {
    let mut num_hand_outcomes: u32 = 0;
    for _suite in HandOutcome::iterator() {
        num_hand_outcomes += 1
    }

    assert_eq!(num_hand_outcomes, 5);
    assert_eq!(HandOutcome::STAND as usize, 0);
    assert_eq!(HandOutcome::BUST as usize, 1);
    assert_eq!(HandOutcome::SURRENDER as usize, 2);
    assert_eq!(HandOutcome::DEALER_BLACKJACK as usize, 3);
    assert_eq!(HandOutcome::IN_PLAY as usize, 4);

    let mut suite: HandOutcome = HandOutcome::STAND;
    assert_eq!(suite.discriminant(), 0);
    suite = HandOutcome::BUST;
    assert_eq!(suite.discriminant(), 1);
    suite = HandOutcome::SURRENDER;
    assert_eq!(suite.discriminant(), 2);
    suite = HandOutcome::DEALER_BLACKJACK;
    assert_eq!(suite.discriminant(), 3);
    suite = HandOutcome::IN_PLAY;
    assert_eq!(suite.discriminant(), 4);

    for hand_outcome in HandOutcome::iterator() {
        // hand_outcome: &HandOutcome
        let roundtrip_hand_outcome: HandOutcome =
            HandOutcome::transmute(hand_outcome.discriminant());
        assert_eq!(hand_outcome, &roundtrip_hand_outcome);

        println!("{}", hand_outcome.to_string());
        println!("{:?}", hand_outcome);
        println!("{:#?}", hand_outcome);
        println!("{hand_outcome:?}");
        println!("{hand_outcome:#?}");
    }
}
