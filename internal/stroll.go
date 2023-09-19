package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Stroll struct {
	cells [][]cell

	registers       []int64
	currentRegister int

	currentPos       point
	currentDirection direction
}

func NewStroll(strollBytes string, args []string) (*Stroll, error) {
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
			cell, err := parseCell(rawCell)
			if err != nil {
				return nil, err
			}
			cells[x][y] = cell
		}
	}

	registers := make([]int64, 10)
	for i := 0; i < 9; i++ {
		if len(args) > i {
			arg := args[i]
			// if len(arg) != 1 {
			// 	return nil, fmt.Errorf("invalid register argument at index %d, length too long (arg: %s, len: %d)", i, arg, len(arg))
			// }
			registers[i+1] = int64(arg[0])
		}
	}

	return &Stroll{
		cells:     cells,
		registers: registers,
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
	s.currentRegister = 0

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
			s.registers[s.currentRegister]++
		} else {
			s.registers[s.currentRegister]--
		}
	case WalkFacingSouth:
		if s.currentDirection == South {
			s.registers[s.currentRegister]++
		} else {
			s.registers[s.currentRegister]--
		}
	case WalkFacingEast:
		if s.currentDirection == East {
			s.registers[s.currentRegister]++
		} else {
			s.registers[s.currentRegister]--
		}
	case WalkFacingWest:
		if s.currentDirection == West {
			s.registers[s.currentRegister]++
		} else {
			s.registers[s.currentRegister]--
		}
	case Home:
		// noop
	case Reg0:
		s.currentRegister = 0
	case Reg1:
		s.currentRegister = 1
	case Reg2:
		s.currentRegister = 2
	case Reg3:
		s.currentRegister = 3
	case Reg4:
		s.currentRegister = 4
	case Reg5:
		s.currentRegister = 5
	case Reg6:
		s.currentRegister = 6
	case Reg7:
		s.currentRegister = 7
	case Reg8:
		s.currentRegister = 8
	case Reg9:
		s.currentRegister = 9
	case TurnNorth, TurnSouth, TurnEast, TurnWest:
		// noop
	case Yell:
		fmt.Printf("%c", s.registers[s.currentRegister])
	case Zero:
		s.registers[s.currentRegister] = 0
	}
}

func (s *Stroll) getNextDirection(c cell) direction {
	switch c {
	case SkipHorizontal, SkipVertical, WalkFacingNorth, WalkFacingSouth, WalkFacingEast, WalkFacingWest:
		// always keep walking straight on an edge
		return s.currentDirection
	case Home:
		return East
	case Reg0, Reg1, Reg2, Reg3, Reg4, Reg5, Reg6, Reg7, Reg8, Reg9, Yell, Zero:
		newDir, foundPath := s.getRandomValidDirection()
		if !foundPath {
			// will probably always be Blank, maybe, probably...
			return s.currentDirection
		}
		return newDir
	case TurnNorth:
		return North
	case TurnSouth:
		return South
	case TurnEast:
		return East
	case TurnWest:
		return West
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

func (s *Stroll) getCellAt(p point) cell {
	if p.x < 0 || p.x >= len(s.cells) || p.y < 0 || p.y >= len(s.cells[0]) {
		return Blank
	}
	return s.cells[p.x][p.y]
}

func (s *Stroll) getHomePoint() (point, error) {
	for x, col := range s.cells {
		for y, cell := range col {
			if cell == Home {
				return point{x: x, y: y}, nil
			}
		}
	}
	return point{}, errors.New("no home found")
}
