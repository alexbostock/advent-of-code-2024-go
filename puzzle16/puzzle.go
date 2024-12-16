package puzzle16

import (
	"bufio"
	"io"
	"slices"
)

type Maze struct {
	height int
	width  int
	walls  map[Coords]bool
	start  Coords
	end    Coords
}
type ReindeerState struct {
	position        Coords
	directionFacing Direction
}
type State struct {
	position        Coords
	directionFacing Direction
	costSoFar       int
	path            []Coords
}
type Coords struct {
	x, y int
}
type Direction = int

const (
	North Direction = iota
	East
	South
	West
)

func ParseInput16(data io.Reader) *Maze {
	maze := Maze{}
	maze.walls = make(map[Coords]bool)
	scanner := bufio.NewScanner(data)
	currentRow := -1
	for scanner.Scan() {
		currentRow++
		maze.height = currentRow + 1
		line := scanner.Text()
		for offset, char := range []rune(line) {
			maze.width = offset + 1
			pos := Coords{offset, currentRow}
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
	return &maze
}

func (maze *Maze) ShortestPath() (shortestPath int, positionsVisited map[Coords]bool) {
	queue := &ArrayPriorityQueue{}
	shortestPath = -1
	positionsSeen := make(map[ReindeerState]int)
	positionsVisited = make(map[Coords]bool)
	state := State{
		position:        maze.start,
		directionFacing: East,
		costSoFar:       0,
		path:            []Coords{maze.start},
	}
	queue.Enqueue(state)

	for {
		state := queue.Dequeue()
		if shortestPath > -1 && state.costSoFar > shortestPath {
			return
		}
		reindeerState := ReindeerState{state.position, state.directionFacing}
		costAlreadySeen, positionSeen := positionsSeen[reindeerState]
		if positionSeen && costAlreadySeen < state.costSoFar {
			continue
		}
		positionsSeen[reindeerState] = state.costSoFar

		if state.position == maze.end {
			shortestPath = state.costSoFar
			for _, pos := range state.path {
				positionsVisited[pos] = true
			}
		}
		possibleNextStates := []State{
			{
				position:        stepForward(state.position, state.directionFacing),
				directionFacing: state.directionFacing,
				costSoFar:       state.costSoFar + 1,
				path:            append(slices.Clone(state.path), stepForward(state.position, state.directionFacing)),
			},
			{
				position:        state.position,
				directionFacing: rotate(state.directionFacing, true),
				costSoFar:       state.costSoFar + 1000,
				path:            state.path,
			},
			{
				position:        state.position,
				directionFacing: rotate(state.directionFacing, false),
				costSoFar:       state.costSoFar + 1000,
				path:            state.path,
			},
		}

		for _, possibleNextState := range possibleNextStates {
			if maze.walls[possibleNextState.position] {
				continue
			}
			queue.Enqueue(possibleNextState)
		}
	}
}

func stepForward(position Coords, direction Direction) Coords {
	switch direction {
	case North:
		return Coords{position.x, position.y - 1}
	case South:
		return Coords{position.x, position.y + 1}
	case East:
		return Coords{position.x + 1, position.y}
	case West:
		return Coords{position.x - 1, position.y}
	}
	panic("Should never happen")
}

func rotate(direction Direction, clockwise bool) Direction {
	if clockwise {
		switch direction {
		case North:
			return East
		case East:
			return South
		case South:
			return West
		case West:
			return North
		}
	} else {
		switch direction {
		case North:
			return West
		case West:
			return South
		case South:
			return East
		case East:
			return North
		}
	}
	panic("Should never happen")
}

type PriorityQueue interface {
	Enqueue(val State)
	Dequeue() State
}

type ArrayPriorityQueue struct {
	elements []State
}

func (queue *ArrayPriorityQueue) Enqueue(val State) {
	insertIndex := 0
	for insertIndex = 0; insertIndex < len(queue.elements) && queue.elements[insertIndex].costSoFar < val.costSoFar; insertIndex++ {
		continue
	}
	queue.elements = slices.Insert(queue.elements, insertIndex, val)
}
func (queue *ArrayPriorityQueue) Dequeue() State {
	head := queue.elements[0]
	queue.elements = slices.Delete(queue.elements, 0, 1)
	return head
}
