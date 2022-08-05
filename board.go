package tictactoe

import (
	"errors"
	"fmt"
	"strings"
)

type Board [3][3]rune

func NewBoard() Board {
	return Board{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
}

func (b *Board) playMove(x, y int, player rune) error {
	if !(x < 3 && y < 3) {
		return errors.New("Invalid location.")
	}

	if b[x][y] != ' ' {
		return errors.New("Position already taken.")
	}

	b[x][y] = player
	return nil
}

func (b Board) String() string {
	var sb strings.Builder
	for i, row := range b {
		if i != 0 {
			fmt.Fprint(&sb, "--+---+--\n")
		}
		fmt.Fprintf(&sb, "%c | %c | %c\n", row[0], row[1], row[2])
	}
	return sb.String()
}

func (b Board) checkRow(i int) bool {
	return b[i][0] == b[i][1] && b[i][1] == b[i][2] && b[i][2] != ' '
}

func (b Board) checkColumn(i int) bool {
	return b[0][i] == b[1][i] && b[1][i] == b[2][i] && b[2][i] != ' '
}

func (b Board) checkDiags() bool {
	return (b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[2][2] != ' ') || (b[2][0] == b[1][1] && b[1][1] == b[0][2] && b[0][2] != ' ')
}
