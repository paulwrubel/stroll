package internal

type direction int

const (
	North direction = iota
	South
	East
	West
)

func (d direction) Opposite() direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("invalid direction")
	}
}

func (d direction) Left() direction {
	switch d {
	case North:
		return West
	case South:
		return East
	case East:
		return North
	case West:
		return South
	default:
		panic("invalid direction")
	}
}

func (d direction) Right() direction {
	return d.Left().Opposite()
}
