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

func CountStonesAfterNumBlinks(stones []int, numBlinks int) int {
	for i := 0; i < numBlinks; i++ {
		stones = transformStonesOnBlink(stones)
	}
	return len(stones)
}

func transformStonesOnBlink(stones []int) []int {
	transformed := make([]int, 0, len(stones)*2)
	for _, stone := range stones {
		if stone == 0 {
			transformed = append(transformed, 1)
			continue
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
	}
	return transformed
}
