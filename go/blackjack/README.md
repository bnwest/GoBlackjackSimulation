### Go (aka GoLang) setup

[Go installation](https://go.dev/doc/install)

I downloaded the MacOS install package and installed it:
```
% open go1.23.3.darwin-amd64.pkg
```
The package installs the Go distribution to /usr/local/go. 

Kick the tires:
```
% /usr/local/go/bin/go version
go version go1.23.3 darwin/amd64
```

In `~/.bash_profile`, add the following:
> export PATH="/usr/local/go/bin:$PATH"

Then this should work:
```
% which go
/usr/local/go/bin/go

% go version
go version go1.23.3 darwin/amd64
```

## Create a Go project

Create a GO source code directory (from my newly cloned github repo)
```
% cd ~/github/
% git clone https://github.com/bnwest/GoBlackjackSimulation.git
% cd GoBlackjackSimulation/
% mkdir go/blackjack
% cd go/blackjack
```

To link my Go module to the github repo:
```
% go mod init github.com/bnwest/GoBlackjackSimulation/go/blackjack
```
which creates go.mod
```
% cat go.mod 
module github.com/bnwest/GoBlackjackSimulation/go/blackjack

go 1.23.3

require github.com/stretchr/testify v1.10.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

Add new module requirements and sums, as needeed:
```
% go mod tidy
```
whcih creates a `go.sum` file.

# Creating Go module and packages therein

See [Packages, Imports and Modules in Go](https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules)
> A package in Go is essentially a named collection of one or more related .go files. 
In Go, the primary purpose of packages is to help you isolate and reuse code.
>
> Every .go file that you write should begin with a package {name} statement 
which indicates the name of the package that the file is a part of.  It's totally OK to have quite a lot of .go files in the same package. 
>
> In Go, one package == one directory.  That is, all .go files for a package should be contained in the same directory, and a directory should contain the .go files for one package only.  For all non-main packages, the directory name that the code lives in should be the same as the package name.  When choosing a name you should pick something that is short, descriptive, lower case and ideally one word. 
>
> Any package with the name main must also contain a main() function somewhere 
in the package which acts as the entry point for the program.  It's conventional for your `main()` function to live in a file with the filename `main.go`. 

I followed the above conventions when creating my go files.  I created the following 
```
% tree
.
├── README.md
├── cards
│   └── cards.go
├── cards_test.go
├── game
│   ├── game.go
│   ├── hand.go
│   └── player.go
├── game_test.go
├── go.mod
├── go.sum
├── main.go
├── rules
│   └── house.go
├── strategy
│   ├── basic.go
│   ├── decisions.go
│   ├── hard.go
│   ├── pairs.go
│   └── soft.go
└── strategy_test.go

5 directories, 17 files
```

To run the main program:
```
% go run main.go
```

# Unit tests and formatting

I ended up writing unit tests, one per package.  The unit test file name ends in `_test.go` and contains unit test function whose names begine with `Test`.

To run the unit tests:
```
% go test -v
```

To run the go formatter to see the differences
```
% gofmt -d main.go
```
and to update with the differences
```
% gofmt -w main.go
```
