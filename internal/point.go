package internal

type point struct {
	x int
	y int
}

func (p point) NeighborInDirection(direction direction) point {
	switch direction {
	case North:
		return p.North()
	case South:
		return p.South()
	case East:
		return p.East()
	case West:
		return p.West()
	}
	panic("invalid direction")
}

func (p point) North() point {
	return point{x: p.x, y: p.y - 1}
}

func (p point) South() point {
	return point{x: p.x, y: p.y + 1}
}

func (p point) East() point {
	return point{x: p.x + 1, y: p.y}
}

func (p point) West() point {
	return point{x: p.x - 1, y: p.y}
}
