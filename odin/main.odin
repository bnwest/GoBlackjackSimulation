package main

import "core:fmt"
import "core:unicode/utf8"
import "core:unicode/utf16"
import "cards"

main :: proc() {
    r := "🚀"
    fmt.printfln("runes in {0} is {1}", "🚀", utf8.rune_count_in_string("🚀"))

    heart: cards.CardSuite
    heart = cards.CardSuite.HEARTS

    heart_rune: rune
    // heart_rune = cards.CardSuite.HEARTS[0]
    // heart_rune = '♥️' // aka U+2665 + U+fe0f
    heart_rune = utf8.MAX_RUNE 
    //              '\U0010ffff'
    // heart_rune = '\U2665fe0f' // aka U+2665 + U+fe0f

    heart_rune = '\u2665'
    fmt.printfln("heart rune is: {0}, is valid? {1}", heart_rune, utf8.valid_rune(heart_rune))

    heart_rune = '\ufe0f'
    fmt.printfln("heart rune is: {0}, is valid? {1}", heart_rune, utf8.valid_rune(heart_rune))

    fmt.printfln("runes in {0} is {1}", cards.to_string(heart), utf8.rune_count_in_string("♥️"))

    ace_of_spades: cards.Card
    ace_of_spades = cards.Card{
        suite=cards.CardSuite.SPADES,
        rank=cards.CardRank.ACE,
    }
    fmt.printfln("card: {0}", cards.to_string(ace_of_spades))

    shoe := cards.create_shoe()
    fmt.printfln("dealer shoe has {0} cards.", len(shoe))
    fmt.printfln("dealer shoe top card is {0}.", cards.to_string(shoe[0]))
}
