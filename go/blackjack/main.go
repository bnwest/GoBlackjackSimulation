// Lesson Learned:
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
// "func (self *<struct>) AddCard()" is a receiver function, self is the receiver arg.
// 9.1 "func (self <struct>) AddCard()" is EVIL, the value the struct instance
// is copied and perhaps modified by call which is NOT what is desired EVER
// the idiom should always be used; "func (self *<struct>) AddCard()" ie
// always define receiver functions with a reference to the struct
// 9.2 paramter copying is the expected behavior for all function arguments
// 9.3 (*p).method() is the same as p.method(), as a convenience
// 10. Tuple type has no intrinsic support ... workaround is to define then use struct???
// 10.1 func swap(x, y string) (string, string) { return y, x } is the counter example
// 11. gofmt is not configurable, no way to disable formatting for a block of code
// 12. python cares about newline and indents; go cares about newlines
// 13. WRT <struct instance?>.<member>, the "." notation is the same
// for both instance and instance reference
// 13.1 (*p).field is the same as p.field, as a convenience
// 13.2 Avoid ":=" in struct creation since type is implicit => do not know
// if an instance or instance reference is being created
// 14. Go Playground is a web service that runs on golang.org's servers.
// 15. stackoverflow default response to a Go question is: roll your on.

package main

import (
	"fmt"

	"github.com/bnwest/GoBlackjackSimulation/go/blackjack/game"
)

func main() {
	var blackjack *game.BlackJack = game.CreateBlackJack()
	for i := 0; i < 100; i++ {
		blackjack.PlayGame()
	}

	// for 1,000,000 games, 3 master hands per game and $2 bets per hand => about $6,000,000 bet
	// Jack: {HandsPlayed:1028420 HandsWon:439691 HandsLost:504276  HandsPushed:84453  Proceeds:6472}
	// Jill: {HandsPlayed:2057493 HandsWon:879804 HandsLost:1009397 HandsPushed:168292 Proceeds:12758}
	// 43% hands won, 49% hands lost, 8% hands pushed

	fmt.Println()
	for playerName, result := range blackjack.Results {
		// result is a copy of blackjack.Results[playerName] ... a pointer in this case
		var playerResult *game.BlackJackPlayerResults = result
		// "%+v" => print the struct field names and values, versus just values
		fmt.Printf("%v: %+v\n", playerName, *playerResult)
	}

	// for 1,000,000 games with 3 master hands per game => about 3,000,000 hands
	// Stats: {DoubleDownCount:302801 SurrenderCount:155506 SplitCount:83061 AcesSplit:32778}
	// roughly 10% hands double down, 5% surrender, 2.5% split, 1% split Aces

	fmt.Println()
	fmt.Printf("Stats: %+v\n", blackjack.Stats)
}
