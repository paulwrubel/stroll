package internal

import "errors"

type cell int

const (
	//-- blank --//
	Blank cell = iota

	//-- edges --//
	SkipHorizontal
	SkipVertical
	WalkFacingNorth
	WalkFacingSouth
	WalkFacingEast
	WalkFacingWest

	//-- nodes --//
	// basic
	Home
	Waypoint
	Crossing
	// advanced
	Portal
	// registers
	Reg0
	Reg1
	Reg2
	Reg3
	Reg4
	Reg5
	Reg6
	Reg7
	Reg8
	Reg9
	// cardinal directions
	TurnNorth
	TurnSouth
	TurnEast
	TurnWest
	// turns
	TurnLeft
	TurnRight
	// actions
	Yell
	Zero
)

func (c cell) String() string {
	switch c {
	case Blank:
		return " "
	case SkipHorizontal:
		return "-"
	case SkipVertical:
		return "|"
	case WalkFacingNorth:
		return "^"
	case WalkFacingSouth:
		return "v"
	case WalkFacingEast:
		return ">"
	case WalkFacingWest:
		return "<"
	case Home:
		return "H"
	case Waypoint:
		return "#"
	case Crossing:
		return "+"
	case Portal:
		return "@"
	case Reg0:
		return "0"
	case Reg1:
		return "1"
	case Reg2:
		return "2"
	case Reg3:
		return "3"
	case Reg4:
		return "4"
	case Reg5:
		return "5"
	case Reg6:
		return "6"
	case Reg7:
		return "7"
	case Reg8:
		return "8"
	case Reg9:
		return "9"
	case TurnNorth:
		return "n"
	case TurnSouth:
		return "s"
	case TurnEast:
		return "e"
	case TurnWest:
		return "w"
	case TurnLeft:
		return "l"
	case TurnRight:
		return "r"
	case Yell:
		return "y"
	case Zero:
		return "z"
	default:
		return "?"
	}
}

func (c cell) IsNode() bool {
	return c == Home ||
		c == Waypoint ||
		c == Crossing ||
		c == Portal ||
		c == Reg0 ||
		c == Reg1 ||
		c == Reg2 ||
		c == Reg3 ||
		c == Reg4 ||
		c == Reg5 ||
		c == Reg6 ||
		c == Reg7 ||
		c == Reg8 ||
		c == Reg9 ||
		c == TurnNorth ||
		c == TurnSouth ||
		c == TurnEast ||
		c == TurnWest ||
		c == TurnLeft ||
		c == TurnRight ||
		c == Yell ||
		c == Zero
}

func parseCell(b byte) (cell, error) {
	switch b {
	//-- blank --//
	case ' ':
		return Blank, nil
	//-- edges --//
	case '-':
		return SkipHorizontal, nil
	case '|':
		return SkipVertical, nil
	case '^':
		return WalkFacingNorth, nil
	case 'v':
		return WalkFacingSouth, nil
	case '>':
		return WalkFacingEast, nil
	case '<':
		return WalkFacingWest, nil
	//-- nodes --//
	// basic
	case 'H':
		return Home, nil
	case '#':
		return Waypoint, nil
	case '+':
		return Crossing, nil
	// advanced
	case '@':
		return Portal, nil
	// registers
	case '0':
		return Reg0, nil
	case '1':
		return Reg1, nil
	case '2':
		return Reg2, nil
	case '3':
		return Reg3, nil
	case '4':
		return Reg4, nil
	case '5':
		return Reg5, nil
	case '6':
		return Reg6, nil
	case '7':
		return Reg7, nil
	case '8':
		return Reg8, nil
	case '9':
		return Reg9, nil
	// cardinal directions
	case 'n':
		return TurnNorth, nil
	case 's':
		return TurnSouth, nil
	case 'e':
		return TurnEast, nil
	case 'w':
		return TurnWest, nil
	// turns
	case 'l':
		return TurnLeft, nil
	case 'r':
		return TurnRight, nil
	// actions
	case 'y':
		return Yell, nil
	case 'z':
		return Zero, nil
	default:
		return 0, errors.New("invalid cell")
	}
}

func isValidTransition(a, b cell, dir direction) bool {
	if b == Blank {
		return true
	}

	switch a {
	case SkipHorizontal:
		return b == SkipHorizontal || b.IsNode()
	case SkipVertical:
		return b == SkipVertical || b.IsNode()
	case WalkFacingNorth:
		return b == WalkFacingNorth || b.IsNode()
	case WalkFacingSouth:
		return b == WalkFacingSouth || b.IsNode()
	case WalkFacingEast:
		return b == WalkFacingEast || b.IsNode()
	case WalkFacingWest:
		return b == WalkFacingWest || b.IsNode()
	case Home:
		return b == SkipHorizontal ||
			b == WalkFacingEast ||
			b == WalkFacingWest
	case Waypoint, Crossing, Portal:
		return !b.IsNode()
	case TurnNorth, TurnSouth:
		return b == SkipVertical || b == WalkFacingNorth || b == WalkFacingSouth
	case TurnEast, TurnWest:
		return b == SkipHorizontal || b == WalkFacingEast || b == WalkFacingWest
	case TurnLeft, TurnRight:
		return !b.IsNode()
	case Reg0, Reg1, Reg2, Reg3, Reg4, Reg5, Reg6, Reg7, Reg8, Reg9, Yell, Zero:
		if dir == North || dir == South {
			return b == SkipVertical || b == WalkFacingNorth || b == WalkFacingSouth
		}
		return b == SkipHorizontal || b == WalkFacingEast || b == WalkFacingWest
	default:
		return false
	}
}
