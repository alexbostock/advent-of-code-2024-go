package puzzle20

import (
	"bufio"
	"io"
	"slices"
)

type Maze struct {
	width  int
	height int
	start  Coords
	end    Coords
	walls  map[Coords]bool
}
type Coords struct {
	x int
	y int
}

func (pos Coords) adjacent4() []Coords {
	return []Coords{
		{pos.x + 1, pos.y},
		{pos.x - 1, pos.y},
		{pos.x, pos.y + 1},
		{pos.x, pos.y - 1},
	}
}
func (pos Coords) manhattanDistance(other Coords) int {
	return abs(pos.x-other.x) + abs(pos.y-other.y)
}
func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func ParseInput20(input io.Reader) Maze {
	maze := Maze{walls: make(map[Coords]bool)}
	y := -1
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		y++
		maze.height = y + 1
		for x, char := range []rune(scanner.Text()) {
			maze.width = x + 1
			pos := Coords{x, y}
			switch char {
			case '#':
				maze.walls[pos] = true
			case 'S':
				maze.start = pos
			case 'E':
				maze.end = pos
			}
		}
	}
	return maze
}

func (maze Maze) NumCheatsSavingAtLeast(allowedCheatLength int, threshold int) (count int) {
	path := maze.pathWithoutCheats()

	for index1, pos1 := range path {
		for index2 := index1 + 1; index2 < len(path); index2++ {
			pos2 := path[index2]
			cheatLength := pos1.manhattanDistance(pos2)
			if cheatLength <= allowedCheatLength {
				cheatTimeSaved := index2 - index1 - cheatLength
				if cheatTimeSaved >= threshold {
					count++
				}
			}
		}
	}
	return
}

func (maze Maze) pathWithoutCheats() []Coords {
	return maze.buildPath(maze.start, Coords{-1, -1})
}
func (maze Maze) buildPath(pos Coords, prev Coords) []Coords {
	if pos == maze.end {
		return []Coords{pos}
	}
	for _, adjacent := range pos.adjacent4() {
		if adjacent.x < 0 || adjacent.x >= maze.width || adjacent.y < 0 || adjacent.y >= maze.height {
			continue
		}
		if maze.walls[adjacent] {
			continue
		}
		if adjacent == prev {
			continue
		}
		pathFromAdjacent := maze.buildPath(adjacent, pos)
		if pathFromAdjacent != nil {
			return append([]Coords{pos}, pathFromAdjacent...)
		}
	}
	return nil
}

func lenPathWithCheat(path []Coords, cheatStart, cheatEnd Coords) int {
	length := 0
	alreadyCheated := false
	for i := 0; i < len(path); i++ {
		length++
		if alreadyCheated || path[i] == cheatEnd {
			continue
		}
		for _, adjacent := range path[i].adjacent4() {
			if adjacent == cheatStart {
				length++
				i = slices.Index(path, cheatEnd)
				alreadyCheated = true
			}
		}
	}
	return length
}
