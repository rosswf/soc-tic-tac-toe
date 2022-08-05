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
		return errors.New("Invalid location")
	}

	if b[x][y] != ' ' {
		return errors.New("Invalid move")
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
