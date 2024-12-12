package puzzle12

import (
	"bufio"
	"io"
)

type Area struct {
	height int
	width  int
	plots  [][]rune
}

type Coords struct {
	i int
	j int
}

type Fence struct {
	perimeter int
	area      int
	numSides  int
}

func ParseInput12(data io.Reader) Area {
	area := Area{}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		area.height++
		area.width = len(line)
		area.plots = append(area.plots, []rune(line))
	}
	return area
}

func CostFences(area Area) int {
	totalCost := 0
	positionsSeen := make(map[Coords]bool)
	for i := 0; i < area.height; i++ {
		for j := 0; j < area.width; j++ {
			pos := Coords{i, j}
			if positionsSeen[pos] {
				continue
			}
			fence := buildFence(area, pos, positionsSeen)
			totalCost += fence.area * fence.perimeter
		}
	}
	return totalCost
}

func CostFencesNumSides(area Area) int {
	totalCost := 0
	positionsSeen := make(map[Coords]bool)
	for i := 0; i < area.height; i++ {
		for j := 0; j < area.width; j++ {
			pos := Coords{i, j}
			if positionsSeen[pos] {
				continue
			}
			fence := buildFence(area, pos, positionsSeen)
			totalCost += fence.area * fence.numSides
		}
	}
	return totalCost
}

func buildFence(area Area, pos Coords, positionsSeen map[Coords]bool) Fence {
	positionsSeen[pos] = true

	fence := Fence{
		area:      1,
		perimeter: 0,
		numSides:  0,
	}

	adjacents := []Coords{
		{pos.i - 1, pos.j},
		{pos.i, pos.j + 1},
		{pos.i + 1, pos.j},
		{pos.i, pos.j - 1},
	}
	inSameRegion := make([]bool, 4)

	for index, adjacent := range adjacents {
		isInSameRegion := adjacent.i >= 0 && adjacent.i < area.height &&
			adjacent.j >= 0 && adjacent.j < area.width &&
			area.plots[pos.i][pos.j] == area.plots[adjacent.i][adjacent.j]
		inSameRegion[index] = isInSameRegion

		if isInSameRegion {
			if !positionsSeen[adjacent] {
				restOfRegion := buildFence(area, adjacent, positionsSeen)
				fence.area += restOfRegion.area
				fence.perimeter += restOfRegion.perimeter
				fence.numSides += restOfRegion.numSides
			}
		} else {
			fence.perimeter++
		}
	}

	for index, isInSameRegion := range inSameRegion {
		nextIndex := index + 1
		if nextIndex == 4 {
			nextIndex = 0
		}

		diagonallyOpposite := Coords{
			i: pos.i,
			j: pos.j,
		}
		if adjacents[index].i == pos.i {
			diagonallyOpposite.i = adjacents[nextIndex].i
		} else {
			diagonallyOpposite.i = adjacents[index].i
		}
		if adjacents[index].j == pos.j {
			diagonallyOpposite.j = adjacents[nextIndex].j
		} else {
			diagonallyOpposite.j = adjacents[index].j
		}

		diagonallyOppositeInSameRegion := diagonallyOpposite.i >= 0 && diagonallyOpposite.j >= 0 &&
			diagonallyOpposite.i < area.height && diagonallyOpposite.j < area.width &&
			area.plots[pos.i][pos.j] == area.plots[diagonallyOpposite.i][diagonallyOpposite.j]

		isCorner := isInSameRegion && inSameRegion[nextIndex] && !diagonallyOppositeInSameRegion ||
			!isInSameRegion && !inSameRegion[nextIndex]

		if isCorner {
			fence.numSides++
		}
	}

	return fence
}
