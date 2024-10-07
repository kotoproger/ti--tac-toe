package game

type FigureName string
type FigureSign rune

const (
	CrossName FigureName = "cross"
	ZeroName  FigureName = "zero"

	CrossSign FigureSign = 'X'
	ZeroSign  FigureSign = '0'
)

type Figure struct {
	Name FigureName
	Sign FigureSign
}
