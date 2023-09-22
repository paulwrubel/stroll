package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Stroll struct {
	cells [][]cell

	tape   *tape
	memory int64

	currentPos       point
	currentDirection direction
}

func NewStroll(strollBytes string, args string) (*Stroll, error) {
	inputLines := strings.Split(strollBytes, "\n")
	if len(inputLines) == 0 {
		return nil, errors.New("no content in input text")
	}

	// get the maximum line length
	// this will make sure all lines are the same length for parsing
	maxLineLength := 0
	for _, line := range inputLines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	lines := [][]byte{}
	for _, line := range inputLines {
		alignedLine := line
		if len(alignedLine) > maxLineLength {
			alignedLine = alignedLine[:maxLineLength]
		} else if len(alignedLine) < maxLineLength {
			alignedLine += strings.Repeat(" ", maxLineLength-len(alignedLine))
		}
		lines = append(lines, []byte(alignedLine))
	}

	// initialize cells
	cells := [][]cell{}
	for i := 0; i < len(lines[0]); i++ {
		cellCol := make([]cell, len(lines))
		cells = append(cells, cellCol)
	}

	// fill cells with parsed data
	for y, line := range lines {
		for x, rawCell := range line {
			cells[x][y] = parseCell(rawCell)
		}
	}

	tape := NewTape()
	tape.MoveTo(0)
	for i, r := range args {
		if i+1 > 9 {
			break
		}
		tape.MoveForward()
		tape.Write(int64(r))
	}

	return &Stroll{
		cells: cells,
		tape:  tape,
	}, nil
}

func (s *Stroll) Execute() error {
	homePoint, err := s.getHomePoint()
	if err != nil {
		return fmt.Errorf("error getting home point: %s", err.Error())
	}

	// we always start by heading east
	s.currentPos = homePoint
	s.currentDirection = East
	s.tape.MoveTo(0)

	for {
		currentCell := s.getCellAt(s.currentPos)
		s.executeCell(currentCell)
		s.currentDirection = s.getNextDirection(currentCell)

		nextCellPos := s.currentPos.NeighborInDirection(s.currentDirection)
		nextCell := s.getCellAt(nextCellPos)

		isValid := isValidTransition(currentCell, nextCell, s.currentDirection)
		if !isValid {
			return fmt.Errorf("invalid transition at [%d, %d] (%s) -> [%d, %d] (%s)",
				s.currentPos.x, s.currentPos.y, currentCell.String(),
				nextCellPos.x, nextCellPos.y, nextCell.String())
		}

		if nextCell == Home {
			return nil
		} else if nextCell == Blank {
			return fmt.Errorf("you're lost (at [%d, %d])! don't forget to return home next time", s.currentPos.x, s.currentPos.y)
		}

		s.currentPos = nextCellPos
	}
}

func (s *Stroll) executeCell(c cell) {
	switch c {
	case Blank, SkipHorizontal, SkipVertical:
		// noop
	case WalkFacingNorth:
		if s.currentDirection == North {
			s.tape.Write(s.tape.Read() + 1)
		} else {
			s.tape.Write(s.tape.Read() - 1)
		}
	case WalkFacingSouth:
		if s.currentDirection == South {
			s.tape.Write(s.tape.Read() + 1)
		} else {
			s.tape.Write(s.tape.Read() - 1)
		}
	case WalkFacingEast:
		if s.currentDirection == East {
			s.tape.Write(s.tape.Read() + 1)
		} else {
			s.tape.Write(s.tape.Read() - 1)
		}
	case WalkFacingWest:
		if s.currentDirection == West {
			s.tape.Write(s.tape.Read() + 1)
		} else {
			s.tape.Write(s.tape.Read() - 1)
		}
	case Home, Waypoint, Crossing:
		// noop
	case Portal:
		otherPortalPos, foundPortal := s.getOtherPortalLocation(s.currentPos)
		if foundPortal {
			s.currentPos = otherPortalPos
		}
		// else
		// noop
		// you're lost, buddy! good luck!
	case Reg0:
		s.tape.MoveTo(0)
	case Reg1:
		s.tape.MoveTo(1)
	case Reg2:
		s.tape.MoveTo(2)
	case Reg3:
		s.tape.MoveTo(3)
	case Reg4:
		s.tape.MoveTo(4)
	case Reg5:
		s.tape.MoveTo(5)
	case Reg6:
		s.tape.MoveTo(6)
	case Reg7:
		s.tape.MoveTo(7)
	case Reg8:
		s.tape.MoveTo(8)
	case Reg9:
		s.tape.MoveTo(9)
	case TurnNorth, TurnSouth, TurnEast, TurnWest, TurnLeft, TurnRight:
		// noop
	case Yell:
		fmt.Printf("%c", s.tape.Read())
	case Zero:
		s.tape.Write(0)
	case Memorize:
		s.memory = s.tape.Read()
	case Recall:
		s.tape.Write(s.memory)
	}
}

