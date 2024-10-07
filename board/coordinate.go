package board

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) GetX() int {
	return c.x
}

func (c *Coordinate) GetY() int {
	return c.y
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x: x, y: y}
}
