package puzzle11

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseInput11(data io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(data)
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")

	var stones []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			return nil, err
		}
		stones = append(stones, num)
	}
	return stones, nil
}

type SubProblem struct {
	stone     int
	numBlinks int
}

func CountAllStonesAfterNumBlinks(stones []int, numBlinks int) int {
	memoised := make(map[SubProblem]int)

	count := 0
	for _, stone := range stones {
		count += CountStonesAfterNumBlinks(stone, numBlinks, memoised)
	}
	return count
}

func CountStonesAfterNumBlinks(stone int, numBlinks int, memoised map[SubProblem]int) int {
	if numBlinks == 0 {
		return 1
	}

	memoisedSolution, isMemoised := memoised[SubProblem{stone, numBlinks}]
	if isMemoised {
		return memoisedSolution
	}

	count := 0
	transformed := transformStoneOnBlink(stone)
	for _, transformedStone := range transformed {
		count += CountStonesAfterNumBlinks(transformedStone, numBlinks-1, memoised)
	}

	memoised[SubProblem{stone, numBlinks}] = count
	return count
}

func transformStoneOnBlink(stone int) []int {
	var transformed []int
	if stone == 0 {
		transformed = append(transformed, 1)
		return transformed
	}

	stoneStr := strconv.Itoa(stone)
	if len(stoneStr)%2 == 0 {
		stone1Str := stoneStr[0 : len(stoneStr)/2]
		stone1, err := strconv.Atoi(stone1Str)
		if err != nil {
			panic(err)
		}
		transformed = append(transformed, stone1)

		stone2Str := stoneStr[len(stoneStr)/2:]
		stone2, err := strconv.Atoi(stone2Str)
		if err != nil {
			panic(err)
		}
		transformed = append(transformed, stone2)
	} else {
		transformed = append(transformed, stone*2024)
	}
	return transformed
}
