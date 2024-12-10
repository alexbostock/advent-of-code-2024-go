package puzzle10

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Coords struct {
	i int
	j int
}

func ParseInput10(data io.Reader) ([][]int, error) {
	var parsed [][]int
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, token := range strings.Split(line, "") {
			num, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		parsed = append(parsed, row)
	}
	return parsed, nil
}

func SumScoresOfTrailheads(area [][]int) int {
	sum := 0
	for i, row := range area {
		for j, cell := range row {
			if cell == 0 {
				sum += scoreTrailhead(area, Coords{i, j})
			}
		}
	}
	return sum
}

func scoreTrailhead(area [][]int, pos Coords) int {
	destinations := make(map[Coords]bool)
	findReachable9s(area, pos, destinations)
	return len(destinations)
}

func findReachable9s(area [][]int, pos Coords, destinations map[Coords]bool) {
	if area[pos.i][pos.j] == 9 {
		destinations[pos] = true
		return
	}

	adjacentPositions := []Coords{
		{pos.i - 1, pos.j},
		{pos.i + 1, pos.j},
		{pos.i, pos.j - 1},
		{pos.i, pos.j + 1},
	}

	for _, adjacentPos := range adjacentPositions {
		if adjacentPos.i < 0 || adjacentPos.i >= len(area) || adjacentPos.j < 0 || adjacentPos.j >= len(area[pos.i]) {
			continue
		}
		if area[adjacentPos.i][adjacentPos.j] == area[pos.i][pos.j]+1 {
			findReachable9s(area, adjacentPos, destinations)
		}
	}
}

func CountDistinctTrails(area [][]int) int {
	sum := 0
	for i, row := range area {
		for j, cell := range row {
			if cell == 0 {
				sum += countDistinctTrailsFromTrailhead(area, Coords{i, j})
			}
		}
	}
	return sum
}

func countDistinctTrailsFromTrailhead(area [][]int, pos Coords) int {
	if area[pos.i][pos.j] == 9 {
		return 1
	}

	adjacentPositions := []Coords{
		{pos.i - 1, pos.j},
		{pos.i + 1, pos.j},
		{pos.i, pos.j - 1},
		{pos.i, pos.j + 1},
	}

	score := 0
	for _, adjacentPos := range adjacentPositions {
		if adjacentPos.i < 0 || adjacentPos.i >= len(area) || adjacentPos.j < 0 || adjacentPos.j >= len(area[pos.i]) {
			continue
		}
		if area[adjacentPos.i][adjacentPos.j] == area[pos.i][pos.j]+1 {
			score += countDistinctTrailsFromTrailhead(area, adjacentPos)
		}
	}
	return score
}
