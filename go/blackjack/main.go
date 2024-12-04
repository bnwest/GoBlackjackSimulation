// Lesson Learned
// 1. Go compiler had a strong opinion on newlines in
//     if conf { stmt } else if { cond } else { stmt }
// 2. Function that start with a lower case letter => private
// 3. Stuct field names that start with a lower case letter => private
// 4. Unit test files must end in "_test.go".  Test functions therein 
// must have the prefix "Test"
// 5. The ":=" declaration + assignment operator does not take a type, FTW.
// 6. ":=" and "=" are maddenly different
// 6.1 ":=" only works in a function context
// 7. Go has no "set" aggregate type.
// 8. Go does not support constant arrays, maps or slices.
// 9. Methods are added outside the struct ... aka a receiver functionAddCard() is a receiver function.
// 10. Tuple type has no intrinsic support ... workaround is to define then use struct
// 11. gofmt is not configurable, no way to disable formatting for a block of code

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

  shoe := cards.CreateShoe()
  fmt.Printf("shoes has %v cards\n", len(shoe))

  cards.ShuffleShoe(shoe)
  fmt.Printf("shoes has %v cards\n", len(shoe))

  cards.DisplayShoe(shoe)
}
