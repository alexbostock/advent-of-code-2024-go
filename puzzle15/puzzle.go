package puzzle15

import (
	"bufio"
	"fmt"
	"io"
)

type Warehouse struct {
	width        int
	height       int
	robot        Coords
	boxes        map[Coords]bool
	walls        map[Coords]bool
	instructions []Instruction
}
type Coords struct {
	x int
	y int
}
type Instruction int

const (
	Up Instruction = iota
	Down
	Left
	Right
)

func ParseInput15(data io.Reader) *Warehouse {
	warehouse := Warehouse{
		0, 0, Coords{}, make(map[Coords]bool), make(map[Coords]bool), nil,
	}
	parsingInstructions := false
	scanner := bufio.NewScanner(data)
	currentRowNum := -1
	for scanner.Scan() {
		currentRowNum++

		line := scanner.Text()
		if len(line) == 0 {
			parsingInstructions = true
			continue
		}

		if parsingInstructions {
			for _, char := range []rune(line) {
				switch char {
				case '^':
					warehouse.instructions = append(warehouse.instructions, Up)
				case 'v':
					warehouse.instructions = append(warehouse.instructions, Down)
				case '<':
					warehouse.instructions = append(warehouse.instructions, Left)
				case '>':
					warehouse.instructions = append(warehouse.instructions, Right)
				default:
					panic(fmt.Sprintf("%q", char))
				}
			}
		} else {
			warehouse.height = currentRowNum + 1
			for offset, char := range []rune(line) {
				warehouse.width = offset + 1
				switch char {
				case '#':
					warehouse.walls[Coords{offset, currentRowNum}] = true
				case 'O':
					warehouse.boxes[Coords{offset, currentRowNum}] = true
				case '@':
					warehouse.robot = Coords{offset, currentRowNum}
				}
			}
		}
	}
	return &warehouse
}

func (warehouse *Warehouse) ExecuteInstructions() {
	for _, instruction := range warehouse.instructions {
		warehouse.attemptToMove(warehouse.robot, instruction)
	}
}

func (warehouse *Warehouse) attemptToMove(pos Coords, instruction Instruction) bool {
	destination := Coords{pos.x, pos.y}
	switch instruction {
	case Up:
		destination.y--
	case Down:
		destination.y++
	case Left:
		destination.x--
	case Right:
		destination.x++
	}

	if destination.x < 0 || destination.x >= warehouse.width || destination.y < 0 || destination.y >= warehouse.height {
		return false
	}

	if warehouse.walls[destination] {
		return false
	}

	if warehouse.boxes[destination] {
		canMove := warehouse.attemptToMove(destination, instruction)
		if !canMove {
			return false
		}
	}

	if pos == warehouse.robot {
		warehouse.robot = destination
	} else {
		warehouse.boxes[destination] = true
		delete(warehouse.boxes, pos)
	}
	return true
}

func (Warehouse *Warehouse) SumGPSCoordsOfBoxes() int {
	sum := 0
	for box := range Warehouse.boxes {
		sum += 100*box.y + box.x
	}
	return sum
}
