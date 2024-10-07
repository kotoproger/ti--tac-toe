package game

type (
	FigureName string
	FigureSign rune
)

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
