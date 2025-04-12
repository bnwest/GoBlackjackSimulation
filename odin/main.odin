package main

import "core:fmt"
import "core:unicode/utf8"
import "core:unicode/utf16"
import "cards"

main :: proc() {
	fmt.println("Hellope!")

    heart: cards.CardSuite
    heart = cards.CardSuite.HEARTS

    r := "🚀"
    heart_rune: rune
    // heart_rune = cards.CardSuite.HEARTS[0]
    // heart_rune = '♥️' // aka U+2665 + U+fe0f
    heart_rune = utf8.MAX_RUNE 
    //              '\U0010ffff'
    // heart_rune = '\U2665fe0f' // aka U+2665 + U+fe0f
    heart_rune = '\u2665'

    is_valid: bool
    is_valid = utf8.valid_rune(heart_rune)

    fmt.printfln("heart rune is: {0}, is valid? {1}", heart_rune, utf8.valid_rune(heart_rune))

    heart_rune = '\ufe0f'
    fmt.printfln("heart rune is: {0}, is valid? {1}", heart_rune, utf8.valid_rune(heart_rune))

    fmt.printfln("runes in {0} is {1}", "♥️", utf8.rune_count_in_string("♥️"))

    foo1: i32 = 0x2665fe0f

    heart_string: string
    heart_string = cards.to_string(heart)

    fmt.printfln("suite: {0}", heart_string)

    ace_of_spades: cards.Card
    ace_of_spades = cards.Card{
        suite=cards.CardSuite.SPADES,
        rank=cards.CardRank.ACE,
    }

    fmt.printfln("card: {0}", cards.to_string(ace_of_spades))
    fmt.printfln("card: {0} {1}", cards.to_string(ace_of_spades.suite), cards.to_string(ace_of_spades.rank))

    shoe := cards.create_shoe()
    fmt.printfln("shoes has {0} cards.", len(shoe))
}
