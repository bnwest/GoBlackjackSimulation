// LESSONS LEARNED:
// 1. odin (like Go) compiler had a strong opinion on newlines in
//     if conf { stmt } else if { cond } else { stmt }
// 2. Global constants at the file level must not have expressions,
// must have value known at compile time, and should look like: 
//     RNG_SEED :: 314159
// 3. Structs can not have methods, which is a big name space problem.
// struct "method" names have to be unique across all structs in package.
// I ended up prefixing all method names.  Alternatively I could have put
// one struct per package which seems a bit extreme.
// 4. no i++
// 5. odin variable name convention is snake case (python) and not camel case (JavaScript, Java)
// 6. "not in" for key in map check is "not_in"
// 7. "Cannot assign to struct field in map". work around is
// to create a struct copy, modify copy and reassign back into map
// 8. tests are easy to write
// 9. tests report memory leaks, which is amazingly useful
// 10. odin has same packge structure as go, a very good thing

package main

import "core:fmt"
import "core:unicode/utf8"
import "core:unicode/utf16"
import "cards"
import "game"
import "strategy"

main :: proc() {
    /*
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

    master_hand := game.create_player_master_hand()
    game.add_start_hand(&master_hand, bet=100)
    game.add_card(&master_hand.hands[0], shoe[0])
    game.add_card(&master_hand.hands[0], shoe[1])
    game.log_hands(&master_hand, "initial hand")
    */

    //
    // FOR REALS
    //

    blackjack := game.create_blackjack()
    defer game.free_blackjack(&blackjack)

    for i in 0..<10 {
        game.play_game(&blackjack)
    }

    fmt.printfln("results: {}", blackjack.results)
    fmt.printfln("stats: {}", blackjack.stats)
}