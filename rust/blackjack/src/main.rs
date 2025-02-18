// Lessons learned:
//
// 1. Coding rust is an ongoing struggle.
// 2. Ownership via moving and borrowing happens implicily.  This is almost a deal breaker.
// 2a. Variables are located either or the stack or on the heap.
// 2b. Gentle coder will need to understand where variable get located on their own.
// 2c. Ownership via moving and borrowing apply only to heap variables.
// 2d. Moving and borrowing errors are hard to grok.
// 2e. Every assignment and argument/parameter linkage should be considered a compiler error
//     in waiting
// 3. Traits to add to types are determined via compiler errors.
// 3a. Coding as a reaction to compiler errors.
// 3b. Traits are Java Interfaces.
// 4. basic enums are integer based.
// 5. to get a working integer enum, you need a lot of code.
// 6. Global variables are nontrivial to initialize at run time (easy at compile time).
// 6a. Had to use the "lazy_static!" macro, which appears to thread safe.
// 7. "consider introducing a named lifetime parameter"
// 8. "cannot apply unary operator `-`" to a u32
// 9. No traditional ANSI C for loop: for i=0; i<N; i++ {}
// 10. for nested structs, ownership is the transitive closure

use rand::prelude::*;
use rand_chacha::ChaCha8Rng;

mod cards;
mod game;
mod hand;
mod player;
mod rules;
mod strategy;

fn main() {
    // println!("Hello, world!");

    let mut blackjack = game::BlackJack::create();
    for _i in 0..5 {
        blackjack.play_game();
    }
}
