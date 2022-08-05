package tictactoe

import (
	"fmt"
)

type TicTacToe struct {
	board         Board
	currentPlayer rune
	moves         int
}

func NewTicTacToe(board Board) *TicTacToe {
	return &TicTacToe{board: board, currentPlayer: 'X'}
}

func (t *TicTacToe) Play() string {
	fmt.Println("Game Starting")
	fmt.Println(t.board)

	for t.moves < 9 {
		fmt.Printf("%c's turn!\n", t.currentPlayer)
		t.move(t.currentPlayer)
		fmt.Println(t.board)

		t.currentPlayer = switchPlayer(t.currentPlayer)

	}

	return "Draw!"

}

func (t *TicTacToe) move(player rune) {
	for {
		l, err := t.getPlayerInput()
		if err != nil {
			fmt.Println(err)
		}

		x, y := l/3, l%3
		err = t.board.playMove(x, y, t.currentPlayer)
		if err == nil {
			t.moves++
			return
		}
		fmt.Printf("Enter a number from 1 to 9. %v\n", err)
	}
}

func (t *TicTacToe) getPlayerInput() (int, error) {
	var location int
	fmt.Print("Enter Location: ")
	_, err := fmt.Scan(&location)
	if err != nil {
		return -1, err
	}
	return location - 1, nil
}

func switchPlayer(currentPlayer rune) rune {
	switch currentPlayer {
	case 'X':
		return 'O'
	default:
		return 'X'
	}
}
