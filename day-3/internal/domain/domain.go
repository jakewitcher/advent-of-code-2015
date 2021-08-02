package domain

type Point struct {
	x int
	y int
}

func (p Point) Up() Point {
	return NewPoint(p.x, p.y+1)
}

func (p Point) Down() Point {
	return NewPoint(p.x, p.y-1)
}

func (p Point) Left() Point {
	return NewPoint(p.x-1, p.y)
}

func (p Point) Right() Point {
	return NewPoint(p.x+1, p.y)
}

func NewPoint(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}
