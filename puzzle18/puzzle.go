package puzzle18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func (coords Coords) String() string {
	return fmt.Sprintf("%v,%v", coords.x, coords.y)
}

type SearchPosition struct {
	position Coords
	cost     int
}

func ParseInput18(data io.Reader) (positions []Coords) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		positions = append(positions, Coords{x, y})
	}
	return
}

func ShortestPathAfterBytesFallen(bytes []Coords, height, width int, numBytes int) int {
	corruptPositions := make(map[Coords]bool)
	for i := 0; i < numBytes; i++ {
		corruptPositions[bytes[i]] = true
	}

	target := Coords{width - 1, height - 1}
	visitedPositions := make(map[Coords]bool)

	queue := []SearchPosition{{Coords{0, 0}, 0}}
	for len(queue) > 0 {
		searchPosition := queue[0]
		queue = queue[1:]

		if searchPosition.position == target {
			return searchPosition.cost
		}

		if visitedPositions[searchPosition.position] {
			continue
		}
		visitedPositions[searchPosition.position] = true

		adjacentPositions := []Coords{
			{searchPosition.position.x, searchPosition.position.y + 1},
			{searchPosition.position.x, searchPosition.position.y - 1},
			{searchPosition.position.x + 1, searchPosition.position.y},
			{searchPosition.position.x - 1, searchPosition.position.y},
		}
		for _, adjacent := range adjacentPositions {
			if corruptPositions[adjacent] {
				continue
			}
			if adjacent.x < 0 || adjacent.x >= width || adjacent.y < 0 || adjacent.y >= height {
				continue
			}
			queue = append(queue, SearchPosition{adjacent, searchPosition.cost + 1})
		}
	}
	return -1 // -1 to indicate no path available
}

func FindFirstByteObstructingExit(bytes []Coords, height, width int) string {
	time := FindTimeUntilNoPathAvailable(bytes, height, width, 0, len(bytes)-1)
	return bytes[time].String()
}

func FindTimeUntilNoPathAvailable(bytes []Coords, height, width int, lowerBound, upperBound int) int {
	if lowerBound == upperBound {
		if ShortestPathAfterBytesFallen(bytes, height, width, upperBound) == -1 {
			return upperBound - 1
		} else {
			return upperBound - 2
		}
	}

	midpoint := (lowerBound + upperBound) / 2
	if ShortestPathAfterBytesFallen(bytes, height, width, midpoint) == -1 {
		return FindTimeUntilNoPathAvailable(bytes, height, width, lowerBound, midpoint)
	} else {
		return FindTimeUntilNoPathAvailable(bytes, height, width, midpoint+1, upperBound)
	}
}
