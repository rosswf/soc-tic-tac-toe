package tictactoe

import "errors"

type Board [3][3]rune

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) PlayMove(x, y int, player rune) error {
	if b[x][y] != 0 {
		return errors.New("Invalid move")
	}
	b[x][y] = player
	return nil
}
