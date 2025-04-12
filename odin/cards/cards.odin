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

to_string :: proc(suite: CardSuite) -> (value: string, ok: bool) {
    switch suite {
    case CardSuite.HEARTS:
        return "♥️", true // aka U+2665 + U+fe0f
    case CardSuite.DIAMONDS:
        return "♦️", true // aka U+2666 + U+fe0f
    case CardSuite.SPADES:
        return "♠️", true // aka U+2660 + U+fe0f
    case CardSuite.CLUBS:
        return "♣️", true // aka U+2663 + U+fe0f
    }
    // Error: Missing return statement at the end of the procedure 'to_string'
    return "dead code: the odin compiler has been fooled!!!", false
}
