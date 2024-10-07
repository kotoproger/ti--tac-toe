package board

import "github.com/kbkondakov/tic-tac-toe/game"

type Board interface {
	GetCell(c Coordinate) game.Figure
	SetCell(c Coordinate, f game.Figure) bool
}
