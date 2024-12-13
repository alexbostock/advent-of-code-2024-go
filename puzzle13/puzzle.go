package puzzle13

import (
	"bufio"
	"fmt"
	"io"
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
	for index, machine := range machines {
		fmt.Println(index, len(machines))
		machine.prize.x += 10000000000000
		machine.prize.y += 10000000000000
		total += searchMinimumSolution(machine)
	}
	return total
}

func searchMinimumSolution(machine Machine) int {
	maxNumButtonAPresses := machine.prize.x/machine.a.x + 1
	maxNumButtonAPressesY := machine.prize.y/machine.a.y + 1
	if maxNumButtonAPressesY < maxNumButtonAPresses {
		maxNumButtonAPresses = maxNumButtonAPressesY
	}

	maxNumButtonBPresses := machine.prize.x/machine.b.x + 1
	maxNumButtonBPressesY := machine.prize.y/machine.b.y + 1
	if maxNumButtonBPressesY < maxNumButtonBPresses {
		maxNumButtonBPresses = maxNumButtonBPressesY
	}

	minimumCost := 0
	for numButtonAPresses := 0; numButtonAPresses < maxNumButtonAPresses; numButtonAPresses++ {
		if minimumCost > 0 && 3*numButtonAPresses > minimumCost {
			continue
		}
		shortfallX := machine.prize.x - numButtonAPresses*machine.a.x
		if shortfallX%machine.b.x != 0 {
			continue
		}
		numButtonBPresses := shortfallX / machine.b.x
		if numButtonAPresses*machine.a.y+numButtonBPresses*machine.b.y != machine.prize.y {
			continue
		}
		cost := 3*numButtonAPresses + numButtonBPresses
		if minimumCost == 0 || cost < minimumCost {
			minimumCost = cost
		}
	}
	return minimumCost
}
