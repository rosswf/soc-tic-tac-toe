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
	displayIntroduction()

	for t.moves < 9 {
		fmt.Printf("\n%c's turn! ", t.currentPlayer)
		t.playMove(t.currentPlayer)
		fmt.Println(t.board)

		if t.isWon() {
			return fmt.Sprintf("%c is the winner!", t.currentPlayer)
		}

		t.currentPlayer = switchPlayer(t.currentPlayer)
	}

	return "It's a Draw!"
}

func (t *TicTacToe) playMove(player rune) {
	for {
		l, err := t.getPlayerInput()
		if err != nil {
			fmt.Printf("Oops something went wrong! %v\n", err)
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

func (t TicTacToe) getPlayerInput() (int, error) {
	var location int
	fmt.Print("Enter Move: ")
	_, err := fmt.Scan(&location)
	if err != nil {
		return -1, err
	}
	return location - 1, nil
}

func (t TicTacToe) isWon() bool {
	for i := 0; i < 3; i++ {
		if t.board.checkRow(i) || t.board.checkColumn(i) {
			return true
		}
	}
	return t.board.checkDiags()
}

func switchPlayer(currentPlayer rune) rune {
	switch currentPlayer {
	case 'X':
		return 'O'
	default:
		return 'X'
	}
}

func displayIntroduction() {
	fmt.Println(`Welcome to tic tac toe!
To make a move, enter the number corresponding to the location as show below
The board and it's playable positions:
1 | 2 | 3
--+---+--
4 | 5 | 6
--+---+--
7 | 8 | 9`)
}
