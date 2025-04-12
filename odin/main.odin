package main

import "core:fmt"
import "cards"

main :: proc() {
	fmt.println("Hellope!")

    heart: cards.CardSuite
    heart = cards.CardSuite.HEARTS

    heart_string: string
    heart_string = cards.to_string(heart)

    fmt.println("suite: {0}", heart_string)
}
