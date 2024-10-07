package game

type Game interface {
	MakeMove(x int, y int, f Figure) bool
	GetWinner() Figure
	IsDraw() bool
}
