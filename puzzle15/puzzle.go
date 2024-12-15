package puzzle15

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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

func (instruction Instruction) String() string {
	switch instruction {
	case Up:
		return "^"
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	default:
		panic(instruction)
	}
}

type WideWarehouse struct {
	width        int
	height       int
	robot        Coords
	boxes        map[Coords]bool // saved pos is LHS of the box
	walls        map[Coords]bool
	instructions []Instruction
}

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

func ParseInput15Wide(data io.Reader) *WideWarehouse {
	warehouse := WideWarehouse{
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
				warehouse.width = 2*offset + 2
				switch char {
				case '#':
					warehouse.walls[Coords{2 * offset, currentRowNum}] = true
					warehouse.walls[Coords{2*offset + 1, currentRowNum}] = true
				case 'O':
					warehouse.boxes[Coords{2 * offset, currentRowNum}] = true
				case '@':
					warehouse.robot = Coords{2 * offset, currentRowNum}
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

func (warehouse *WideWarehouse) ExecuteInstructions() {
	for _, instruction := range warehouse.instructions {
		warehouse.attemptToMove(warehouse.robot, instruction, false)
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

func (warehouse *WideWarehouse) attemptToMove(pos Coords, instruction Instruction, dryRun bool) bool {
	isBox := warehouse.boxes[pos]

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

	secondaryDestination := Coords{destination.x + 1, destination.y}

	if destination.x < 0 || destination.x >= warehouse.width || destination.y < 0 || destination.y >= warehouse.height {
		return false
	}
	if isBox && (secondaryDestination.x < 0 || secondaryDestination.x >= warehouse.width || secondaryDestination.y < 0 || secondaryDestination.y >= warehouse.height) {
		return false
	}

	if warehouse.walls[destination] {
		return false
	}
	if isBox && warehouse.walls[secondaryDestination] {
		return false
	}

	potentialBoxObstructionPosition := Coords{destination.x - 1, destination.y}

	// Dry runs to check the way is clear
	if warehouse.boxes[destination] {
		canMove := warehouse.attemptToMove(destination, instruction, true)
		if !canMove {
			return false
		}
	}
	if instruction != Right && warehouse.boxes[potentialBoxObstructionPosition] {
		canMove := warehouse.attemptToMove(potentialBoxObstructionPosition, instruction, true)
		if !canMove {
			return false
		}
	}
	if isBox && instruction != Left && warehouse.boxes[secondaryDestination] {
		canMove := warehouse.attemptToMove(secondaryDestination, instruction, true)
		if !canMove {
			return false
		}
	}

	if !dryRun {
		// Now that we have checked the way is clear, actually move whichever boxes are in the way
		if warehouse.boxes[destination] {
			warehouse.attemptToMove(destination, instruction, false)
		}
		if instruction != Right && warehouse.boxes[potentialBoxObstructionPosition] {
			warehouse.attemptToMove(potentialBoxObstructionPosition, instruction, false)
		}
		if isBox && instruction != Left && warehouse.boxes[secondaryDestination] {
			warehouse.attemptToMove(secondaryDestination, instruction, false)
		}

		if pos == warehouse.robot {
			warehouse.robot = destination
		} else {
			warehouse.boxes[destination] = true
			delete(warehouse.boxes, pos)
		}
	}
	return true
}

func (warehouse *Warehouse) SumGPSCoordsOfBoxes() int {
	sum := 0
	for box := range warehouse.boxes {
		sum += 100*box.y + box.x
	}
	return sum
}

func (warehouse *WideWarehouse) SumGPSCoordsOfBoxes() int {
	sum := 0
	for box := range warehouse.boxes {
		sum += 100*box.y + box.x
	}
	return sum
}

func (warehouse *WideWarehouse) String() string {
	var sb strings.Builder
	for y := 0; y < warehouse.height; y++ {
		for x := 0; x < warehouse.width; x++ {
			pos := Coords{x, y}
			posToLeft := Coords{x - 1, y}
			if pos == warehouse.robot {
				sb.WriteRune('@')
			} else if warehouse.walls[pos] {
				sb.WriteRune('#')
			} else if warehouse.boxes[pos] {
				sb.WriteRune('[')
			} else if warehouse.boxes[posToLeft] {
				sb.WriteRune(']')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
