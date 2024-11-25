package main

import (
	"fmt"

  "github.com/bnwest/GoBlackjackSimulation/go/blackjack/cards"
)

func main() {
	fmt.Println("Hello World.")

  var spade cards.CardSuite
  spade = cards.SPADES 
	fmt.Printf("spade enum value: %v\n", spade)
	fmt.Printf("spade enum string value: %v\n", cards.CardSuiteValue[spade])

  heart := cards.HEARTS
	fmt.Printf("heart enum value: %v\n", heart)
	fmt.Printf("heart enum string value: %v\n", cards.CardSuiteValue[heart])
}
