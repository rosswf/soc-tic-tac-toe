package tictactoe

import (
	"errors"
)

type Board [3][3]rune

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) playMove(x, y int, player rune) error {
	if !(x < 3 && y < 3) {
		return errors.New("Invalid location")
	}

	if b[x][y] != 0 {
		return errors.New("Invalid move")
	}

	b[x][y] = player
	return nil
}
