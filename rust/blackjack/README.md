## Introduction
Simulates a handful blackjack games.

## Getting Started
Need rust.

See:
https://doc.rust-lang.org/book/ch01-01-installation.html

```commandline
$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
```

## Creating a New Project Using Cargo

```commandline
% cd $repo-head
% cd rust
% cargo new blackjack
% tree
.
└── blackjack
    ├── Cargo.toml
    └── src
        └── main.rs
% cd blackjack
% cargo build
   Compiling blackjack v0.1.0 (/Users/bwest/github/GoBlackjackSimulation/rust/blackjack)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 1.35s
% cargo run
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.00s
     Running `target/debug/blackjack`
Hello, world!
```

## Build and Test

To compile/build the rust code:
```commandline
% cargo build
```

To run the rust code:
```commandline
% cargo run
```

To run the unit tests
```commandline
% cargo test
```

To format the rust code
```commandline
% cargo fmt
```