func (s *Stroll) getNextDirection(c cell) direction {
	switch c {
	case SkipHorizontal, SkipVertical, WalkFacingNorth, WalkFacingSouth, WalkFacingEast, WalkFacingWest:
		// always keep walking straight on an edge
		return s.currentDirection
	case Home:
		return East
	case Waypoint, Reg0, Reg1, Reg2, Reg3, Reg4, Reg5, Reg6, Reg7, Reg8, Reg9, Yell, Zero, Memorize, Recall:
		newDir, foundPath := s.getRandomValidDirection()
		if !foundPath {
			// will probably always be Blank, maybe, probably...
			return s.currentDirection
		}
		return newDir
	case Portal:
		// we ignore which direction we came from with portals
		newDir, foundPath := s.getRandomDirection()
		if !foundPath {
			// will probably always be Blank, maybe, probably...
			return s.currentDirection
		}
		return newDir
	case Crossing:
		return s.currentDirection
	case TurnNorth:
		return North
	case TurnSouth:
		return South
	case TurnEast:
		return East
	case TurnWest:
		return West
	case TurnLeft:
		if s.tape.Read() == 0 {
			return s.currentDirection
		}
		return s.currentDirection.Left()
	case TurnRight:
		if s.tape.Read() == 0 {
			return s.currentDirection
		}
		return s.currentDirection.Right()
	default:
		panic("invalid cell")
	}
}

func (s *Stroll) getRandomValidDirection() (direction, bool) {
	candidates := []direction{}
	if s.currentDirection != North.Opposite() && s.getCellAt(s.currentPos.North()) != Blank {
		candidates = append(candidates, North)
	}
	if s.currentDirection != South.Opposite() && s.getCellAt(s.currentPos.South()) != Blank {
		candidates = append(candidates, South)
	}
	if s.currentDirection != East.Opposite() && s.getCellAt(s.currentPos.East()) != Blank {
		candidates = append(candidates, East)
	}
	if s.currentDirection != West.Opposite() && s.getCellAt(s.currentPos.West()) != Blank {
		candidates = append(candidates, West)
	}
	if len(candidates) == 0 {
		return direction(0), false
	}
	return candidates[rand.Intn(len(candidates))], true
}

func (s *Stroll) getRandomDirection() (direction, bool) {
	candidates := []direction{}
	if s.getCellAt(s.currentPos.North()) != Blank {
		candidates = append(candidates, North)
	}
	if s.getCellAt(s.currentPos.South()) != Blank {
		candidates = append(candidates, South)
	}
	if s.getCellAt(s.currentPos.East()) != Blank {
		candidates = append(candidates, East)
	}
	if s.getCellAt(s.currentPos.West()) != Blank {
		candidates = append(candidates, West)
	}
	if len(candidates) == 0 {
		return direction(0), false
	}
	return candidates[rand.Intn(len(candidates))], true
}

func (s *Stroll) getOtherPortalLocation(thisPos point) (point, bool) {
	candidates := []point{}
	for x, col := range s.cells {
		for y, cell := range col {
			thatPos := point{
				x: x,
				y: y,
			}
			if cell == Portal && thatPos != thisPos {
				candidates = append(candidates, thatPos)
			}
		}
	}
	if len(candidates) == 0 {
		return point{}, false
	}
	return candidates[rand.Intn(len(candidates))], true
}

func (s *Stroll) getCellAt(p point) cell {
	if p.x < 0 || p.x >= len(s.cells) || p.y < 0 || p.y >= len(s.cells[0]) {
		return Blank
	}
	return s.cells[p.x][p.y]
}

func (s *Stroll) getHomePoint() (point, error) {
	candidates := []point{}
	for x, col := range s.cells {
		for y, cell := range col {
			if cell == Home {
				candidates = append(candidates, point{x: x, y: y})
			}
		}
	}
	if len(candidates) == 0 {
		return point{}, errors.New("no home found")
	}
	if len(candidates) > 1 {
		return point{}, errors.New("multiple homes found")
	}
	return candidates[0], nil
}
