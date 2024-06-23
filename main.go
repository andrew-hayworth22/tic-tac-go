package main

import (
	"os"
	"tictacgo/board"
	"tictacgo/game"
)

func main() {
	game := game.New(board.New())
	game.Run(os.Stdin, os.Stdout)
}
