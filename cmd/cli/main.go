package main

import (
	"fmt"

	tictactoe "github.com/rosswf/soc-tic-tac-toe"
)

func main() {
	board := tictactoe.NewBoard()
	game := tictactoe.NewTicTacToe(board)

	result := game.Play()

	fmt.Println(result)
}
