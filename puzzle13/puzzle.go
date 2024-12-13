package puzzle13

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	a     Coords
	b     Coords
	prize Coords
}
type Coords struct {
	x int
	y int
}

func ParseInput13(data io.Reader) ([]Machine, error) {
	var machines []Machine
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		a, err := parseButtonLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		scanner.Scan()
		b, err := parseButtonLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		scanner.Scan()
		prize, err := parsePrizeLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		machines = append(machines, Machine{a, b, prize})
		scanner.Scan()
	}
	return machines, nil
}

func parseButtonLine(line string) (Coords, error) {
	tokens := strings.Split(line[10:], ", ")
	xVal := strings.Split(tokens[0], "+")[1]
	yVal := strings.Split(tokens[1], "+")[1]

	x, err := strconv.Atoi(xVal)
	if err != nil {
		return Coords{}, err
	}
	y, err := strconv.Atoi(yVal)
	if err != nil {
		return Coords{}, err
	}
	return Coords{x, y}, nil
}

func parsePrizeLine(line string) (Coords, error) {
	tokens := strings.Split(line[7:], ", ")
	xVal := strings.Split(tokens[0], "=")[1]
	yVal := strings.Split(tokens[1], "=")[1]

	x, err := strconv.Atoi(xVal)
	if err != nil {
		return Coords{}, err
	}
	y, err := strconv.Atoi(yVal)
	if err != nil {
		return Coords{}, err
	}
	return Coords{x, y}, nil
}

func SearchMinimumTokensToWinAllPrizes(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		total += searchMinimumSolution(machine)
	}
	return total
}

func SearchMinimumTokensToWinAllPrizesWithPrizeError(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		machine.prize.x += 10000000000000
		machine.prize.y += 10000000000000
		total += searchMinimumSolution(machine)
	}
	return total
}

func searchMinimumSolution(machine Machine) int {
	// Solve simultaneous equations in a, b:
	// a * machine.a.x + b * machine.b.x == machine.prize.x (1)
	// a * machine.a.y + b * machine.b.y == machine.prize.y (2)

	// (2) => b = (machine.prize.y - a * machine.a.y) / machine.b.y
	// Substituting in (1): a * machine.a.x + machine.b.x * (machine.prize.y - a * machine.a.y) / machine.b.y == machine.prize.x
	// => a (machine.a.x - machine.a.y * machine.b.x / machine.b.y) == machine.prize.x - machine.prize.y * machine.b.x / machine.b.y
	// => a == ( machine.prize.x - machine.prize.y * machine.b.x / machine.b.y ) / ( machine.a.x - machine.a.y * machine.b.x / machine.b.y )

	numerator := float64(machine.prize.x) - float64(machine.prize.y*machine.b.x)/float64(machine.b.y)
	denominator := float64(machine.a.x) - float64(machine.a.y*machine.b.x)/float64(machine.b.y)
	a := int(math.Round(numerator / denominator))
	b := (machine.prize.y - a*machine.a.y) / machine.b.y

	isValidSolution := a*machine.a.x+b*machine.b.x == machine.prize.x &&
		a*machine.a.y+b*machine.b.y == machine.prize.y
	if !isValidSolution {
		return 0
	}

	return 3*a + b
}
