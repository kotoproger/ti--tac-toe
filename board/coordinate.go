package board

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) getX() int {
	return c.x
}

func (c *Coordinate) getY() int {
	return c.y
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x: x, y: y}
}
