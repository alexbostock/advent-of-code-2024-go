package puzzle6

import (
	"bufio"
	"io"
	"strings"
)

type Map struct {
	height         int
	width          int
	obstacles      map[Coords]bool
	guardPosition  Coords
	guardDirection Direction
}
type Coords struct {
	i int
	j int
}
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func ParseInput6(data io.Reader) *Map {
	scanner := bufio.NewScanner(data)

	area := Map{
		obstacles: make(map[Coords]bool),
	}

	currentLineNum := -1
	for scanner.Scan() {
		currentLineNum++
		area.height = currentLineNum + 1
		line := scanner.Text()
		cells := strings.Split(line, "")
		for indexInLine, cell := range cells {
			area.width = indexInLine + 1
			if cell == "#" {
				area.obstacles[Coords{currentLineNum, indexInLine}] = true
			} else if cell == "^" {
				area.guardPosition = Coords{currentLineNum, indexInLine}
				area.guardDirection = Up
			}
		}
	}

	return &area
}

func (area *Map) Clone() *Map {
	return &Map{
		height:    area.height,
		width:     area.width,
		obstacles: area.obstacles,
		guardPosition: Coords{
			area.guardPosition.i,
			area.guardPosition.j,
		},
		guardDirection: area.guardDirection,
	}
}

func CountGuardPositionsVisited(area *Map) int {
	positionsVisited := make(map[Coords]bool)
	for area.guardPosition.i >= 0 && area.guardPosition.i < area.height && area.guardPosition.j >= 0 && area.guardPosition.j < area.width {
		positionsVisited[area.guardPosition] = true
		nextPosition := oneStepForward(area.guardPosition, area.guardDirection)
		if area.obstacles[nextPosition] {
			area.guardDirection = turnRight(area.guardDirection)
		} else {
			area.guardPosition = nextPosition
		}
	}
	return len(positionsVisited)
}

func oneStepForward(position Coords, direction Direction) Coords {
	switch direction {
	case Up:
		return Coords{position.i - 1, position.j}
	case Down:
		return Coords{position.i + 1, position.j}
	case Right:
		return Coords{position.i, position.j + 1}
	case Left:
		return Coords{position.i, position.j - 1}
	}
	panic("Unexpected direction")
}

func turnRight(direction Direction) Direction {
	switch direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	panic("Unexpected direction")
}

func CountPossibleNewObstaclesCausingLoops(area *Map) int {
	initialGuardPosition := Coords{area.guardPosition.i, area.guardPosition.j}
	initialGuardDirection := area.guardDirection
	count := 0
	for i := 0; i < area.height; i++ {
		for j := 0; j < area.width; j++ {
			pos := Coords{i, j}
			if area.obstacles[pos] {
				continue
			}
			if pos == initialGuardPosition {
				continue
			}

			area.obstacles[pos] = true
			if hasGuardLoop(area) {
				count++
			}

			delete(area.obstacles, pos)
			area.guardPosition.i = initialGuardPosition.i
			area.guardPosition.j = initialGuardPosition.j
			area.guardDirection = initialGuardDirection
		}
	}
	return count
}

type PositionAndDirection struct {
	position  Coords
	direction Direction
}

func hasGuardLoop(area *Map) bool {
	visitedPositions := make(map[PositionAndDirection]bool)
	for area.guardPosition.i >= 0 && area.guardPosition.i < area.height && area.guardPosition.j >= 0 && area.guardPosition.j < area.width {
		guardState := PositionAndDirection{area.guardPosition, area.guardDirection}
		if visitedPositions[guardState] {
			return true
		}
		visitedPositions[guardState] = true
		nextPosition := oneStepForward(area.guardPosition, area.guardDirection)
		if area.obstacles[nextPosition] {
			area.guardDirection = turnRight(area.guardDirection)
		} else {
			area.guardPosition = nextPosition
		}
	}
	return false
}
