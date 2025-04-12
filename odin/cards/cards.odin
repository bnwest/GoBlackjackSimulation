package cards

CardSuite :: enum{
    HEARTS,    // == 0
    DIAMONDS,  // == 1
    SPADES,    // == 2
    CLUBS,     // == 3
}

/*
// Error: Compound literals of dynamic types are disabled by default
CardSuiteValue := map[CardSuite]string{
    .HEARTS   = "♥️", // aka U+2665 + U+fe0f
	.DIAMONDS = "♦️", // aka U+2666 + U+fe0f
	.SPADES   = "♠️", // aka U+2660 + U+fe0f
	.CLUBS    = "♣️", // aka U+2663 + U+fe0f
}
*/

card_suite_value := [CardSuite]string {
    .HEARTS   = "♥️", // aka U+2665 + U+fe0f
	.DIAMONDS = "♦️", // aka U+2666 + U+fe0f
	.SPADES   = "♠️", // aka U+2660 + U+fe0f
	.CLUBS    = "♣️", // aka U+2663 + U+fe0f
}

to_string :: proc(
    suite: CardSuite
) -> string {
    return card_suite_value[suite]
}
