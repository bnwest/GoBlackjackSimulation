// Lesson Learned
// 1. Go compiler had a strong opinion on newlines in
//     if conf { stmt } else if { cond } else { stmt }
// 2. Function that start with a lower case letter => private
// 3. Stuct field names that start with a lower case letter => private
// 4. Unit test files must end in "_test.go".  Test functions therein
// must have the prefix "Test"
// 5. The ":=" is the declaration + assignment operator does not take a type, FTW.
// 6. ":=" and "=" are maddenly different
// 6.1 ":=" only works in a function context
// 6.2 ":=" has no explicit type => avoid for aggregate types and references thereto
// 7. Go has no "set" aggregate type.
// 8. Go does not support constant arrays, maps or slices.
// 9. Methods are added outside the struct ... aka a receiver function ...
// "func (self *<struct>) AddCard()" is a receiver function.
// 9.1 "func (self <struct>) AddCard()" is EVIL, the value the struct instance
// is copied and perhaps modified by call which is NOT what is desired EVER
// the idiom should always be used; "func (self *<struct>) AddCard()" ie
// always define receiver functions with a reference tot he struct
// 10. Tuple type has no intrinsic support ... workaround is to define then use struct
// 11. gofmt is not configurable, no way to disable formatting for a block of code
// 12. python cares about newline and indents; go cares about newlines
// 13. WRT <struct instance?>.<member>, the "." notation is the same
// for both instance and instance reference
// 13.1 Avoid ":=" since type is implicit => do not know if an instance
// or instance reference is being created

package main

import (
	"fmt"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/game"
)

func main() {
	var blackjack *game.BlackJack = game.CreateBlackJack()
	for i := 0; i < 1000_000; i++ {
		blackjack.PlayGame()
	}

	// for 100,000 games and $2 bets per hand:
	// John: {HandsPlayed:102848 HandsWon:44180 HandsLost:50248  HandsPushed:8420  Proceeds:1365}
	// Jane: {HandsPlayed:205678 HandsWon:88299 HandsLost:100572 HandsPushed:16807 Proceeds:3267}

	// for 1,000,000 games and $2 bets per hand:
	// John: {HandsPlayed:1028420 HandsWon:439691 HandsLost:504276  HandsPushed:84453  Proceeds:6472}
	// Jane: {HandsPlayed:2057493 HandsWon:879804 HandsLost:1009397 HandsPushed:168292 Proceeds:12758}

	for playerName, result := range blackjack.Results {
		var playerResult *game.BlackJackResults = result
		// "%+v" => print the struct field names and values, versus just values
		fmt.Printf("%v: %+v\n", playerName, *playerResult)
	}
}
